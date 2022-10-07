import Button from 'react-bootstrap/Button';
import router, { useRouter } from "next/router";

function Room() {
    const roomDetail = async () => {
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/joinroom/${room.id}/${user.id}`, {
          method: 'GET',
        });
        const data = await res.json();
        setData(data);
      } catch (error) {
        setError(error)
      }
    }
    return (
        <div className='container pt-5'>
            <Button type='submit'>Close</Button>
            <div className="d-flex justify-content-between">
                <div className='pt-5'>
                    <h2>Detail Produk</h2>
                    <h4> Berikut adalah detail produk di room ini :</h4>
                </div>
                <div className='pt-5'>
                    <Button type='submit'>Edit Produk</Button>
                </div>
            </div>
            <div className='pt-5'>
                <h5> NAMA PRODUK : </h5>
                <h5> Razer </h5>
            </div>
            <div className="d-flex justify-content-between pt-5">
                <div className="col">
                    <div className="row">
                        <div className="col">
                            <h5> KUANTITAS PRODUK : </h5>
                            <h5> 1 pcs </h5>
                        </div>
                        <div className="col">
                            <h5> HARGA PRODUK : </h5>
                            <h5> Rp. 1.000.000 </h5>
                        </div>
                    </div>
                </div>
            </div>
            <div className='pt-5'>
                <h5> DESKRIPSI PRODUK : </h5>
                <h5> Lorem Ipsum is simply dummy text of the printing and typesetting industry.</h5>
            </div>
            <div className="row pt-5">
                <Button type='submit'>Checkout</Button>
            </div>
        </div>
    )
}

export default Room;
