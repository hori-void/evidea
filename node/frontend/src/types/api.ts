export type LoginRequest = {
    inputId: string;
    inputPassword: string;
  };
  
export type LoginResponse = {
    token: string;
    userName: string;
    bio: string;
  };