import axios from "axios";
import { tokenConfig } from "./utils";


const API_URL = import.meta.env.REACT_APP_API_URL;

export const createUser = (
  firstname,
  lastname,
  username,
  avatar_url,
  token
) => {
  let data = {
    firstname: firstname,
    lastname: lastname,
    username: username,
    avatar_url,
  };
  let config = {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
  return axios.post(`${API_URL}/users`, data, config);
};

export const getUserById = (token) => {
  let config = {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
  return axios.get(`${API_URL}/users/profile/me`, config);
};

export const searchUsers = (searchTerm, page, limit) => {
  const config = tokenConfig()
  return axios.get(
    `${API_URL}/users/profile?username=${searchTerm}&page=${page}&limit=${limit}`,
    config
  );
};
