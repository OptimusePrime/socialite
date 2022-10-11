import { api } from "./api_service";
import { accessTokenStore } from "../../stores";
import * as jose from "jose";
import { isPast } from "date-fns";
import { Preferences } from "@capacitor/preferences";

export enum RegisterUserErrors {
    USERNAME_NOT_UNIQUE = "user with the specified username already exists",
    EMAIL_NOT_UNIQUE = "user with the specified email already exists",
    INVALID_EMAIL = "invalid email address",
    INVALID_NAME = "invalid name",
    INVALID_USERNAME = "invalid username",
    INVALID_PASSWORD = "invalid password",
    INTERNAL_SERVER_ERROR = "internal server error",
}

export enum LoginUserErrors {
    INVALID_EMAIL = "invalid email address",
    INVALID_PASSWORD = "invalid password",
    INTERNAL_SERVER_ERROR = "internal server error",
}

export enum RefreshAccessTokenErrors {
    INVALID_REFRESH_TOKEN = "invalid refresh token",

}

export async function registerUser(data: { username: string, email: string, name: string, password: string }): Promise<RegisterUserErrors> {
    return (await api.post("/users", data).catch(resp => {
        console.log(resp.response.data);
        return resp.response;
    }).then(resp => {
        console.log(resp.data);
        return resp;
    })).data.message;
}

export async function loginUser(data: { email: string, password: string }): Promise<LoginUserErrors> {
    const response = await api.post("/auth/login", data).catch(resp => resp.response);

    if (response.data.message) {
        return response.data.message as LoginUserErrors;
    }

    accessTokenStore.set(response.data.accessToken);
    await Preferences.set({
        key: "socialite_refreshToken",
        value: response.data.refreshToken,
    });
}

export async function refreshAccessToken(): Promise<RefreshAccessTokenErrors> {
    accessTokenStore.set("");
    const { value: refreshToken } = await Preferences.get({
        key: "socialite_refreshToken",
    });

    const response = await api.post("/auth/refresh", {
        refreshToken,
    }).catch(resp => resp.response);
    if (response.data.message) {
        return response.data.message as RefreshAccessTokenErrors;
    }
    accessTokenStore.set(response.data.accessToken);
}

export async function getAccessToken(): Promise<string> {
    let accessToken: string;
    const unsubscribe = accessTokenStore.subscribe(token => accessToken = token);

    const expirationTime = jose.decodeJwt(accessToken).exp;
    if (isPast(expirationTime) || !accessToken) {
        await refreshAccessToken();
    }

    unsubscribe();
    return accessToken;
}