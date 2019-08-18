<template>
  <div>
    <a @click="onClick" v-if="isWhite">
      <el-card class="card-style" shadow="hover" :body-style="bodyStyle">
        {{item.text}}
      </el-card>
    </a>
    <a @click="onClick" v-else>
      <el-card class="card-style states-style" shadow="never" :body-style="bodyStyle">
        {{item.text}}
      </el-card>
    </a>
  </div>
</template>

<script>
    export default {
        name: 'Item',
        props: {
            item: {
                type: Object,
                default: () => ({
                    id: {
                        type: String,
                        default: ''
                    },
                    text: {
                        type: String,
                        default: ''
                    },
                    status: {
                        type: String,
                        default: 'white'
                    }
                })
            }
        },
        data() {
            return {
                bodyStyle: {
                    // padding: '2px 0 0 0',
                    'box-sizing': 'border-box',
                    height: '52px',
                    display: 'flex',
                    'align-items': 'center',
                    'justify-content': 'center'
                }
            }
        },
        computed: {
            isWhite() {
                return this.item.status === 'white'
            }
        },
        methods: {
            onClick() {
                this.$emit('click', this.item);
            }
        }
    }
</script>

<style scoped>
  .states-style {
    background-color: #EEEEEE;
    text-decoration: line-through;
  }

  .card-style {
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
    width: 100%;
    height: 52px;
    padding: 0px 0px;
    margin: 10px 0;
  }
</style>
