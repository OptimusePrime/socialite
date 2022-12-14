module.exports = {
  env: {
    browser: true,
    es2021: true
  },
  extends: ["eslint:recommended", "plugin:@typescript-eslint/recommended", "plugin:storybook/recommended"],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    ecmaVersion: "latest",
    sourceType: "module"
  },
  plugins: ["@typescript-eslint", "svelte3"],
  overrides: [{
    files: ["*.svelte"],
    processor: "svelte3/svelte3"
  }],
  settings: {
    "svelte3/typescript": () => require("typescript"),
    "svelte3/typescript": true
  },
  rules: {
    indent: ["error", 4],
    "linebreak-style": ["error", "unix"],
    quotes: ["error", "double"],
    semi: ["error", "always"],
    "space-before-blocks": ["error", "always"]
  }
};