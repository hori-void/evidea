import { LoginInfo } from "../../types/state.ts";
import "../../css/Header.css"

type Props = {
    loginInfo: LoginInfo;
};

const Header = ({ loginInfo }: Props) => {
    return (
        <div className="header">
            <div className="title-button">Evidea</div>
            {/* <div className="login-open-button" >Login</div> */}
            <div className="login-open-button" >{loginInfo.name}</div>            
        </div>
    ) 
}
export default Header