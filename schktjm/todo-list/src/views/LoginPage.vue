<template>
  <div id="login-page">
    <logout-btn v-if="isLoggedin"></logout-btn>
    <login-btn v-else></login-btn>
  </div>
</template>

<script>
    import firebase from '../js/firebase'
    import LoginBtn from '../components/LoginBtn'
    import LogoutBtn from '../components/LogoutBtn'

    export default {
        name: 'LoginPage',
        components: {
            LogoutBtn,
            LoginBtn
        },
        data() {
            return {
                isLoggedin: false
            }
        },
        mounted() {
            firebase.auth().onAuthStateChanged(user => {
                if (user) {
                    this.isLoggedin = true;
                    this.$router.push('/', () => {
                    });
                } else {
                    this.isLoggedin = false
                }
            })
        }
    }
</script>

<style scoped>
  #login-page {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
  }
</style>
