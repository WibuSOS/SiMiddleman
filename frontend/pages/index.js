import { Button } from 'react-bootstrap';
import { signOut, signIn, getSession } from "next-auth/react";
import CreateRoom from './CreateRoom';
import JoinRoom from './JoinRoom';
import RegisterForm from './register';
import jwt from "jsonwebtoken";
import CardRoom from './CardRoom';
import { useEffect, useState } from 'react';

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
        <div className='welcome-banner'>
          <div className='row'>
            <div className='col-lg-6 banner-text-wrap'>
              <div className='banner-text'>
                <h2>Akses yang mudah dan aman untuk bertransaksi dengan pihak ke 3</h2>
                <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo.</p>
                <Button variant="white" onClick={() => signIn()}>Login</Button><br/>
                <p>Belum punya akun?<RegisterForm /></p>
              </div>
            </div>
            <div className='col-lg-6 banner-image'></div>
          </div>
        </div>
        <div className='alasan-simiddleman'>
          <div className='container'>
            <h2>Kenapa menggunakan SiMiddleman+ ?</h2><br/>
            <div className='row'>
              <div className='col-lg-6 alasan-text-wrap'>
                <div className='alasan-text'>
                  <h3><strong>To build for people.</strong></h3><br/>
                  <p>Whether you want to edit your Google Docs, resolve Jira issues, or collaborate over Zoom.<br/><br/>Circle has 100+ integrations with tools you already use and love.</p>
                </div>
              </div>
              <div className='col-lg-6 alasan-image'></div>
            </div>
          </div>
        </div>
        <div className='simiddleman-keamanan'>
          <div className='container'>
            <h2>Keamanan yang unggul</h2><br/><br/>
            <div className='keamanan-image'></div><br/>
            <div className='row keamanan-summaries'>
              <div className='col'>
                <h3><i class="fa fa-check"></i> 18281 <span>signed up last month</span></h3>
              </div>
              <div className='col'>
                <h3><i class="fa fa-check"></i> GPDR-&amp; CCPA- <span>ready</span></h3>
              </div>
              <div className='col'>
                <h3><i class="fa fa-check"></i> Leader@G2 <span>Summer</span></h3>
              </div>
            </div>
          </div>
        </div>
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
