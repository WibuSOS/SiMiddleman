import { Button } from 'react-bootstrap';
import { useState } from 'react';
import Swal from 'sweetalert2';
import { signIn } from "next-auth/react";
import RegisterModal from './registerModal';
import useTranslation from 'next-translate/useTranslation';;

export default function RegisterForm() {
  const [registerModal, setRegisterModal] = useState(false);
  const openRegisterModal = () => setRegisterModal(true);
  const closeRegisterModal = () => setRegisterModal(false);
  const { t } = useTranslation('common');

  const handleSubmitRegister = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = { nama: formData.get("nama"), noHp: formData.get("noHp"), noRek: formData.get("noRek"), email: formData.get("email"), password: formData.get("password"),}

    if (formData.get("confirmPassword") == formData.get("password")){
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/register`, {
          method: 'POST',
          body: JSON.stringify(body)
        });
        const data = await res.json();
  
        if (data.message === "success") 
          Swal.fire({ icon: 'success', title: 'Register berhasil, silahkan login', confirmButtonText: 'Login', }).then((result) => { result.isConfirmed ? signIn : null })
        else if (data.message === "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)")
          Swal.fire({ icon: 'error', title: 'Register gagal', text: 'Email sudah digunakan', })
        else Swal.fire({ icon: 'error', title: 'Register gagal', text: data.message, })
      }
      catch (error) { console.log(error) }
    } else Swal.fire({ icon: 'error', title: 'Register gagal', text: 'Password dan Confirm Password berbeda', })
  }

  return (
    <>
      <Button variant='link' onClick={openRegisterModal} data-testid="buttonRegisterForm">{t("banner.button-signup")}</Button>
      <RegisterModal handleSubmitRegister={handleSubmitRegister} closeRegisterModal={closeRegisterModal} registerModal={registerModal} />
    </>
  )
}