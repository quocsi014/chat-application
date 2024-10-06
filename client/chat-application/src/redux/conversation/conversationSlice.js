import { createSlice } from "@reduxjs/toolkit";

const currentChat = createSlice({
  name: 'currentChat',
  initialState:{
    chat: null
  },
  reducers:{
    setCurrentChat: (state, action)=>{
      state.chat = action.payload.chat
    }
  }
})

export default currentChat.reducer

export const {setCurrentChat} = currentChat.actions