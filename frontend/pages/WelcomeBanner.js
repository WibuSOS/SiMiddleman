import { Button } from "react-bootstrap";
import RegisterForm from './register';
import { signIn } from "next-auth/react";

export default function WelcomeBanner() {
  return (
    <div className='welcome-banner'>
      <div className='container'>
        <div className='row'>
          <div className='col-lg-6 banner-text-wrap'>
            <div className='banner-text'>
              <h2>Transaksi yang mudah &amp; aman</h2><br/>
              <p>Anda dapat menikmati kemudahan dalam melakukan transaksi jual beli dengan rasa aman dengan sistem rekening bersama dalam satu platform</p><br/>
              <Button variant="white" onClick={() => signIn()}>Masuk ke SiMiddleman+</Button><br/>
              <p>atau Anda belum punya akun?<RegisterForm /></p>
            </div>
          </div>
          <div className='col-lg-6 banner-image'></div>
        </div>
      </div>
    </div>
  )
}