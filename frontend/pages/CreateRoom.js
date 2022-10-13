import { useState } from "react";
import { Button } from "react-bootstrap";
import ModalCreateRoom from "./ModalCreateRoom";

export default function CreateRoom(props) {
  const [createRoomModal, setCreateRoomModal] = useState(false);
  const openCreateRoomModal = () => setCreateRoomModal(true)
  const closeCreateRoomModal = () => setCreateRoomModal(false)
  return (
    <>
      <Button onClick={() => openCreateRoomModal()} data-testid="createRoomButton">Create Room</Button>
      <ModalCreateRoom idPenjual={props.idPenjual} sessionToken={props.sessionToken} closeCreateRoomModal={closeCreateRoomModal} createRoomModal={createRoomModal} />
    </>
  )
}