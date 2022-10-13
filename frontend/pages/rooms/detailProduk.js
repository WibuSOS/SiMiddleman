import Button from 'react-bootstrap/Button';
import ShowRoomCode from '../ShowRoomCode';

const STATUS_TRANSAKSI = [
  { status: 1, text: "mulai transaksi" },
  { status: 2, text: "barang dibayar" },
  { status: 3, text: "barang dikirim" },
  { status: 4, text: "konfirmasi barang sampai" },
];

function capitalizeFirstLetter(string) { return string[0].toUpperCase() + string.slice(1) }

export default function DetailProduk({ data, error, decoded, router, kirimBarang, handleConfirmation }) {
  let currentStatus = true;
  return (
    <div className='content'>
      <div className="detail-produk-header">
        <div className='container'>
          <h2>Detail Produk</h2>
          <ShowRoomCode roomCode={data?.data.roomCode} />
        </div>
      </div>
      <div className='container pt-5'>
        <div className='detail-produk'>
          {error && <div>Failed to load {error.toString()}</div>}
          {!data ? <div>Loading...</div> : ((data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>Data Kosong</p>)}
          <h3 className='nama-produk'>{data?.data.product.nama}</h3>
          <p>{data?.data.product.deskripsi}</p>
          <div className="py-5">
            <div className="container row d-flex justify-content-around">
              <div className="col-lg-6 harga-produk">
                <div className='row'>
                  <div className='col-lg-8 col-sm-12'>
                    <h3>Harga</h3>
                    <h3>Rp{data?.data.product.harga}</h3>
                  </div>
                  <div className='col-lg-4 col-sm-12'>
                    {data?.data.pembeliID === decoded.ID && data?.statuses.slice(0, -1).includes(data.data.status) && <Button className='me-3' onClick={() => { router.push({ pathname: '/rooms/payment/[idRoom]', query: { idRoom: `${data?.data.ID}` } }, '/rooms/payment/[idRoom]') }}>Beli</Button>}
                  </div>
                </div>                                 
              </div>
              <div className="col-lg-4 col-sm-12 kuantitas-produk">
                <h3>Kuantitas</h3>
                <p> {data?.data.product.kuantitas} PCS</p>
              </div>
            </div>
          </div>
          <div className='status-transaksi-wrapper'>
            <h3>Status Transaksi</h3>
            {data?.data.penjualID === decoded.ID && data?.statuses.slice(1, -1).includes(data.data.status) && <Button variant='simiddleman' className='mb-4' onClick={() => kirimBarang()}>Kirim Barang</Button>}
            {data?.data.pembeliID === decoded.ID && data?.statuses.slice(2, -1).includes(data.data.status) && <Button variant='simiddleman' className='mb-4' onClick={() => handleConfirmation()}>Konfirmasi Barang Telah Sampai</Button>}
            <div className='row d-flex justify-content-between status-transaksi'>
              {STATUS_TRANSAKSI.map((value, key) => {      
                if (data?.data.status === value.text) {
                  currentStatus = false;
                  return (
                    <div key={key} className='col order-tracking completed'>
                      <span class="is-complete"></span>
                      <p>{capitalizeFirstLetter(value.text)}</p>
                    </div>)}
                if (currentStatus) {
                  return (
                    <div key={key} className='col order-tracking completed'>
                      <span class="is-complete"></span>
                      <p>{capitalizeFirstLetter(value.text)}</p>
                    </div>)}
                else {
                  return (
                    <div key={key} className='col order-tracking'>
                        <span class="is-complete"></span>
                        <p>{capitalizeFirstLetter(value.text)}<br/></p>
                    </div>)}
              })}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}