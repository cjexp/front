import Vue from 'vue';
import App from './randomHash.vue'

let application = Vue.extend(App);
let vm = new application({
    el: "#vueRandomHash"
});