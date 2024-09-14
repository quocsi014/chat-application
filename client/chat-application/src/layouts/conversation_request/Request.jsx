import { NavLink, Outlet } from "react-router-dom";

function Request() {
  const handleActiveClass = ({ isActive }) => {
    return `w-full text-xl font-bold ${isActive ? "mb-3 text-gray-700 border-b-4 border-gray-700" : "mb-3 text-gray-300"}`;
  };
  return (
    <div
      className={`h-full shrink-0 w-120 bg-white rounded-2xl mx-2 p-5`}
    >
      <h1 className="text-2xl font-bold">Requests</h1>
      <div className="flex mt-4">
        <NavLink className={handleActiveClass} to={"received"}>Received</NavLink>
        <NavLink className={handleActiveClass} to={"sent"}>Sent</NavLink>
      </div>
      <Outlet/>
    </div>
  );
}

export default Request;
