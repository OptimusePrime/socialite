<script lang="ts">
    import Button from "./design_system/Button.svelte";
    import { followUser, unfollowUser } from "../services/api/users_service";
    import { findFollow, getSignedInUserId } from "../services/api/users_service.js";
    import { createEventDispatcher, onMount } from "svelte";

    export let userProfileId: string;
    let isFollowing: boolean;

    let signedInUserId: string;
    (async function () {
        signedInUserId = await getSignedInUserId();
    })();

    const dispatch = createEventDispatcher();

    $: if (userProfileId) {
        (async function () {
            isFollowing = await findFollow(signedInUserId, userProfileId);
        })();
    }

    async function onFollow() {
        if (signedInUserId === userProfileId) {
            return;
        }
        await followUser(userProfileId);
        isFollowing = !isFollowing;
        dispatch("follow");
    }

    async function onUnfollow() {
        await unfollowUser(userProfileId);
        isFollowing = !isFollowing;
        dispatch("unfollow");
    }

</script>

{#if signedInUserId !== userProfileId}
    {#if (isFollowing)}
        <Button on:click={onUnfollow} size="sm" color="indigo">
            Unfollow
        </Button>
    {:else}
        <Button on:click={onFollow} size="sm" color="indigo">
            Follow
        </Button>
    {/if}
{/if}


