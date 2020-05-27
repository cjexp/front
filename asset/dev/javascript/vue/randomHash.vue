<template>
    <div class="root">
        <div v-for="result in results">
            {{ result }}<br>
        </div>

        <span class="mt-3 btn btn-primary" v-on:click="generate"><i class="fas fa-cog"></i> Generate 5 more</span>
    </div>
</template>

<script type="module">
    import {simpleFetch} from "../functions/misc";

    function generate(object) {
        object.results = object.results.concat(JSON.parse(simpleFetch(location.pathname, 'PUT')));
    }

    export default {
        data() {
            return {
                results: []
            }
        },
        mounted() {
            generate(this);
        },
        methods: {
            generate() {
                generate(this);
            }
        }
    }
</script>