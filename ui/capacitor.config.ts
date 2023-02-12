import { CapacitorConfig } from "@capacitor/cli";

const config: CapacitorConfig = {
    appId: "tech.socialit.app",
    appName: "socialite",
    webDir: "www",
    bundledWebRuntime: false,
    plugins: {
        CapacitorHttp: {
            enabled: true,
        }
    }
/*    server: {
        url: "http://192.168.1.102:8100/",
        cleartext: true
    },*/
};

export default config;
