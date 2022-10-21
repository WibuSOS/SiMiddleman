import { Form, Modal, Button } from 'react-bootstrap';
import logo from './assets/logo.png';
import Swal from 'sweetalert2';
import useTranslation from 'next-translate/useTranslation';
import { useRouter } from 'next/router';

export default function ModalJoinRoom({ idPembeli, sessionToken, closeJoinRoomModal, joinRoomModal, GetAllRoom }) {
    const { t, lang } = useTranslation('joinRoom');
    const router = useRouter();

    const handleSubmitJoinRoom = async (e) => {
        closeJoinRoomModal();
        e.preventDefault();

        const formData = new FormData(e.currentTarget);
        const body = {
            id: idPembeli,
            roomcode: formData.get("kodeRuangan"),
        }
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/rooms/join/${body.roomcode}/${body.id}`, {
                method: 'PUT',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + sessionToken,
                }
            });
            const data = await res.json();
            console.log(data);

            if (data?.message === "Success join seller room" || data?.message === "Berhasil masuk ruangan penjual") {
                Swal.fire({ icon: 'success', title: t("successJoin"), text: data.message, showConfirmButton: false, timer: 1500, })
                GetAllRoom();
            } else {
                Swal.fire({ icon: 'error', title: t("failJoin"), text: data.message, })
            }
        }
        catch (error) {
            console.log(error);
        }
    }
    return (
        <Modal show={joinRoomModal} onHide={closeJoinRoomModal}
            aria-labelledby="contained-modal-title-vcenter"
            data-testid="ModalJoinRoom"
            centered>
            <Modal.Header data-testid="ModalHeader" closeButton>
                <div className="avatar" data-testid="avatar">
                    <img src={logo.src} alt="logo SiMiddleman+" data-testid="logo" />
                </div>
                <Modal.Title className="ms-auto mt-4" data-testid="title">{t("modalTitle")}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={handleSubmitJoinRoom} id="joinRoomForm">
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="text"
                            placeholder={t("placeholder")}
                            data-testid="kodeRuangan"
                            name="kodeRuangan"
                            autoFocus
                        />
                    </Form.Group>
                    <Button variant='merah'
                        className='w-100'
                        type='submit'
                        form='joinRoomForm'
                        data-testid="buttonJoinRoom">{t("joinRoomBtnModal")}
                    </Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}