import React from "react";
import { loginAuth } from "./utils/api/usersApi.ts";

type Props = {
    inputUserId:string;
    inputPassword:string;
    isOpen:boolean;
    onClose : () => void;
}

export const LoginModal: React.FC<Props> = ({ inputUserId,inputPassword,isOpen,onClose }) => {
    const 
    if (!isOpen) return null;
    return(
        <>
            <div className={"form-back"} style={{display:isOpen ? "block":"none"}} onClick={onClose}>
            </div>
            <div className={"login-form"} style={{display:isOpen ? "block":"none"}}>
            <div style={{marginBottom:"20px",fontSize:"25px"}}>EVIDEA</div>
                <input className={"login-input"} type="text" placeholder="username" value={inputUserId} onChange={(e) => setInputUserId(e.target.value)}/>
                <input className={"login-input"} type="password" placeholder="password" value={inputPassword} onChange={(e) => setInputPassword(e.target.value)}/>
                <button className={"login-button"} onClick={authLogin}>login</button>
            </div>
        </>
    )
}


// export default LoginModal