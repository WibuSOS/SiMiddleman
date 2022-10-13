import Button from 'react-bootstrap/Button';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { getSession } from 'next-auth/react';
import Swal from 'sweetalert2';
import jwt from "jsonwebtoken";

function Pembayaran({ user }) {
    const router = useRouter();
    const [data, setData] = useState(null);
    const [dataRoom, setDataRoom] = useState(null);
    const [dataAfterChangeStatus, setDataAfterChangeStatus] = useState(null);
    const [error, setError] = useState(null);
    const decoded = jwt.verify(user, process.env.NEXT_PUBLIC_JWT_SECRET);

    useEffect(() => {
        getHarga();
        getRoomDetails();
    }, [])

    const getHarga = async () => {
        const idRoom = router.query.idRoom;
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/getHarga/${idRoom}`, {
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

    const getRoomDetails = async () => {
        const idRoom = router.query.idRoom;
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/joinroom/${idRoom}/${decoded.ID}`, {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer ' + user,
                }
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
                    headers: {
                        'Authorization': 'Bearer ' + user,
                    },
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
        <div className='content container pt-5' style={{ backgroundColor: "#FFFFFF" }}>
            {"current: " + data?.data?.status}
            <br />
            {"before: " + dataRoom?.statuses[0]}
            <br />
            {"after: " + dataRoom?.statuses[1]}
            <br />
            <Button type='submit'>Back</Button>
            <div className='d-flex flex-column justify-content-center'>
                <h2 className='mx-auto mb-4' style={{ fontSize: "48px" }}>Sinarmas</h2>
                <h3 className='mx-auto' style={{ fontSize: "36px" }}>0123456789</h3>
                <div className='mx-auto mb-4' style={{ fontSize: "24px" }}>
                    a/n Admin SiMiddleman
                </div>
                {error && <div>Failed to load {error.toString()}</div>}
                {
                    !data ? <div>Loading...</div>
                        : (
                            (data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>Data Kosong</p>
                        )
                }
                <div className='mx-auto mb-4' style={{ fontSize: "30px" }}>
                    Total Pembayaran :
                    <b> Rp. {data?.data.total}</b>
                </div>
                <div className='mx-auto mb-4' style={{ fontSize: "24px" }}>
                    Silahkan lakukan pembayaran ke nomor yang ada diatas.
                </div>
                <div className='mx-auto' style={{ fontSize: "24px" }}>
                    Anda sudah melakukan pembayaran?
                </div>
                <div className='mx-auto mb-3' style={{ fontSize: "24px" }}>
                    Silahkan upload bukti pembayaran anda!
                </div>
                <Button onClick={() => changeStatus()} className='mx-auto'>Upload Bukti Pembayaran</Button>
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

export default Pembayaran;
