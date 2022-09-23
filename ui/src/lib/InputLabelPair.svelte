<script lang="ts">
	import { Helper, Iconinput, Label } from "flowbite-svelte";

	export let
		labelClass: string,
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
	<Label for="input" class="block mb-2 ml-1 {labelClass}">{label}</Label>
	<Iconinput noBorder id="input" class={inputClass} {placeholder} {type} {size} bind:value={value}>
		<slot/>
	</Iconinput>
	<Helper bind:color={color}><span class="font-medium">{helperText || ""}</span></Helper>
</div>
