import {Button } from 'react-bootstrap';
import { signOut, signIn, useSession, getSession } from "next-auth/react";
import CreateRoom from './CreateRoom';
import RegisterForm from './register';
import jwt from "jsonwebtoken";
import CardRoom from './CardRoom';

function Home( dataRoom) {    
  const { data: session } = useSession();
  if (session) {
    const token = session['user'];
    const decoded = jwt.verify(token, process.env.JWT_SECRET);
    let getAllRoom = dataRoom.dataRoom['data'];
    console.log(getAllRoom);
    let AllRoom = [];
    getAllRoom == null ? "data kosong" :
    getAllRoom.map((item, index) => (
      AllRoom.push(
        <CardRoom 
        key={index}
        idPenjual={item.penjualID}
        idRoom={item.ID}
        kodeRuangan={item.roomCode}
        namaProduk={item.product.nama} 
        deskripsiProduk={item.product.deskripsi} 
        hargaProduk={item.product.harga} 
        kuantitasProduk={item.product.kuantitas}/>
      )
    ))
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
          <div className='row d-flex justify-content-between p-3'>
            {AllRoom}
          </div>
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

export async function getServerSideProps({req}) {
  const session  = await getSession({req});
  let dataRoom = null;
  if (session) {
    const token = await session['user'];
    const decoded = jwt.verify(token, process.env.JWT_SECRET);

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/rooms/${decoded.ID}`, {
        method: 'GET',
      });
      dataRoom = await res.json();
    } catch (error) {
      console.error();
    }
  }
  else {
    return {
      redirect: {
        destination: '/api/auth/signin',
        permanent: false,
      }
    };
  }
  
  return {
    props: {
      dataRoom,
      session,
    }, // will be passed to the page component as props
  } 
}

export default Home;
