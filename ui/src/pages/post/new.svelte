<script lang="ts">
    import { Camera, CameraResultType, Photo } from "@capacitor/camera";
    import Button from "../../lib/components/design_system/Button.svelte";
    import Header from "../../lib/components/Header.svelte";
    import { goto } from "@roxi/routify";
    import PostImage from "../../lib/components/PostImage.svelte";
    import LineSeparator from "../../lib/components/design_system/LineSeparator.svelte";
    import { createPost } from "../../lib/services/api/users_service";

    let imagePath = "";
    let caption = "";
    let location = "";
    let blob: Blob;

    async function takePicture() {
        Camera.getPhoto({
            quality: 90,
            // allowEditing: true,
            resultType: CameraResultType.Uri
        }).then(async img => {
            imagePath = img.webPath;
            console.log("Image: " + imagePath);
            blob = await fetch(imagePath).then(r => r.blob());
        });
    }

    async function share() {
        createPost(blob, caption, location).then(id => {
            $goto(`/p/${id}`);
        });
    }
</script>

<Header title="New Post">
    <span slot="first" class="hover:cursor-pointer" on:click={() => $goto("/home")}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-12 h-12">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
        </svg>
    </span>
    <span on:click={share} slot="second" class=" mr-2 font-sans text-indigo text-xl hover:cursor-pointer">
        Share
    </span>
</Header>

<main>
    <PostImage autoSize={true} imgSrc={imagePath} className="" background={false} on:click={takePicture}/>
    <div>
        <input bind:value={caption} type="text" placeholder="Write a caption..." class="font-sans w-full outline-none border-none placeholder-primary-lighter placeholder-opacity-50 bg-transparent pt-5 px-5 pb-2 text-xl">
        <LineSeparator/>
    </div>
    <div>
        <span class="p-[0.9rem] ml-[20rem] absolute">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
            </svg>
        </span>
        <input bind:value={location} type="text" placeholder="Location" class="font-sans w-full outline-none border-none placeholder-primary-lighter placeholder-opacity-50 bg-transparent pt-3 px-5 text-xl">
        <LineSeparator/>
    </div>
</main>

<style>
    span {
        cursor: pointer;
    }
</style>