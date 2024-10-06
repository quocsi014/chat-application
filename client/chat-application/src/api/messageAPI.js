import axios from "axios";
import { tokenConfig } from "./utils";
const API_URL = import.meta.env.REACT_APP_API_URL;

export const getMessages = (conversationId)=>{
  const config = tokenConfig()
  return axios.get(`${API_URL}/conversations/${conversationId}/messages`, config)
}