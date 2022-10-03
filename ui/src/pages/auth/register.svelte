<script lang="ts">
    import InputLabelPair from "../../lib/components/InputLabelPair.svelte";
    import { InputLabelPairProps } from "../../lib/types/form_types";
    import Fa from "svelte-fa/src/fa.svelte";
    import { faEnvelope, faKey, faUser, faUserTie } from "@fortawesome/free-solid-svg-icons";
    import { Button, Helper } from "flowbite-svelte";
    import {
        generatePassword,
        validateEmail,
        validateName,
        validatePassword,
        validateUsername
    } from "../../lib/services/validation_service";
    import {
        getAccessToken,
        loginUser,
        LoginUserErrors,
        registerUser,
        RegisterUserErrors
    } from "../../lib/services/api/users_service";
    import { onMount } from "svelte";

    onMount(() => console.log(getAccessToken()));

    let username = "", usernameColor: string, usernameHelperText: string,
        name = "", nameColor: string, nameHelperText: string,
        password = "", passwordColor: string, passwordHelperText: string,
        email = "", emailColor: string, emailHelperText: string,
        registerError: RegisterUserErrors, errorHelperText = "";

    function validateAll(username: string, name: string, password: string, email: string): boolean {
        return (validateUsername(username) && validateName(name) && validatePassword(password) && validateEmail(email)) || (!username && !name && !password && !email);
    }

    $: {
        if (validateName(name) || !name || registerError === RegisterUserErrors.INVALID_NAME) {
            nameColor = "";
            nameHelperText = "";
        } else {
            nameColor = "red";
            nameHelperText = "Name length must be between 3 and 16.";
        }

        if (validateUsername(username) || !username || registerError === RegisterUserErrors.INVALID_USERNAME) {
            usernameColor = "";
            usernameHelperText = "";
        } else {
            usernameColor = "red";
            usernameHelperText = "Username length must be between 3 and 16.";
        }

        if (validatePassword(password) || !password || registerError === RegisterUserErrors.INVALID_PASSWORD) {
            passwordColor = "";
            passwordHelperText = "";
        } else {
            passwordColor = "red";
            passwordHelperText = "Password length must be from 8 to 16 and include at least one: uppercase and lowercase letter, special character, and number.";
        }

        if (validateEmail(email) || !email || registerError === RegisterUserErrors.INVALID_EMAIL) {
            emailColor = "";
            emailHelperText = "";
        } else {
            emailColor = "red";
            emailHelperText = "Email must be valid and its length must not exceed 48.";
        }
    }

    async function handleRegister() {
        if (!validateAll(username, name, password, email)) {
            return;
        }

        /*        registerError = await registerUser({
            email,
            password,
            username,
            name,
        });
        console.log(typeof registerError);
        console.log(registerError === RegisterUserErrors.EMAIL_NOT_UNIQUE);

        switch (registerError) {
        case RegisterUserErrors.EMAIL_NOT_UNIQUE:
            emailHelperText = "An account with that email already exists.";
            emailColor = "red";
            break;
        case RegisterUserErrors.USERNAME_NOT_UNIQUE:
            usernameHelperText = "An account with that username already exists.";
            usernameColor = "red";
            break;
        case RegisterUserErrors.INTERNAL_SERVER_ERROR:
            errorHelperText = "Something went wrong. Please try again later.";
            break;
        }*/

        const loginError = await loginUser({
            email,
            password,
        });

        if (loginError) {
            errorHelperText = "Something went wrong. Please try to login manually.";
        }
        
    }


    const inputs: InputLabelPairProps[] = [
        {
            label: "Username",
            placeholder: "janed",
            size: "md",
            type: "text",
            icon: faUser,
            // color: nameColor
        },
        {
            label: "Name",
            placeholder: "Jane Doe",
            size: "md",
            type: "text",
            icon: faUserTie,
            // color: nameColor,
        },
        {
            label: "Email",
            placeholder: "jane.doe@example.com",
            size: "md",
            type: "email",
            icon: faEnvelope,
            // color: nameColor
        },
        {
            label: "Password",
            placeholder: generatePassword(16),
            size: "md",
            type: "password",
            icon: faKey,
            // color: nameColor
        },
    ];
</script>

<a href="/">Back</a>
<main>
    <form>
        <InputLabelPair inputClass="pr-20" {...inputs[0]} bind:value={username} bind:color={usernameColor}
                        bind:helperText={usernameHelperText}>
            <Fa icon={inputs[0].icon}/>
        </InputLabelPair>
        <InputLabelPair inputClass="pr-20" {...inputs[1]} bind:value={name} bind:color={nameColor}
                        bind:helperText={nameHelperText}>
            <Fa icon={inputs[1].icon}/>
        </InputLabelPair>
        <InputLabelPair inputClass="pr-20" {...inputs[2]} bind:value={email} bind:color={emailColor}
                        bind:helperText={emailHelperText}>
            <Fa icon={inputs[2].icon}/>
        </InputLabelPair>
        <InputLabelPair inputClass="pr-20" {...inputs[3]} bind:value={password} bind:color={passwordColor}
                        bind:helperText={passwordHelperText}>
            <Fa icon={inputs[3].icon}/>
        </InputLabelPair>
        {#if registerError}
            <Helper class="text-sm mb-3" color="red">{errorHelperText}</Helper>
        {/if}
        <Button on:click={handleRegister} color="purple">Register</Button>
    </form>
</main>




