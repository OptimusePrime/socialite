import { CapacitorConfig } from '@capacitor/cli';

const config: CapacitorConfig = {
  appId: 'tech.socialit.app',
  appName: 'socialite',
  webDir: 'www',
  bundledWebRuntime: false,
  server: {
    url: "http://172.20.48.1:8100",
    cleartext: true
  }
};

export default config;
