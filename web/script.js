Vue.component("v-select", VueSelect.VueSelect);
Vue.component("v-range", VueRange);

window.app = new Vue({
    el: "main",
    data: {
        options: {

        },
        pages: ["a"],
        scanners: ["test"],
    }
})
