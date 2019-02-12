import React, { Component } from "react";

import Sede from "./container/Sede";

const sedeRoute = "http://ec2-18-191-125-159.us-east-2.compute.amazonaws.com/api/sedes";
class Sedes extends Component {
    constructor(props) {
        super(props);
        this.state = {
            Sedes: []
        };
    }
    componentDidMount() {
        fetch(sedeRoute)
            .then(res => res.json())
            .then(
                result => {
                    this.setState({
                        Sedes: result
                    });
                },
                error => {
                    console.log("Error!!", error);
                }
            );
    }
    render() {
        return (
            <div className="flex w-full mt-4 flex-wrap">
                {this.state.Sedes.map(x => (
                    <div className="w-1/3 min-h-12 max-h-1/3 px-2" key={x.ID}>
                        <Sede
                            ciudad={x.ciudad}
                            url_img={x.url_img}
                            description={x.description}
                            name={x.name}
                            place={x.place}
                            capacity={x.capacity}
                        />
                    </div>
                ))}
            </div>
        );
    }
}

export default Sedes;

//ciudad, capacity, description, name, place, url_img
