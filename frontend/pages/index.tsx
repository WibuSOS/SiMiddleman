import Head from 'next/head';
import { useState, useEffect } from 'react';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import logo from './assets/logo.png'
import LoginForm from './Login';
import RegisterForm from './register';
import { useRouter } from 'next/router';
import { signOut, signIn, useSession } from "next-auth/react";
import CreateRoom from './CreateRoom';

function Home() {
  const [show, setShow] = useState(false);
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
      if (data.message == "success") {
        router.push('./home')
      }
      else {
        alert("salah")
      }
    }
    catch (error) {
      alert("salah")
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

  const router = useRouter();
  const { data: session } = useSession();

  console.log("session", session);

  if (session) {
    return (
      <>
        <Button onClick={() => signOut()}>Sign out</Button>
        <CreateRoom />
      </>
    )
  }
  else {
    return (
      <div className='container mx-10 my-7'>
        <Button variant="primary" onClick={() => signIn()}>
          Login
        </Button>

        <Button variant="primary" onClick={openRegisterModal}>
          Register
        </Button>

        <RegisterForm handleSubmit={handleSubmitRegister}
          closeRegisterModal={closeRegisterModal}
          registerModal={registerModal} />
      </div>
    );
  }
}

export default Home;
