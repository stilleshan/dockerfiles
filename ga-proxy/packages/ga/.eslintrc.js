module.exports = {
  extends: ["plugin:prettier/recommended"],
  rules: {
    "prettier/prettier": "error",
  },
  env: {
    browser: true,
    es6: true
  },
  parserOptions: {
    sourceType: "module"
  }
};
