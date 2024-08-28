import axios from "axios"

const API_URL = import.meta.env.REACT_APP_API_URL;

export const createUser = (firstname, lastname, avatar_url, token)=>{
  let data = {
    firstname: firstname,
    lastname: lastname,
    avatar_url
  }
  let config = {
    headers:{
      "Authorization": `Bearer ${token}`
    }
  }
  return axios.post(`${API_URL}/users`, data, config)
}