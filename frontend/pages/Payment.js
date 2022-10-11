import Button from 'react-bootstrap/Button';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { getSession } from 'next-auth/react';
import jwt from "jsonwebtoken";

function Pembayaran({ user }) {
    const router = useRouter();
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);
    const idRoom = router.query.idRoom;

    useEffect(() => {
        getHarga();
    }, [])

    const decoded = jwt.verify(user, process.env.NEXT_PUBLIC_JWT_SECRET);

    const getHarga = async () => {

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
    return (
        <div className='container pt-5' style={{ backgroundColor: "#FFFFFF" }}>
            <Button type='submit' data-testid='back_button'>Back</Button>
            <div className='d-flex flex-column justify-content-center'>
                <h2 className='mx-auto mb-4' style={{ fontSize: "48px" }} data-testid='sinarmas'>Sinarmas</h2>
                <h3 className='mx-auto' style={{ fontSize: "36px" }} data-testid='no_rek'>0123456789</h3>
                <div className='mx-auto mb-4' style={{ fontSize: "24px" }} data-testid='simiddleman'>
                    a/n Admin SiMiddleman
                </div>
                <div className='mx-auto mb-4' style={{ fontSize: "30px" }}>
                    Total Pembayaran :
                    <b data-testid='harga'> Rp. {data?.data.total}</b>
                </div>
                <div className='mx-auto mb-4' style={{ fontSize: "24px" }} data-testid='instruction_no_rek'>
                    Silahkan lakukan pembayaran ke nomor yang ada diatas.
                </div>
                <div className='mx-auto' style={{ fontSize: "24px" }} data-testid='question_payment'>
                    Anda sudah melakukan pembayaran?
                </div>
                <div className='mx-auto mb-3' style={{ fontSize: "24px" }} data-testid='instruction_receipt'>
                    Silahkan upload bukti pembayaran anda!
                </div>
                <Link href={`https://forms.gle/yAtYBvu583nuVqmN6`}>
                    <Button className='mx-auto' data-testid='upload_receipt_button'>Upload Bukti Pembayaran</Button>
                </Link>
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
