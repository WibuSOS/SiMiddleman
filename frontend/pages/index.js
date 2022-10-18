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
import useTranslation from 'next-translate/useTranslation';

function Home({ user }) {
  const [data, setData] = useState(null)
  const [error] = useState(null)

  const { t } = useTranslation('common');

  useEffect(() => {
    if (user) {
      GetAllRoom();
    }
  }, [user])

  if (!user) {
    return (
      <div className='content'>
        <h1>{t('testing')}</h1>
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
        decoded={decoded.ID}
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
      <div className='pt-5' style={{ backgroundColor: "#CC0F0F", paddingBottom: "150px", borderBottomRightRadius: "20%", borderBottomLeftRadius: "20%", background: "linear-gradient(76.81deg,#CC0F0F 15.71%,#ff0025 68.97%,#fd195e 94.61%)" }}>
        <h2 className='text-center' style={{ color: "white", fontFamily: "Ubuntu" }}>{t('logged-in.banner.title')}</h2>
        <h3 className='text-center' style={{ color: "white", fontFamily: "Ubuntu" }}>{t('logged-in.banner.text')}</h3>
        <div className='position-relative'>
          <div className='d-flex position-absolute start-50 translate-middle' style={{ paddingTop: "300px" }}>
            <Card className='user-action mx-3'>
              <Card.Body>
                <Card.Title className='mb-5'>{t('logged-in.user-action.title.0')}</Card.Title>
                <Card.Text>
                {t('logged-in.user-action.text.0')}
                </Card.Text>
                <CreateRoom idPenjual={decoded.ID} sessionToken={user} />
              </Card.Body>
            </Card>
            <Card className='user-action mx-3'>
              <Card.Body>
                <Card.Title className='mb-5'>{t('logged-in.user-action.title.1')}</Card.Title>
                <Card.Text>
                {t('logged-in.user-action.text.1')}
                </Card.Text>
                <JoinRoom idPembeli={decoded.ID} sessionToken={user} />
              </Card.Body>
            </Card>
            <Card className='user-action mx-3'>
              <Card.Body>
                <Card.Title className='mb-5'>{t('logged-in.user-action.title.2')}</Card.Title>
                <Card.Text>
                {t('logged-in.user-action.text.2')}
                </Card.Text>
                <Button onClick={() => signOut()} className='w-100 btn-simiddleman'>{t('logged-in.user-action.title.2')}</Button>
              </Card.Body>
            </Card>
          </div>
        </div>
      </div>
      <div className='container'>
        <div className='pb-5' style={{ paddingTop: "175px" }}>
          <h2 className='room-anda'>{t('logged-in.room-list.title')}</h2>
          <div className='row d-flex justify-content-center'>
            {error && <div> {t('logged-in.error.load-fail')} {error.toString()}</div>}
            {
              !data ? <div>{t('logged-in.error.loadingt')}</div>
                : (
                  (data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>{t('logged-in.error.list-empty')}</p>
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
