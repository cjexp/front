import Vue from 'vue';
import App from './alert.vue'

let application = Vue.extend(App);
let vm = new application({
    el: "#vueAlert"
});