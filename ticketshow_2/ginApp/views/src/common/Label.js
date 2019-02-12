import React from "react";
const labelStyle =
    "text-base text-80 leading-tight w-full block font-sans text-left text-grey-darker pt-2";

const Label = props => (
    <div className="w-1/4">
        <label className={labelStyle}>{props.name}</label>
    </div>
);
export default Label;
