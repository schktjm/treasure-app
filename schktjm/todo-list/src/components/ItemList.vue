<template>
  <div id="item-list">
    <Item v-for="item in items" :key="item.id" :item="item" @click="o => $emit('click',o)"></Item>
  </div>
</template>

<script>
    import Item from './Item'

    export default {
        name: 'ItemList',
        components: {Item},
        props: {
            items: {
                type: Array,
                default: () => ([])
            }
        }
    }
</script>

<style scoped>
  #item-list {
    width: 80%;
  }
</style>
