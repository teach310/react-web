import React, { useState } from 'react';
import ReactDOM from 'react-dom';
import './App.css';
// import butachan from './img/butachan.png';
import * as serviceWorker from './serviceWorker';
import axios from 'axios';

// const API_URL = 'localhost:8080'

// const GithubProfile = () => {

//     const [content, setContent] = useState("api result")
//     const userName = "teach310"
//     const getProfile = async () => {
//         try {
//             const result = await axios.get(`https://api.github.com/users/${userName}`);
//             console.log(result)
//             console.log(result.data)
//             setContent(JSON.stringify(result.data))
//         } catch (error) {
//             console.log(error);
//         }
//     }

//     return (
//         <div>
//             <button onClick={() => getProfile()}>get profile!</button>
//             {content}
//         </div>
//     )
// }

// const ThemeContext = React.createContext();

// function Toolbar(props){
//     return (
//         <div>
//             <ThemedButton />
//         </div>
//     );
// }

// const ThemedButton = () => (
//     <ThemeContext.Consumer>
//         {context=>
//             <button onClick={context.toggleTheme}>{context.theme}</button>
//         }
//     </ThemeContext.Consumer>
// )

// function HookExample(){
//     const [count, setCount] = useState(0)

//     return (
//         <div>
//             <p>You clicked {count} times</p>
//             <button onClick={()=>setCount(count+1)}>
//                 Click me
//             </button>
//         </div>
//     );
// }

// class Popup extends React.Component {
//     constructor(props){
//         super(props);
//         this.state = {
//             value: 'coconut'
//         }
//     }

//     handleChange = ev => {
//         this.setState({value: ev.target.value});
//     }

//     handleSubmit = ev => {
//         alert("selected: " + this.state.value);
//         ev.preventDefault();
//     }

//     render(){
//         return (
//             <form onSubmit={this.handleSubmit}>
//                 <label>
//                     Flavor　
//                     <select value={this.state.value} onChange={this.handleChange}>
//                         <option value="grapefruit">Grapefruit</option>
//                         <option value="lime">Lime</option>
//                         <option value="coconut">Coconut</option>
//                         <option value="mango">Mango</option>
//                     </select>
//                 </label>
//                 <input type="submit" value="Submit"/>
//             </form>
//         );
//     }
// }


class InputField extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            value: 'content'
        }
    }

    handleChange = ev => {
        this.setState({ value: ev.target.value });
    }

    handleSubmit = ev => {
        alert("submitted: " + this.state.value);
        ev.preventDefault();
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    TextArea
                    <textarea value={this.state.value} onChange={this.handleChange} />
                </label>
                <input type="submit" value="Submit" />
            </form>
        );
    }
}

class NameForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = { value: '' }
    }

    handleChange = ev => {
        this.setState({ value: ev.target.value });
    }

    handleSubmit = ev => {
        alert('A name was submitted: ' + this.state.value);
        ev.preventDefault();
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    Name:
                    <input type="text" value={this.state.value} onChange={this.handleChange} />
                </label>
                <input type="submit" value="Submit" />
            </form>
        );
    }
}



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
            <ul>{listItems}</ul>
        </div>
    );
}

const host = "http://localhost:8080"

function App() {

    const initialModels = [
        { id: 1, isDone: false, name: "anyTodo1" },
        { id: 2, isDone: true, name: "anyTodo2" },
        { id: 3, isDone: true, name: "anyTodo3" },
    ];

    const [todoModels, setTodoModels] = useState(initialModels);

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
            const result = await axios.get(`http://go:8080/load`);
            console.log(result.data)
            console.log(result)
            // setContent(JSON.stringify(result.data))
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <div>
            <h1>TODOアプリ</h1>
            <div>
                <button onClick={loadTodo}>LOAD</button>
                <button>SAVE</button>
                <button onClick={handleAdd}>ADD</button>
            </div>
            <TodoList models={todoModels} handleClick={handleClick} handleRemove={handleRemove} handleChangeName={handleChangeName}/>
            {/* <Toggle />
            <NameForm />
            <InputField />
            <Popup />
            <HookExample />
            <ThemeContext.Provider value={{ theme: theme, toggleTheme: toggleTheme }}>
                <Toolbar />
            </ThemeContext.Provider>
            <GithubProfile /> */}
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

// ReactDOM.render(<App />, document.getElementById('root'));
ReactDOM.render(
    <App />,
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
