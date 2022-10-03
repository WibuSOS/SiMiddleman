import Head from 'next/head';
import { useState, useEffect } from 'react';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import logo from './assets/logo.png'

export default function Home() {
  const [registerModal, setRegisterModal] = useState(false);
  useEffect(() => {
    getData()
  }, []);

  const getData = async () => {
    
  }

  const openRegisterModal = () => setRegisterModal(true)

  const closeRegisterModal = () => setRegisterModal(false)

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      nama: formData.get("nama"),
      noHp: formData.get("noHp"),
      NoRek: formData.get("NoRek"),
      email: formData.get("email"),
      password: formData.get("password"),
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/register`, {
        method: 'POST',
        body: JSON.stringify(body)
      });
      const data = await res.json();

      // alihkan kehalaman home sekaligus login
      console.log(data)
    }
    catch (error) {
      console.log(error);
    }
  }

  return (
    <div className='container mx-10 my-7'>
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
                <Form onSubmit={handleSubmit}>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
                        <Form.Control
                            type="text"
                            placeholder="Nama"
                            name='nama'
                            autoFocus
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput2">
                        <Form.Control
                            type="text"
                            placeholder="No HP"
                            name='noHp'
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput3">
                        <Form.Control
                            type="text"
                            placeholder="No Rekening"
                            name='NoRek'
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput4">
                        <Form.Control
                            type="email"
                            placeholder="Email"
                            name='email'
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput5">
                        <Form.Control
                            type="password"
                            placeholder="Password"
                            name='password'
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="exampleForm.ControlInput6">
                        <Form.Control
                            type="password"
                            placeholder="Confirm Password"
                            name='confirmPassword'
                        />
                    </Form.Group>
                    <Button variant='merah' type='submit'>Submit</Button>
                </Form>
            </Modal.Body>
        </Modal>
    </div>
  )
}