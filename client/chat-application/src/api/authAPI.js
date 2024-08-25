import axios from "axios"

const API_URL = import.meta.env.REACT_APP_API_URL

export const API_Login = (account, password)=>{
  let data = {
    "account": account,
    "password": password
  }
  return axios.post(`${API_URL}/auth/login`, data)
}