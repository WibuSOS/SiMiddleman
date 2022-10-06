import Button from 'react-bootstrap/Button';
import router, { useRouter } from "next/router";

function Room() {
    const router = useRouter()
    const query = router.query()
    
    const handleClose = async (e) => {
        e.preventDefault();
    }

    return (
        <div className='container pt-5'>
            <Button type='submit' onSubmit={handleClose}>Close</Button>
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
                <h5> {query.product.nama} </h5>
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
                <h5> Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</h5>
            </div>
            <div className="row pt-5">
                <Button type='submit'>Checkout</Button>
            </div>
        </div>
    )
}

export default Room;
