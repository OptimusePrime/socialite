<script lang="ts">
    import Helper from "./design_system/Helper.svelte";
    import LineSeparator from "./design_system/LineSeparator.svelte";

    export let title: string;
</script>

<header>
    <aside class="flex items-center justify-between mx-4 pt-3">
        <span>
            <slot name="first"/>
        </span>
        {#if title}
            <Helper color="white" bold="true" size="lg">{title}</Helper>
        {/if}
        <span class="text-xl font-sans">
            <slot name="second"/>
        </span>
    </aside>
    <LineSeparator className="mt-3"/>
</header>