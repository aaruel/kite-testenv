import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

const getSquare = (num) => {
  return fetch(`http://localhost:3579/math/square/${num}`)
    .then((r) => r.json())
}

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      number: 0.0
    }
  }

  componentDidMount() {
    getSquare(16).then((r) => {
      this.setState({number: r})
    })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <p className="App-intro">
          Square: {this.state.number}
        </p>
      </div>
    );
  }
}

export default App;
