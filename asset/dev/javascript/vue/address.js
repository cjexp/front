import Vue from 'vue';
import App from './address.vue'

let application = Vue.extend(App);
let vm = new application({
    el: "#vueAddress"
});

document.getElementById("siteHeader").classList.add("d-print-none");