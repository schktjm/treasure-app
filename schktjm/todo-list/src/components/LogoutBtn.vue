<template>
  <el-button type="info" @click="logout">Log Out</el-button>
</template>

<script>
    import firebase from '../js/firebase'

    export default {
        methods: {
            logout() {
                firebase.logout();
            }
        }
    }
</script>
