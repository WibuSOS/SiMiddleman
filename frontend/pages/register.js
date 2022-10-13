import { Button } from 'react-bootstrap';
import { useState } from 'react';
import { useRouter } from 'next/router';
import Swal from 'sweetalert2';
import { signIn } from "next-auth/react";
import RegisterModal from './registerModal'

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

      if (data.message === "success") {
        Swal.fire({ icon: 'success', title: 'Register berhasil, silahkan login', confirmButtonText: 'Login', }).then((result) => {
          if (result.isConfirmed) {
            signIn()
          }
        })
      } else if (data.message === "no hp tidak memenuhi syarat") {
        Swal.fire({ icon: 'error', title: 'Register gagal', text: 'No Hp tidak memenuhi syarat', })
      } else {
        Swal.fire({ icon: 'error', title: 'Register gagal', text: 'Terjadi kesalahan pada input anda', })
      }
    }
    catch (error) {
      console.log(error);
    }
  }

  return (
    <>
      <Button variant='link' onClick={openRegisterModal} data-testid="buttonRegisterForm">
        Daftar Sekarang
      </Button>

      <RegisterModal handleSubmitRegister={handleSubmitRegister}
        closeRegisterModal={closeRegisterModal}
        registerModal={registerModal} />
    </>
  )
}