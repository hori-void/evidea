// import { useState } from "react";
// import { LoginInfo } from "../types/state";

// export const useLogin = () => {
//       // ログイン状態
//       const [loginState,setLoginState] = useState<boolean>(false);
//       // ログイン情報
//       const [loginInfo,setLoginInfo] = useState<LoginInfo>({
//           id: "",
//           name: "",
//         }
//       )

//     // トークン保存 (ローカルストレージ)
//     localStorage.setItem("token",token);

//     // ログイン関係stateの更新
//     setLoginInfo({"id":inputUserId,"name":userName})
//     setLoginState(true);

//     // 入力値リセット
//     setInputUserId("");
//     setInputPassword("");

//     // ログインフォーム非表示
//     setAddLoginForm(false);
// }