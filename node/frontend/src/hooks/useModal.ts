import { useState } from "react";

export const useModal = (initialStatus:boolean) => {
    const [modalStatus,setModalStatus] = useState(initialStatus);

    // モーダルを開く
    const openModal = () => {
        setModalStatus(true);
    }

    // モーダルを閉じる
    const closeModal = () => {
        setModalStatus(false);
    }

    return {
        modalStatus,
        closeModal,
        openModal,
    };
}