import { build } from 'esbuild';
import vue from 'esbuild-plugin-vue';

build({
    entryPoints: ['dev/javascript/main.js'],
    bundle: true,
    plugins: [vue()],
    outfile: 'live/javascript/main.js',
})