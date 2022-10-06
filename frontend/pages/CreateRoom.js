import { useState } from "react";
import { Button, Modal, Form } from "react-bootstrap";
import logo from './assets/logo.png'

export default function CreateRoom () {
  const [createRoomModal, setCreateRoomModal] = useState(false);
  const openCreateRoomModal = () => setCreateRoomModal(true)
  const closeCreateRoomModal = () => setCreateRoomModal(false)
  return (
    <>
      <Button onClick={() => openCreateRoomModal()}>Create Room</Button>
      <Modal show={createRoomModal} onHide={closeCreateRoomModal}
      aria-labelledby="contained-modal-title-vcenter"
      centered>
        <Modal.Header closeButton>	
          <div className="avatar">
            <img src={logo.src} alt="logo SiMiddleman+"/>
          </div>
          <Modal.Title className="ms-auto">Create Room</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Control
                type="text"
                placeholder="Nama Produk"
                autoFocus
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Control
                type="number"
                placeholder="Harga Awal Produk"
                autoFocus
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Control
                type="number"
                placeholder="Kuantitas Produk"
                autoFocus
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Control
                placeholder="Deskripsi Produk"
                as="textarea" 
                rows={5}
                autoFocus
              />
            </Form.Group>
          </Form>
          <Button variant='merah' onClick={closeCreateRoomModal} className='w-100'>Buat Room</Button>
        </Modal.Body>
      </Modal>
    </>
  )
}