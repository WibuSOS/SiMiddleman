import { Form, Modal, Button } from 'react-bootstrap';
import logo from './assets/logo.png';

export default function ModalCreateRoom({ closeCreateRoomModal, createRoomModal }) {
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
            <Form>
                <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
                <Form.Control
                    type="text"
                    placeholder="Nama Produk"
                    data-testid="namaProduk"
                    autoFocus
                />
                </Form.Group>
                <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
                <Form.Control
                    type="number"
                    placeholder="Harga Awal Produk"
                    data-testid="hargaProduk"
                    autoFocus
                />
                </Form.Group>
                <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
                <Form.Control
                    type="number"
                    placeholder="Kuantitas Produk"
                    data-testid="kuantitasProduk"
                    autoFocus
                />
                </Form.Group>
                <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
                <Form.Control
                    placeholder="Deskripsi Produk"
                    as="textarea" 
                    rows={5}
                    data-testid="deskripsiProduk"
                    autoFocus
                />
                </Form.Group>
            </Form>
            <Button variant='merah'
                onClick={closeCreateRoomModal}
                className='w-100'
                data-testid="buttonCreateRoom">Buat Room</Button>
            </Modal.Body>
        </Modal>
    )
}