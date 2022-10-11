import { Form, Modal, Button } from 'react-bootstrap';
import logo from './assets/logo.png';
import Swal from 'sweetalert2';
import Home from '.';
import { router } from 'next/router';

export default function ModalJoinRoom({ idPembeli, sessionToken, closeJoinRoomModal, joinRoomModal }) {
    const handleSubmitJoinRoom = async (e) => {
        closeJoinRoomModal();
        e.preventDefault();

        const formData = new FormData(e.currentTarget);
        const body = {
            id: idPembeli,
            roomcode: formData.get("kodeRuangan"),
        }
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/joinroom/${body.roomcode}/${body.id}`, {
                method: 'PUT',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + sessionToken,
                }
            });
            const data = await res.json();

            if (data.message === "success") {
                Swal.fire({
                    icon: 'success',
                    title: 'Berhasil join room',
                    text: 'Silahkan refresh untuk melihat room',
                    showConfirmButton: false,
                    timer: 1500,
                })
            } else if (data.message === "Sudah ada pembeli pada ruangan") {
                Swal.fire({
                    icon: 'error',
                    title: 'Join Room gagal',
                    text: 'Sudah ada pembeli pada ruang tersebut',
                })
            } else {
                Swal.fire({
                    icon: 'error',
                    title: 'Join Room gagal',
                    text: 'Terjadi kesalahan pada inputan',
                })
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
                <Modal.Title className="ms-auto mt-4" data-testid="title">Join Room</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={handleSubmitJoinRoom} id="joinRoomForm">
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="text"
                            placeholder="Kode Ruangan"
                            data-testid="kodeRuangan"
                            name="kodeRuangan"
                            autoFocus
                        />
                    </Form.Group>
                    <Button variant='merah'
                        className='w-100'
                        type='submit'
                        form='joinRoomForm'
                        data-testid="buttonJoinRoom">Join Room
                    </Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}