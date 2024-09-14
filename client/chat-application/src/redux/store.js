import { configureStore } from "@reduxjs/toolkit";
import searchUser from "./SearchUser/searchUserSlice";

export const store = configureStore({
  reducer: {searchUser},
})