<script lang="ts">
    import InputLabelPair from "../../lib/components/InputLabelPair.svelte";
    import { faEnvelope, faKey } from "@fortawesome/free-solid-svg-icons";
    import Fa from "svelte-fa/src/fa.svelte";
    import { generatePassword, validateEmail, validatePassword } from "../../lib/services/validation_service.ts";
    import { loginUser, LoginUserErrors } from "../../lib/services/api/users_service";
    import Button from "../../lib/components/design_system/Button.svelte";
    import Input from "../../lib/components/design_system/Input.svelte";

    let email = "", emailError = "AHGAHGAHA", emailColor: string, password = "", passwordColor: string, showPassword = false, loginError: LoginUserErrors;
    let passwordError = "";
    // let passwordHelperText = "gaga";

    $: {
        console.log(email);
        if (validatePassword(password) || !password) {
            passwordError = "";
        } else {
            passwordError = "Password length must be from 8 to 16 and <br>include at least one: uppercase and<br> lowercase letter, special character, and number.";
        }

        if (validateEmail(email) || !email) {
            emailError = "";
        } else {
            emailError = "Email must be valid and its length must not exceed 48.";
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

<a href="/">Back</a>

<main>
    <Input className="mt-4" placeholder="Email" size="lg" type="email" bind:value={email} bind:helperText={emailError} />
    <Input className="mt-4" placeholder="Password" size="lg" type={showPassword ? "text" : "password"} bind:value={password} bind:helperText={passwordError}>
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
    <Button on:click={handleLogin} color="indigo" size="lg" className="mt-10">Login</Button>
</main>



