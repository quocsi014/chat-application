import { BiLoaderAlt } from "react-icons/bi";

function LoadingButton(props) {
  const { handleFunction, className, disabled, value, isLoading } = props;

  return !disabled ? (
    <button
      className={`w-full bg-blue-500 py-2 text-xl font-semibold text-white rounded-xl flex justify-center items-center ${className}`}
      onClick={() => {
        handleFunction();
      }}
    >
      {isLoading ? <div className="animate-spin"><BiLoaderAlt size={28}/></div> : value}
    </button>
  ) : (
    <button
      className="w-full bg-gray-200 py-2 text-xl font-semibold text-white rounded-xl flex justify-center items-center"
      onClick={() => {
        handleFunction();
      }}
      disabled
    >
      {isLoading ? <div className="animate-spin"><BiLoaderAlt size={28}/></div> : value}
    </button>
    
  );  
}

export default LoadingButton;
