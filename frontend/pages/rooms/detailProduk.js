import Button from 'react-bootstrap/Button';
import ShowRoomCode from '../ShowRoomCode';
import UpdateProduct from '../UpdateProduct';
import { useState } from 'react';
import useTranslation from 'next-translate/useTranslation';

const STATUS_TRANSAKSI = [
  { status: 1, text: "mulai transaksi" },
  { status: 2, text: "barang dibayar" },
  { status: 3, text: "barang dikirim" },
  { status: 4, text: "konfirmasi barang sampai" },
];

function capitalizeFirstLetter(string) { return string[0].toUpperCase() + string.slice(1) }

export default function DetailProduk({ data, error, decoded, router, kirimBarang, handleConfirmation, user, namaProduk, setNamaProduk, hargaProduk, setHargaProduk, deskripsiProduk, setDeskripsiProduk, kuantitasProduk, setKuantitasProduk, getRoomDetails }) {
  const [updateProductModal, setupdateProductModal] = useState(false);
  const openUpdateProductModal = () => setupdateProductModal(true);
  const closeUpdateProductModal = () => setupdateProductModal(false);
  const { t, lang } = useTranslation('detailProduct');
  let currentStatus = true;

  const contactNumber = (idPenjual) => {

    if (idPenjual === decoded) {
      let text = data?.data?.pembeli.NoHp + ""
      let hp = text.substring(1)
      let href = "https://wa.me/" + hp + "?text=I'm%20interested%20in%20your%20product%20for%20sale"
      return href
    }
    else {
      let text = data?.data?.penjual.NoHp + ""
      let hp = text.substring(1)
      let href = "https://wa.me/" + hp + "?text=I'm%20interested%20in%20your%20product%20for%20sale"
      return href
    }
  }

  return (
    <div className='content'>
      <div className="detail-produk-header">
        <div className='container col'>
          <h2>{t("header")}</h2>
          <ShowRoomCode roomCode={data?.data.roomCode} />
          {data?.data?.pembeli && data?.data.penjualID === decoded?.ID && <a href={contactNumber(data?.data.idPenjual)} className='ms-5 btn-simiddleman wa' target={'_blank'}>Chat Pembeli</a>}
          {data?.data?.pembeli && data?.data.pembeliID === decoded?.ID && <a href={contactNumber(data?.data.idPenjual)} className='ms-5 btn-simiddleman wa' target={'_blank'}>Chat Penjual</a>}
        </div>
      </div>
      <div className='container pt-5'>
        <div className='detail-produk'>
          {error && <div>Failed to load {error.toString()}</div>}
          {!data ? <div>Loading...</div> : ((data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>Data Kosong</p>)}
          <h3 className='nama-produk'>
            {data?.data.product.nama}
            <Button className='ms-5 btn-simiddleman' onClick={openUpdateProductModal}>{t("updateProductButton")}</Button>
            <UpdateProduct closeUpdateProductModal={closeUpdateProductModal} updateProductModal={updateProductModal} data={data} user={user} namaProduk={namaProduk} setNamaProduk={setNamaProduk} hargaProduk={hargaProduk} setHargaProduk={setHargaProduk} deskripsiProduk={deskripsiProduk} setDeskripsiProduk={setDeskripsiProduk} kuantitasProduk={kuantitasProduk} setKuantitasProduk={setKuantitasProduk} getRoomDetails={getRoomDetails} />
          </h3>
          <p>{data?.data.product.deskripsi}</p>
          <div className="py-5">
            <div className="container row d-flex justify-content-around">
              <div className="col-lg-6 harga-produk">
                <div className='row'>
                  <div className='col-lg-8 col-sm-12'>
                    <h3>{t("price")}</h3>
                    <h3>Rp{data?.data.product.harga}</h3>
                  </div>
                  <div className='col-lg-4 col-sm-12'>
                    {data?.data.pembeliID === decoded?.ID && data?.statuses.slice(0, -1).includes(data.data.status) && <Button className='w-100' variant='simiddleman' onClick={() => { router.push({ pathname: '/rooms/payment/[idRoom]', query: { idRoom: `${data?.data.ID}` } }, '/rooms/payment/[idRoom]') }}>Beli</Button>}
                  </div>
                </div>
              </div>
              <div className="col-lg-4 col-sm-12 kuantitas-produk">
                <h3>{t("quantity")}</h3>
                <p> {data?.data.product.kuantitas} PCS</p>
              </div>
            </div>
          </div>
          <div className='status-transaksi-wrapper'>
            <h3>{t("transactionStatus")}</h3>
            {data?.data.penjualID === decoded?.ID && data?.statuses.slice(1, -1).includes(data.data.status) && <Button variant='simiddleman' className='mb-4' onClick={() => kirimBarang()}>Kirim Barang</Button>}
            {data?.data.pembeliID === decoded?.ID && data?.statuses.slice(2, -1).includes(data.data.status) && <Button variant='simiddleman' className='mb-4' onClick={() => handleConfirmation()}>Konfirmasi Barang Telah Sampai</Button>}
            <div className='row d-flex justify-content-between status-transaksi'>
              {STATUS_TRANSAKSI.map((value, key) => {
                if (data?.data.status === value.text) {
                  currentStatus = false;
                  return (
                    <div key={key} className='col order-tracking completed'>
                      <span className="is-complete"></span>
                      <p>{capitalizeFirstLetter(t("transaction." + key))}</p>
                    </div>)
                }
                if (currentStatus) {
                  return (
                    <div key={key} className='col order-tracking completed'>
                      <span className="is-complete"></span>
                      <p>{capitalizeFirstLetter(t("transaction." + key))}</p>
                    </div>)
                }
                else {
                  return (
                    <div key={key} className='col order-tracking'>
                      <span className="is-complete"></span>
                      <p>{capitalizeFirstLetter(t("transaction." + key))}<br /></p>
                    </div>)
                }
              })}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}