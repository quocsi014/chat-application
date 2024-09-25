import { useEffect, useState } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import { getUserById } from "../api/userAPI";
import { getCookie } from "../utils/cookie";
import { MdArrowForwardIos, MdArrowBackIos } from "react-icons/md";
import Chat from "../pages/conversation/Chat";
import SearchUsersModal from "./SearchUsersModal";
import { BiSearch } from "react-icons/bi";
import ConversationList from "../pages/conversation/ConversationList";
import Navigator from "./Navigator";
import { useSelector } from "react-redux";
import ChatBox from "./ChatBox";

function AppLayout() {
  const navigate = useNavigate();
  
  const isSearchModalOpen = useSelector(
    (state) => state.searchUser.isSearchUserOpen
  );

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


  return (
    <div className="w-screen h-screen bg-bg py-4 px-2 flex">
      <Navigator />
      <Outlet />
      <ChatBox />
      <SearchUsersModal />
    </div>
  );
}

export default AppLayout;
