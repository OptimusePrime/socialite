<script lang="ts">
    export let
        color: "indigo" | "transparent" = "indigo",
        className: string,
        size: "sm" | "md" | "lg" = "sm";

    const sizes = new Map<string, string>([
        ["sm", "px-6 py-0.5"],
        ["md", "px-10 py-3"],
        ["lg", "px-32 py-2.5"],
    ]);

    const colors = new Map<string, string>([
        ["indigo", "bg-accent"],
        ["transparent", "border border-primary-light"],
    ]);
</script>

<button on:click class="{sizes.get(size)} {colors.get(color)} {className} rounded-lg font-bold font-sans">
    <slot/>
</button>
