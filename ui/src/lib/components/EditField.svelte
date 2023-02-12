<script lang="ts">
    export let label: string;
    export let value: string;
</script>

<div class="flex gap-20">
    <span class="font-sans text-lg font-bold">{label}</span>
    <input class="border-none outline-none bg-transparent text-primary-lighter" value={value ? value : ""} placeholder={value ? "" : "N/A"}>
</div>