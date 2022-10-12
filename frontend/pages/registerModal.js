import { Form, Modal, Button } from 'react-bootstrap';
import logo from './assets/logo.png';

export default function FormRegister({ handleSubmitRegister, closeRegisterModal, registerModal }) {
    return (
        <Modal show={registerModal} onHide={closeRegisterModal}
            aria-labelledby="contained-modal-title-vcenter"
            data-testid="modalRegister"
            centered>
            <Modal.Header closeButton data-testid="modalHeader">
                <div className="avatar">
                    <img src={logo.src} alt="logo SiMiddleman+" data-testid="logo" />
                </div>
                <Modal.Title className="ms-auto" data-testid="title">Register</Modal.Title>
            </Modal.Header>

            <Modal.Body data-testid="modalBody">
                <Form onSubmit={handleSubmitRegister} data-testid="form">
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="text"
                            placeholder="Nama"
                            name='nama'
                            data-testid="nama"
                            autoFocus
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="number"
                            placeholder="No HP"
                            name='noHp'
                            data-testid="noHp"
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="number"
                            placeholder="No Rekening"
                            name='noRek'
                            data-testid="noRek"
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="email"
                            placeholder="Email"
                            name='email'
                            data-testid="email"
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="password"
                            placeholder="Password"
                            name='password'
                            data-testid="password"
                            required
                            minLength={8}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                            type="password"
                            placeholder="Confirm Password"
                            name='confirmPassword'
                            data-testid="confirm"
                            required
                            minLength={8}
                        />
                    </Form.Group>
                    <Button variant='merah'
                        type='submit'
                        className='w-100'
                        data-testid="submitButton">Register</Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}