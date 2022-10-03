import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
import logo from './assets/logo.png'

function Example() {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  console.log(logo.src)

  return (
    <div id='login'>
      <Button variant="primary" onClick={handleShow}>
        Login
      </Button>
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
          <Form>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Control
                type="email"
                placeholder="Email"
                autoFocus
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput2">
              <Form.Control
                type="password"
                placeholder="Password"
                autoFocus
              />
            </Form.Group>
          </Form>
          <div className='d-flex justify-content-between'>
            <a href="javascript:alert('not yet implemented')">Lupa Password?</a>
            <Button variant='merah' onClick={handleClose}>Masuk</Button>
          </div>
          <p className='or'>OR</p>
          <Button variant='merah' onClick={handleClose} className='w-100'>Daftar Akun</Button>
        </Modal.Body>
      </Modal>
    </div>
  );
}

export default Example;