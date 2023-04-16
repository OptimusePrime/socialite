<script lang="ts">
    import Header from "../../lib/components/Header.svelte";
    import { goto } from "@roxi/routify";
    import Avatar from "../../lib/components/design_system/Avatar.svelte";
    import DateSeparator from "../../lib/components/DateSeparator.svelte";


</script>

<Header>
    <div slot="first" class="flex items-center gap-3" on:click={() => $goto("/")}>
        <span>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-12 h-12">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5L8.25 12l7.5-7.5" />
            </svg>
        </span>
        <div class="flex items-center font-sans font-bold gap-2">
            <Avatar size="xsm"/>
            <div class="flex flex-col items-start flex-row">
                <span class="text-white text-xl">OptimusePrime</span>
                <div class="flex text-primary-lighter text-sm gap-4">
                    <span>He/Him</span>
                    <span>Active 2h ago</span>
                </div>
            </div>
        </div>
    </div>
    <span slot="second" class="opacity-0">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-10 h-10">
            <path fill-rule="evenodd" d="M12 3.75a.75.75 0 01.75.75v6.75h6.75a.75.75 0 010 1.5h-6.75v6.75a.75.75 0 01-1.5 0v-6.75H4.5a.75.75 0 010-1.5h6.75V4.5a.75.75 0 01.75-.75z" clip-rule="evenodd" />
        </svg>
    </span>
</Header>

<main class="flex justify-center mt-6">
    <DateSeparator date={new Date(Date.now())}></DateSeparator>
    <div>
        <input placeholder="Message" class="w-[90vw] right-0 left-[5vw] bottom-[8.5vh] font-bold px-5 absolute text-primary-lighterest text-sm h-12 bg-primary-light rounded-xl outline-none border-radius-12" type="text">
    </div>
</main>
