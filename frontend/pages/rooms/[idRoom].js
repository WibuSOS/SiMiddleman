import Button from 'react-bootstrap/Button';
import { useRouter } from "next/router";
import { useEffect, useState } from 'react';
import { getSession} from 'next-auth/react';
import jwt from "jsonwebtoken";

export default function Room( {user} ) {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);

  const router = useRouter();
  
  useEffect(() => {
    getRoomDetails();
  }, [])

  const decoded = jwt.verify(user, process.env.JWT_SECRET);

  const getRoomDetails = async () => {
    const idRoom = router.query.id;
    const idPenjual = decoded.ID;

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/joinroom/${idRoom}/${idPenjual}`, {
        method: 'GET',
        headers: {
          'Authorization': 'Bearer ' + user,
        }
      });
      const data = await res.json();
      setData(data);
    } catch (error) {
      console.error();
    }
  }

  return (
    <div className='container pt-5'>
      <Button type='submit'>Close</Button>
        <hr></hr>
        <h2>Detail Session: </h2>
        <p><strong>Token</strong>: { user }</p>
        <p><strong>ID</strong>: { decoded.ID }</p>
        <p><strong>Email</strong>: { decoded.Email }</p>
        <p><strong>Role</strong>: { decoded.Role }</p>
        <hr></hr>
        <h2>Detail Room: </h2>
        <p><strong>Kode Room</strong>: { data?.data.roomCode }</p>
        <p><strong>ID Penjual</strong>: { data?.data.penjualID }</p>
        <p><strong>ID Pembeli</strong>: { data?.data.pembeliID }</p>
        <hr></hr>
      <div className="d-flex justify-content-between">
        <div className='pt-5'>
            <h2>Detail Produk</h2>
            {error && <div>Failed to load {error.toString()}</div>}
            {
              !data ? <div>Loading...</div>
                : (
                  (data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>Data Kosong</p>
                )
            }
            <p>{  data?.data.product.deskripsi }</p>
        </div>
        <div className='pt-5'>
          {decoded.ID === data?.data.penjualID ? (
            <Button type='submit'>Edit Produk</Button>
          ) : (
            <Button onClick={() => {router.push(
              {
                pathname: '/Payment',
                query: {
                  idRoom: `${data?.data.ID}`,
                },
              }, '/Payment'
            )}}>Beli</Button>
          )}
        </div>
      </div>
      <div className='pt-5'>
        <h5> NAMA PRODUK : </h5>
        <p>{ data?.data.product.nama }</p>
      </div>
      <div className="d-flex justify-content-between pt-5">
        <div className="col">
          <div className="row">
            <div className="col">
              <h5> KUANTITAS PRODUK : </h5>
              <p> { data?.data.product.kuantitas } </p>
            </div>
            <div className="col">
              <h5> HARGA PRODUK : </h5>
              <p> { data?.data.product.harga } </p>
            </div>
          </div>
        </div>
      </div>
      <div className='pt-5'>
        <h5> DESKRIPSI PRODUK : </h5>
        <p>{ data?.data.product.deskripsi }</p>
      </div>
      <div className="row pt-5">
        <Button type='submit'>Checkout</Button>
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
