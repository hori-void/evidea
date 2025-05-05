import { useState } from "react"
import { useModal } from "../../hooks/useModal.ts"
import { loginAuth } from "../../utils/api/usersApi.ts";
import { LoginInfo, OrgInfo } from "../../types/state.ts";
import '../../css/LoginModal.css'

type Props = {
    setLoginState: React.Dispatch<React.SetStateAction<boolean>>;
    setLoginInfo: React.Dispatch<React.SetStateAction<LoginInfo>>;
};

const LoginModal = ({ setLoginState,setLoginInfo }: Props) => {
    const [inputId,setInputId] = useState<string>("");
    const [inputPassword,setInputPassword] = useState<string>("");
    const { modalStatus, closeModal, openModal } = useModal(true);

    // ログインクリック
    const handleClickLogin = async () => {
        try {
            // APIリクエスト
            const {token,userName,bio} = await loginAuth({inputId,inputPassword});

            // トークン保存
            localStorage.setItem("token",token);

            setLoginInfo(
                {
                    id: inputId,
                    name: userName,
                    bio: bio,
                }
            );

            setLoginState(true);
            closeModal();
        } catch(err) {

        }
    }

    if (!modalStatus) return null;

    if (modalStatus) return (
        <div className="login-background">
            <div className="login-form">
                <input className="login-input" type="text" placeholder="ID" onChange={(e) => setInputId(e.target.value)}></input>
                <input className="login-input" type="password" placeholder="PASSWORD" onChange={(e) => setInputPassword(e.target.value)}></input>
                <div className="login-button" onClick={handleClickLogin}>LOGIN</div>
            </div>
        </div>
    )
}

export default LoginModal