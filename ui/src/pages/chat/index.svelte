<script lang="ts">
    import Header from "../../lib/components/Header.svelte";
    import { goto } from "@roxi/routify";
    import Avatar from "../../lib/components/design_system/Avatar.svelte";

    let users = [];
</script>

<Header title="Messages">
    <span slot="first" class="cursor-pointer" on:click={() => $goto("/home")}>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-12 h-12">
          <path fill-rule="evenodd" d="M7.72 12.53a.75.75 0 010-1.06l7.5-7.5a.75.75 0 111.06 1.06L9.31 12l6.97 6.97a.75.75 0 11-1.06 1.06l-7.5-7.5z" clip-rule="evenodd" />
        </svg>
    </span>
    <span slot="second" class="opacity-0">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-12 h-12">
            <path fill-rule="evenodd" d="M12 3.75a.75.75 0 01.75.75v6.75h6.75a.75.75 0 010 1.5h-6.75v6.75a.75.75 0 01-1.5 0v-6.75H4.5a.75.75 0 010-1.5h6.75V4.5a.75.75 0 01.75-.75z" clip-rule="evenodd" />
        </svg>
    </span>
</Header>

<main class="flex flex-col justify-start gap-6">
    {#each users as user, _}
        <div class="ml-[5vw] flex items-center gap-7">
            <Avatar src={user.avatar} size="vsm"/>
            <div class="flex flex-col font-bold">
                <a href="{user.id}">{user.username}</a>
                <span class="text-xs text-primary-lighterest">{user.name}</span>
            </div>
        </div>
    {/each}
</main>