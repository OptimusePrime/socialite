<script lang="ts">
    import Button from "../lib/components/design_system/Button.svelte";
    import { getRefreshToken, isSignedIn } from "../lib/services/api/users_service";
    import { goto } from "@roxi/routify";
    import { signOut } from "../lib/services/api/users_service.js";

    let user: string;
    (async function() {
        user = await getRefreshToken();
    })();

    if (!isSignedIn()) {
        $goto("/auth/register");
    }
</script>

<main>
    <div>
        <a href="/auth/login">Login</a>
        <a href="/auth/register">Register</a>
        <a href="/u/jeff">Profile</a>
        <Button on:click={signOut}>Sign Out</Button>
    </div>
    <p>Token: {user}</p>
<!--    <p>{getUser()}</p>-->
</main>