<template>
  <div id="home">
    <h1>todos</h1>
    <ItemList class="item-list-style" :items="items" @click="changeState"></ItemList>
    <ItemFrom class="item-list-style" @click="addItems"></ItemFrom>
  </div>
</template>

<script>
    // @ is an alias to /src
    import firebase from '../js/firebase'
    import database from '../js/database'
    import ItemList from '../components/ItemList'
    import ItemFrom from '../components/ItemFrom'

    export default {
        name: 'home',
        data() {
            return {
                user: null,
                items: []
            }
        },
        components: {
            ItemList,
            ItemFrom
        },
        mounted() {
            firebase.auth().onAuthStateChanged(async user => {
                if (user) {
                    this.user = user;
                    this.items = await this.getAllItems();
                } else {
                    this.$router.push('/login', () => {
                    })
                }
            });
        },
        methods: {
            async addItems(text) {
                const newObj = {text, status: 'white'};
                if (this.user) {
                    database.putItem({uid: this.user.uid, item: newObj})
                }
                this.items = await this.getAllItems();
            },
            async changeState(i) {
                const newState = (i.status === 'white') ? 'grey' : 'white';
                database.setStatus({uid: this.user.uid, item: {...i, status: newState}});
                this.items = await this.getAllItems();
            },
            async getAllItems() {
                return database.getAll(this.user)
                    .then(res => {
                        return Object.values(res).map(x => x)
                    })
                    .catch(err => {
                        return [];
                    })
            }
        }
    }
</script>

<style scoped>
  #home {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 100%;
  }
</style>
