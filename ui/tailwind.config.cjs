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
				"primary-lightbit": "#353535",
				"primary-light": "#3a3a3a",
				"primary-lighter": "#878787",
				"primary-lightest": "#d1d5db",
				"primary-lighterest": "#929292",
				"primary": "#242424",
				"red": "#ef4444",
				"indigo": "#6875F5",
			},
		},
	},

	darkMode: "class",
};

module.exports = config;
