import axios from "axios";

const API_URL = import.meta.env.REACT_APP_API_URL;

export const API_Login = (account, password) => {
  let data = {
    account: account,
    password: password,
  };
  return axios.post(`${API_URL}/auth/login`, data);
};

export const API_Register = (email, password) => {
  let data = {
    email: email,
    password: password,
  };
  return axios.post(`${API_URL}/auth/register`, data);
};

export const API_Verify = (token) => {
  return axios.post(`${API_URL}/auth/verify`,{}, {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
}