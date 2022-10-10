import Button from 'react-bootstrap/Button';
import Link from 'next/link';

function Pembayaran() {
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
                    <b> Rp. 150.000</b>
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
