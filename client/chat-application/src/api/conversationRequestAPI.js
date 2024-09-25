import axios from "axios";
import { tokenConfig } from "./utils";

const API_URL = import.meta.env.REACT_APP_API_URL;


export const getRequestSent = () => {
  let config = tokenConfig();
  return axios.get(`${API_URL}/conversations/requests/sent`, config);
};

export const getRequestReceived = () => {
  let config = tokenConfig();
  return axios.get(`${API_URL}/conversations/requests/received`, config);
};

export const deleteRequest = (receiverId) => {
  let config = tokenConfig();
  return axios.delete(
    `${API_URL}/conversations/requests/sent/${receiverId}`,
    config
  );
};

export const rejectRequest = (senderId) => {
  let config = tokenConfig();
  return axios.post(
    `${API_URL}/conversations/requests/received/${senderId}/reject`,
    {},
    config
  );
};

export const acceptRequest = (senderId) => {
  let config = tokenConfig();
  return axios.post(
    `${API_URL}/conversations/requests/received/${senderId}/accept`,
    {},
    config
  );
};
