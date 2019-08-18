<template>
  <div id="app">
    <logout-btn></logout-btn>
    <div id="app-box">
      <div id="base">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
    import LogoutBtn from './components/LogoutBtn'

    export default {
        components: {LogoutBtn}
    }
</script>

<style lang="scss">
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
  }

  #app-box {
    display: flex;
    justify-content: center;
  }

  #base {
    width: 60%;
    min-width: 310px;
    max-width: 1011px;
    height: 85vh;
    border: #ffcdd2 5px solid;
    margin: 10px 0 30px 0;
  }

</style>
