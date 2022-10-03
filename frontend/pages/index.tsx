import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
import logo from './assets/logo.png'

function Example() {
  const [show, setShow] = useState(false);
  const [error, setError] = useState(null)

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);
  const handlLogin = async (e) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)
    const body = {
      email: formData.get("email"),
      password: formData.get("password")
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/login`, {
        method: 'POST',
        body: JSON.stringify(body)
      })
      const data = await res.json()
      console.log(data);

    } catch (error) {
      setError(error)
    }
  }

  console.log(logo.src)

  return (
    <div id='login'>
      <Button variant="primary" onClick={handleShow}>
        Login
      </Button>

      {/* Form Login */}
      <Modal show={show} onHide={handleClose}
        aria-labelledby="contained-modal-title-vcenter"
        centered>
        <Modal.Header closeButton>
          <div className="avatar">
            <img src={logo.src} alt="logo SiMiddleman+" />
          </div>
          <Modal.Title className="ms-auto">Login</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handlLogin}>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Control
                type="email"
                placeholder="Email"
                name="email"
                autoFocus
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput2">
              <Form.Control
                type="password"
                placeholder="Password"
                name="password"
                autoFocus
              />
            </Form.Group>
            <div className='d-flex justify-content-between'>
              <a>Lupa Password?</a>
              <Button variant='merah' type='submit'>Masuk</Button>
            </div>
          </Form>
          <p className='or'>OR</p>
          <Button variant='merah' onClick={handleClose} className='w-100'>Daftar Akun</Button>
        </Modal.Body>
      </Modal>
    </div>
  );
}

export default Example;