import Button from 'react-bootstrap/Button';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { getSession } from 'next-auth/react';
import Swal from 'sweetalert2';
import jwt from "jsonwebtoken";
import LogoBankSinarmas from '../../assets/logoBankSinarmas.png'
import useTranslation from 'next-translate/useTranslation';

function Pembayaran({ user }) {
    const router = useRouter();
    const [data, setData] = useState(null);
    const [dataRoom, setDataRoom] = useState(null);
    const [dataAfterChangeStatus, setDataAfterChangeStatus] = useState(null);
    const [error, setError] = useState(null);
    const decoded = jwt.verify(user, process.env.NEXT_PUBLIC_JWT_SECRET);
    const { t, lang } = useTranslation('payment');

    useEffect(() => {
        getHarga();
        getRoomDetails();
    }, [])

    const getHarga = async () => {
        const idRoom = router.query.idRoom;
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/getHarga/${idRoom}`, {
                method: 'GET',
                headers: { 'Authorization': 'Bearer ' + user, }
            });
            const data = await res.json();
            setData(data);
        } catch (error) {
            console.error();
        }
    }

    const getRoomDetails = async () => {
        const idRoom = router.query.idRoom;
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/joinroom/${idRoom}/${decoded.ID}`, {
                method: 'GET',
                headers: { 'Authorization': 'Bearer ' + user, }
            });
            const data = await res.json();
            setDataRoom(data);
        } catch (error) {
            console.error();
        }
    }

    const changeStatus = async () => {
        const idRoom = router.query.idRoom;
        if (data.data.status != dataRoom.statuses[1] && dataRoom.statuses[0] == data.data.status) {
            try {
                const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/updatestatus/${idRoom}`, {
                    method: 'PUT',
                    headers: { 'Authorization': 'Bearer ' + user, },
                    body: JSON.stringify({ status: dataRoom.statuses[1] })
                });
                const data = await res.json();
                setDataAfterChangeStatus(data);
            } catch (error) {
                console.error();
            }
        }
        if (dataAfterChangeStatus?.message !== null) {
            Swal.fire({ icon: 'success', title: 'Status Pembelian Berhasil diubah', showConfirmButton: false, timer: 1500, })
            router.push("https://forms.gle/yAtYBvu583nuVqmN6")
        }
    }

    return (
        <div className='content'>
            <div className="detail-produk-header">
                <div className='container'>
                    <h2>{t("header")}</h2>
                </div>
            </div>
            <div className='container'>
                <div className='d-flex flex-column justify-content-left detail-pembayaran align-items-center mx-auto'>
                    <img className='logo-bank-sinarmas' src={LogoBankSinarmas.src}></img>
                    {/* <h2>Bank Sinarmas</h2> */}
                    <h3 className='mt-5'>0056221875</h3>
                    <p>(Admin SiMiddleman)</p>
                    {error && <div>Failed to load {error.toString()}</div>}
                    {
                        !data ? <div>Loading...</div> : (
                            (data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>Data Kosong</p>
                        )
                    }
                    <p>{t("totalPembayaran")} : <b>Rp{data?.data.total}</b></p>
                    <p className='mt-5'>{t("detail.0")}</p>
                    <p className='mt-5'>{t("detail.1")}</p>
                    <p>{t("detail.2")}</p>
                    <Button variant='simiddleman' onClick={() => changeStatus()}>{t("buttonUpload")}</Button>
                </div>
            </div>
        </div>
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

export default Pembayaran;
