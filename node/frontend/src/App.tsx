import { useState, useEffect } from 'react'
import './css/App.css'
// import RoomNavi from './etc/RoomNavi.tsx'
// import StickyRoom from './etc/StickyRoom.tsx'
// import { loginAuth } from "./utils/api/usersApi.ts";
import {LoginInfo, OrgInfo} from "./types/state.ts";
// import { LoginModal } from './etc/LoginModal.tsx';
// import { useModal } from './hooks/useModal.ts';

import Header from "./components/Header/Header.tsx"
import SideWindow from "./components/SideWindow/SideWindow.tsx"
import StickyRoom from "./components/StickyRoom/StickyRoom.tsx"
import LoginModal from './components/Modal/LoginModal.tsx'

const App = () => {

  // ログイン状態
  const [loginState,setLoginState] = useState<boolean>(false);

  // ログイン情報
  const [loginInfo, setLoginInfo] = useState<LoginInfo>(
    {
      id: "",
      name: "",
      bio:"",
    }
  );

  // ユーザーが属する組織情報
  const [orgInfo, setOrgInfo] = useState<OrgInfo[]>([]);

  useEffect(() => {
    if (loginState) {
      // 組織情報取得
    }
  }, [loginState]); 
  
  // ユーザーIDをキーに組織情報を取得して、orgInfoを更新する
  const fetchAndSetOrgInfo = () =>{
    
  }

  return (
    <>
      {/* モーダル画面 */}
      <LoginModal setLoginState={setLoginState} setLoginInfo={setLoginInfo}/>

      {/* メイン画面 */}
      <div className="base">
        <Header loginInfo={loginInfo} />
        <div className="main">
          <SideWindow />
          <StickyRoom />
        </div>
      </div>
    </>
  )
}

export default App
