<template>
    <div class="root">
        <div class="form-group d-print-none">
            <label for="size">Size</label>
            <select class="form-control" id="size" v-model="addressSize">
                <option value="addressFontSmall">Small</option>
                <option value="addressFontMedium">Medium</option>
                <option value="addressFontLarge">Large</option>
            </select>
        </div>

        <div class="form-group d-print-none">
            <label for="address">Address</label>
            <textarea class="form-control" id="address" v-model="address"></textarea>
        </div>

        <div class="form-group d-print-none">
            <label for="returnAddress">Return Address</label>
            <textarea class="form-control" id="returnAddress" v-model="returnAddress"></textarea>
            <small class="form-text text-muted">
                Automatically saves to local storage, save you from typing it in again.
            </small>
        </div>

        <p class="d-print-none">When you're done just press Ctrl+P to Print. It's will only show the address
            and the dashed line below.</p>

        <div :class="addressSize">{{address}}</div>

        <div class="dottedLine mt-5"></div>

        <div class="returnAddress mt-5">{{returnAddress}}</div>
    </div>
</template>

<script type="module">
    export default {
        data() {
            return {
                address: "",
                addressSize: "addressFontSmall",
                returnAddress: "Return to:",
            }
        },
        mounted() {
            if (localStorage.returnAddress) {
                this.returnAddress = localStorage.returnAddress;
            }
        },
        watch: {
            returnAddress(newReturnAddress) {
                localStorage.returnAddress = newReturnAddress;
            }
        }
    };
</script>

<style type="text/css">
    .addressFontSmall {
        font-size: 2em;
        white-space: pre;
    }

    .addressFontMedium {
        font-size: 4em;
        white-space: pre;
    }

    .addressFontLarge {
        font-size: 6em;
        white-space: pre;
    }

    .returnAddress {
        font-size: 1em;
        white-space: pre;
    }

    .dottedLine {
        width: 100%;
        border-bottom-style: dashed;
    }
</style>