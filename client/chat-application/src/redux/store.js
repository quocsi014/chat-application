import { configureStore } from "@reduxjs/toolkit";
import searchUser from "./SearchUser/searchUserSlice";
import currentChat from "./conversation/conversationSlice";
import me from "./me/meSlice"
export const store = configureStore({
  reducer: { searchUser, currentChat, me },
});
