<script lang="ts">
	import InputLabelPair from "../../lib/InputLabelPair.svelte";
	import { InputLabelPairProps } from "../../types/form_types";
	import Fa from "svelte-fa/src/fa.svelte";
	import { faUser, faUserTie, faKey,  faEnvelope} from "@fortawesome/free-solid-svg-icons";
	import { onMount } from "svelte";

	let username,
		name = "", nameColor = "", nameHelperText = "",
		password,
		email;

	function generatePassword(length: number): string {

		const charset = "!\"#&'()*,-./:;?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]abcdefghijklmnopqrstuvwxyz0123456789";
		let password = "";
		for (let i = 0; i < length; i++) {
			password += charset[Math.trunc(Math.random() * charset.length)];
		}
		return password;
	}
	
	$: if (name.length < 3 || name.length > 16) {
		nameColor = "red";
		nameHelperText = "Wrong!";
		console.log("Error");
	} else {
		nameColor = "base";
		nameHelperText = "";
	}

	onMount(() => {
		nameColor = "base";
		nameHelperText = "";
	});

	function handleRegister() {
		fetch("localhost:");
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
<form>
	{#each inputs as input}
		<InputLabelPair bind:value={name} bind:color={nameColor} bind:helperText={nameHelperText} inputClass="pr-20" {...input}>
			<Fa icon={input.icon}/>
		</InputLabelPair>
	{/each}
<!--	<Button on:click={handleButtonClick} color="purple">Register</Button>-->
</form>



