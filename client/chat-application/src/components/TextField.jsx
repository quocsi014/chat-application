import { useEffect, useRef } from "react";

function TextField(props) {
  const { title, type, value, setValue, errorMessage, setErrorMessage, id } =
    props;

  return (
    <div className="flex flex-col mb-2 w-full">
      <label className="mb-1 font-semibold" htmlFor={id}>
        {title}
      </label>
      <input
        className={`outline-none px-2 border-0 border-b-2 focus:border-gray-400 text-lg w-96 ${
          errorMessage != "" ? "border-red-300" : "border-gray-100"
        }`}

        type={type}
        value={value}

        onChange={(e) => {
          setValue(e.target.value);
        }}

        onFocus={()=>{
          setErrorMessage("")
        }}
      />
      <span className="text-red-400 text-sm w-full h-5">{errorMessage}</span>
    </div>
  );
}

export default TextField;
