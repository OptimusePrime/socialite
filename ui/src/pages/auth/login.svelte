<script lang="ts">
    import { validateEmail, validatePassword } from "../../lib/services/validation_service.ts";
    import { loginUser, LoginUserErrors } from "../../lib/services/api/users_service";
    import Button from "../../lib/components/design_system/Button.svelte";
    import Input from "../../lib/components/design_system/Input.svelte";
    import { Body } from "svelte-body";
    import { goto } from "@roxi/routify";

    let email = "", emailError = "", emailColor: string, password = "", passwordColor: string, showPassword = false, loginError: LoginUserErrors;
    let passwordError = "";

    $: {
        console.log(email);
        if (validatePassword(password) || !password) {
            passwordError = "";
        } else {
            passwordError = "Password length must be from 8 to 32.";
        }

        if (validateEmail(email) || !email) {
            emailError = "";
        } else {
            emailError = "Email must be valid and its length can't <br> exceed 48.";
        }
    }

    async function handleLogin() {
        if (!validateEmail(email) || !validatePassword(password)) {
            return;
        }

        loginError = await loginUser({
            email,
            password,
        });

        switch (loginError) {
        case LoginUserErrors.INVALID_EMAIL:
            emailError = "An account with this email doesn't exist.";
            emailColor = "red";
            break;
        case LoginUserErrors.INVALID_PASSWORD:
            passwordError = "Wrong password.";
            passwordColor = "red";
            break;
        }
    }
</script>

<Body class="flex justify-center items-center"/>

<div class="wrap z-10">
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
            <h1 class="font-extrabold text-[2.6rem] mb-3">Log In</h1>
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
            <Button on:click={handleLogin} className="mt-6" color="indigo" size="lg">Log In</Button>
        </main>
        <div class="fixed left-0 bottom-0 w-full">
            <hr class="mt-3 border-t-[2px] border-primary-light">
            <footer class="bg-primary p-5 text-center">
                <p class="font-bold font-sans">Don't have an account? <a href="/auth/register" class="text-accent font-bold">Register</a></p>
            </footer>
        </div>
    </div>
</div>

<style lang="postcss">

    html, body {
        margin: 0;
        height: 100%;
        overflow: hidden;
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
    }
</style>





