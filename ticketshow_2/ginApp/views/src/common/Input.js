import React from "react";

const inputStyle =
    "form-control form-input form-input-bordered w-full p-2 rounded-lg pl-2 shadow-inner";

const Input = props => (
    <div className={props.className || "flex-1"}>
        <input
            className={inputStyle}
            type={props.type}
            defaultValue={props.defaultValue}
            name={props.name}
            onChange={props.onChange}
            placeholder={props.placeholder}
        />
        {props.error ? (
            <small className="float-right mt-2 text-red-light">{props.errorText}</small>
        ) : null}
    </div>
);

export default Input;
