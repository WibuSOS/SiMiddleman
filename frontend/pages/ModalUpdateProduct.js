import { Form, Modal, Button } from 'react-bootstrap';
import logo from './assets/logo.png';

export default function ModalUpdateProduct({ closeUpdateProductModal, updateProductModal, handleSubmitUpdateProduct, namaProduk, hargaProduk, kuantitasProduk, deskripsiProduk, onChangeText }) {
    return (
        <Modal show={updateProductModal} onHide={closeUpdateProductModal}
            aria-labelledby="contained-modal-title-vcenter"
            data-testid="modalRegister"
            centered>
            <Modal.Header closeButton data-testid="modalHeader">
                <div className="avatar">
                    <img src={logo.src} alt="logo SiMiddleman+" data-testid="logo" />
                </div>
                <Modal.Title className="ms-auto" data-testid="title">Update Product</Modal.Title>
            </Modal.Header>

            <Modal.Body data-testid="modalBody">
                <Form onSubmit={handleSubmitUpdateProduct} id="createRoomForm">
                    <Form.Group className="mb-3">
                        <Form.Control
                        type="text"
                        placeholder="Nama Produk"
                        data-testid="namaProduk"
                        name="namaProduk"
                        value={namaProduk}
                        onChange={(e) => onChangeText(e , "namaProduk")}
                        autoFocus
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                        type="number"
                        placeholder="Harga Produk"
                        data-testid="hargaProduk"
                        name="hargaProduk"
                        value={hargaProduk}
                        onChange={(e) => onChangeText(e , "hargaProduk")}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                        type="number"
                        placeholder="Kuantitas Produk"
                        data-testid="kuantitasProduk"
                        name="kuantitasProduk"
                        value={kuantitasProduk}
                        onChange={(e) => onChangeText(e , "kuantitasProduk")}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                        placeholder="Deskripsi Produk"
                        as="textarea"
                        rows={5}
                        data-testid="deskripsiProduk"
                        name="deskripsiProduk"
                        value={deskripsiProduk}
                        onChange={(e) => onChangeText(e , "deskripsiProduk")}
                        />
                    </Form.Group>
                    <Button variant='merah'
                        className='w-100'
                        type='submit'
                        form='createRoomForm'
                        data-testid="buttonCreateRoom">Update Product</Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}