import '../../css/StickyRoom.css'
import Sticky from './Sticky'

const StickyRoom = () => {
    return (
        <div className="stickyroom">
            <div>期限 : 2025/04/28</div>
            <div className="stickyroom-container">
                <Sticky />
                <Sticky />
                <Sticky />
                <Sticky />
                <Sticky />
                <Sticky />
                <Sticky />
                <Sticky />
                <Sticky />
                <Sticky />
                <Sticky />
                <Sticky />
            </div>
        </div>
    )
}

export default StickyRoom