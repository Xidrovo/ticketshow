import React from "react";
import DownIcon from "@icons/DownIcon";
const selectStyle = "form-control form-input form-input-bordered form-select-bordered w-full";
const selectStyleSmall =
    "form-control form-input form-input--small form-input-bordered form-select-bordered w-full";

const Select = props => {
    <div className={props.className || "inline-block relative flex-1"}>
        <select
            className={props.small ? selectStyleSmall : selectStyle}
            value={props.value}
            defaultValue={props.defaultValue}
            onChange={props.onChange}
            name={props.name}
        >
            <option key="-1" value="-1" disabled>
                {props.placeHolder || "Escoja una empresa"}
            </option>
            {props.options.map(option => (
                <option key={option.id} value={option.id}>
                    {option.name}
                </option>
            ))}
        </select>
        <div className="pointer-events-none absolute pin-y pin-r flex items-center mx-2 text-grey">
            | <DownIcon className=" ml-1 fill-current text-grey" width="22px" height="22" />
        </div>
    </div>;
};

export default Select;
