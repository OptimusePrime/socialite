<script lang="ts">
    import { onMount } from "svelte";

    export let className: string;
    export let imgSrc: string;
    export let autoSize: boolean;


    onMount(() => {
        $: if (!imgSrc) {
            (document.body.querySelector(".post-image-div") as HTMLElement).style.height = `${screen.width}px`;
        } else {
            (document.body.querySelector(".post-image-div") as HTMLElement).style.maxHeight = `${screen.width}px`;
        }
    });
</script>

<div on:click class="post-image-div {className} {(!imgSrc && autoSize) ? `bg-primary-lighter w-full` : ``} overflow-hidden">
    <img src={imgSrc}/>
</div>

