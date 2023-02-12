<script lang="ts">
    import Button from "../../lib/components/design_system/Button.svelte";
    import { goto } from "@roxi/routify";
</script>

<main>
    <div class="flex flex-col items-center gap-5 mt-72">
        <h1 class="font-extrabold text-[2.6rem] page-title-shadow">Socialite</h1>
        <div class="flex flex-col items-center gap-4">
            <Button on:click={() => $goto("/auth/login")} className="shadow-xl" size="xl">Log In</Button>
            <Button on:click={() => $goto("/auth/register")} className="shadow-xl" size="xl">Register</Button>
        </div>
    </div>
</main>

<style>
    :global(html) {
        overflow: hidden;
        background-image: url("/auth_index_background.svg");
    }
</style>