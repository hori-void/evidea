import { type Dispatch,useState } from "react"
import './css/App.css'
import './css/StickyRoom.css'
import user from "../public/user.svg"
import plus from "../public/plus.svg"
import React from "react"
import {LoginInfo} from "../types/state";
import { useModal } from "../hooks/useModal"
import { LoginModal } from "./LoginModal"

type Props  = {
  setaddStickyForm:Dispatch<React.SetStateAction<boolean>>;
  setaddLoginForm:Dispatch<React.SetStateAction<boolean>>;
  loginState:boolean;
  loginInfo:LoginInfo;
};

const StickyRoom = ({setaddStickyForm,setaddLoginForm,loginInfo,loginState}:Props) => {
  const { isOpen,open,close } = useModal();

  function addStickyTrue() {
    setaddStickyForm(true);
  };

  function addLoginTrue() {
    setaddLoginForm(true);
  };

  return (
      <div style={{height:"100vh",width:"100%"}}>
        {/* ルーム内ヘッダー */}
        <div style={{display:"flex",alignItems:"center",padding:"10px",width:"100%"}}>
          <div style={{paddingLeft: "10px"}}>
            業務改善案
          </div>
          <div style={{backgroundColor:"#84919e",color:"white",height:"40px",marginLeft:"auto",paddingRight:"25px",display:loginState ? "none":"flex"}} onClick={open}>
            Login
          </div>
          <div style={{height:"40px",marginLeft:"auto",paddingRight:"25px",alignItems:"center",display:loginState ? "flex":"none"}}>
            <img src={user} style={{height:"40px"}} />
            <p style={{fontSize:"20px"}}>{loginInfo.name}</p>
          </div>
        </div>

        {/* ルーム本体 */}
        {/* <div style={{display:"grid",gridTemplateColumns:"1fr 1fr 1fr",gridTemplateRows:"150px 150px",gap:"20px",padding:"20px"}}> */}
        <div style={{display:"grid",gridTemplateColumns:"repeat(auto-fit,minmax(273px,1fr))",gridTemplateRows:"150px 150px",gap:"clamp(25px, 4vw, 40px)",padding:"clamp(10px, 4vw, 20px)"}}>

          {/* <div style={{backgroundColor:"red"}}> */}
          <div className={"sticky-note"}>
            a
          </div>
          <div className={"sticky-note"}>
            b
          </div>
          <div className={"sticky-note"}>
            c
          </div>
          <div className={"sticky-note"}>
            d
          </div>
          <div className={"sticky-note"}>
            d
          </div>
          <div className={"sticky-note"}>
            d
          </div>
        </div>

        {/* <div style={{height:"65px",width:"65px",borderRadius:"50%",backgroundColor:"#84919e",position:"absolute",display:"flex",color:"white",right:"0",bottom:"0",marginRight:"30px",marginBottom:"30px",alignItems:"center",textAlign:"center"}}> */}
        <div onClick={addStickyTrue} className='sticky-add'>  
          <img src={plus} style={{height:"30px",marginLeft:"auto",margin:"0 auto"}}/>
        </div>
      </div>

      
        {/* ログインフォーム */}
        <LoginModal isOpen={isOpen} onClose={close} />
  )
}

export default StickyRoom