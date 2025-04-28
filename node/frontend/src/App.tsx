import { useState } from 'react'
import './css/App.css'
// import RoomNavi from './etc/RoomNavi.tsx'
// import StickyRoom from './etc/StickyRoom.tsx'
// import { loginAuth } from "./utils/api/usersApi.ts";
import {LoginInfo} from "./types/state.ts";
// import { LoginModal } from './etc/LoginModal.tsx';
// import { useModal } from './hooks/useModal.ts';

import Header from "./components/Header/Header.tsx"
import SideWindow from "./components/SideWindow/SideWindow.tsx"
import StickyRoom from "./components/StickyRoom/StickyRoom.tsx"
import LoginModal from './components/Modal/LoginModal.tsx'

const App = () => {

  const [loginState,setLoginState] = useState<boolean>(false);
  const [loginInfo, setLoginInfo] = useState<LoginInfo>(
    {
      id: "",
      name: "",
      bio:"",
    }
  );

  return (
    <>
      {/* モーダル画面 */}
      <LoginModal setLoginInfo={setLoginInfo}/>

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
