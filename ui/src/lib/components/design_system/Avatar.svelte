<script lang="ts">
    import { createAvatar } from "@dicebear/core";
    import { adventurerNeutral } from "@dicebear/collection";

    export let
        src: string,
        size: "sm" | "md" | "lg" | "vsm" = "sm",
        border = false,
        className: string;

    const sizes = new Map<string, string>([
        ["xsm", "w-11"],
        ["vsm", "w-20"],
        ["sm", "w-24"],
        ["md", "w-36"],
        ["lg", "w-44"],
    ]);

    const avatar = createAvatar(adventurerNeutral);
</script>

<div on:click class="{sizes.get(size)} {border ? 'border-accent' : 'border-transparent'} p-[0.175rem] {className}">
    <img src={src ? src : avatar.toDataUriSync()} alt="User avatar" class="rounded-full">
</div>