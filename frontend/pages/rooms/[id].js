import Button from 'react-bootstrap/Button';
import { useRouter } from "next/router";
import { useEffect, useState } from 'react';
import { getSession} from 'next-auth/react';

export default function Room( {user} ) {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);

  const router = useRouter();
  
  useEffect(() => {
    getRoomDetails();
  }, [])

  const getRoomDetails = async () => {
    const idRoom = router.query.id;
    const idPenjual = router.query.idPenjual;

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
            <Button type='submit'>Edit Produk</Button>
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
