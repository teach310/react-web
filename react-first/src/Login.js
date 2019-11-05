import React, { useContext } from 'react'
import { makeStyles, Button, Card, CardActions, CardContent, Typography, Box } from '@material-ui/core';
import { AuthContext } from './auth'

const useStyles = makeStyles(theme => ({
    root: {
        marginTop: 200,
    },
    card: {
        margin: 'auto',
        width: 300
    },
    content: {
        textAlign: 'center',
    },
    header: {
        margin: theme.spacing(1),
        flexGrow: 1,
    },
    title: {
        fontSize: 14,
    },
    signinButton: {
        margin: 'auto',
        width: 280,
        textTransform: 'none'
    }
}));

export default () => {
    const classes = useStyles();
    const { signin } = useContext(AuthContext)

    return (
        <Box className={classes.root}>
            <Card className={classes.card}>
                <CardContent className={classes.content}>
                    <Typography component="h5" variant="h5">
                        Sign in to TODOアプリ
                    </Typography>
                </CardContent>
                <CardActions>
                    <Button variant="contained" color="primary" onClick={signin} size="small" className={classes.signinButton}>Sign in with Google</Button>
                </CardActions>
            </Card>
        </Box>
    )
}