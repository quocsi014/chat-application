import { useEffect, useState } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import { getUserById } from "./api/userAPI";
import { getCookie } from "./utils/cookie";
import { MdArrowForwardIos, MdArrowBackIos } from "react-icons/md";
import Chat from "./components/Chat";
import SearchUsersModal from "./components/SearchUsersModal";
import { BiSearch } from "react-icons/bi";
import ConversationList from "./layouts/ConversationList";
import Navigator from "./layouts/Navigator";
import { useSelector } from "react-redux";

function App() {
  const navigate = useNavigate();
  const [isChatInfoOpen, setIsChatInfoOpen] = useState(true);
  const isSearchModalOpen = useSelector(state => state.searchUser.isSearchUserOpen)

  const sectionPadding = "p-5";

  // Tạo fake data cho chats
  
  useEffect(() => {
    const token = getCookie("access_token");
    if (token) {
      getUserById(token)
        .then(() => {
          // Nếu thành công, không làm gì cả
        })
        .catch((error) => {
          if (error.response && error.response.status === 401) {
            navigate("/login");
          } else if (error.response && error.response.status === 404) {
            navigate("/onboarding");
          }
        });
    } else {
      navigate("/login");
    }
  }, [navigate]);

  const handleCloseChatInfo = () => {
    setIsChatInfoOpen(!isChatInfoOpen);
  };

  return (
    <div className="w-screen h-screen bg-bg py-4 px-2 flex">
      <Navigator/>
      <Outlet/>

      <div
        className={`h-full w-full bg-white rounded-2xl mx-2 ${sectionPadding} relative`}
      >
        <div className="absolute top-0 right-0 translate-x-1/2 h-full flex items-center opacity-0 hover:opacity-100 transition-opacity duration-300">
          <button
            className="bg-gray-200 p-2 rounded-full border hover:opacity-100 opacity-60"
            onClick={handleCloseChatInfo}
          >
            {isChatInfoOpen ? (
              <MdArrowForwardIos size={30} />
            ) : (
              <MdArrowBackIos size={30} />
            )}
          </button>
        </div>
      </div>

      <div
        className={`h-full shrink-0 ${
          isChatInfoOpen ? `w-120 ${sectionPadding}` : "w-0"
        } bg-white rounded-2xl mx-2 flex flex-col transition-all duration-300 overflow-hidden`}
      ></div>
      <SearchUsersModal
      />
    </div>
  );
}

export default App;
