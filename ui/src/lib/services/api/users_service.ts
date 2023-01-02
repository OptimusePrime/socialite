import { api } from "./api_service";
import { accessTokenStore } from "../../stores";
import * as jose from "jose";
import { isPast } from "date-fns";
import { Preferences } from "@capacitor/preferences";
import type { User } from "../../types/user_types";
import { goto } from "@roxi/routify";

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

export async function registerUser(data: { username: string; email: string; name: string; password: string }): Promise<RegisterUserErrors> {
    const response = await api.post("/users", data);
    if (response.data.message) {
        return response.data.message as RegisterUserErrors;
    }
/*    return (await api.post("/users", data).catch(resp => {
        console.log(resp.response.data);
        return resp.response;
    }).then(resp => {
        console.log(resp.data);
        return resp;
    })).data.message as RegisterUserErrors;*/
}

export async function loginUser(data: { email: string; password: string }): Promise<LoginUserErrors> {
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
    }).catch(reason => {
        console.log("Reason: " + reason);
    }).then(response => response);
    console.log("Response: " + response);
    /*    if (response.data.message) {
        return response.data.message as RefreshAccessTokenErrors;
    }
    accessTokenStore.set(response.data.accessToken);*/
    return null;
}

export async function isValidRefreshToken(): Promise<boolean> {
    const { value: refreshToken } = await Preferences.get({
        key: "socialite_refreshToken",
    });

    if(!refreshToken || refreshToken.length <= 0) {
        return false;
    }
    const payload = jose.decodeJwt(refreshToken);
    return Date.now() < payload.exp * 1000;
}

export async function getUser(): Promise<string> {
    let accessToken: string;
    const unsubscribe = accessTokenStore.subscribe(token => accessToken = token);
    console.log("Hello, world!");
    if(!accessToken) {
        console.log("Test 1");
        await refreshAccessToken();
        console.log("Test 4");
    }

    const payload = jose.decodeJwt(accessToken);
    console.log(payload);
    unsubscribe();
    return accessToken;
}

export async function getRefreshToken(): Promise<string> {
    const refreshToken = await Preferences.get({
        key: "socialite_refreshToken",
    });

    return refreshToken.value;
}

export async function isSignedIn(): Promise<boolean> {
    return !(await Preferences.get({
        key: "socialite_refreshToken",
    }));
}

export async function getUserById(id: string): Promise<User> {
    const res = await api.get(`/users/${id}`);
    return res.data as User;
}

export async function signOut() {
    await Preferences.remove({
        key: "socialite_refreshToken",
    });

    $goto("/");
}