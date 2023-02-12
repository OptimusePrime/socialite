<script lang="ts">
    import Button from "../lib/components/design_system/Button.svelte";
    import {
        getRefreshToken,
        getSignedInUser, isSignedIn,
    } from "../lib/services/api/users_service";
    import { goto } from "@roxi/routify";
    import { signOut } from "../lib/services/api/users_service.js";
    import { User } from "../lib/types/user_types";
    import { api } from "../lib/services/api/api_service.js";
    import axios from "axios";
    import { CapacitorHttp } from "@capacitor/core";

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
            $goto("/auth/login");
            return;
        }
        $goto("/home");
        userData = await getSignedInUser();
    })();
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
    </div>
    <p>Token: {user}</p>
    <p>Name: {userData?.name}</p>
<!--    <p>{getUser()}</p>-->
</main>