// src/api/userApi.ts
import apiClient from "../apiClient";
import { LoginRequest,LoginResponse} from "../../types/api";

// ログイン認証API
export const loginAuth = async (data: LoginRequest): Promise<LoginResponse> => {
  const res = await apiClient.post<LoginResponse>("users/login", data);
  return res.data;
};

// 組織情報取得API
export const fetchOrgInfo = async (token: string) => {
  try {
    const res = await apiClient.get(`/users/organizations`, {
      params: {
        token: token,
      },
    });
    return res.data;
  } catch (error) {
    console.error("Failed to fetch user organizations:", error);
    throw error;
  }
};