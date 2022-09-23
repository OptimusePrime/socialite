import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [svelte()],
	build: {
		outDir: "www"
	},
	server: {
		origin: "http://192.168.0.1",
		port: 8100
	}
});
