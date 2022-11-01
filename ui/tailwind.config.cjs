const config = {
	content: [
		"./src/**/*.{html,js,svelte,ts}",
		// "./node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}",
	],

	theme: {
		fontFamily: {
			sans: ["Lato", "sans-serif"],
		},
		extend: {
			colors: {
				"accent": "#6875f5",
				"primary-light": "#3a3a3a",
				"primary-lighter": "#878787",
				"red": "#ef4444",
			},
		},
	},

	darkMode: "class",
};

module.exports = config;
