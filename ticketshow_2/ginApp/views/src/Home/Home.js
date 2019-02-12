import React from "react";
import Input from "../common/Input";
import Label from "../common/Label";

class Home extends React.Component {
    render() {
        return (
            <div className="bg-blue-lightest absolute pin h-1/2 w-2/5 mt-32 mb-auto mx-auto rounded-lg shadow">
                <div className="mt-16 text-center">
                    <h1 className="text-blue-darker mb-20">Log in</h1>
                    <div className="flex items-center mb-12 w-3/5 justify-center mx-auto">
                        <Label name="Username:" />
                        <Input
                            type="text"
                            onChange={this.props.handleChange}
                            placeholder="Nombre completo"
                            name="name"
                        />
                    </div>
                    <div className="flex items-center mb-12 w-3/5 justify-center mx-auto">
                        <Label name="Password:" />
                        <Input
                            type="text"
                            onChange={this.props.handleChange}
                            placeholder="Nombre completo"
                            name="name"
                        />
                    </div>
                    <button
                        onClick={this.authenticate}
                        className="bg-blue-light hover:bg-blue text-blue-darkest font-bold py-2 px-4 rounded-lg mt-12"
                    >
                        Sign In
                    </button>
                </div>
            </div>
        );
    }
}

export default Home;

{
    /* <Label name="Canal" />
    <Input type="text" onChange={this.props.handleChange} placeholder="Nombre completo" name="name" /> */
}
