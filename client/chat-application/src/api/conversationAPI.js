import axios from "axios";
import { tokenConfig } from "./utils";
const API_URL = import.meta.env.REACT_APP_API_URL;

export const searchUsers = () => {
  return axios.get(
    `${API_URL}/users/profile?username=${searchTerm}&page=${page}&limit=${limit}`
  );
};

export const getConversations = () => {
  const config = tokenConfig()
  return axios.get(
    `${API_URL}/conversations`,
    config
  )
}