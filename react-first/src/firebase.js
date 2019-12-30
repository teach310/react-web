import * as firebase from "firebase/app";
import "firebase/auth"

var firebaseConfig = {
    apiKey: "AIzaSyBu-ChZQ4_JFOVcYwhnYYsJotCQoy63Mc8",
    authDomain: "todolist-65933.firebaseapp.com",
    databaseURL: "https://todolist-65933.firebaseio.com",
    projectId: "todolist-65933",
    storageBucket: "todolist-65933.appspot.com",
    messagingSenderId: "117984386438",
    appId: "1:117984386438:web:4ada3859a20f91c6932007",
    measurementId: "G-HSY1SSLM82"
};

firebase.initializeApp(firebaseConfig);

export const googleProvider = new firebase.auth.GoogleAuthProvider();
export default firebase;