// Snowpack Configuration File
// See all supported options: https://www.snowpack.dev/reference/configuration
const httpProxy = require("http-proxy");
const proxy = httpProxy.createServer({ target: "http://localhost:8080" });

/** @type {import("snowpack").SnowpackUserConfig } */
module.exports = {
  mount: {
    public: { url: "/", static: true },
    src: { url: "/dist" },
  },
  plugins: [
    "@snowpack/plugin-dotenv",
    "@snowpack/plugin-typescript",
    "@prefresh/snowpack",
  ],
  packageOptions: {
    /* ... */
  },
  devOptions: {
    port: 3000,
  },
  buildOptions: {
    /* ... */
  },
  routes: [
    { match: "all", src: "/oto/.*", dest: (req, res) => proxy.web(req, res) },
    { match: "routes", src: ".*", dest: "/index.html" },
  ],
};
