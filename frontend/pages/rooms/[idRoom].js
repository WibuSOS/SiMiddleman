import Button from 'react-bootstrap/Button';
import { useRouter } from "next/router";
import { useEffect, useState } from 'react';
import { getSession } from 'next-auth/react';
import jwt from "jsonwebtoken";
import ShowRoomCode from '../ShowRoomCode';
import Swal from 'sweetalert2';

export default function Room({ user }) {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const router = useRouter();
  useEffect(() => {
    getRoomDetails();
  }, [])
  const decoded = jwt.verify(user, process.env.NEXT_PUBLIC_JWT_SECRET);

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

  const handleConfirmation = async () => {
    const idRoom = router.query.id;
    let res = null;

    if (data.data.status != data.statuses.at(-2)) {
      Swal.fire({ icon: 'error', title: 'Status Barang Tidak Dapat Diubah', showConfirmButton: false, timer: 1500 });
      return
    }

    await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/updatestatus/${idRoom}`, {
      method: 'PUT',
      headers: {
        'Authorization': 'Bearer ' + user,
      },
      body: JSON.stringify({ status: data.statuses.at(-1) })
    }).then(response => response.json()).then(data => res = data).catch(error => console.error('Error:', error));

    if (res?.message == `success update status ${data.statuses.at(-1)}`) {
      Swal.fire({ icon: 'success', title: 'Status Barang Berhasil Diubah', showConfirmButton: false, timer: 1500, })
      getRoomDetails();
    } else { Swal.fire({ icon: 'error', title: 'Status Barang Tidak Dapat Diubah', showConfirmButton: false, timer: 1500 }) }
  }

  const kirimBarang = async () => {
    const idRoom = router.query.id;
    let res = null;

    if (data.data.status != data.statuses.at(-3)) {
      router.push('https://forms.gle/4uFn5cDSnYLW88ek9');
      return
    }

    await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/updatestatus/${idRoom}`, {
      method: 'PUT',
      headers: {
        'Authorization': 'Bearer ' + user,
      },
      body: JSON.stringify({ status: data.statuses.at(-2) })
    }).then(response => response.json()).then(data => res = data).catch(error => console.error('Error:', error));

    if (res?.message == `success update status ${data.statuses.at(-2)}`) {
      Swal.fire({ icon: 'success', title: 'Status Barang Berhasil Diubah', showConfirmButton: false, timer: 1500 });
      router.push('https://forms.gle/4uFn5cDSnYLW88ek9');
    } else { Swal.fire({ icon: 'error', title: 'Status Barang Tidak Dapat Diubah', showConfirmButton: false, timer: 1500 }) }
  }

  return (
    <div className='content container pt-5'>
      <Button type='submit' className='me-3'>Close</Button>
      <ShowRoomCode roomCode={data?.data.roomCode} />
      <div className="d-flex justify-content-between">
        <div className='pt-5'>
          <h2>Detail Produk</h2>
          {error && <div>Failed to load {error.toString()}</div>}
          {!data ? <div>Loading...</div> : ((data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>Data Kosong</p>)}
          <p>{data?.data.product.deskripsi}</p>
        </div>
        <div className='pt-5'>
          {data?.data.pembeliID === decoded.ID && data?.statuses.slice(0, -1).includes(data.data.status) && <Button className='me-3' onClick={() => { router.push({ pathname: '/rooms/payment/[idRoom]', query: { idRoom: `${data?.data.ID}`, statusAfter: data?.statuses[1], statusBefore: data?.statuses[0] } }, '/rooms/payment/[idRoom]') }}>Beli</Button>}
          {data?.data.penjualID === decoded.ID && data?.statuses.slice(1, -1).includes(data.data.status) && <Button className='me-3' onClick={() => kirimBarang()}>Kirim Barang</Button>}
          {data?.data.pembeliID === decoded.ID && data?.statuses.slice(2, -1).includes(data.data.status) && <Button className='me-3' onClick={() => handleConfirmation()}>Barang Telah Sampai</Button>}
        </div>
      </div>
      <div className='pt-5'>
        <h5> NAMA PRODUK : </h5>
        <p>{data?.data.product.nama}</p>
      </div>
      <div className="d-flex justify-content-between pt-5">
        <div className="col">
          <div className="row">
            <div className="col">
              <h5> KUANTITAS PRODUK : </h5>
              <p> {data?.data.product.kuantitas} </p>
            </div>
            <div className="col">
              <h5> HARGA PRODUK : </h5>
              <p> {data?.data.product.harga} </p>
            </div>
          </div>
        </div>
      </div>
      <div className='pt-5'>
        <h5> DESKRIPSI PRODUK : </h5>
        <p>{data?.data.product.deskripsi}</p>
      </div>
      <div className="row pt-5">
        <Button type='submit'>Checkout</Button>
      </div>
    </div >
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
