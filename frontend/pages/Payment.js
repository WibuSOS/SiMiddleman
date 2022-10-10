import Button from 'react-bootstrap/Button';
import Link from 'next/link';
import { useRouter } from 'next/router';

function Pembayaran() {
    const router = useRouter();
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);
    const idRoom = router.query.idRoom;

    useEffect(() => {
        getHarga();
    }, [])

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
        <div className='container pt-5' style={{backgroundColor: "#FFFFFF"}}>
            <Button type='submit'>Back</Button>
            <div className='d-flex flex-column justify-content-center'>
                <h2 className='mx-auto mb-4' style={{fontSize: "48px"}}>Sinarmas</h2>
                <h3 className='mx-auto' style={{fontSize: "36px"}}>0123456789</h3>
                <div className='mx-auto mb-4' style={{fontSize: "24px"}}>
                    a/n Admin SiMiddleman
                </div>
                <div className='mx-auto mb-4' style={{fontSize: "30px"}}>
                    Total Pembayaran :
                    <b> Rp. {data.data.harga}</b>
                </div>
                <div className='mx-auto mb-4' style={{fontSize: "24px"}}>
                    Silahkan lakukan pembayaran ke nomor yang ada diatas.
                </div>
                <div className='mx-auto'style={{fontSize: "24px"}}>
                    Anda sudah melakukan pembayaran?
                </div>
                <div className='mx-auto mb-3'style={{fontSize: "24px"}}>
                    Silahkan upload bukti pembayaran anda!
                </div>
                <Link href={`https://forms.gle/yAtYBvu583nuVqmN6`}>
                    <Button className='mx-auto'>Upload Bukti Pembayaran</Button>
                </Link>
            </div>
        </div>
    )
}

export default Pembayaran;
