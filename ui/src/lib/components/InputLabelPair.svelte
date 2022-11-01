<!--
<script lang="ts">
    import { Helper, Iconinput, Label } from "flowbite-svelte";

    export let labelClass: string,
        inputClass: string,
        placeholder: string,
        label: string,
        type: string,
        size: "sm" | "md" | "lg" = "md",
        value: string,
        helperText: string,
        color: "base" | "green" | "red" = "base";
</script>

<div class="mb-6 flex flex-col content-start text-left">
    <Label for="input" class="block mb-2 ml-1 {labelClass}" bind:color={color}>{label}</Label>
    <Iconinput noBorder id="input" class={inputClass} {placeholder} {type} bind:color={color} {size} bind:value={value}>
        <slot/>
    </Iconinput>
    {#if helperText}
        <Helper bind:color={color} class="mt-1"><span class="font-medium">{helperText || ""}</span></Helper>
    {/if}
</div>
-->
