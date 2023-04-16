import { api } from "./api_service";
import { accessTokenStore } from "../../stores";
import * as jose from "jose";
import { isPast } from "date-fns";
import { Preferences } from "@capacitor/preferences";
import { User } from "../../types/user_types";
import type { Photo } from "@capacitor/camera";
import { Post } from "../../types/post";
import { goto } from "@roxi/routify";
import type { Favourite } from "../../types/favourite";

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

export enum FollowUserErrors {
    FOLLOWEE_NOT_FOUND = "followee not found",
    FOLLOWER_NOT_FOUND = "follower not found",
    CANNOT_FOLLOW_TWICE_SAME_USER = "cannot follow twice the same user",
    INTERNAL_SERVER_ERROR = "internal server error"
}

export enum FindFollowsErrors {
    INVALID_ID = "invalid id",
    FOLLOW_NOT_FOUND = "follow not found",
}

export enum DeletePostErrors {
    INTERNAL_SERVER_ERROR = "internal server error"
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

    const response = (await api.post("/auth/refresh", {
        refreshToken,
    }).catch(err => {
        if (err.response) {
            return err.response;
        }
    }));

    if (response.data.message) {
        const error = response.data.message as RefreshAccessTokenErrors;
        if(error === RefreshAccessTokenErrors.INVALID_REFRESH_TOKEN) {
            await signOut();
        }

        return error;
    }

    accessTokenStore.set(response.data.accessToken);
    return null;
}

export async function isValidRefreshToken(): Promise<boolean> {
    const { value: refreshToken } = await Preferences.get({
        key: "socialite_refreshToken",
    });

    if(!refreshToken || refreshToken.length <= 0) {
        return false;
    }

    let payload;
    try {
        payload = jose.decodeJwt(refreshToken);
    } catch (err) {
        if (err) {
            return false;
        }
    }

    return Date.now() < payload.exp * 1000;
}

export async function getSignedInUserId(): Promise<string> {
    let accessToken: string;
    const unsubscribe = accessTokenStore.subscribe(token => accessToken = token);
    if(!accessToken) {
        await refreshAccessToken();
    }

    let payload;
    try {
        payload = jose.decodeJwt(accessToken);
    } catch (err) {
        if (err) {
            await signOut();
            return null;
        }
    }

    unsubscribe();
    return payload.userId.toString();
}

export async function getRefreshToken(): Promise<string> {
    const refreshToken = await Preferences.get({
        key: "socialite_refreshToken",
    });

    return refreshToken.value;
}

export async function isSignedIn(): Promise<boolean> {
    return (
        await Preferences.get({
            key: "socialite_refreshToken",
        })
    ).value as unknown as boolean;
}

export async function getUserById(id: string): Promise<User> {
    if (!id) {
        return null;
    }

    const res = await api.get(`/users/${id}`);
    console.log("User id:" + id);
    console.log("Response data: " + res.data.name);


    const user: User = new User();
    for (const key of Object.keys(res.data)) {
        user[key] = res.data[key];
    }
    console.log("User: " + user.name);

    return user;
}

export async function getSignedInUser(): Promise<User> {
    return await getUserById(await getSignedInUserId());
}

export async function signOut() {
    await Preferences.remove({
        key: "socialite_refreshToken",
    });
    accessTokenStore.set("");
    // $goto("/auth");
}

export async function findWhoUserFollows(userId: string): Promise<{ error: FindFollowsErrors; users: User[] }> {
    const response = await api.get(`/follows/userFollows/${userId}`);

    if (response?.data?.message) {
        return {
            error: response.data.message as FindFollowsErrors,
            users: null,
        };
    }

    return {
        error: null,
        users: response.data as User[],
    };
}

export async function findFollowersOfUser(followerId: string): Promise<{ error: FindFollowsErrors; users: User[]; }> {
    const response = await api.get(`/follows/followersOfUser/${followerId}`);

    if (response?.data?.message) {
        return {
            error: response.data.message as FindFollowsErrors,
            users: null,
        };
    }

    return {
        error: null,
        users: response.data as User[],
    };
}

export async function findFollow(followerId: string, followeeId: string): Promise<boolean> {
    const response = await api.put( "/follows", {
        follower: followerId,
        followee: followeeId,
    }).catch(err => err.response);

    if(response.status === 200) {
        return true;
    }

    return false;
}

export async function followUser(followeeId: string): Promise<FollowUserErrors> {
    const response = await api.post("/follows", {
        follower: await getSignedInUserId(),
        followee: followeeId,
    }).catch(err => err.response);

    if (response.data.message) {
        return response.data.message as FollowUserErrors;
    }

    return null;
}

