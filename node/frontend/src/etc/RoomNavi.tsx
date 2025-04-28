// import { useState } from 'react'
// import reactLogo from './assets/react.svg'
// import viteLogo from '/vite.svg'
import React from "react"
import "./css/App.css"
import "./css/RoomNavi.css"

function RoomNavi() {
  // const [count, setCount] = useState(0)

  return (
      <div style={{height:"100vh",padding:"5px",fontSize:"15px"}}>
        {/* 組織名 */}
        <div style={{fontWeight:"bolder",padding:"5px",margin:"3px"}}>
          株式会社XXXX システム開発部
        </div>
        
        {/* 募集中 */}
        <div className="room-cat">
          募集中
        </div>
        <div className="room-unit">
          新機能実装
        </div >
        <div className="room-unit">
          業務改善案
        </div>
        <div className="room-unit">
          忘年会店決め
        </div>

        {/* 保留中 */}
        <div className="room-cat">
          保留中
        </div>
        <div className="room-unit">
          AA社要望システム改修
        </div >
        <div className="room-unit">
          職場環境改善
        </div>
        <div className="room-unit">
          上司に言いたいこと
        </div>


        {/* 終了 */}
        <div className="room-cat">
          終了
        </div>
        <div className="room-unit">
          2022年度 新規プロジェクトアイデア
        </div >
        <div className="room-unit">
          2023年度 新規プロジェクトアイデア
        </div>
        <div className="room-unit">
          BB社不具合対応案
        </div>

      </div>
  )
}

export default RoomNavi
