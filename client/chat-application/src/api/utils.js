import { getCookie } from "../utils/cookie";

export const tokenConfig = () => {
  const token = getCookie("access_token");
  return {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
};