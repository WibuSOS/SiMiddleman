import {Button } from 'react-bootstrap';
import { signOut, signIn, useSession } from "next-auth/react";
import CreateRoom from './CreateRoom';
import RegisterForm from './register';
import jwt from "jsonwebtoken";
import CardRoom from './CardRoom';

function Home() {    
  const { data: session } = useSession();

  if (session) {
    const decoded = jwt.verify(session['user'], process.env.JWT_SECRET);
    const token = session['user'];
    return (
      <div className='container'>
        <div className='pt-5'>
          <h2>Halo Selamat Datang, berikut merupakan data anda:</h2>
          <ul>
            <li>Id: { decoded.ID }</li>
            <li>Email: { decoded.Email }</li>
            <li>Role: { decoded.Role }</li>
          </ul>
          <div className='d-flex justify-content-left'>
            <Button onClick={() => signOut()} className="mx-3">Sign out</Button>
            <CreateRoom idPenjual={decoded.ID} sessionToken={token}/>
          </div>
        </div>
        <div className='pt-5'>
          <h2>Berikut merupakan room yang telah anda buat</h2>
          <CardRoom/>
        </div>
      </div>
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
