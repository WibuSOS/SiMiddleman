import { Form } from "react-bootstrap";
import Swal from 'sweetalert2';
import router from "next/router";
import { signIn } from "next-auth/react";
import useTranslation from 'next-translate/useTranslation';

export default function LoginForm () {
  const handleSubmitLogin = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = { email: formData.get("email"), password: formData.get("password")}
    signIn("credentials", { email: body.email, password: body.password, redirect: false, callbackUrl: "/"},).then(({ ok, error }) => {
      if (ok) {
        Swal.fire({ icon: 'success', title: t('login-page.swal.success'), showConfirmButton: false, timer: 1500,})
        router.push("/");
      } else Swal.fire({ icon: 'error', title: t('login-page.swal.failed.title'), text: t('login-page.swal.failed.text'),})
    })
  }
  const { t } = useTranslation('common');

  return (
    <Form onSubmit={handleSubmitLogin} id="loginForm" className="pt-3" data-testid="login">
      <Form.Group className="mb-3">
        <Form.Control type="email" placeholder={t('login-page.email-placeholder')} name="email" data-testid="email" required autoFocus/>
      </Form.Group>
      <Form.Group className="mb-3">
        <Form.Control type="password" placeholder={t('login-page.password-placeholder')} name="password" data-testid="password" required minLength={8}/>
      </Form.Group>
    </Form>
  )
}