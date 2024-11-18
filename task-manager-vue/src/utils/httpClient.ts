import axios, { AxiosResponse, AxiosError } from "axios";
const { VITE_BASE_URL } = import.meta.env;
import { getCookie } from "./jsCookie";
import { AUTH } from "@/Constants";

const axiosInstance = axios.create({
  baseURL: VITE_BASE_URL,
  headers: {
    "Content-type": "application/json",
  },
});

// Add a request interceptor
axiosInstance.interceptors.request.use(
  (config: any) => {
    const userSignIn: any = getCookie(AUTH.userSignIn);
    const { token } = JSON.parse(userSignIn);

    if (!userSignIn || !token) {
      window.location.href = "/auth/signIn";
      return false;
    }

    if (!config.headers) config.headers = {};
    config.headers.Authorization = `Bearer ${token}`;

    return config;
  },
  (error: AxiosError) => {
    return Promise.reject(error);
  }
);

// Add a response interceptor
axiosInstance.interceptors.response.use(
  (response: AxiosResponse) => {
    return response;
  },
  (error: AxiosError) => {
    // Handle error responses
    return Promise.reject(error);
  }
);

export default axiosInstance;
