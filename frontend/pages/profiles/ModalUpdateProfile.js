import { Form, Modal, Button } from 'react-bootstrap';
import logo from '../assets/logo.png';
import useTranslation from 'next-translate/useTranslation';

export default function FormRegister( props ) {
    const { t, lang } = useTranslation('userProfile');
    return (
        <Modal show={props.updateProfileModal} onHide={props.closeUpdateProfileModal}
            aria-labelledby="contained-modal-title-vcenter"
            data-testid="modalRegister"
            centered>
            <Modal.Header closeButton data-testid="modalHeader">
                <div className="avatar">
                    <img src={logo.src} alt="logo SiMiddleman+" data-testid="logo" />
                </div>
                <Modal.Title className="ms-auto" data-testid="title">{t("modal.header")}</Modal.Title>
            </Modal.Header>

            <Modal.Body data-testid="modalBody">
                <Form onSubmit={props.handleSubmitUpdateProfile} data-testid="form">
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="text"
                            placeholder={t("modal.placeholder.0")}
                            name='nama'
                            value={props.nama}
                            onChange={(e) => props.onChangeText(e , "nama")}
                            autoFocus
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="number"
                            placeholder={t("modal.placeholder.1")}
                            name='noHp'
                            value={props.noHp}
                            onChange={(e) => props.onChangeText(e , "noHp")}
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="number"
                            placeholder={t("modal.placeholder.2")}
                            name='noRek'
                            value={props.noRek}
                            onChange={(e) => props.onChangeText(e , "noRek")}
                            required
                        />
                    </Form.Group>
                    <Button variant='merah'
                        type='submit'
                        className='w-100'>{t("modal.btnUpdateProfile")}</Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}