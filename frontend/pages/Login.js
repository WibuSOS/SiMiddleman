import { Form } from "react-bootstrap";
import logo from './assets/logo.png';
import Button from 'react-bootstrap/Button';
import { signIn } from "next-auth/react";

export default function LoginForm() {
  const handleSubmitLogin = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      email: formData.get("email"),
      password: formData.get("password")
    }
    signIn("credentials", {email: body.email, password: body.password, callbackUrl: "/"})
  }
    return (   
      <div className="test">
        <div className="centered-form text-center">
          <div className="avatar text-center">
              <img src={logo.src} alt="logo SiMiddleman+"/>
          </div>
          <h2>Login</h2>
          <Form onSubmit={handleSubmitLogin} id="loginForm" className="pt-3">
            <Form.Group className="mb-3">
              <Form.Control
              type="email"
              placeholder="Email"
              name="email"
              required
              autoFocus/>
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Control
              type="password"
              placeholder="Password"
              name="password"
              required
              minLength={8}/>
            </Form.Group>
          </Form>
          <div className='d-flex justify-content-between'>
            <a>Lupa Password?</a>
            <Button type='submit' variant='merah' form='loginForm'>Masuk</Button>
          </div>
          <p className='or'>OR</p>
          <Button variant='merah' className='w-100'>Daftar Akun</Button>
        </div>
      </div>
    )
}