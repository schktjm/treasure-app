import firebase from 'firebase/app'
import 'firebase/database'

const FirebaseDatabase = () => {
    let database = firebase.database();
    return {
        putItem({uid, item}) {
            const newPostKey = firebase.database().ref().child(`todos/${uid}`).push().key;
            let newItem = {};
            newItem[`todos/${uid}/${newPostKey}`] = {...item, id: newPostKey};
            database.ref().update(newItem)
        },
        getAll({uid}) {
            return database.ref(`todos/${uid}`).once('value').then(snapshot => {
                return snapshot.val();
            })
        },
        setStatus({uid, item}) {
            let newItem = {};
            newItem[`todos/${uid}/${item.id}`] = {...item};
            database.ref().update(newItem)
        },
        getById({uid, itemID}) {
            return database.ref(`todos/${uid}/${itemID}`).once('value').then(ss => ss.val())
        }
    }
};

export default FirebaseDatabase()
