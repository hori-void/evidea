// src/api/userApi.ts
import apiClient from "../apiClient";
import { LoginRequest,LoginResponse} from "../../types/api";

export const loginAuth = async (data: LoginRequest): Promise<LoginResponse> => {
  const res = await apiClient.post<LoginResponse>("users/login", data);
  return res.data;
};
