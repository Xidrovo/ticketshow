import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";

class App extends React.Component {
    render() {
        if (this.loggedIn) {
            return <LoggedIn />;
        } else {
            return <Home />;
        }
    }
}

export default App;
