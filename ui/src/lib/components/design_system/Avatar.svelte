<script lang="ts">
    export let
        src: string,
        size: "sm" | "md" | "lg" = "sm",
        border = false,
        className: string;

    const sizes = new Map<string, string>([
        ["xsm", "w-9"],
        ["sm", "w-24"],
        ["md", "w-36"],
        ["lg", "w-44"],
    ]);
</script>

<div on:click class="{sizes.get(size)} {border ? 'border-accent' : 'border-transparent'} p-[0.175rem] rounded-full border-4 {className}">
    <img {src} alt="User avatar">
</div>