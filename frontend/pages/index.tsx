import Head from 'next/head';
import { useState, useEffect } from 'react';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import logo from './assets/logo.png'
import LoginForm from './Login';
import RegisterForm from './register';
import { useRouter } from 'next/router';

function Home() {
  const router = useRouter();
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const [show, setShow] = useState(false);
  const handleShow = () => setShow(true)
  const handleClose = () => setShow(false)
  const [registerModal, setRegisterModal] = useState(false);
  const openRegisterModal = () => setRegisterModal(true)
  const closeRegisterModal = () => setRegisterModal(false)

  const getData = async () => {

  }

  const handleSubmitLogin = async (e) => {
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

  const handleSubmitRegister = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);

    const body = {
      nama: formData.get("nama"),
      noHp: formData.get("noHp"),
      noRek: formData.get("noRek"),
      email: formData.get("email"),
      password: formData.get("password"),
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/register`, {
        method: 'POST',
        body: JSON.stringify(body)
      });
      const data = await res.json();

      console.log(data)
      if (data.message === "success") {
        router.push('/')
      } else {
        alert("Register Gagal!")
      }
    }
    catch (error) {
      console.log(error);
    }
  }

  return (
    <div className='container mx-10 my-7'>
      <Button variant="primary" onClick={handleShow}>
        Login
      </Button>

      <Button variant="primary" onClick={openRegisterModal}>
        Register
      </Button>

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
          <LoginForm handleSubmit={handleSubmitLogin} />

          <div className='d-flex justify-content-between'>
            <a>Lupa Password?</a>
            <Button type='submit' variant='merah' form='loginForm'>Masuk</Button>
          </div>

          <p className='or'>OR</p>

          <Button variant='merah' onClick={handleClose} className='w-100'>Daftar Akun</Button>
        </Modal.Body>
      </Modal>

      <Modal show={registerModal} onHide={closeRegisterModal}
        aria-labelledby="contained-modal-title-vcenter"
        centered>

        <Modal.Header closeButton>
          <div className="avatar">
            <img src={logo.src} alt="logo SiMiddleman+" />
          </div>
          <Modal.Title className="ms-auto">Register</Modal.Title>
        </Modal.Header>

        <Modal.Body>
          <RegisterForm handleSubmit={handleSubmitRegister} />
        </Modal.Body>
      </Modal>
    </div>
  );
}

export default Home;
