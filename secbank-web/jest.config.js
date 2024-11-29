module.exports = {
    watchPlugins: [],
    collectCoverage: true,
    collectCoverageFrom: [
      "src/**/*.{js,jsx,ts,tsx}", // Ajuste conforme necessário
      "!src/index.js",           // Exclua arquivos não relevantes
      "!src/**/*.test.js"        // Exclua arquivos de teste
    ],
    coverageDirectory: "coverage",
  };
  