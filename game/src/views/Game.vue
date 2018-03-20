<template>
    <div>
        {{code}}
    </div>
</template>

<script>
export default {
    data() {
        return {
            code: "0000"
        };
    },
    methods: {
        init() {
            if (window.WebSocket && !window.conn) {
                window.conn = new WebSocket("ws://localhost:38080/init");
                conn.onopen = event => {
                    console.log("init success.");
                };
                conn.onclose = event => {
                    console.log("disconnect.");
                };
                conn.onmessage = event => {
                    console.log("message", event);
                };
            }
        }
    },
    created() {
        this.code = this.$route.params.code;
    }
};
</script>

<style scoped>

</style>

