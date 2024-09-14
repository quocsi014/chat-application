import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  isSearchUserOpen: false,
}

export const searchUser = createSlice({
  name: 'searchUser',
  initialState,
  reducers: {
    toggleSearchUser: (state) => {
      state.isSearchUserOpen = !state.isSearchUserOpen
    },
  },
})

export const {toggleSearchUser} = searchUser.actions

export default searchUser.reducer