import build from 'esbuild';
import vuePlugin from 'esbuild-vue';

build({
    entryPoints: ['dev/javascript/main.js'],
    bundle: true,
    plugins: [vuePlugin()],
    outfile: 'live/javascript/main.js',
})