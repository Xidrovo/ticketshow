import React from "react";
import ReactDOM from "react-dom";
import "./styles.css";

import LoggedIn from "./Login/LoggedIn";
import Home from "./Home/Home";

class Index extends React.Component {
    render() {
        if (!this.loggedIn) {
            return <LoggedIn />;
        } else {
            return <Home />;
        }
    }
}

ReactDOM.render(<Index />, document.getElementById("index"));
