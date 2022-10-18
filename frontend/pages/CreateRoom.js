import { useState } from "react";
import { Button } from "react-bootstrap";
import ModalCreateRoom from "./ModalCreateRoom";
import useTranslation from 'next-translate/useTranslation';

export default function CreateRoom(props) {
  const [createRoomModal, setCreateRoomModal] = useState(false);
  const openCreateRoomModal = () => setCreateRoomModal(true);
  const closeCreateRoomModal = () => setCreateRoomModal(false);
  const { t, lang } = useTranslation('createRoom');
  return (
    <>
      <Button onClick={() => openCreateRoomModal()} data-testid="createRoomButton" className='w-100 btn-simiddleman'>{t("createRoomBtnCard")}</Button>
      <ModalCreateRoom idPenjual={props.idPenjual} sessionToken={props.sessionToken} closeCreateRoomModal={closeCreateRoomModal} createRoomModal={createRoomModal} />
    </>
  )
}