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
        <div className='container pt-5' data-testid="container">
            <Button type='submit' data-testid="Close">Close</Button>
            <div className="d-flex justify-content-between" data-testid="container2">
                <div className='pt-5'>
                    <h2 data-testid="title">Detail Produk</h2>
                    <h4 data-testid="subTitle"> Berikut adalah detail produk di room ini :</h4>
                </div>
                <div className='pt-5'>
                    <Button type='submit' data-testid="buttonEdit">Edit Produk</Button>
                </div>
            </div>
            <div className='pt-5'>
                <h5 data-testid="titleNamaProduk"> NAMA PRODUK : </h5>
                <h5 data-testid="namaProduk"> Razer </h5>
            </div>
            <div className="d-flex justify-content-between pt-5">
                <div className="col">
                    <div className="row">
                        <div className="col">
                            <h5 data-testid="titleKuantitas"> KUANTITAS PRODUK : </h5>
                            <h5 data-testid="kuantitasProduk"> 1 pcs </h5>
                        </div>
                        <div className="col">
                            <h5 data-testid="titleHarga"> HARGA PRODUK : </h5>
                            <h5 data-testid="hargaProduk"> Rp. 1.000.000 </h5>
                        </div>
                    </div>
                </div>
            </div>
            <div className='pt-5'>
                <h5 data-testid="titleDeskripsi"> DESKRIPSI PRODUK : </h5>
                <h5 data-testid="deskripsiProduk"> Lorem Ipsum is simply dummy text of the printing and typesetting industry.</h5>
            </div>
            <div className="row pt-5">
                <Button type='submit' data-testid="buttonCheckout">Checkout</Button>
            </div>
        </div>
    )
}

export default Room;
