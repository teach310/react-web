import React, { useState } from 'react';
import ReactDOM from 'react-dom';
import './App.css';
import butachan from './img/butachan.png';
import * as serviceWorker from './serviceWorker';

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
    return (
        <div>
            <Clock />
            <Toggle />
            <NameForm />
            <InputField />
            <Popup />
            <HookExample />
            {/* <Butachan text="はーい" /> */}
        </div>
    );
}

function Butachan(props){
    return (
        <div id="footer">
            <div>{props.text}</div> 
            <div><img src={butachan} alt="ぶたちゃん"/></div>
        </div>            
    );
}

// ReactDOM.render(<App />, document.getElementById('root'));
ReactDOM.render(
    App(),
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
