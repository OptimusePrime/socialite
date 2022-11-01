<script lang="ts">
    export let color: string, className: string, size: string, bold: boolean;

    const colors = new Map<string, string>([
        ["red", "text-red"]
    ]);
</script>

<p class="{colors.get(color)} {className} font-sans text-{size} {bold ? 'font-bold' : ''}">
    <slot />
</p>