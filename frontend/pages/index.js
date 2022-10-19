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
import { useRouter } from "next/router";

function Home({ user }) {
  const [data, setData] = useState(null)
  const [error] = useState(null)
  const [showSeller, setShowSeller] = useState(false);
  const [showBuyer, setShowBuyer] = useState(false);
  const router = useRouter();

  const { t } = useTranslation('common');

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
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/rooms/${decoded.ID}`, {
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

  console.log(showSeller)
  console.log(decoded.ID)

  let dataList = []
  data?.data && data?.data?.map((item, index) => (
    (showSeller && decoded.ID === item.penjualID) ?
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
        kuantitasProduk={item.product.kuantitas}/>
    ) : (showBuyer && decoded.ID !== item.penjualID) ? dataList.push(
      <CardRoom
        key={index}
        decoded={decoded.ID}
        idPenjual={item.penjualID}
        idRoom={item.ID}
        kodeRuangan={item.roomCode}
        namaProduk={item.product.nama}
        deskripsiProduk={item.product.deskripsi}
        hargaProduk={item.product.harga}
        kuantitasProduk={item.product.kuantitas}/>
    ) : (!showBuyer && !showSeller) ? dataList.push(
      <CardRoom
        key={index}
        decoded={decoded.ID}
        idPenjual={item.penjualID}
        idRoom={item.ID}
        kodeRuangan={item.roomCode}
        namaProduk={item.product.nama}
        deskripsiProduk={item.product.deskripsi}
        hargaProduk={item.product.harga}
        kuantitasProduk={item.product.kuantitas}/>
    ) : ""
  ))

  return (
    <div className='content'>
      <div className='home-banner text-center'>
        <h2>{t('logged-in.banner.title')}</h2>
        <h3>{t('logged-in.banner.text')}</h3>
        <Button onClick={() => signOut()} className='btn-simiddleman'>{t('logged-in.user-action.title.2')}</Button>
      </div>
      <div className='user-action-wrapper'>
        <div className='row d-flex justify-content-around p-2'>
          <Card className='user-action col-lg-4 col-md-5 col-sm-12'>
            <Card.Body className='d-flex flex-column justify-content-around'>
              <Card.Title className='mb-5'>{t('logged-in.user-action.title.0')}</Card.Title>
              <Card.Text>
              {t('logged-in.user-action.text.0')}
              </Card.Text>
              <CreateRoom idPenjual={decoded.ID} sessionToken={user} />
            </Card.Body>
          </Card>
          <Card className='user-action col-lg-4 col-md-5 col-sm-12'>
            <Card.Body className='d-flex flex-column justify-content-around'>
              <Card.Title className='mb-5'>{t('logged-in.user-action.title.1')}</Card.Title>
              <Card.Text>
              {t('logged-in.user-action.text.1')}
              </Card.Text>
              <JoinRoom idPembeli={decoded.ID} sessionToken={user} />
            </Card.Body>
          </Card>
        </div>
      </div>
      <div className='container'>
        <div className='pb-5 pt-5'>
          <h2 className='room-anda'>{t('logged-in.room-list.title')}</h2>
          <div className='d-flex justify-content-center'>
              <div className='row d-flex justify-content-around room-role'>
                  <Button className={showSeller ? "active" : ""} variant='simiddleman' onClick={() => {setShowSeller(current => !current);setShowBuyer(false)}}>{t("buyer")}</Button>
                  <Button className={showBuyer ? "active" : ""} variant='simiddleman' onClick={() => {setShowBuyer(current => !current);setShowSeller(false)}}>{t("seller")}</Button>
              </div>
          </div>
          <div className='row d-flex justify-content-around px-3'>
            {error && <div> {t('logged-in.error.load-fail')} {error.toString()}</div>}
            {
              !data ? <div>{t('logged-in.error.loading')}</div>
                : (
                  (data?.data ?? []).length === 0 && t('logged-in.error.list-empty')
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
