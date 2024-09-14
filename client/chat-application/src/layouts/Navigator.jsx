import { BiConversation, BiMailSend } from "react-icons/bi";
import { MdChat, MdSend, MdSettings } from "react-icons/md";
import { Link, NavLink } from "react-router-dom";

function Navigator(props) {
  const handleActiveClass = ({ isActive }) => {
    return isActive ? "mb-3 text-gray-700" : "mb-3 text-gray-300";
  };
  return (
    <div className="h-full w-16 p-5 rounded-2xl mx-2 bg-white grow-0 shrink-0 flex flex-col items-center">
      <NavLink className={handleActiveClass} to={"/conversations"}>
        <MdChat size={30} />
      </NavLink>
      <NavLink 
        className={({ isActive }) => 
          isActive || location.pathname.startsWith('/requests') 
            ? "mb-3 text-gray-700" 
            : "mb-3 text-gray-300"
        } 
        to={"/requests/received"}
      >
        <MdSend size={30} />
      </NavLink>
      <NavLink className={handleActiveClass} to={"/settings"}>
        <MdSettings size={30} />
      </NavLink>
    </div>
  );
}

export default Navigator;
