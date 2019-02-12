import React from "react";

const ConciertoItem = props => {
    return (
        <div className="mt-6 w-1/3 shadow-inner">
            <h2 className="text-grey-darkest pt-8 ml-8">Concierto Name</h2>
            <div className="w-full flex">
                <img
                    className="mt-8 mx-auto"
                    src="https://comicconecuador.com/wp-content/uploads/2019/01/logo.png"
                />
            </div>
            <div className="mt-8 text-lg ml-4">
                <span>Dónde?</span> Centro de convenciones de guayaquil
            </div>
            <div className="mt-8 text-lg ml-4">
                Qué esperas quedan <span className="font-bold">200</span> cupos!
            </div>
        </div>
    );
};

export default ConciertoItem;
