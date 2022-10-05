import { Form, Modal, Button } from 'react-bootstrap';
import { useState } from 'react';
import { useRouter } from 'next/router';
import logo from './assets/logo.png';
import Swal from 'sweetalert2';
import { signIn } from "next-auth/react";

export default function RegisterForm() {
  const [registerModal, setRegisterModal] = useState(false);
  const openRegisterModal = () => setRegisterModal(true);
  const closeRegisterModal = () => setRegisterModal(false);
  const router = useRouter();

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
      
      if (data.message === "success"){
        Swal.fire({
          icon: 'success',
          title: 'Register berhasil, silahkan login',
          confirmButtonText: 'Login',
        }).then((result) => {
          if (result.isConfirmed) {
            signIn()
          }
        })
      }else {
        Swal.fire({
          icon: 'error',
          title: 'Register gagal',
          text: 'Terjadi kesalahan pada input anda',
        })
      }
    }
    catch (error) {
        console.log(error);
    }
  }

  return (
    <>
      <Button variant="primary" onClick={openRegisterModal}>
        Register
      </Button>
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
          <Form onSubmit={handleSubmitRegister}>
            <Form.Group className="mb-3">
              <Form.Control
                  type="text"
                  placeholder="Nama"
                  name='nama'
                  autoFocus
                  required
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
                  type="text"
                  placeholder="No HP"
                  name='noHp'
                  required
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
                  type="text"
                  placeholder="No Rekening"
                  name='noRek'
                  required
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
                  type="email"
                  placeholder="Email"
                  name='email'
                  required
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
                  type="password"
                  placeholder="Password"
                  name='password'
                  required
                  minLength={8}
              />
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
                  type="password"
                  placeholder="Confirm Password"
                  name='confirmPassword'
                  required
                  minLength={8}
              />
            </Form.Group>
            <Button variant='merah' type='submit' className='w-100'>Register</Button>
          </Form>
        </Modal.Body>
      </Modal> 
    </>
  )
}