import { Form, Modal, Button } from 'react-bootstrap';
import logo from './assets/logo.png';
import useTranslation from 'next-translate/useTranslation';

export default function FormRegister({ handleSubmitRegister, closeRegisterModal, registerModal }) {
  const { t } = useTranslation('common');
    return (
        <Modal show={registerModal} onHide={closeRegisterModal}
            aria-labelledby="contained-modal-title-vcenter"
            data-testid="modalRegister"
            centered>
            <Modal.Header closeButton data-testid="modalHeader">
                <div className="avatar">
                    <img src={logo.src} alt="logo SiMiddleman+" data-testid="logo" />
                </div>
                <Modal.Title className="ms-auto" data-testid="title">{t('register-modal.btn-register')}</Modal.Title>
            </Modal.Header>

            <Modal.Body data-testid="modalBody">
                <Form onSubmit={handleSubmitRegister} data-testid="form">
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="text"
                            placeholder={t('register-modal.placeholder.name')}
                            name='nama'
                            data-testid="nama"
                            autoFocus
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="number"
                            placeholder={t('register-modal.placeholder.phone-number')}
                            name='noHp'
                            data-testid="noHp"
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="number"
                            placeholder={t('register-modal.placeholder.account-number')}
                            name='noRek'
                            data-testid="noRek"
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="email"
                            placeholder={t('register-modal.placeholder.email')}
                            name='email'
                            data-testid="email"
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="password"
                            placeholder={t('register-modal.placeholder.password')}
                            name='password'
                            data-testid="password"
                            required
                            minLength={8}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="password"
                            placeholder={t('register-modal.placeholder.confirm-password')}
                            name='confirmPassword'
                            data-testid="confirm"
                            required
                            minLength={8}
                        />
                    </Form.Group>
                    <Button variant='merah'
                        type='submit'
                        className='w-100'
                        data-testid="submitButton">{t('register-modal.btn-register')}</Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}