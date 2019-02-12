import React, { Component } from "react";
import { Link } from "react-router-dom";

class Header extends Component {
    render() {
        const linkStyle = "no-underline px-4 pt-1";
        return (
            <div className="flex justify-start flex-row bg-blue-light py-6 shadow-lg mb-4 align-baseline">
                <Link to="/" className={linkStyle}>
                    <div className="pl-8 pr-8 text-blue-darker font-bold">Not-ticketshow</div>
                </Link>
                <div className="pr-8 text-lg text-grey-darker">|</div>
                <Link to="/" className={linkStyle}>
                    <div>Home</div>
                </Link>
                <Link to="/eventos" className={linkStyle}>
                    <div>Próximos eventos</div>
                </Link>
                <Link to="/about/" className={linkStyle}>
                    About
                </Link>
            </div>
        );
    }
}

export default Header;
