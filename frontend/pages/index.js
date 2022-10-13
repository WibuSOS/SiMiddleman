import { Button } from 'react-bootstrap';
import { signOut, getSession } from "next-auth/react";
import CreateRoom from './CreateRoom';
import JoinRoom from './JoinRoom';
import jwt from "jsonwebtoken";
import CardRoom from './CardRoom';
import { useEffect, useState } from 'react';
import WelcomeBanner from './WelcomeBanner';
import AlasanSimiddleman from './AlasanSimiddleman';
import SimiddlemanSummaries from './SimiddlemanSummaries';
import Card from 'react-bootstrap/Card';

function Home({ user }) {
  const [data, setData] = useState(null)
  const [error] = useState(null)

  useEffect(() => {
    if (user) {
      GetAllRoom();
    }
  }, [user])

  if (!user) {
    return (
      <div className='content'>
        <WelcomeBanner />
        <AlasanSimiddleman />
        <SimiddlemanSummaries />
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
      <div className='pt-5' style={{ backgroundColor: "#CC0F0F", paddingBottom: "150px" }}>
        <h2 className='text-center' style={{ color: "white" }}>Halo, Selamat Datang di aplikasi SiMiddleman+</h2>
        <h3 className='text-center' style={{ color: "white" }}>Buat atau join Ruang obrolan pada tombol dibawah</h3>
        <div className='position-relative'>
          <div className='d-flex position-absolute start-50 translate-middle' style={{ paddingTop: "300px" }}>
            <Card style={{ width: '358px', boxShadow: "0px 32px 50px -9px rgba(0, 0, 0, 0.25)", borderRadius: "10px", border: "none" }}>
              <Card.Body>
                <Card.Title className='mb-5'>Create Room</Card.Title>
                <Card.Text>
                  Kamu dapat membuat room untuk melakukan transaksi dengan pembeli.
                </Card.Text>
                <CreateRoom idPenjual={decoded.ID} sessionToken={user} />
              </Card.Body>
            </Card>
            <Card className='mx-5' style={{ width: '358px', boxShadow: "0px 32px 50px -9px rgba(0, 0, 0, 0.25)", borderRadius: "10px", border: "none" }}>
              <Card.Body>
                <Card.Title className='mb-5'>Join Room</Card.Title>
                <Card.Text>
                  Mendaftarkan room yang telah dibuat oleh penjual kedalam list room kamu.
                </Card.Text>
                <JoinRoom idPembeli={decoded.ID} sessionToken={user} />
              </Card.Body>
            </Card>
            <Card style={{ width: '358px', boxShadow: "0px 32px 50px -9px rgba(0, 0, 0, 0.25)", borderRadius: "10px", border: "none" }}>
              <Card.Body>
                <Card.Title className='mb-5'>Sign Out</Card.Title>
                <Card.Text>
                  Melakukan signout untuk keluar dari akun kamu.
                </Card.Text>
                <Button onClick={() => signOut()} className='w-100 btn-simiddleman'>Sign out</Button>
              </Card.Body>
            </Card>
          </div>
        </div>
      </div>
      <div className='container'>
        <div className='pb-5' style={{ paddingTop: "175px" }}>
          <h2>Berikut merupakan room yang telah anda buat</h2>
          <div className='row d-flex justify-content-between p-3'>
            {error && <div>Failed to load {error.toString()}</div>}
            {
              !data ? <div>Loading...</div>
                : (
                  (data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>Data Kosong</p>
                )
            }
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
