import { useState } from "react";
import { Button, Modal, Form } from "react-bootstrap";
import logo from './assets/logo.png'
import ModalJoinRoom from "./ModalJoinRoom";

export default function JoinRoom(props) {
    const [joinRoomModal, setJoinRoomModal] = useState(false);
    const openJoinRoomModal = () => setJoinRoomModal(true)
    const closeJoinRoomModal = () => setJoinRoomModal(false)
    return (
        <>
            <Button onClick={() => openJoinRoomModal()} data-testid="joinRoomButton" className="mx-3">Join Room</Button>
            <ModalJoinRoom idPembeli={props.idPembeli} sessionToken={props.sessionToken} closeJoinRoomModal={closeJoinRoomModal} joinRoomModal={joinRoomModal} />
        </>
    )
}