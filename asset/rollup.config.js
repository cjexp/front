import nodeResolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs';
import { terser } from 'rollup-plugin-terser';
import json from '@rollup/plugin-json';
import vue from 'rollup-plugin-vue';
import replace from '@rollup/plugin-replace';

export default {
    input: {
        'main': 'dev/javascript/main.js',

        'vue/alert': 'dev/javascript/vue/alert.js'
    },
    output: {
        dir: 'live/javascript',
        format: 'es',
        globals: [],
    },
    plugins: [
        vue(),
        json(),
        nodeResolve(),
        commonjs(),
        terser(),
        replace({
            'process.env.NODE_ENV': JSON.stringify( process.env.BUILD )
        })
    ],
    external: []
}