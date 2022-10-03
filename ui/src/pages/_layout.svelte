<script lang="ts">
    import { onMount } from "svelte";

    onMount(() => {
        window.document.documentElement.classList.add("dark");
        localStorage.setItem("color-theme", "dark");
    });
</script>

<slot/>