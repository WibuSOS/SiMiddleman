import { useState } from "react";
import { Button } from "react-bootstrap";
import ModalShowRoomCode from "./ModalShowRoomCode";

const ShowRoomCode = ( props ) => {
  const [ShowRoomCodeModal, setShowRoomCodeModal] = useState(false);
    const openShowRoomCodeModal = () => setShowRoomCodeModal(true)
    const closeShowRoomCodeModal = () => setShowRoomCodeModal(false)
    return (
      <>
        <Button variant='simiddleman' onClick={() => openShowRoomCodeModal()}data-testid="buttonShowRoomCode">Tunjukkan Kode Ruangan</Button>
        <ModalShowRoomCode roomCode={props.roomCode} closeShowRoomCodeModal={closeShowRoomCodeModal} showRoomCodeModal={ShowRoomCodeModal} />
      </>
    )
}

export default ShowRoomCode;