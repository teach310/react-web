import React, { useState } from 'react';
import ReactDOM from 'react-dom';
import Button from '@material-ui/core/Button';
import AppBar from '@material-ui/core/AppBar';
import * as serviceWorker from './serviceWorker';
import axios from 'axios';
import { makeStyles, Typography, Checkbox, TextField, IconButton, List, ListItem, Card } from '@material-ui/core';
import DeleteIcon from '@material-ui/icons/Delete';


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
                props.models.length != 0 ? <List>{listItems}</List> : 'no task'
            }
        </div>
    );
}

const host = "http://localhost:8080"

function App() {

    const classes = useStyles();
    const [todoModels, setTodoModels] = useState([]);

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
            const result = await axios.get(host+`/load`);
            setTodoModels(result.data)
        } catch (error) {
            console.log(error);
        }
    }

    const saveTodo = async () => {
        try {
            const result = await axios.post(host+`/save`, todoModels);
            if (result.status == 200) {
                alert("save succeeded")
            }
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <div>
            <AppBar position="static">
                <Typography variant="h6" className={classes.title}>TODOアプリ</Typography>
            </AppBar>
            <div className={classes.container}>
                <Button className={classes.horizontalButton} variant="contained" onClick={loadTodo} size="small">LOAD</Button>
                <Button className={classes.horizontalButton} variant="contained" onClick={saveTodo} size="small">SAVE</Button>
                <Button className={classes.horizontalButton} variant="contained" onClick={handleAdd} size="small">ADD</Button>
            </div>
            <TodoList models={todoModels} handleClick={handleClick} handleRemove={handleRemove} handleChangeName={handleChangeName}/>
        </div>
    );
}

ReactDOM.render(
    <App />,
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
