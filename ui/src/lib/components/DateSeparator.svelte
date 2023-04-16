<script lang="ts">
    export let date: Date;

    const monthNames = ["January", "February", "March", "April", "May", "June",
        "July", "August", "September", "October", "November", "December"
    ];
</script>

{#if date}
    <div class="flex flex-row items-center gap-3">
        <hr class="border-t-[2px] w-[32vw] border-primary-lightbit">
        <span class="font-bold font-sans text-sm">{`${monthNames[date.getMonth()]} ${date.getUTCDate()}, ${date.getFullYear()}`}</span>
        <hr class="border-t-[2px] w-[32vw] border-primary-lightbit">
    </div>
{/if}


