import { Link } from "react-router-dom";
import logo from "../../assets/croped-logo.png";
function LandingPage() {
  return (
    <div className="flex justify-center items-center h-screen w-screen bg-bg">
      <div className="flex h-full flex-row justify-center items-center pb-20">
        <div className="">
          <img src={logo} alt="" />
        </div>
        <div className="max-w-[580px] flex flex-col p-10">
          <h3 className="text-8xl font-bold">Chat</h3>
          <p className="text-gray-700 font-semibold">
            Connect and communicate effortlessly with our chat application,
            designed for seamless conversations, group chats, and file
            sharing—bringing people closer, one message at a time.
          </p>
          <div className="py-4">
            <Link to={'register'} className="text-2xl bg-gray-400 px-4 py-2 font-bold mr-4 text-white rounded-lg">Register</Link>
            <Link to={'login'} className="text-2xl bg-blue-500 px-4 py-2 font-bold mr-4 text-white rounded-lg">Login</Link>
          </div>
        </div>
      </div>
    </div>
  );
}

export default LandingPage;
