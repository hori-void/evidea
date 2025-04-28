import '../../css/SideWindow.css'
import StickyGroup from './StickyGroup'

const SideWindow = () => {
    return (
        <div className="side-window">
            <div className="org-name">
                株式会社XXXXX システム開発部
            </div>

            <div className="stickygroup-block">
                <div className="stickygroup-name">
                    募集中
                </div>
                <StickyGroup />
            </div>

            <div>
                <div className="stickygroup-name">
                    保留中
                </div>
                <StickyGroup />
            </div>

        </div>
    )
}

export default SideWindow