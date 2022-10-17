import { useState } from "react";
import { Button } from "react-bootstrap";
import ModalJoinRoom from "./ModalJoinRoom";

export default function JoinRoom(props) {
    const [joinRoomModal, setJoinRoomModal] = useState(false);
    const openJoinRoomModal = () => setJoinRoomModal(true)
    const closeJoinRoomModal = () => setJoinRoomModal(false)
    return (
        <>
            <Button onClick={() => openJoinRoomModal()} data-testid="joinRoomButton" className='w-100 btn-simiddleman'>Join Room</Button>
            <ModalJoinRoom idPembeli={props.idPembeli} sessionToken={props.sessionToken} closeJoinRoomModal={closeJoinRoomModal} joinRoomModal={joinRoomModal} />
        </>
    )
}