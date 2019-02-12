import React, { Component } from "react";

import Concierto from "../concierto/Concierto";
import Sedes from "../Sedes/Sedes";

const sedeRoute = "http://ec2-18-191-125-159.us-east-2.compute.amazonaws.com/api/eventos";

class Eventos extends Component {
    constructor(props) {
        super(props);
        this.state = {
            Eventos: []
        };
    }
    componentDidMount() {
        fetch(sedeRoute)
            .then(res => res.json())
            .then(
                result => {
                    this.setState({
                        Eventos: result
                    });
                },
                error => {
                    console.log("Error!!", error);
                }
            );
    }
    render() {
        return (
            <div className="w-4/5 mx-auto pt-8">
                <h3 className="text-grey-darkest mb-4">Próximos eventos</h3>
                {this.state.Eventos.map(evento => (
                    <div key={evento.id}>
                        <Concierto name={evento.name} />
                    </div>
                ))}
            </div>
        );
    }
}

export default Eventos;
