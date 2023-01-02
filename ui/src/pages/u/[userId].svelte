<script>
    import Helper from "../../lib/components/design_system/Helper.svelte";
    import { goto } from "@roxi/routify";
    import Avatar from "../../lib/components/design_system/Avatar.svelte";
    import Button from "../../lib/components/design_system/Button.svelte";
    import { onMount } from "svelte";
    import { api } from "../../lib/services/api/api_service.ts";
    import { getUser } from "../../lib/services/api/users_service.ts";

    onMount(async () => {
        console.log("JWT: " + (await getUser()));
        // await api.get(`/users/${(await getUser())}`);
    });
</script>

<aside class="flex items-center justify-between mx-4 mt-3">
    <span on:click={() => $goto("/")}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-10 h-10">
             <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5L8.25 12l7.5-7.5" />
        </svg>
    </span>
    <Helper color="white" bold="true" size="lg">optimuse.prime</Helper>
    <span class="text-xl">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 12a.75.75 0 11-1.5 0 .75.75 0 011.5 0zM12.75 12a.75.75 0 11-1.5 0 .75.75 0 011.5 0zM18.75 12a.75.75 0 11-1.5 0 .75.75 0 011.5 0z" />
        </svg>
    </span>
</aside>
<hr class="mt-3 border-t-[2px] border-primary-light">
<article class="mx-5 mt-4">
    <div class="flex gap-7">
        <Avatar size="sm" border={true} src="https://bsnl.ch/wp-content/uploads/2019/03/avatar-default-circle.png"/>
        <div class="flex flex-col justify-center">
            <Helper color="primary-lightest" size="xl" weight="bold">
                optimuse.prime
            </Helper>
            <div class="flex flex-row gap-3 mt-1.5">
                <Button size="sm" color="transparent">Message</Button>
                <Button size="sm" color="indigo">Follow</Button>
            </div>
        </div>
    </div>
    <Helper className="mt-2.5" color="white" weight="bold" size="xl">OptimusePrime</Helper>
    <div>
        <Helper color="primary-lightest">Software Engineer</Helper>
        <Helper color="primary-lightest" className="inline">CEO of </Helper><Helper className="inline" color="indigo">@pekkl.technologies</Helper>
        <Helper color="indigo">www.optimuseprime.com</Helper>
    </div>
    <div class="flex mt-1 gap-6 items-center">
        <div>
            <Helper size="sm" color="white" weight="bold" className="inline">69</Helper>
            <Helper size="sm" color="primary-lighter" className="inline">Following</Helper>
        </div>
        <div>
            <Helper size="sm" color="white" weight="bold" className="inline">420</Helper>
            <Helper size="sm" color="primary-lighter" className="inline">Followers</Helper>
        </div>
    </div>
    <div class="flex gap-2 items-center">
        <div class="flex mx-[-0.5rem]">
            <Avatar className="mr-[-1.75rem]" size="xsm" src="https://bsnl.ch/wp-content/uploads/2019/03/avatar-default-circle.png"/>
            <Avatar size="xsm" src="https://bsnl.ch/wp-content/uploads/2019/03/avatar-default-circle.png"/>
        </div>
        <span>
            <Helper size="sm" className="inline" color="primary-lighter">Followed by</Helper>
            <Helper size="sm" className="inline" color="primary-lightest">me</Helper>
            <Helper size="sm" className="inline" color="primary-lighter">and</Helper>
            <Helper size="sm" className="inline" color="primary-lightest">you</Helper>
        </span>
    </div>
</article>
<hr class="mt-3 border-t-[2px] border-primary-light">
<div class="mt-2 flex items-center justify-around">
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
        <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 15.75l5.159-5.159a2.25 2.25 0 013.182 0l5.159 5.159m-1.5-1.5l1.409-1.409a2.25 2.25 0 013.182 0l2.909 2.909m-18 3.75h16.5a1.5 1.5 0 001.5-1.5V6a1.5 1.5 0 00-1.5-1.5H3.75A1.5 1.5 0 002.25 6v12a1.5 1.5 0 001.5 1.5zm10.5-11.25h.008v.008h-.008V8.25zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z" />
    </svg>
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
        <path stroke-linecap="round" d="M15.75 10.5l4.72-4.72a.75.75 0 011.28.53v11.38a.75.75 0 01-1.28.53l-4.72-4.72M4.5 18.75h9a2.25 2.25 0 002.25-2.25v-9a2.25 2.25 0 00-2.25-2.25h-9A2.25 2.25 0 002.25 7.5v9a2.25 2.25 0 002.25 2.25z" />
    </svg>
</div>
<hr class="mt-2 border-t-[2px] border-primary-light">
<main class="grid grid-cols-3 gap-3 m-3 mt-0">
    <div class="w-32 h-32 bg-primary-lighter"/>
    <div class="w-32 h-32 bg-primary-lighter"/>
    <div class="w-32 h-32 bg-primary-lighter"/>
    <div class="w-32 h-32 bg-primary-lighter"/>
    <div class="w-32 h-32 bg-primary-lighter"/>
    <div class="w-32 h-32 bg-primary-lighter"/>
    <div class="w-32 h-32 bg-primary-lighter"/>
    <div class="w-32 h-32 bg-primary-lighter"/>
    <div class="w-32 h-32 bg-primary-lighter"/>
</main>
fa

<style>
    .overlay {
        gap: -3rem;
    }
</style>