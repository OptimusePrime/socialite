<script lang="ts">
    import Header from "../../lib/components/Header.svelte";
    import Post from "../../lib/components/Post.svelte";
    import { Post as PostType } from "../../lib/types/post";
    import { onMount } from "svelte";
    import PostImage from "../../lib/components/PostImage.svelte";
    import { goto } from "@roxi/routify";
    import LineSeparator from "../../lib/components/design_system/LineSeparator.svelte";
    import { findFavouritePostsByUserId, findPosts, getSignedInUserId } from "../../lib/services/api/users_service.js";
    import { api } from "../../lib/services/api/api_service.js";

    let posts = new Array<PostType>();

    onMount(async () => {
        posts = await findFavouritePostsByUserId(await getSignedInUserId());
    });
</script>

<Header title="Favourites">
    <span slot="first" on:click={() => $goto("/")}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-12 h-12">
             <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5L8.25 12l7.5-7.5" />
        </svg>
    </span>
    <span slot="second" class="opacity-0">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-12 h-12">
            <path fill-rule="evenodd" d="M12 3.75a.75.75 0 01.75.75v6.75h6.75a.75.75 0 010 1.5h-6.75v6.75a.75.75 0 01-1.5 0v-6.75H4.5a.75.75 0 010-1.5h6.75V4.5a.75.75 0 01.75-.75z" clip-rule="evenodd" />
        </svg>
    </span>
</Header>

<main class="grid-rows">
    {#each posts as post, i}
        <a href={`/p/${post?.id}`}>
            <PostImage imgSrc={`${api.defaults.baseURL}/cdn/images/${post?.images[0]}`} className="h-36"/>
        </a>
    {/each}
</main>

<style>
    .grid-rows {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-auto-rows: minmax(min-content, max-content);
        grid-auto-flow: dense;
    }
</style>