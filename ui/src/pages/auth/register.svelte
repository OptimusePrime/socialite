<script lang="ts">
    import Input from "../../lib/components/design_system/Input.svelte";
    import {
        validateEmail,
        validateName,
        validatePassword,
        validateUsername
    } from "../../lib/services/validation_service";
    import {
        loginUser,
        registerUser,
        RegisterUserErrors
    } from "../../lib/services/api/users_service";
    import Button from "../../lib/components/design_system/Button.svelte";
    import Helper from "../../lib/components/design_system/Helper.svelte";
    import { Body } from "svelte-body";
    import { onMount } from "svelte";
    import { faker } from "@faker-js/faker";
    import { goto } from "@roxi/routify";

    let username = "", usernameError: string,
        name = "", nameError: string,
        password = "", passwordError: string,
        email = "", emailError: string,
        registerError: RegisterUserErrors, errorHelperText = "",
        showPassword: boolean;

    function validateAll(username: string, name: string, password: string, email: string): boolean {
        return (validateUsername(username) && validateName(name) && validatePassword(password) && validateEmail(email)) || (!username && !name && !password && !email);
    }

    $: {
        if (validateName(name) || !name || registerError === RegisterUserErrors.INVALID_NAME) {
            nameError = "";
        } else {
            nameError = "Name length must be between 3 and 16";
        }

        if (validateUsername(username) || !username || registerError === RegisterUserErrors.INVALID_USERNAME) {
            usernameError = "";
        } else {
            usernameError = "Username length must be between 3 and 16";
        }

        if (validatePassword(password) || !password || registerError === RegisterUserErrors.INVALID_PASSWORD) {
            passwordError = "";
        } else {
            passwordError = "Password length must be from 8 to 32";
        }

        if (validateEmail(email) || !email || registerError === RegisterUserErrors.INVALID_EMAIL) {
            emailError = "";
        } else {
            emailError = "Email must be valid and its length can't exceed 48";
        }
    }

    async function handleRegister() {
        if (!validateAll(username, name, password, email)) {
            return;
        }

        console.log(`Email: ${email}, Username: ${username}, Password: ${password}, Name: ${name}`);

        registerError = await registerUser({
            email,
            password,
            username,
            name,
        });

        switch (registerError) {
        case RegisterUserErrors.EMAIL_NOT_UNIQUE:
            emailError = "An account with that email already exists.";
            break;
        case RegisterUserErrors.USERNAME_NOT_UNIQUE:
            usernameError = "An account with that username already exists.";
            break;
        case RegisterUserErrors.INTERNAL_SERVER_ERROR:
            errorHelperText = "Something went wrong. Please try again later.";
            break;
        }

        const loginError = await loginUser({
            email,
            password,
        });

        if (loginError) {
            errorHelperText = "Something went wrong. Please try to login manually.";
            return;
        }

        $goto("/home");
    }

    onMount(() => {
        username = faker.name.firstName();
        name = faker.name.fullName().substring(0, 15);
        password = faker.internet.password();
        email = faker.internet.email();
    });
</script>

<Body class="flex justify-center items-center"/>

<div class="wrap">
    <img class="background-img" src="/auth_index_background.svg" alt="Bla bla">
    <aside class="flex items-start justify-between mx-1">
            <span class="absolute top-5 bottom-0 " on:click={() => $goto("/")}>
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-10 h-10">
                     <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5L8.25 12l7.5-7.5" />
                </svg>
            </span>
    </aside>
    <div class="main-content">
        <main class="flex items-center flex-col gap-1.5">
            <h1 class="font-extrabold text-[2.6rem] mb-3 page-title-shadow">Register</h1>
            <Input className="mt-1" placeholder="Email" size="lg" type="email" bind:value={email} bind:helperText={emailError} />
            <Input className="mt-1" placeholder="Password" size="lg" type={showPassword ? "text" : "password"} bind:value={password} bind:helperText={passwordError}>
                <span class="hover:cursor-pointer" on:click={() => showPassword = !showPassword}>
                    {#if showPassword}
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
                            <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                        </svg>
                    {:else}
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M3.98 8.223A10.477 10.477 0 001.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.45 10.45 0 0112 4.5c4.756 0 8.773 3.162 10.065 7.498a10.523 10.523 0 01-4.293 5.774M6.228 6.228L3 3m3.228 3.228l3.65 3.65m7.894 7.894L21 21m-3.228-3.228l-3.65-3.65m0 0a3 3 0 10-4.243-4.243m4.242 4.242L9.88 9.88" />
                        </svg>
                    {/if}
                </span>
            </Input>
            <Input className="mt-1" placeholder="Username" size="lg" type="text" bind:value={username}
                   bind:helperText={usernameError} />
            <Input className="mt-1" placeholder="Name" size="lg" type="text" bind:value={name}
                   bind:helperText={nameError} />
            {#if registerError}
                <Helper className="mt-1.5" size="base" color="red" bold="true">{errorHelperText}</Helper>
            {/if}
            <Button on:click={handleRegister} className="mt-6" color="indigo" size="lg">Register</Button>
        </main>
        <div class="fixed left-0 bottom-0 w-full">
            <hr class="mt-3 border-t-[2px] border-primary-light">
            <footer class="bg-primary p-5 text-center">
                <p class="font-bold font-sans">Have an account? <a href="/auth/login" class="text-accent font-bold">Log In</a></p>
            </footer>
        </div>
    </div>
</div>

<style lang="postcss">

    html, body {
        margin: 0;
        height: 100%;
        overflow: hidden !important;
    }

    .main-content {
        position: relative;
    }

    .wrap {
        //position: relative;
        height: 100%;
        width: 100%;
    }

    .background-img {
        //opacity: 0.4;
        filter: brightness(0.55);
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: auto;
        overflow: hidden;
    }
</style>

<!--
<style>
    :global(body) {
        display: flex;
        place-items: center;
    }
</style>
-->




