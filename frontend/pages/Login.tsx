import { Form } from "react-bootstrap";
import Modal from 'react-bootstrap/Modal';
import logo from './assets/logo.png';
import Button from 'react-bootstrap/Button';

export default function LoginForm({ handleSubmit, handleClose, show }) {
    return (
      <Modal show={show} onHide={handleClose}
      aria-labelledby="contained-modal-title-vcenter"
      centered>
        <Modal.Header closeButton>	
          <div className="avatar">
              <img src={logo.src} alt="logo SiMiddleman+"/>
          </div>
          <Modal.Title className="ms-auto">Login</Modal.Title>
        </Modal.Header>

        <Modal.Body>            
          <Form onSubmit={handleSubmit} id="loginForm">
            <Form.Group className="mb-3" controlId='userEmail'>
              <Form.Control
              type="email"
              placeholder="Email"
              name="email"
              required
              autoFocus/>
            </Form.Group>
            <Form.Group className="mb-3" controlId='userPassword'>
              <Form.Control
              type="password"
              placeholder="Password"
              name="password"
              required
              minLength={8}/>
            </Form.Group>
          </Form>

          <div className='d-flex justify-content-between'>
              <a>Lupa Password?</a>
              <Button type='submit' variant='merah' form='loginForm'>Masuk</Button>
          </div>

          <p className='or'>OR</p>

          <Button variant='merah' onClick={handleClose} className='w-100'>Daftar Akun</Button>
        </Modal.Body>
      </Modal>
    )
}