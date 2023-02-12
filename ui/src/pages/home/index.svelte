<script lang="ts">
    import LineSeparator from "../../lib/components/design_system/LineSeparator.svelte";
    import { findPosts, findWhoUserFollows, getSignedInUserId } from "../../lib/services/api/users_service.js";
    import Post from "../../lib/components/Post.svelte";
    import PostImage from "../../lib/components/PostImage.svelte";
    import Header from "../../lib/components/Header.svelte";

    // let isUserFollowingSomeone = false;
    let posts = new Array<Post>();
    (async function() {
/*        let [following] = await findWhoUserFollows(await getSignedInUserId());
        isUserFollowingSomeone = following.length > 0;*/

        posts = await findPosts(0);
    })();

/*    if(isUserFollowingSomeone) {
        document.body.style.backgroundImage = "/home_no_friends_background.svg";
    }*/
</script>

<Header title="Socialite">
    <a slot="first" href="/favourites">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="#fff" class="w-8 h-8">
            <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 00-.182.557l1.285 5.385a.562.562 0 01-.84.61l-4.725-2.885a.563.563 0 00-.586 0L6.982 20.54a.562.562 0 01-.84-.61l1.285-5.386a.562.562 0 00-.182-.557l-4.204-3.602a.563.563 0 01.321-.988l5.518-.442a.563.563 0 00.475-.345L11.48 3.5z" />
        </svg>
    </a>
    <button slot="second">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0" />
        </svg>
    </button>
</Header>
<!--<header class="mx-6 my-4 flex justify-between items-center">
    <h2 class="font-sans font-[750] tracking-wide text-xl">Socialite</h2>
    <div class="flex gap-5 items-center">
        <a href="/favourites">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="#fff" class="w-8 h-8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 00-.182.557l1.285 5.385a.562.562 0 01-.84.61l-4.725-2.885a.563.563 0 00-.586 0L6.982 20.54a.562.562 0 01-.84-.61l1.285-5.386a.562.562 0 00-.182-.557l-4.204-3.602a.563.563 0 01.321-.988l5.518-.442a.563.563 0 00.475-.345L11.48 3.5z" />
            </svg>
        </a>
        <button>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
                <path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0" />
            </svg>
        </button>
    </div>
</header>
<LineSeparator/>-->

<main class="w-full h-full">
    {#if (!true)}
        <div class="mt-44">
            <div class="w-full flex flex-col items-center">
                <img src="/no_users_icon.svg" class="w-7/12 h-full" alt="No followers icon">
                    <h1 class="page-title-shadow text-4xl text-center font-sans font-extrabold">You need to follow at least <span class="text-indigo">1 person</span> to view posts</h1>
                </div>
            </div>
    {:else }
        {#each posts as post, i}
            <Post {post}>
                <div slot="posterIcon">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="#fff" class="w-11 h-11">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M17.982 18.725A7.488 7.488 0 0012 15.75a7.488 7.488 0 00-5.982 2.975m11.963 0a9 9 0 10-11.963 0m11.963 0A8.966 8.966 0 0112 21a8.966 8.966 0 01-5.982-2.275M15 9.75a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                </div>
                <PostImage slot="image" imgSrc={`http://192.168.1.102:3000/cdn/images/${post?.images?.at(0)}`}/>
                <div slot="signedInUserIcon">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="#ffff" class="w-7 h-7">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M17.982 18.725A7.488 7.488 0 0012 15.75a7.488 7.488 0 00-5.982 2.975m11.963 0a9 9 0 10-11.963 0m11.963 0A8.966 8.966 0 0112 21a8.966 8.966 0 01-5.982-2.275M15 9.75a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                </div>
            </Post>
        {/each}
    {/if}
</main>
<style>
    :global(html) {
        /*background-image: url("/home_no_friends_background.svg");*/
    }
</style>



