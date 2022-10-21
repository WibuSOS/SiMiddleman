import { useRouter } from "next/router";
import { useEffect, useState } from 'react';
import { getSession, signIn } from 'next-auth/react';
import jwt from "jsonwebtoken";
import Swal from 'sweetalert2';
import DetailProduk from './detailProduk';
import useTranslation from 'next-translate/useTranslation';

export default function Room({ user }) {
  const [data, setData] = useState(null);
  const [error] = useState(null);
  const [namaProduk, setNamaProduk] = useState("");
  const [kuantitasProduk, setKuantitasProduk] = useState("");
  const [deskripsiProduk, setDeskripsiProduk] = useState("");
  const [hargaProduk, setHargaProduk] = useState("");
  const router = useRouter();
  const { t, lang } = useTranslation('detailProduct');
  useEffect(() => {
    getRoomDetails();
  }, [])
  const decoded = jwt.verify(user, process.env.NEXT_PUBLIC_JWT_SECRET);

  const getRoomDetails = async () => {
    const idRoom = router.query.id;
    const idPenjual = decoded.ID;
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/rooms/join/${idRoom}/${idPenjual}`, {
        method: 'GET',
        headers: {
          'Authorization': 'Bearer ' + user,
        }
      });
      const data = await res.json();
      setData(data);
      setNamaProduk(data?.data.product.nama);
      setKuantitasProduk(data?.data.product.kuantitas);
      setDeskripsiProduk(data?.data.product.deskripsi);
      setHargaProduk(data?.data.product.harga);
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

    await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/rooms/details/updatestatus/${idRoom}`, {
      method: 'PUT',
      headers: {
        'Authorization': 'Bearer ' + user,
      },
      body: JSON.stringify({ status: data.statuses.at(-1) })
    }).then(response => response.json()).then(data => res = data).catch(error => console.error('Error:', error));

    if (res?.message === "Success Update Status" || res?.message === "Sukses mengubah status") {
      Swal.fire({ icon: 'success', title: t("success"), text: res?.message, showConfirmButton: false, timer: 1500, })
      getRoomDetails();
    } else { Swal.fire({ icon: 'error', title: t("fail"), text: res?.message, showConfirmButton: false, timer: 1500 }) }
  }

  const kirimBarang = async () => {
    const idRoom = router.query.id;
    let res = null;
    if (data.data.status != data.statuses.at(-3)) {
      router.push('https://forms.gle/4uFn5cDSnYLW88ek9');
      return
    }

    await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/rooms/details/updatestatus/${idRoom}`, {
      method: 'PUT',
      headers: {
        'Authorization': 'Bearer ' + user,
      },
      body: JSON.stringify({ status: data.statuses.at(-2) })
    }).then(response => response.json()).then(data => res = data).catch(error => console.error('Error:', error));

    if (res?.message === "Success Update Status" || res?.message === "Sukses mengubah status") {
      Swal.fire({ icon: 'success', title: t("success"), text: res?.message, showConfirmButton: false, timer: 1500 });
      router.push('https://forms.gle/4uFn5cDSnYLW88ek9');
    } else { Swal.fire({ icon: 'error', title: t("fail"), text: res?.message, showConfirmButton: false, timer: 1500 }) }
  }

  return (
    <div className='content'>
      <DetailProduk data={data} error={error} decoded={decoded} router={router} kirimBarang={kirimBarang} handleConfirmation={handleConfirmation} user={user} namaProduk={namaProduk} setNamaProduk={setNamaProduk} hargaProduk={hargaProduk} setHargaProduk={setHargaProduk} deskripsiProduk={deskripsiProduk} setDeskripsiProduk={setDeskripsiProduk} kuantitasProduk={kuantitasProduk} setKuantitasProduk={setKuantitasProduk} getRoomDetails={getRoomDetails} />
    </div >
  )
}

export async function getServerSideProps(ctx) {
  const session = await getSession(ctx)
  if (!session) {
    return {
      redirect: { permanent: false, destination: "/api/auth/signin" }
    }
  }
  const { user } = session;
  return {
    props: { user },
  }
}
