module.exports = {
  root: true,
  env: {
    node: true
  },
  //extends: ["plugin:vue/essential", "@vue/prettier"],
  extends: ["plugin:vue/essential"],
  parserOptions: {
    parser: "babel-eslint"
  },
  rules: {
    "no-console": process.env.NODE_ENV === "production" ? "error" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "error" : "off"
  }
};
