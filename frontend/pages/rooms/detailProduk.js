import Button from 'react-bootstrap/Button';

export default function DetailProduk({ data, error, decoded, router, kirimBarang, handleConfirmation }) {
    return (
        <div>
            <div className="d-flex justify-content-between">
                <div className='pt-5'>
                <h2>Detail Produk</h2>
                {error && <div>Failed to load {error.toString()}</div>}
                {!data ? <div>Loading...</div> : ((data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>Data Kosong</p>)}
                <p>{data?.data.product.deskripsi}</p>
                </div>
                <div className='pt-5'>
                {data?.data.pembeliID === decoded.ID && data?.statuses.slice(0, -1).includes(data.data.status) && <Button className='me-3' onClick={() => { router.push({ pathname: '/rooms/payment/[idRoom]', query: { idRoom: `${data?.data.ID}` } }, '/rooms/payment/[idRoom]') }}>Beli</Button>}
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
                <h5> STATUS TRANSAKSI : </h5>
                <p>{data?.data.status}</p>
            </div>
            <div className='pt-5'>
                <h5> DESKRIPSI PRODUK : </h5>
                <p>{data?.data.product.deskripsi}</p>
            </div>
        </div>
    )
}