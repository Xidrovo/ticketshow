import React from "react";

const Sede = props => {
    return (
        <React.Fragment>
            <div className="m-2 hover:shadow-lg p-6 shadow">
                <div className="text-lg text-grey-darker my-2">{props.name}</div>
                <img src={props.url_img} />
                <div className="text-grey-darker mt-4 p-2">{props.description}</div>
                <div className="flex flex-row justify-left mt-2 text-grey-darker text-xs">
                    <label className="mr-2">Direccion:</label>
                    <div className="pr-2">{props.place}</div>
                </div>
                <div className="text-grey-darker mt-2 text-right pr-2">
                    Capacidad para <b>{props.capacity}</b>
                </div>
            </div>
        </React.Fragment>
    );
};

export default Sede;

// ciudad, capacity, description, name, place, url_img;
