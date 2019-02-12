import React from "react";

const Concierto = props => {
    return (
        <React.Fragment>
            <div className="m-2 hover:shadow-lg p-6 shadow">
                <div className="text-lg text-grey-darker my-2">{props.name}</div>
            </div>
        </React.Fragment>
    );
};

export default Concierto;

// creador: 1
// end_time: "2019-07-12"
// id: 1
// name: "Comic-Con"
// sede: 1
// start_time: "2019-07-10"
// valor: "{"VIP": "$25.00", "normal": "$5.00"}"
