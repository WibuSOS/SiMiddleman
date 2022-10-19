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
      let hp = "62" + text.substring(1)
      let href = "https://wa.me/" + hp
      return href
    }
    else {
      let text = data?.data?.penjual.NoHp + ""
      let hp = "62" + text.substring(1)
      let href = "https://wa.me/" + hp + "?text=I'm%20interested%20in%20your%20product%20for%20sale"
      return href
    }
  }

  return (
    <div className='content'>
      <div className="detail-produk-header">
        <div className='container'>
          <h2>{t("header")}</h2>
          <div className='row d-flex justify-content-start'>
            <div className='col-lg-4 col-md-5 col-sm-12'>
              <ShowRoomCode roomCode={data?.data.roomCode} />
              {data?.data.penjualID === decoded?.ID && data?.statuses[0] === data?.data.status ? <Button className='btn-simiddleman' onClick={openUpdateProductModal}>{t("updateProductButton")}</Button> : "" }
            </div>
            <div className='col-lg-4 col-md-5 col-sm-12'>
              <UpdateProduct closeUpdateProductModal={closeUpdateProductModal} updateProductModal={updateProductModal} data={data} user={user} namaProduk={namaProduk} setNamaProduk={setNamaProduk} hargaProduk={hargaProduk} setHargaProduk={setHargaProduk} deskripsiProduk={deskripsiProduk} setDeskripsiProduk={setDeskripsiProduk} kuantitasProduk={kuantitasProduk} setKuantitasProduk={setKuantitasProduk} getRoomDetails={getRoomDetails} />
              {data?.data?.pembeli && data?.data.penjualID === decoded?.ID && <Button href={contactNumber(data?.data.idPenjual)} variant="simiddleman" className='wa' rel='noreferrer' target={'_blank'}>Chat Pembeli</Button>}
              {data?.data?.pembeli && data?.data.pembeliID === decoded?.ID && <Button href={contactNumber(data?.data.idPenjual)} variant="simiddleman" className='wa' rel='noreferrer' target={'_blank'}>Chat Penjual</Button>}
            </div>
          </div>
        </div>
      </div>
      <div className='container pt-5'>
        <div className='detail-produk'>
          {error && <div>{t("load-fail")} {error.toString()}</div>}
          {!data ? <div>{t("load.loading")}</div> : ((data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>{t("list-empty")}</p>)}
          <h3 className='nama-produk'>
            {data?.data.product.nama}
          </h3>
          <p>{data?.data.product.deskripsi}</p>
          <div className="py-3">
            <div className='harga-produk'>
              <div className='row pt-3'>
                <div className='col-6 harga text-center'>
                  <h3>{t("price")}</h3>
                  <p>Rp {data?.data.product.harga.toLocaleString()}</p>
                </div>
                <div className='col-6 kuantitas text-center'>
                  <h3>{t("quantity")}</h3>
                  <p>{data?.data.product.kuantitas} PCS</p>
                </div>
              </div>
              <div className='row pt-3 d-flex justify-content-center'>
                <div className='col-lg-4 col-md-6 col-sm-8'>
                  {data?.data.pembeliID === decoded?.ID && data?.statuses[0] === data?.data.status && <Button className='w-100' variant='simiddleman' onClick={() => { router.push({ pathname: '/rooms/payment/[idRoom]', query: { idRoom: `${data?.data.ID}` } }) }}>{t("buyButton")}</Button>}
                    {data?.data.pembeliID === decoded?.ID && data?.statuses[0] != data?.data.status && data?.statuses[3] != data?.data.status && <Button className='w-100' variant='simiddleman' onClick={() => { router.push("https://forms.gle/yAtYBvu583nuVqmN6") }}>{t("buttonUpload")}</Button>}
                </div>
              </div>
            </div>
            {/* <div className="container row d-flex justify-content-around">
              <div className="col-lg-6 harga-produk">
                <div className='row'>
                  <div className='col-lg-8 col-sm-12'>
                    <h3>{t("price")}</h3>
                    <h3>Rp{data?.data.product.harga.toLocaleString()}</h3>
                  </div>
                  <div className='col-lg-4 col-sm-12'>
                    {data?.data.pembeliID === decoded?.ID && data?.statuses[0] === data?.data.status && <Button className='w-100' variant='simiddleman' onClick={() => { router.push({ pathname: '/rooms/payment/[idRoom]', query: { idRoom: `${data?.data.ID}` } }) }}>{t("buyButton")}</Button>}
                    {data?.data.pembeliID === decoded?.ID && data?.statuses[0] != data?.data.status && data?.statuses[3] != data?.data.status && <Button className='w-100' variant='simiddleman' onClick={() => { router.push("https://forms.gle/yAtYBvu583nuVqmN6") }}>{t("buttonUpload")}</Button>}
                  </div>
                </div>
              </div>
              <div className="col-lg-4 col-sm-12 kuantitas-produk">
                <h3>{t("quantity")}</h3>
                <p> {data?.data.product.kuantitas} PCS</p>
              </div>
            </div> */}
          </div>
          <div className='status-transaksi-wrapper'>
            <div className='d-flex flex-row justify-content-between mb-2'>
              <h3 className='align-self-center'>{t("transactionStatus")}</h3>
              {data?.data.penjualID === decoded?.ID && data?.statuses.slice(1, -1).includes(data.data.status) && <Button variant='simiddleman' onClick={() => kirimBarang()}>{t("sentButton")}</Button>}
              {data?.data.pembeliID === decoded?.ID && data?.statuses.slice(2, -1).includes(data.data.status) && <Button variant='simiddleman' onClick={() => handleConfirmation()}>{t("confirmButton")}</Button>}
            </div>
            {data?.data.pembeliID === decoded?.ID && data?.statuses[0] === data?.data.status ? <div className='mb-4'><p>{t("statusExplanationBuyer.0")}</p></div> : data?.data.pembeliID === decoded?.ID && data?.statuses[1] === data?.data.status ? <div className='mb-4'><p>{t("statusExplanationBuyer.1")}</p></div> : data?.data.pembeliID === decoded?.ID && data?.statuses[2] === data?.data.status ? <div className='mb-4'><p>{t("statusExplanationBuyer.2")}</p></div> : data?.data.pembeliID === decoded?.ID && data?.statuses[3] === data?.data.status ? <div className='mb-4'><p>{t("statusExplanationBuyer.3")}</p></div> : ""}
            {data?.data.penjualID === decoded?.ID && data?.statuses[0] === data?.data.status ? <div className='mb-4'><p>{t("statusExplanationSeller.0")}</p></div> : data?.data.penjualID === decoded?.ID && data?.statuses[1] === data?.data.status ? <div className='mb-4'><p>{t("statusExplanationSeller.1")}</p></div> : data?.data.penjualID === decoded?.ID && data?.statuses[2] === data?.data.status ? <div className='mb-4'><p>{t("statusExplanationSeller.2")}</p></div> : data?.data.penjualID === decoded?.ID && data?.statuses[3] === data?.data.status ? <div className='mb-4'><p>{t("statusExplanationSeller.3")}</p></div> : ""}
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