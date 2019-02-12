import React, { Component } from "react";

import Concierto from "../concierto/Concierto";
import Sedes from "../Sedes/Sedes";

class Index extends Component {
    render() {
        return (
            <div className="w-4/5 mx-auto pt-8">
                <h3 className="text-grey-darkest">Nuestras Sedes</h3>
                <Sedes />
            </div>
        );
    }
}

export default Index;
