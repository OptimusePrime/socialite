<script lang="ts">
    import { Post } from "../types/post";
    import { ActionSheet, ActionSheetButtonStyle } from "@capacitor/action-sheet";
    import { Clipboard } from "@capacitor/clipboard";
    import {
        countPostLikes,
        createLike, deleteLike,
        deletePost,
        DeletePostErrors,
        getSignedInUserId,
        isPostLiked
    } from "../services/api/users_service";
    import { goto } from "@roxi/routify";
    import { Toast } from "@capacitor/toast";
    import { onMount } from "svelte";

    export let post: Post;

    async function getPostConfigOptions(): Promise<({title: string; style: ActionSheetButtonStyle} | {title: string;})[]> {
        if (post.poster.id === await getSignedInUserId()) {
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
        }

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

        switch (result.index) {
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
        }
    }

    function getRelativeTime(post: Post): string {
        const diff = (Date.now() - new Date(post?.createdAt).getTime()) / (1000*60*60);

        if (diff < 1) {
            return "Less than an hour ago";
        }
        if (diff >= 1 && diff < 24) {
            return `${Math.round(diff)} hours ago`;
        }
        if (diff >= 24 && diff < (24*365)) {
            return `${Math.trunc(diff / (24*30))} months ago`;
        }
        if (diff >= (24*365)) {
            return `${Math.trunc(diff / (24*365))} years ago`;
        }
    }

    let relativeTime = "";
    let isLiked = false;
    let likeCount = 0;

    onMount(async () => {
        if (await isPostLiked(post.id, await getSignedInUserId())) {
            isLiked = true;
        }

        relativeTime = getRelativeTime(post);

        $: likeCount = await countPostLikes(post.id);
    });

    async function likePost() {
        await createLike(post.id, await getSignedInUserId());
        isLiked = !isLiked;
        likeCount += 1;
    }

    async function unLikePost() {
        await deleteLike(post.id, await getSignedInUserId());
        isLiked = !isLiked;
        likeCount -= 1;
    }
</script>

<aside class="flex justify-between items-center mx-4 my-2.5">
    <div class="flex gap-3 items-center">
        <slot name="posterIcon"/>
        <a class="text-white" href={`/u/${post?.poster?.id}`}>{post?.poster?.name}</a>
    </div>
    <div>
        <button class="hover:cursor-pointer" on:click={async () => await showPostOptions()}>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 12a.75.75 0 11-1.5 0 .75.75 0 011.5 0zM12.75 12a.75.75 0 11-1.5 0 .75.75 0 011.5 0zM18.75 12a.75.75 0 11-1.5 0 .75.75 0 011.5 0z" />
            </svg>
        </button>
    </div>
</aside>
<main class="w-full">
    <slot name="image"/>
    <div>
        <div class="flex justify-between items-center m-4">
            <div class="flex flex-row items-center gap-4">
                <div>
                    {#if !isLiked}
                        <span class="cursor-pointer" on:click={async () => likePost()}>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12z" />
                            </svg>
                        </span>
                    {:else}
                        <span class="cursor-pointer" on:click={async () => unLikePost()}>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="#ed4956" viewBox="0 0 24 24" stroke-width="1.5" stroke="#ed4956" class="w-8 h-8">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12z" />
                            </svg>
                        </span>
                    {/if}
                </div>
                <div>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M8.625 12a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H8.25m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H12m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0h-.375M21 12c0 4.556-4.03 8.25-9 8.25a9.764 9.764 0 01-2.555-.337A5.972 5.972 0 015.41 20.97a5.969 5.969 0 01-.474-.065 4.48 4.48 0 00.978-2.025c.09-.457-.133-.901-.467-1.226C3.93 16.178 3 14.189 3 12c0-4.556 4.03-8.25 9-8.25s9 3.694 9 8.25z" />
                    </svg>
                </div>
                <div>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M7.217 10.907a2.25 2.25 0 100 2.186m0-2.186c.18.324.283.696.283 1.093s-.103.77-.283 1.093m0-2.186l9.566-5.314m-9.566 7.5l9.566 5.314m0 0a2.25 2.25 0 103.935 2.186 2.25 2.25 0 00-3.935-2.186zm0-12.814a2.25 2.25 0 103.933-2.185 2.25 2.25 0 00-3.933 2.185z" />
                    </svg>
                </div>
            </div>
            <div>
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 00-.182.557l1.285 5.385a.562.562 0 01-.84.61l-4.725-2.885a.563.563 0 00-.586 0L6.982 20.54a.562.562 0 01-.84-.61l1.285-5.386a.562.562 0 00-.182-.557l-4.204-3.602a.563.563 0 01.321-.988l5.518-.442a.563.563 0 00.475-.345L11.48 3.5z" />
                </svg>
            </div>
        </div>
        <div class="mt-3 mx-4">
            <p class="font-sans"><span class="font-semibold">{likeCount}</span> {likeCount === 1 ? "like" : "likes"}</p>
        </div>
    </div>
</main>

<footer class="mx-4">
    <div>
        <p class="font-sans"><span class="font-bold">{post?.poster?.username}</span> {post?.caption}</p>
    </div>
    <div>
        <p class="text-primary-lighter text-sm">View all 180 comments</p>
    </div>
    <div class="mt-2 flex items-center gap-1">
        <div>
            <slot name="signedInUserIcon"/>
        </div>
        <input class="h-5 font-sans p-0 w-96 bg-primary placeholder-primary-lightest placeholder-opacity-60 border-none text-[0.8rem] focus:outline-none" placeholder="Add a comment" type="text">
    </div>
    <p class="mt-1.5 font-sans text-[0.75rem] text-primary-lighter">{relativeTime}</p>
</footer>