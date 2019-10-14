import React, { useState } from 'react';
import ReactDOM from 'react-dom';
import './App.css';
// import butachan from './img/butachan.png';
import * as serviceWorker from './serviceWorker';
import axios from 'axios';

const Toggle = props => (
    <button onClick={props.handleClick}>
        {!props.isDone ? '□' : '■'}
    </button>
);

const TodoItem = props => (
    <div>
        <Toggle isDone={props.model.isDone} handleClick={props.handleClick}/>
        <input type="text" value={props.model.name} onChange={ev => props.handleChangeName(ev.target.value)}/> 
        <button onClick={props.handleRemove}>×</button>
    </div>
)

const TodoList = props => {

    const listItems = props.models.map((m, i) =>
        <li key={m.id}>
             <TodoItem model={m} handleClick={() => props.handleClick(i)} handleRemove={() => props.handleRemove(i)} handleChangeName={value => props.handleChangeName(i, value)}/>
        </li>
    );

    return (
        <div>
            {
                props.models.length != 0 ? <ul>{listItems}</ul> : 'no task'
            }
        </div>
    );
}

const host = "http://localhost:8080"

function App() {

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
            console.log("save")
            if (result.status == 200) {
                alert("save succeeded")
            }
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <div>
            <h1>TODOアプリ</h1>
            <div>
                <button onClick={loadTodo}>LOAD</button>
                <button onClick={saveTodo}>SAVE</button>
                <button onClick={handleAdd}>ADD</button>
            </div>
            <TodoList models={todoModels} handleClick={handleClick} handleRemove={handleRemove} handleChangeName={handleChangeName}/>
            {/* <Butachan text="はーい" /> */}
        </div>
    );
}

// function Butachan(props){
//     return (
//         <div id="footer">
//             <div>{props.text}</div> 
//             <div><img src={butachan} alt="ぶたちゃん"/></div>
//         </div>            
//     );
// }

ReactDOM.render(
    <App />,
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
