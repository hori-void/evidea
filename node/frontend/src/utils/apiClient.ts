// src/lib/apiClient.ts
import axios from "axios";

const apiClient = axios.create({
  baseURL: "http://localhost:8080/api/v1",
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
});

// リクエスト/レスポンスの共通処理（例：トークン付与やログ出力など）
apiClient.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    console.error("API error:", error);
    return Promise.reject(error);
  }
);

export default apiClient;
