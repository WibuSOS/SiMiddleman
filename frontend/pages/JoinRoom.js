import { useState } from "react";
import { Button } from "react-bootstrap";
import ModalJoinRoom from "./ModalJoinRoom";
import useTranslation from 'next-translate/useTranslation';

export default function JoinRoom(props) {
    const [joinRoomModal, setJoinRoomModal] = useState(false);
    const openJoinRoomModal = () => setJoinRoomModal(true)
    const closeJoinRoomModal = () => setJoinRoomModal(false)
    const { t } = useTranslation('joinRoom');
    return (
        <>
            <Button onClick={() => openJoinRoomModal()} data-testid="joinRoomButton" className='w-100 btn-simiddleman'>{t("joinRoomBtnCard")}</Button>
            <ModalJoinRoom idPembeli={props.idPembeli} sessionToken={props.sessionToken} closeJoinRoomModal={closeJoinRoomModal} joinRoomModal={joinRoomModal} />
        </>
    )
}