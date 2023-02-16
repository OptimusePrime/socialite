<script lang="ts">
    import {
        getSignedInUserId,
        isSignedIn,
        isValidRefreshToken,
        refreshAccessToken
    } from "../lib/services/api/users_service";
    import { afterPageLoad, goto } from "@roxi/routify";
    import { onMount } from "svelte";

    let isRefreshTokenValid = false;
    let userId;

    onMount(async () => {
        userId = await getSignedInUserId();
    });

/*    if (!$accessTokenStore || $accessTokenStore === "") {
        $goto("/auth");
    }*/

    let url = "";
    $afterPageLoad(async () => {
        url = window.location.href.replace(window.location.protocol + "//" + window.location.host, "");
        isRefreshTokenValid = await isSignedIn();
    });

    async function goToUserProfilePage() {
        let userId = await getSignedInUserId();

        $goto(`/u/${userId}`);
    }
</script>

<slot/>

{#if (isRefreshTokenValid)}
    <div class="fixed left-0 bottom-0 w-full">
        <hr class="mt-3 border-t-[2px] border-primary-light">
        <nav class="p-3 flex justify-between items-center w-full bg-primary">
            <a href="/home">
                {#if (url === "/home")}
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#fff" class="w-8 h-8">
                        <path d="M11.47 3.84a.75.75 0 011.06 0l8.69 8.69a.75.75 0 101.06-1.06l-8.689-8.69a2.25 2.25 0 00-3.182 0l-8.69 8.69a.75.75 0 001.061 1.06l8.69-8.69z" />
                        <path d="M12 5.432l8.159 8.159c.03.03.06.058.091.086v6.198c0 1.035-.84 1.875-1.875 1.875H15a.75.75 0 01-.75-.75v-4.5a.75.75 0 00-.75-.75h-3a.75.75 0 00-.75.75V21a.75.75 0 01-.75.75H5.625a1.875 1.875 0 01-1.875-1.875v-6.198a2.29 2.29 0 00.091-.086L12 5.43z" />
                    </svg>
                {:else}
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="#fff" class="w-8 h-8">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 12l8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25" />
                    </svg>
                {/if}
            </a>

            <a href="/search">
                {#if (url === "/search")}
                    <svg class="w-8 h-8" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M21.75 21.75L16.1199 16.1199M16.1199 16.1199C17.6437 14.5961 18.4998 12.5294 18.4998 10.3744C18.4998 8.21949 17.6437 6.15278 16.1199 4.62899C14.5961 3.1052 12.5294 2.24915 10.3744 2.24915C8.21949 2.24915 6.15278 3.1052 4.62899 4.62899C3.1052 6.15278 2.24915 8.21949 2.24915 10.3744C2.24915 12.5294 3.1052 14.5961 4.62899 16.1199C6.15278 17.6437 8.21949 18.4998 10.3744 18.4998C12.5294 18.4998 14.5961 17.6437 16.1199 16.1199Z" stroke="white" stroke-width="3.5" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>
                {:else}
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#fff" class="w-8 h-8">
                        <path fill-rule="evenodd" d="M10.5 3.75a6.75 6.75 0 100 13.5 6.75 6.75 0 000-13.5zM2.25 10.5a8.25 8.25 0 1114.59 5.28l4.69 4.69a.75.75 0 11-1.06 1.06l-4.69-4.69A8.25 8.25 0 012.25 10.5z" clip-rule="evenodd" />
                    </svg>
                {/if}
            </a>
            <a href="/post/new">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="#fff" class="w-8 h-8">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v6m3-3H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
            </a>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="#fff" class="w-8 h-8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M8.625 12a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H8.25m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H12m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0h-.375M21 12c0 4.556-4.03 8.25-9 8.25a9.764 9.764 0 01-2.555-.337A5.972 5.972 0 015.41 20.97a5.969 5.969 0 01-.474-.065 4.48 4.48 0 00.978-2.025c.09-.457-.133-.901-.467-1.226C3.93 16.178 3 14.189 3 12c0-4.556 4.03-8.25 9-8.25s9 3.694 9 8.25z" />
            </svg>
            <span on:click={goToUserProfilePage}>
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="#fff" class="cursor-pointer w-8 h-8">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z" />
                </svg>
            </span>
        </nav>
    </div>
{/if}
