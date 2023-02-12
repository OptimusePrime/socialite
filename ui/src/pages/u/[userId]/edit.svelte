<script lang="ts">
    import Header from "../../../lib/components/Header.svelte";
    import Avatar from "../../../lib/components/design_system/Avatar.svelte";
    import { User } from "../../../lib/types/user_types";
    import { getSignedInUser, updateUser } from "../../../lib/services/api/users_service";

    type Field = {
        label: string;
        value: string;
    }

    let fields: Field[] = [];
    let user: User;

    let bio = "";
    let username = "";
    let name = "";
    let gender = "";
    let pronouns = "";

    (async function() {
        user = await getSignedInUser();
        bio = user.biography;
        username = user.username;
        name = user.name;
        gender = user.gender;
    })();

    async function onSave() {
        await updateUser({
            biography: bio,
            gender,
            name,
            username,
            pronouns,
        });
    }
</script>

<Header title="Edit Profile">
    <a slot="first" class="font-sans cursor-pointer" href="/home">Cancel</a>
    <button slot="second" class="text-indigo font-bold text-lg" on:click={async () => await onSave()}>Save</button>
</Header>

<div class="flex flex-col items-center gap-7 mt-10">
    <Avatar size="md" src={user?.avatar}/>
</div>
<div class="mt-8 flex flex-col ml-7 gap-6">
        <div class="flex flex-row">
            <div class="flex flex-col">
                <span class="absolute font-sans text-lg font-bold">Name</span>
            </div>
            <div class="flex flex-col ml-40">
                <input bind:value={name} class="border-none outline-none bg-transparent text-primary-lighter placeholder-primary-lighter"
                       placeholder={user?.name ? "" : "N/A"}>
            </div>
        </div>

    <div class="flex flex-row">
        <div class="flex flex-col">
            <span class="absolute font-sans text-lg font-bold">Username</span>
        </div>
        <div class="flex flex-col ml-40">
            <input bind:value={username} class="border-none outline-none bg-transparent text-primary-lighter placeholder-primary-lighter"
                   placeholder={user?.username ? "" : "N/A"}>
        </div>
    </div>

    <div class="flex flex-row">
        <div class="flex flex-col">
            <span class="absolute font-sans text-lg font-bold">Biography</span>
        </div>
        <div class="flex flex-col ml-40">
            <input bind:value={bio} class="border-none outline-none bg-transparent text-primary-lighter placeholder-primary-lighter"
                   placeholder={user?.biography ? "" : "N/A"}>
        </div>
    </div>

    <div class="flex flex-row">
        <div class="flex flex-col">
            <span class="absolute font-sans text-lg font-bold">Gender</span>
        </div>
        <div class="flex flex-col ml-40">
            <input bind:value={gender} class="border-none outline-none bg-transparent text-primary-lighter placeholder-primary-lighter"
                   placeholder={user?.gender ? "" : "N/A"}>
        </div>
    </div>

    <div class="flex flex-row">
        <div class="flex flex-col">
            <span class="absolute font-sans text-lg font-bold">Pronouns</span>
        </div>
        <div class="flex flex-col ml-40">
            <input bind:value={pronouns} class="border-none outline-none bg-transparent text-primary-lighter placeholder-primary-lighter"
                   placeholder={user?.pronouns ? "" : "N/A"}>
        </div>
    </div>
</div>

<style>
    a {
        cursor: pointer;
    }
</style>