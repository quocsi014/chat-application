import { createSlice } from "@reduxjs/toolkit";

const me = createSlice({
  name: 'me',
  initialState: {
    me: null
  },
  reducers:{
    setMe: (state, action)=>{
      state.me = action.payload.me
    },
  }
})

export default me.reducer

export const {setMe} = me.actions