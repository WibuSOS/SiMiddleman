import {Button } from 'react-bootstrap';
import { signOut, signIn, useSession } from "next-auth/react";
import CreateRoom from './CreateRoom';
import RegisterForm from './register';
import jwt from "jsonwebtoken";

function Home() {    
  const { data: session } = useSession();

  if (session) {
    const decoded = jwt.verify(session['user'], process.env.JWT_SECRET);
    return (
      <>
        <div className='container pt-5'>
          <h2>Halo Selamat Datang, berikut merupakan data anda:</h2>
          <ul>
            <li>Id: { decoded.ID }</li>
            <li>Email: { decoded.Email }</li>
            <li>Role: { decoded.Role }</li>
          </ul>
          <div className='d-flex justify-content-left'>
            <Button onClick={() => signOut()} className="mx-3">Sign out</Button>
            <CreateRoom/>
          </div>
        </div>
      </>
    )
  }
  else {
    return (
      <div className='container mx-10 my-7'>
          <Button variant="primary" onClick={() => signIn()}>
              Login
          </Button>
          <RegisterForm/>
      </div>
    );
  }
}

export default Home;
