import React from "react";
import { Switch, Route } from "react-router-dom";
import { BrowserRouter } from "react-router-dom";

import Header from "../common/Header";
import Index from "../dashboard/Index";
import About from "../dashboard/About";
import Eventos from "../dashboard/Eventos";

class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <BrowserRouter>
                <div className="bg-blue-lightest min-h-screen">
                    <Header />
                    <Switch>
                        <Route exact path="/" component={Index} />
                        <Route exact path="/about" component={About} />
                        <Route exact path="/eventos" component={Eventos} />
                    </Switch>
                </div>
            </BrowserRouter>
        );
    }
}
export default LoggedIn;
