import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {

    if (window.WebSocket) {
        const socket = new WebSocket("ws://localhost:3000/ws")
        socket.addEventListener('open', (event) => {
            socket.send("Hello from the frontend")
        })
    }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
