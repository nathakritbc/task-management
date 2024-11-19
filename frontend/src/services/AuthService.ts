import { AuthSignIn, AuthSignUp } from "@/types/Auth";
import axios, { AxiosResponse } from "axios";
import axiosInstance from "@/utils/httpClient";
const { VITE_BASE_URL } = import.meta.env;
const PATH = `${VITE_BASE_URL}/auth`;

const signIn = async (data: AuthSignIn): Promise<AxiosResponse> => {
  return await axios.post(`${PATH}/signIn`, data);
};

const signUp = async (data: AuthSignUp): Promise<AxiosResponse> => {
  return await axios.post(`${PATH}/signUp`, data);
};

const enable2FA = async (data: string): Promise<AxiosResponse> => {
  return await axios.post(`${PATH}/enable-2fa/${data}`, data);
};

const findUserSignInById = async (username: string): Promise<AxiosResponse> => {
  return await axiosInstance.get(`${VITE_BASE_URL}/users/${username}`);
};

const AuthService = {
  signIn,
  signUp,
  enable2FA,
  findUserSignInById,
};

export default AuthService;
