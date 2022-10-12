import { Button } from 'react-bootstrap';
import { signOut, signIn, useSession, getSession } from "next-auth/react";
import CreateRoom from './CreateRoom';
import JoinRoom from './JoinRoom';
import RegisterForm from './register';
import jwt from "jsonwebtoken";
import CardRoom from './CardRoom';
import { useEffect, useState } from 'react';

function Home({ user }) {
  const [data, setData] = useState(null)
  const [error, setError] = useState(null)

  useEffect(() => {
    if (user) {
      GetAllRoom();
    }
  }, [user])

  if (!user) {
    return (
      <div className='content container mx-10 my-7'>
        <Button variant="primary" onClick={() => signIn()}>
          Login
        </Button>
        <RegisterForm />
      </div>
    )
  }

  const decoded = jwt.verify(user, process.env.NEXT_PUBLIC_JWT_SECRET);

  const GetAllRoom = async () => {

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/rooms/${decoded.ID}`, {
        method: 'GET',
        headers: {
          'Authorization': 'Bearer ' + user
        },
      });
      const data = await res.json();
      setData(data);
    } catch (error) {
      console.error();
    }
  }


  let dataList = []
  data?.data && data?.data?.map((item, index) => (
    dataList.push(
      <CardRoom
        key={index}
        idPenjual={item.penjualID}
        idRoom={item.ID}
        kodeRuangan={item.roomCode}
        namaProduk={item.product.nama}
        deskripsiProduk={item.product.deskripsi}
        hargaProduk={item.product.harga}
        kuantitasProduk={item.product.kuantitas} />
    )))

  return (
    <div className='content'>
      {error && <div>Failed to load {error.toString()}</div>}
      {
        !data ? <div>Loading...</div>
          : (
            (data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>Data Kosong</p>
          )
      }

      <div className='container'>
        <div className='pt-5'>
          <h2>Halo Selamat Datang, berikut merupakan data anda:</h2>
          <ul>
            <li>Id: {decoded.ID}</li>
            <li>Email: {decoded.Email}</li>
            <li>Role: {decoded.Role}</li>
          </ul>
          <div className='d-flex justify-content-left'>
            <Button onClick={() => signOut()} className="mx-3">Sign out</Button>
            <CreateRoom idPenjual={decoded.ID} sessionToken={user} />
            <JoinRoom idPembeli={decoded.ID} sessionToken={user} />
          </div>
        </div>
        <div className='pt-5'>
          <h2>Berikut merupakan room yang telah anda buat</h2>
          <div className='row d-flex justify-content-between p-3'>
            {dataList}
          </div>
        </div>
      </div>
    </div>
  )

}

export async function getServerSideProps(ctx) {
  const session = await getSession(ctx)
  if (!session) {
    return {
      props: {}
    }
  }
  const { user } = session;
  return {
    props: { user },
  }
}

export default Home;
