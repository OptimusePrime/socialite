<script lang="ts">
    import InputLabelPair from "../../lib/components/InputLabelPair.svelte";
    import { faEnvelope, faKey } from "@fortawesome/free-solid-svg-icons";
    import Fa from "svelte-fa/src/fa.svelte";
    import { generatePassword, validateEmail, validatePassword } from "../../lib/services/validation_service.js";
    import { loginUser, LoginUserErrors } from "../../lib/services/api/users_service";
    import { Button } from "flowbite-svelte";

    let email = "", emailHelperText: string, emailColor: string, password = "", passwordHelperText: string, passwordColor: string, loginError: LoginUserErrors;

    $: {
        if (validatePassword(password) || !password) {
            passwordColor = "";
            passwordHelperText = "";
        } else {
            passwordColor = "red";
            passwordHelperText = "Password length must be from 8 to 16 and include at least one: uppercase and lowercase letter, special character, and number.";
        }

        if (validateEmail(email) || !email) {
            emailColor = "";
            emailHelperText = "";
        } else {
            emailColor = "red";
            emailHelperText = "Email must be valid and its length must not exceed 48.";
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
            emailHelperText = "An account with this email doesn't exist.";
            emailColor = "red";
            break;
        case LoginUserErrors.INVALID_PASSWORD:
            passwordHelperText = "Wrong password.";
            passwordColor = "red";
            break;
        }
    }
</script>

<a href="/">Back</a>

<InputLabelPair inputClass="pr-20" label="Email" placeholder="jane.doe@example.com" size="md" type="email" icon={faEnvelope} bind:value={email} bind:color={emailColor}
                bind:helperText={emailHelperText}>
    <Fa icon={faEnvelope}/>
</InputLabelPair>
<InputLabelPair inputClass="pr-20" label="Password" placeholder={generatePassword(16)} size="md" type="password" icon={faKey} bind:value={password} bind:color={passwordColor}
                bind:helperText={passwordHelperText}>
    <Fa icon={faKey}/>
</InputLabelPair>
<Button on:click={handleLogin} color="purple">Login</Button>

