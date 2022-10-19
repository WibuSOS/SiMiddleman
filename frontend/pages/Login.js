import { Form } from "react-bootstrap";
import Button from 'react-bootstrap/Button';
import { signIn } from "next-auth/react";
import RegisterForm from './register';
import Swal from 'sweetalert2';
import router from "next/router";
import useTranslation from 'next-translate/useTranslation';

export default function LoginForm() {
  const { t } = useTranslation('common');
  const handleSubmitLogin = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      email: formData.get("email"),
      password: formData.get("password")
    }
    signIn("credentials", {
      email: body.email, 
      password: body.password,
      redirect: false,
      callbackUrl: "/"},)
      .then(({ ok, error }) => {
        if (ok) {
          Swal.fire({
            icon: 'success',
            title: t('login-page.swal.success'),
            showConfirmButton: false,
            timer: 1500,
          })
          router.push("/");
        } else {
          Swal.fire({
            icon: 'error',
            title: t('login-page.swal.failed.title'),
            text: t('login-page.swal.failed.text'),
          })
        }
      })
  }
    return (   
      <div className="content" data-testid="main">
        <div className="container">
          <div className="row d-flex justify-content-center align-items-center">
            <div className="col-lg-8 col-sm-12 login-illustration"></div>
            <div className="col-lg-4 col-sm-12 mt-3">
              <div className="login-form-wrapper text-center">
                <h2 data-testid="Title">{t('login-page.title')}</h2>
                <Form onSubmit={handleSubmitLogin} id="loginForm" className="pt-3" data-testid="login">
                  <Form.Group className="mb-3">
                    <Form.Control
                    type="email"
                    placeholder={t('login-page.email-placeholder')}
                    name="email"
                    data-testid="email"
                    required
                    autoFocus/>
                  </Form.Group>
                  <Form.Group className="mb-3">
                    <Form.Control
                    type="password"
                    placeholder={t('login-page.password-placeholder')}
                    name="password"
                    data-testid="password"
                    required
                    minLength={8}/>
                  </Form.Group>
                </Form>
                <div className='d-flex justify-content-between align-items-center'>
                  {/* <a onClick={() => alert("not yet implemented")} data-testid="LupaPassword">{t('login-page.forget-password')}</a> */}
                  <Button type='submit' className="w-100 mb-3" variant='simiddleman' form='loginForm' data-testid="masukButton">{t('login-page.btn-login')}</Button>
                </div>
                <p className='or' data-testid="OR">{t('login-page.or')}</p>
                <RegisterForm />
              </div>
            </div>
          </div>
        </div>
      </div>
    )
}