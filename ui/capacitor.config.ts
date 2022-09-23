import { CapacitorConfig } from "@capacitor/cli";

const config: CapacitorConfig = {
	appId: "tech.socialit.app",
	appName: "socialite",
	webDir: "www",
	bundledWebRuntime: false,
	server: {
		url: "http://192.168.1.108:8100",
		cleartext: true
	}
};

export default config;
