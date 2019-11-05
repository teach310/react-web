import * as firebase from "firebase/app";
import "firebase/auth"

const firebaseConfig = {
    apiKey: "AIzaSyDPUd40mlqJUHMurOuHxc3-ihOlcYD1EsQ",
    authDomain: "todo-web-8e31b.firebaseapp.com",
    databaseURL: "https://todo-web-8e31b.firebaseio.com",
    projectId: "todo-web-8e31b",
    storageBucket: "todo-web-8e31b.appspot.com",
    messagingSenderId: "43613592062",
    appId: "1:43613592062:web:496946f7d317529c5648af",
    measurementId: "G-ZGXXP08L0S"
};

firebase.initializeApp(firebaseConfig);

export const googleProvider = new firebase.auth.GoogleAuthProvider();
export default firebase;