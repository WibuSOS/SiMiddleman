import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
import logo from './assets/logo.png'

export default function RegisterForm({ handleSubmit, closeRegisterModal, registerModal }) {
    return (       
        <Modal show={registerModal} onHide={closeRegisterModal}
            aria-labelledby="contained-modal-title-vcenter"
            centered>

            <Modal.Header closeButton>	
                <div className="avatar">
                    <img src={logo.src} alt="logo SiMiddleman+"/>
                </div>
                <Modal.Title className="ms-auto">Register</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form onSubmit={handleSubmit}>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
                        <Form.Control
                            type="text"
                            placeholder="Nama"
                            name='nama'
                            autoFocus
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput2">
                        <Form.Control
                            type="text"
                            placeholder="No HP"
                            name='noHp'
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput3">
                        <Form.Control
                            type="text"
                            placeholder="No Rekening"
                            name='noRek'
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput4">
                        <Form.Control
                            type="email"
                            placeholder="Email"
                            name='email'
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput5">
                        <Form.Control
                            type="password"
                            placeholder="Password"
                            name='password'
                            required
                            minLength={8}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput6">
                        <Form.Control
                            type="password"
                            placeholder="Confirm Password"
                            name='confirmPassword'
                            required
                            minLength={8}
                        />
                    </Form.Group>
                    <Button variant='merah' type='submit'>Submit</Button>
                </Form>
            </Modal.Body>
        </Modal> 
    )
}