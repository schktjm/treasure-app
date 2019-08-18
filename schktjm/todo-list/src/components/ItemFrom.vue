<template>
  <div id="input-form">
    <el-input id="input-area" placeholder="Please input" v-model="text"></el-input>
    <span style="width: 12px;"></span>
    <el-button id="input-button" type="danger" @click="onSubmit">add</el-button>
  </div>
</template>

<script>
    export default {
        name: 'ItemFrom',
        data() {
            return {
                text: ''
            }
        },
        methods: {
            onSubmit() {
                if (this.text.length !== 0) {
                    this.$emit('click', this.text);
                    this.text = '';
                }
            }
        }
    }
</script>

<style scoped>
  #input-form {
    display: flex;
    align-items: center;
    box-sizing: border-box;
    width: 80%;
    height: 52px;
  }

  #input-area {

  }
</style>
