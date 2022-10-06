import { Form, Modal, Button } from 'react-bootstrap';
import logo from './assets/logo.png';
import Swal from 'sweetalert2';
import Home from '.';
import { router } from 'next/router';

export default function ModalCreateRoom({idPenjual, sessionToken, closeCreateRoomModal, createRoomModal }) {
  const handleSubmitCreateRoom = async (e) => {
    e.preventDefault();
  
    const formData = new FormData(e.currentTarget);
    const body = {
      id: idPenjual,
      product: {
        nama: formData.get("namaProduk"),
        deskripsi: formData.get("deskripsiProduk"),
        harga: parseInt(formData.get("hargaProduk")),
        kuantitas: parseInt(formData.get("kuantitasProduk")),
      }
    }
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/rooms`, {
        method: 'POST',
        body: JSON.stringify(body),
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ' + sessionToken,
        }
      });
      const data = await res.json();

      if (data.message === "success"){
        Swal.fire({
          icon: 'success',
          title: 'Room berhasil dibuat',
          showConfirmButton: false,
          timer: 1500,
        })
      }else {
        Swal.fire({
          icon: 'error',
          title: 'Buat Room gagal',
          text: 'Terjadi kesalahan pada input anda',
        })
      }
    }
    catch(error) {
      console.log(error);
    }
  }
    return (
        <Modal show={createRoomModal} onHide={closeCreateRoomModal}
        aria-labelledby="contained-modal-title-vcenter"
        data-testid="ModalCreateRoom"
        centered>
            <Modal.Header data-testid="ModalHeader" closeButton>	
            <div className="avatar" data-testid="avatar">
                <img src={logo.src} alt="logo SiMiddleman+" data-testid="logo" />
            </div>
            <Modal.Title className="ms-auto" data-testid="title">Create Room</Modal.Title>
            </Modal.Header>
            <Modal.Body>
            <Form onSubmit={handleSubmitCreateRoom} id="createRoomForm">
                <Form.Group className="mb-3">
                <Form.Control
                    type="text"
                    placeholder="Nama Produk"
                    data-testid="namaProduk"
                    name="namaProduk"
                    autoFocus
                />
                </Form.Group>
                <Form.Group className="mb-3">
                <Form.Control
                    type="number"
                    placeholder="Harga Produk"
                    data-testid="hargaProduk"
                    name="hargaProduk"
                    autoFocus
                />
                </Form.Group>
                <Form.Group className="mb-3">
                <Form.Control
                    type="number"
                    placeholder="Kuantitas Produk"
                    data-testid="kuantitasProduk"
                    name="kuantitasProduk"
                    autoFocus
                />
                </Form.Group>
                <Form.Group className="mb-3">
                <Form.Control
                    placeholder="Deskripsi Produk"
                    as="textarea" 
                    rows={5}
                    data-testid="deskripsiProduk"
                    name="deskripsiProduk"
                    autoFocus
                />
                </Form.Group>
                <Button variant='merah'
                    className='w-100'
                    type='submit'
                    form='createRoomForm'
                    data-testid="buttonCreateRoom">Buat Room</Button>
            </Form>
            </Modal.Body>
        </Modal>
    )
}