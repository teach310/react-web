import React, { useContext } from 'react'
import { AuthContext } from './auth'

// あとでreact routerに置き換え
export default ({ renderLoading, renderLogin, renderTodos }) => {
    const { currentUser, loading } = useContext(AuthContext)

    return (
        <>
            { loading ? renderLoading() : currentUser ? renderTodos() : renderLogin()}
        </>
    )
}