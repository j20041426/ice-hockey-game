<template>
    <div class="main">
        <div class="title text-reflect">冰球对战</div>
        <mu-raised-button @click="createGame()" label="创建游戏" class="demo-raised-button" primary/>
        <mu-raised-button @click="joinGame()" label="加入游戏" class="demo-raised-button" primary/>
        <mu-dialog :open="dialog" title="房间码">
            这是一个简单的弹出框
            <mu-flat-button slot="actions" @click="dialog=false" primary label="取消"/>
            <mu-flat-button slot="actions" primary @click="confirm" label="确定"/>
        </mu-dialog>
    </div>
</template>

<script>
export default {
    data() {
        return {
            dialog: false
        };
    },
    methods: {
        createGame() {
            this.$axios.get("http://localhost:38080/create").then(res => {
                if (res.status == 200 && res.data) {
                    this.$router.push({ path: `/game/${res.data}` });
                }
            });
        },
        joinGame() {
            this.dialog = true;
        },
        confirm() {

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
