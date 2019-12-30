import React, { useState, createContext, useCallback, useEffect } from 'react'
import firebase, { googleProvider } from './firebase'

const AuthContext = createContext()

const AuthProvider = ({ children }) => {
    const [loading, setLoading] = useState(true)
    const [currentUser, setCurrentUser] = useState(null)
    const [accessToken, setAccessToken] = useState(null)
    const [authState, setAuthState] = useState("dummy")
    
    const auth = firebase.auth()

    const signin = useCallback(async () => {
        setAuthState("start signin")
        try {
            setLoading(true)
            await auth.signInWithRedirect(googleProvider)
            setAuthState("end signin")
        } catch (e) {
            console.error(e.code, e.message)
            setLoading(false)
            setAuthState("error signin: " + e.message)
        }
        
    }, [auth])

    const signout = useCallback(async () => {
        try {
            setLoading(true)
            await auth.signOut()
        } catch (e) {
            console.error(e.code, e.message)
        }
    }, [auth])

    useEffect(() => {
        auth.onAuthStateChanged(async user => {
            setAuthState("start onAuthStateChanged")
            if(user){
                const token = await user.getIdToken();
                console.log(token)
                setAccessToken(token)
                setAuthState("onAuthStateChanged user: not null")
            }else{
                setAccessToken(null)
                setAuthState("onAuthStateChanged user: null")
            }
            setLoading(false)
            setCurrentUser(user)
            
        })
    }, [auth])

    const refreshAccessToken = useCallback(async () => {
        const token = await currentUser.getIdToken();
        setAccessToken(token)
        return token
    }, [currentUser])

    return (
        <AuthContext.Provider value={{ currentUser, accessToken, signin, signout, loading, refreshAccessToken, authState}}>
            {children}
        </AuthContext.Provider>
    )
}

export { AuthContext, AuthProvider }