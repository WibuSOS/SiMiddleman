import { useState } from "react";
import { Button } from "react-bootstrap";
import ModalShowRoomCode from "./ModalShowRoomCode";
import useTranslation from 'next-translate/useTranslation';

const ShowRoomCode = ( props ) => {
  const [ShowRoomCodeModal, setShowRoomCodeModal] = useState(false);
    const openShowRoomCodeModal = () => setShowRoomCodeModal(true)
    const closeShowRoomCodeModal = () => setShowRoomCodeModal(false)
    const { t, lang } = useTranslation('detailProduct');
    return (
      <>
        <Button variant='simiddleman' onClick={() => openShowRoomCodeModal()}data-testid="buttonShowRoomCode">{t("showRoomCodeButton")}</Button>
        <ModalShowRoomCode roomCode={props.roomCode} closeShowRoomCodeModal={closeShowRoomCodeModal} showRoomCodeModal={ShowRoomCodeModal} />
      </>
    )
}

export default ShowRoomCode;