import { Form, Modal, Button } from 'react-bootstrap';
import logo from './assets/logo.png';
import useTranslation from 'next-translate/useTranslation';

export default function ModalUpdateProduct({ closeUpdateProductModal, updateProductModal, handleSubmitUpdateProduct, namaProduk, hargaProduk, kuantitasProduk, deskripsiProduk, onChangeText }) {
    const { t, lang } = useTranslation('detailProduct');

    return (
        <Modal show={updateProductModal} onHide={closeUpdateProductModal}
            aria-labelledby="contained-modal-title-vcenter"
            data-testid="modalRegister"
            centered>
            <Modal.Header closeButton data-testid="modalHeader">
                <div className="avatar">
                    <img src={logo.src} alt="logo SiMiddleman+" data-testid="logo" />
                </div>
                <Modal.Title className="ms-auto" data-testid="title">{t("updateProductModal.modalTitle")}</Modal.Title>
            </Modal.Header>

            <Modal.Body data-testid="modalBody">
                <Form onSubmit={handleSubmitUpdateProduct} id="createRoomForm">
                    <Form.Group className="mb-3">
                        <Form.Control
                        type="text"
                        placeholder={t("updateProductModal.placeholder.0")}
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
                        placeholder={t("updateProductModal.placeholder.1")}
                        data-testid="hargaProduk"
                        name="hargaProduk"
                        min={1}
                        value={hargaProduk}
                        onChange={(e) => onChangeText(e , "hargaProduk")}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                        type="number"
                        placeholder={t("updateProductModal.placeholder.2")}
                        data-testid="kuantitasProduk"
                        name="kuantitasProduk"
                        min={1}
                        value={kuantitasProduk}
                        onChange={(e) => onChangeText(e , "kuantitasProduk")}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                        placeholder={t("updateProductModal.placeholder.3")}
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
                        data-testid="buttonCreateRoom">{t("updateProductModal.updateProductBtnModal")}</Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}