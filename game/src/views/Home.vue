<template>
    <div class="main">
        <div class="title text-reflect">冰球对战</div>
        <mu-raised-button @click="createGame()" label="创建游戏" class="demo-raised-button" primary/>
        <mu-raised-button @click="joinGame()" label="加入游戏" class="demo-raised-button" primary/>
    </div>
</template>

<script>
export default {
    data() {
        return {};
    },
    methods: {
        createGame() {
            if (window.WebSocket && !window.conn) {
                window.conn = new WebSocket("ws://localhost:38080/init");
                conn.onclose = event => {
                    console.log('disconnect.')
                }
                conn.onmessage = event => {
                    console.log('message', event);
                }
            }
        },
        joinGame() {
        }
    }
};
</script>

<style scoped>
.main {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-around;
    height: 40%;
    width: 100%;
    position: absolute;
    top: 30%;
}
.text-reflect {
    font-size: 30px;
    float: left;
    -webkit-box-reflect: below 0px -webkit-linear-gradient(transparent, transparent 50%, rgba(255, 255, 255, 0.5));
}
</style>
