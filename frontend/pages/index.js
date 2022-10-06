import {Button } from 'react-bootstrap';
import { signOut, signIn, useSession } from "next-auth/react";
import CreateRoom from './CreateRoom';
import RegisterForm from './register';
import jwt from "jsonwebtoken";
import CardRoom from './CardRoom';
import { useState, useEffect } from 'react'

function Home() {    
  const { data: session } = useSession();
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);

  if (session) {
    const decoded = jwt.verify(session['user'], process.env.JWT_SECRET);
    const token = session['user'];
    const roomList = [];
    
    const getAllRoom = async () => {
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/rooms/${decoded.ID}`, {
          method: 'GET',
        });
        const data = await res.json();
        setData(data);
      } catch (error) {
        setError(error)
      }
    }
    useEffect(() => {
      getAllRoom();
    },[]);

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
          {error && <div>Failed to load {error.toString()}</div>}
          {
            !data ? <div>Loading...</div>
              : (
                (data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>Data Kosong</p>
              )
          }
          {data?.data && data?.data?.map((item, index) => (
            roomList.push(
              <CardRoom 
                key={index}
                kodeRuangan={item.roomCode}
                namaProduk={item.product.nama} 
                deskripsiProduk={item.product.deskripsi} 
                hargaProduk={item.product.harga} 
                kuantitasProduk={item.product.kuantitas}/>
            )
          ))}
          <div className='row d-flex justify-content-between p-3'>
            {roomList}
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

export default Home;
