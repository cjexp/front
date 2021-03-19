const vuePlugin = require('esbuild-vue');

require('esbuild').build({
    entryPoints: ['dev/javascript/main.mjs'],
    bundle: true,
    plugins: [vuePlugin()],
    outfile: 'live/javascript/main.js',
})