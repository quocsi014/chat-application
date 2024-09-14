import axios from "axios";
const API_URL = import.meta.env.REACT_APP_API_URL;

export const getRequestSent = (token) => {
  let config = {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
  return axios.get(`${API_URL}/conversations/requests/sent`, config);
};