export async function unfollowUser(followeeId: string): Promise<FollowUserErrors> {
    const response = await api.delete("/follows", {
        data: {
            follower: await getSignedInUserId(),
            followee: followeeId,
        }
    }).catch(err => err.response);

    if (response.data.message) {
        return response.data.message as FollowUserErrors;
    }

    return null;
}

export async function createPost(image: Blob, caption: string, location: string): Promise<string> {
    const formData = new FormData();
    formData.append("image", image);
    formData.set("caption", caption);
    formData.set("location", location);
    formData.set("poster", await getSignedInUserId());

    const id = (await api.post("/posts",
        formData,
        {
            headers: {
                "Content-Type": "multipart/form-data",
            }
        }
    )).data.id;

    return id;
}

function retrievedPostToPost(retrievedPost: any): Post {
    const user = new User();
    user.id = retrievedPost.poster.id;
    user.username = retrievedPost.poster.username;
    user.email = retrievedPost.poster.email;
    user.name = retrievedPost.poster.name;
    user.password = retrievedPost.poster.password;
    user.birthDate = new Date(retrievedPost.poster.birthDate);

    return new Post(retrievedPost.id, retrievedPost.createdAt, retrievedPost.updatedAt, retrievedPost.caption, retrievedPost.images, user, retrievedPost.location);
}

export async function findPostById(id: string): Promise<Post> {
    return retrievedPostToPost((await api.get(`/posts/one/${id}`)).data.posts[0]);
}

export async function findPostsByPosterId(posterId: string, limit: number): Promise<Post[]> {
    const retrievedPosts = (await api.get(`/posts/many/${posterId}`, {
        data: {
            limit,
        }
    })).data.posts;

    if (!retrievedPosts) {
        return null;
    }

    const posts = new Array<Post>();
    for (const retrievedPost of retrievedPosts) {
        posts.push(retrievedPostToPost(retrievedPost));
    }

    return posts;
}

export async function findPosts(limit: number): Promise<Post[]> {
    const retrievedPosts = (await api.get("/posts", {
        data: {
            limit,
        }
    })).data.posts;

    const posts = new Array<Post>();
    for (const retrievedPost of retrievedPosts) {
        posts.push(retrievedPostToPost(retrievedPost));
    }

    return posts;
}

export async function deletePost(id: string): Promise<DeletePostErrors> {
    const response = (await api.delete(`/posts/${id}`).catch(err => err.response));

    if (response.data.message) {
        return response.data.message as DeletePostErrors;
    }

    return null;
}

export async function createLike(postId: string, userId: string) {
    (await api.post("/likes", {
        post: postId,
        user: userId,
    }).catch(err => err.response));
}

export async function deleteLike(postId: string, userId: string) {
    (await api.delete("/likes", {
        data: {
            post: postId,
            user: userId,
        },
    }).catch(err => err.response));
}

export async function isPostLiked(postId: string, userId: string): Promise<boolean> {
    const response = (await api.get(`/likes/${postId}`).catch(err => err.response));

    if (response.data.likes) {
        for (const like of response.data.likes) {
            if (like.user === userId) {
                return true;
            }
        }
    }

    return false;
}

export async function countPostLikes(postId: string): Promise<number> {
    const response = (await api.get(`/likes/${postId}`).catch(err => err.response));

    return response?.data?.likes?.length || 0;
}

export async function updateUser(updateData: { biography: string; gender: string; username: string; name: string; pronouns: string; }) {
    // await refreshAccessToken();
    let accessToken: string;
    const unsubscribe = accessTokenStore.subscribe(token => accessToken = token);
    await api.patch("/users", {
        ...updateData,
        accessToken,
    });
    unsubscribe();
}

export async function findFavouritesByUserId(userId: string): Promise<Favourite[]> {
    const response = await api.get(`/favourites/${userId}`);

    // console.log(response.data);
    return response.data as Favourite[];
}

export async function findFavouritePostsByUserId(userId: string): Promise<Post[]> {
    const favourites = await findFavouritesByUserId(userId);
    if (!favourites) {
        return null;
    }

    if (favourites.length  <= 0) {
        return null;
    }
    // console.log(favourites);

    const posts = new Array<Post>();
    for (const favourite of favourites) {
        posts.push(favourite.post);
    }

    return posts;
    // console.log(posts);
}

export async function isPostFavourited(userId: string, postId: string): Promise<boolean> {
    const favouritePosts = await findFavouritePostsByUserId(userId);
    if (!favouritePosts) {
        return false;
    }
    console.log(favouritePosts);

    for (const favouritePost of favouritePosts) {
        if (favouritePost.id === postId) {
            return true;
        }
    }

    return false;
}

export async function createFavourite(userId: string, postId: string) {
    await api.post("/favourites", {
        userId,
        postId,
    });
}

export async function deleteFavourite(userId: string, postId: string) {
    await api.delete("/favourites", {
        data: {
            userId,
            postId,
        }
    });
}