import React, { useState, useContext } from 'react';
import Button from '@material-ui/core/Button';
import AppBar from '@material-ui/core/AppBar';
import { makeStyles, Typography, Checkbox, TextField, IconButton, List, ListItem, Card, Toolbar, Menu, MenuItem } from '@material-ui/core';
import DeleteIcon from '@material-ui/icons/Delete';
import KeyboardArrowDownIcon from '@material-ui/icons/KeyboardArrowDown';
import { AuthProvider, AuthContext } from './auth'
import Loading from './Loading'
import Router from './Router'
import Login from './Login'
import webrequest from './webrequest'


const useStyles = makeStyles(theme => ({
    horizontalButton: {
        marginRight : theme.spacing(1),
    },
    container: {
        margin: theme.spacing(1),
    },
    title: {
        margin: theme.spacing(1),
        flexGrow: 1,
    },
    task: {
        padding: theme.spacing(1),
    }
}));


const TodoItem = props => {
    const classes = useStyles();
    return (
        <div>
            <Card className={classes.task}>
            <Checkbox checked={props.model.isDone} onChange={props.handleClick} />
            <TextField value={props.model.name} onChange={ev => props.handleChangeName(ev.target.value)}/>
            <IconButton onClick={props.handleRemove} size='small'><DeleteIcon fontSize="small"/></IconButton>
            </Card>
        </div>
    );
}

const TodoList = props => {

    const listItems = props.models.map((m, i) =>
        <ListItem key={m.id}>
             <TodoItem model={m} handleClick={() => props.handleClick(i)} handleRemove={() => props.handleRemove(i)} handleChangeName={value => props.handleChangeName(i, value)}/>
        </ListItem>
    );

    return (
        <div>
            {
                props.models.length !== 0 ? <List>{listItems}</List> : 'no task'
            }
        </div>
    );
}

const AccountButton = () => {
    const [anchorEl, setAnchorEl] = useState(null);
    const { currentUser, signout } = useContext(AuthContext);

    const handleClick = event => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    return (
        <>
            <Button onClick={handleClick} color="inherit">{currentUser.displayName}<KeyboardArrowDownIcon fontSize="small"/></Button>
            <Menu
                anchorEl={anchorEl}
                keepMounted
                open={Boolean(anchorEl)}
                onClose={handleClose}
            >
                <MenuItem onClick={signout}>signout</MenuItem>
            </Menu>
        </>
    );
}

function Todos() {

    const classes = useStyles();
    const [todoModels, setTodoModels] = useState([]);
    const authContext = useContext(AuthContext);

    const handleClick = i => {
        const temp = todoModels.slice();
        temp[i] = Object.assign({}, temp[i], { isDone: !temp[i].isDone});
        setTodoModels(temp);
    };

    const handleChangeName = (i, value) => {
        const temp = todoModels.slice();
        temp[i] = Object.assign({}, temp[i], { name: value });
        setTodoModels(temp);
    }

    const handleRemove = i => {
        const temp = todoModels.slice();
        temp.splice(i, 1)
        setTodoModels(temp);
    };

    const handleAdd = () => {
        setTodoModels(prev => prev.concat([{id: prev.length+1, isDone: false, name: "task" }]));
    };

    const loadTodo = async () => {
        try {
            const result = await webrequest.get(`/load`, authContext);
            setTodoModels(result.data.todoList)
        } catch (e) {
            console.error(e)
        }
    }

    const saveTodo = async () => {
        try {
            const result = await webrequest.post(`/save`, { todoList: todoModels}, authContext);
            if (result.status === 200) {
                alert("save succeeded")
            }
        } catch (e) {
            console.error(e)
        }
    }

    return (
        <>
            <AppBar position="static">
               <Toolbar>
                    <Typography variant="h6" className={classes.title}>TODOアプリ</Typography>
                    <AccountButton />
                </Toolbar>
            </AppBar>
            
            <div className={classes.container}>
                <Button className={classes.horizontalButton} variant="contained" onClick={loadTodo} size="small">LOAD</Button>
                <Button className={classes.horizontalButton} variant="contained" onClick={saveTodo} size="small">SAVE</Button>
                <Button className={classes.horizontalButton} variant="contained" onClick={handleAdd} size="small">ADD</Button>
            </div>
            <TodoList models={todoModels} handleClick={handleClick} handleRemove={handleRemove} handleChangeName={handleChangeName}/>
        </>
    );
}

export default () => (
    <AuthProvider>
        <Router
            renderLoading={() => <Loading />}
            renderTodos={() => <Todos />}
            renderLogin={() => <Login />}
        />
    </AuthProvider>
)