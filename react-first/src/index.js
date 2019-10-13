import React, { useState } from 'react';
import ReactDOM from 'react-dom';
import './App.css';
// import butachan from './img/butachan.png';
import * as serviceWorker from './serviceWorker';
import axios from 'axios';

// const API_URL = 'localhost:8080'

const GithubProfile = () => {

    const [content, setContent] = useState("api result")
    const userName = "teach310"
    const getProfile = async () => {
        try {
            const result = await axios.get(`https://api.github.com/users/${userName}`);
            console.log(result)
            console.log(result.data)
            setContent(JSON.stringify(result.data))
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <div>
            <button onClick={() => getProfile()}>get profile!</button>
            {content}
        </div>
    )
}

const ThemeContext = React.createContext();

function Toolbar(props){
    return (
        <div>
            <ThemedButton />
        </div>
    );
}

const ThemedButton = () => (
    <ThemeContext.Consumer>
        {context=>
            <button onClick={context.toggleTheme}>{context.theme}</button>
        }
    </ThemeContext.Consumer>
)

function HookExample(){
    const [count, setCount] = useState(0)

    return (
        <div>
            <p>You clicked {count} times</p>
            <button onClick={()=>setCount(count+1)}>
                Click me
            </button>
        </div>
    );
}

class Popup extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            value: 'coconut'
        }
    }

    handleChange = ev => {
        this.setState({value: ev.target.value});
    }

    handleSubmit = ev => {
        alert("selected: " + this.state.value);
        ev.preventDefault();
    }

    render(){
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    Flavor　
                    <select value={this.state.value} onChange={this.handleChange}>
                        <option value="grapefruit">Grapefruit</option>
                        <option value="lime">Lime</option>
                        <option value="coconut">Coconut</option>
                        <option value="mango">Mango</option>
                    </select>
                </label>
                <input type="submit" value="Submit"/>
            </form>
        );
    }
}


class InputField extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            value: 'content'
        }
    }

    handleChange = ev => {
        this.setState({value: ev.target.value});
    }

    handleSubmit = ev => {
        alert("submitted: " + this.state.value);
        ev.preventDefault();
    }

    render(){
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    TextArea　
                    <textarea value={this.state.value} onChange={this.handleChange}/>
                </label>
                <input type="submit" value="Submit"/>
            </form>
        );
    }
}

class NameForm extends React.Component {
    constructor(props){
        super(props);
        this.state = {value: ''}
    }

    handleChange = ev => {
        this.setState({value: ev.target.value});
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
                <input type="submit" value="Submit"/>
            </form>
        );
    }
}

class Toggle extends React.Component {
    constructor(props){
        super(props);
        this.state = {isToggleOn: true};
    }

    handleClick = () => {
        this.setState(st => ({
            isToggleOn: !st.isToggleOn
        }));
    }

    render() {
        return(
             <button onClick={this.handleClick}>
                { this.state.isToggleOn ? 'ON' : 'OFF' }
            </button>
        );
    }
}

class Clock extends React.Component {
    constructor(props) {
        super(props);
        this.state = {date: new Date()};
    }

    componentDidMount() {
        this.timerID = setInterval(
            () => this.tick(),
            1000
        );
    }

    componentWillUnmount(){
        clearInterval(this.timerID)
    }

    tick() {
        this.setState({
            date: new Date()
        });
    }

    render() {
        return (
            <div>
                <h1>Hello, world!</h1>
                <h2>It is {this.state.date.toLocaleTimeString()}</h2>
            </div>
        );    
    }
}

// const element = <Welcome name="Sara" />;

function App() {
    console.log("App");
    const [theme, setTheme] = useState("dark")
    const toggleTheme = () => setTheme(prev => prev==="dark" ? "light" : "dark" )

    return (
        <div>
            <Clock />
            <Toggle />
            <NameForm />
            <InputField />
            <Popup />
            <HookExample />
            <ThemeContext.Provider value={{theme: theme, toggleTheme: toggleTheme}}>
                <Toolbar  />
            </ThemeContext.Provider>
            <GithubProfile />
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
