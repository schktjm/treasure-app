<template>
  <el-button type="primary" @click="login">Log In</el-button>
</template>

<script>
    import firebase from '../js/firebase'

    export default {
        methods: {
            login() {
                firebase.login();
                this.$router.push('/', () => {
                })
            },
        }
    }
</script>
