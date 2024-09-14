import { BiLoaderAlt } from "react-icons/bi";

function Loading(props){
  const {size, className} = props

  return(
    <div className={`origin-center animate-spin size-fit ${className}`}>
      <BiLoaderAlt size={size}/>
    </div>
  )
}
export default Loading