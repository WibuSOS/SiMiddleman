import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
import logo from './assets/logo.png'
import LoginForm from './Login';

function Home() {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const [data, setData] = useState(null);
  const [error, setError] = useState(null);

  let hasil = ""

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      email: formData.get("email"),
      password: formData.get("password")
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/login`, {
        method: 'POST',
        body: JSON.stringify(body)
      });
      const data = await res.json();
      console.log(data);
      setData(data.message);
      if (data.message == "success") {
        alert("benar")
      }
      else {
        alert("salah")
      }
    }
    catch (error) {
      setError(error);
    }
  }

  return (
    <>
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
          <LoginForm handleSubmit={handleSubmit}/>
          <div className='d-flex justify-content-between'>
            <a>Lupa Password?</a>
            <Button type='submit' variant='merah' form='loginForm'>Masuk</Button>
          </div>
          <p className='or'>OR</p>
          <Button variant='merah' onClick={handleClose} className='w-100'>Daftar Akun</Button>
        </Modal.Body>
      </Modal>
    </>
  );
}

export default Home;