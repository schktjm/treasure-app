import firebase from 'firebase/app'
import 'firebase/auth'

const firebaseConfig = {
    apiKey: process.env.VUE_APP_FIREBASE_APIKEY,
    authDomain: process.env.VUE_APP_FIREBASE_AUTHDOMAIN,
    databaseURL: process.env.VUE_APP_FIREBASE_DATABASEURL,
    projectId: process.env.VUE_APP_FIREBASE_PROJECTID,
    messagingSenderId: process.env.VUE_APP_FIREBASE_MESSAGINGSENDERID,
    appId: process.env.VUE_APP_FIREBASE_APPID
};
// https://firebase.google.com/docs/auth/web/start
// see: GitHub認証の統合 https://firebase.google.com/docs/auth/web/github-auth?hl=ja
const githubProvider = new firebase.auth.GithubAuthProvider()

const FirebaseFactory = () => {
    firebase.initializeApp(firebaseConfig);
    // let database = firebase.database();
    let auth = firebase.auth();
    return {
        auth() {
            return auth
        },

        login() {
            return auth.signInWithPopup(githubProvider)
        },

        logout() {
            return auth.signOut()
        }
    }
};

export default FirebaseFactory()
