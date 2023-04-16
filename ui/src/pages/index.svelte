<script lang="ts">
    import Button from "../lib/components/design_system/Button.svelte";
    import {
        deletePost, DeletePostErrors,
        getRefreshToken,
        getSignedInUser, getSignedInUserId, isSignedIn,
    } from "../lib/services/api/users_service";
    import { goto } from "@roxi/routify";
    import { signOut } from "../lib/services/api/users_service.js";
    import { User } from "../lib/types/user_types";
    import { api } from "../lib/services/api/api_service.js";
    import axios from "axios";
    import { CapacitorHttp } from "@capacitor/core";
    import { ActionSheet, ActionSheetButtonStyle } from "@capacitor/action-sheet";
    import { Clipboard } from "@capacitor/clipboard";
    import { Toast } from "@capacitor/toast";
    // import { Post } from "../lib/components/Post.svelte";

    let userData = new User();
    let user: string;
    (async function() {
        user = await getRefreshToken();
        userData =  await getSignedInUser();
    })();

    let signedIn: boolean;
    (async function () {
        signedIn = await isSignedIn();
        if (!signedIn) {
            $goto("/auth");
            return;
        }
        $goto("/home");
        userData = await getSignedInUser();
    })();

    async function getPostConfigOptions(): Promise<({title: string; style: ActionSheetButtonStyle} | {title: string;})[]> {
/*        // if (post.poster.id === await getSignedInUserId()) {
            return [
                {
                    title: "Go to post",
                },
                {
                    title: "Link",
                },
                {
                    title: "Delete",
                    style: ActionSheetButtonStyle.Destructive,
                }
            ];
        }*/

        return [
            {
                title: "Go to post",
            },
            {
                title: "Link",
            }
        ];
    }

    async function showPostOptions() {
        const result = await ActionSheet.showActions({
            title: "Post Options",
            options: await getPostConfigOptions(),
        });

/*        switch (0) {
        case 0:
            $goto(`/p/${post.id}`);
            break;
        case 1:
            await Clipboard.write({
                string: `/p/${post.id}`,
            });
            break;
        case 2:
            const err = await deletePost(post.id);
            if (err) {
                if (err === DeletePostErrors.INTERNAL_SERVER_ERROR) {
                    await Toast.show({
                        text: "An internal server error occurred. Please try again later."
                    });
                }
            }
            break;
        }*/
    }
</script>

<main>
    <div>
        <a href="/auth/login">Login</a>
        <a href="/auth/register">Register</a>
        <a href="/u/jeff">Profile</a>
        <a href="/auth">Auth</a>
        <a href="/post/new">New Post</a>
        <a href="/home">Home</a>
        <Button on:click={signOut}>Sign Out</Button>
<!--        <Button on:click={async () => {
        }}>Test</Button>-->
        <br>
        <Button on:click={showPostOptions}>Test</Button>
    </div>
    <p>Token: {user}</p>
    <p>Name: {userData?.name}</p>
<!--    <p>{getUser()}</p>-->
</main>