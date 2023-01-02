<script lang="ts">
    export let color: string, className: string, size: string, bold: boolean, weight: string;

    const colors = new Map<string, string>([
        ["red", "text-red"],
        ["white", "text-white"],
        ["primary-lighter", "text-primary-lighter"],
        ["primary-light", "text-primary-light"],
        ["primary-lightest", "text-primary-lightest"],
        ["indigo", "text-accent"],
    ]);

    const sizes = new Map<string, string>([
        ["sm", "text-sm"],
        ["md", "text-md"],
        ["lg", "text-lg"],
        ["xl", "text-xl"],
        ["2xl", "text-2xl"],
    ]);

    const weights = new Map<string, string>([
        ["light", "font-light"],
        ["normal", "font-normal"],
        ["semibold", "font-semibold"],
        ["bold", "font-bold"],
    ]);
</script>

<p class="{colors.get(color)} {className} font-sans {weights.get(weight)} {sizes.get(size)} {bold ? 'font-bold' : ''}">
    <slot />
</p>