<script lang="ts">
    import instantsearch, { Widget } from "instantsearch.js";
    import { instantMeiliSearch } from "@meilisearch/instant-meilisearch";
    import { MeiliSearch } from "meilisearch";
    import { searchBox } from "instantsearch.js/es/widgets";
    import { onMount } from "svelte";
    import { connectHits } from "instantsearch.js/es/connectors";
    import { Post } from "../../lib/types/post";
    import {
        findFavouritePostsByUserId,
        findPosts,
        getSignedInUserId,
        getUserById
    } from "../../lib/services/api/users_service";
    import PostImage from "../../lib/components/PostImage.svelte";
    import { api } from "../../lib/services/api/api_service.js";
    import LineSeparator from "../../lib/components/design_system/LineSeparator.svelte";
    import { User } from "../../lib/types/user_types";
    import Avatar from "../../lib/components/design_system/Avatar.svelte";

    let posts = new Array<Post>();
    (async function() {
        posts = await findPosts(0);
    })();

    const search = instantsearch({
        indexName: "users",
        searchClient: instantMeiliSearch(
            "http://localhost:7700",
            ""
        )
    });
    const client = new MeiliSearch(({
        host: "http://localhost:7700",
        apiKey: "",
    }));


    let searchQuery = "";
    let prevQuery = "";
    let users = new Array<User>();
    $: {
        if (searchQuery !== prevQuery && searchQuery.trim()) {
            prevQuery = searchQuery;
            client.index("users").search(searchQuery).then(results => {
                users.length = 0;
                for (const hit of results.hits) {
                    getUserById(hit.id).then(user => {
                        users = [...users, user];
                    });
                }
            });
        }
    }

/*    function renderHits(renderOptions, isFirstRender) {
        const { hits: currentHits, widgetParams } = renderOptions;
        hits = currentHits;
        console.log(`${currentHits}`);
    }

    const customHits = connectHits(renderHits);*/

    onMount(() => {
        search.addWidgets([
            searchBox({
                searchAsYouType: true,
                showReset: false,
                container: "#searchbox",
                placeholder: "Search",
                queryHook: (query, hook) => {
                    searchQuery = query;
                },
                cssClasses: {
                    form: "text-center font-bold text-primary-lighterest",
                    input: "w-[90vw] px-14 my-[5vw] h-12 bg-primary-light rounded-xl outline-none border-radius-12",
                },
                templates: {
                    submit({ cssClasses }, { html }) {
                        return html
                        `
                            <span class="top-[1.95rem] left-10 absolute">
                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="#929292" class="w-7 h-7">
                                    <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
                                </svg>
                            </span>
                        `;
                    },
                },
            }),
            // customHits()
        ]);

        search.start();
    });
</script>

<header>
    <div class="flex items-center focus:outline-primary-lighter justify-center rounded" id="searchbox"></div>
    <LineSeparator/>
</header>

{#if !(searchQuery.trim())}
    <main class="grid-rows">
        {#each posts as post, i}
            <a href={`/p/${post?.id}`}>
                <PostImage imgSrc={`${api.defaults.baseURL}/cdn/images/${post?.images[0]}`} className="h-36"/>
            </a>
        {/each}
    </main>
{:else}
    <main class="flex flex-col justify-start gap-6">
        <div class="ml-[8vw] mt-6 flex items-center gap-7">
            <span>
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="#fff" class="w-12 h-12">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
                </svg>
            </span>
            <span class="font-sans font-bold ">{searchQuery}</span>
        </div>
        {#each users as user, _}
            <div class="ml-[5vw] flex items-center gap-7">
                <Avatar src={user.avatar} size="vsm"/>
                <div class="flex flex-col font-bold">
                    <a href="../u/{user.id}" class="text-white">{user.username}</a>
                    <span class="text-xs text-primary-lighterest">{user.name}</span>
                </div>
            </div>
        {/each}
    </main>
{/if}

<style>
    .grid-rows {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-auto-rows: minmax(min-content, max-content);
        grid-auto-flow: dense;
    }
</style>
