<script lang="ts">
    import PostImage from "../../lib/components/PostImage.svelte";
    import Post from "../../lib/components/Post.svelte";
    import Header from "../../lib/components/Header.svelte";
    import { goto } from "@roxi/routify";
    import { onMount } from "svelte";
    import { findPostById } from "../../lib/services/api/users_service";
    // import { params } from "@roxi/routify";
    import { Post as PostType } from "../../lib/types/post";
    import { params } from "@roxi/routify";
    import { api } from "../../lib/services/api/api_service.js";

    let post: PostType;
    onMount(async () => {
        post = await findPostById($params.postId);
    });
</script>

<Header title="Post">
    <span slot="first" on:click={() => $goto("/")}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-10 h-10">
            <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5L8.25 12l7.5-7.5" />
        </svg>
    </span>
</Header>

<Post {post}>
    <div slot="posterIcon">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="w-11 h-11">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17.982 18.725A7.488 7.488 0 0012 15.75a7.488 7.488 0 00-5.982 2.975m11.963 0a9 9 0 10-11.963 0m11.963 0A8.966 8.966 0 0112 21a8.966 8.966 0 01-5.982-2.275M15 9.75a3 3 0 11-6 0 3 3 0 016 0z" />
        </svg>
    </div>
    <PostImage slot="image" imgSrc={`${api.defaults.baseURL}/cdn/images/${post?.images?.at(0)}`}/>
    <div slot="signedInUserIcon">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-7 h-7">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17.982 18.725A7.488 7.488 0 0012 15.75a7.488 7.488 0 00-5.982 2.975m11.963 0a9 9 0 10-11.963 0m11.963 0A8.966 8.966 0 0112 21a8.966 8.966 0 01-5.982-2.275M15 9.75a3 3 0 11-6 0 3 3 0 016 0z" />
        </svg>
    </div>
</Post>