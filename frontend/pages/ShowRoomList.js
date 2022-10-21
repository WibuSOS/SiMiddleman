import CardRoom from './CardRoom';
import { useState } from 'react';
import useTranslation from 'next-translate/useTranslation';
import { Button } from 'react-bootstrap';

export default function ShowRoomList ( props ) {
  const [error] = useState(null)
  const [showSeller, setShowSeller] = useState(false);
  const [showBuyer, setShowBuyer] = useState(false);
  let dataList = []
  const { t } = useTranslation('common');

  const singleRoom = (item, index) => {
    return (<CardRoom key={index} decoded={props.decoded.ID} idPenjual={item.penjualID} idRoom={item.ID} kodeRuangan={item.roomCode}namaProduk={item.product.nama} deskripsiProduk={item.product.deskripsi} hargaProduk={item.product.harga} kuantitasProduk={item.product.kuantitas}/>)
  }

  props.data?.data && props.data?.data?.map((item, index) => (
    (showSeller && props.decoded.ID === item.penjualID) ? dataList.push(singleRoom(item, index)) : 
    (showBuyer && props.decoded.ID !== item.penjualID) ? dataList.push(singleRoom(item, index)) : 
    (!showBuyer && !showSeller) ? dataList.push(singleRoom(item, index)) : ""
  ))

  return (
    <div className='container'>
      <div className='pb-5 pt-5'>
        <h2 className='room-anda'>{t('logged-in.room-list.title')}</h2>
        <div className='d-flex justify-content-center'>
          <div className='row d-flex justify-content-around room-role'>
            <Button className={showSeller ? "active" : ""} variant='simiddleman' onClick={() => {setShowSeller(current => !current);setShowBuyer(false)}}>{t("seller")}</Button>
            <Button className={showBuyer ? "active" : ""} variant='simiddleman' onClick={() => {setShowBuyer(current => !current);setShowSeller(false)}}>{t("buyer")}</Button>
          </div>
        </div>
        <div className='row d-flex justify-content-around px-3'>
          {error && <div> {t('logged-in.error.load-fail')} {error.toString()}</div>}
          {!props.data ? <div>{t('logged-in.error.loading')}</div> : (
            (props.data?.data ?? []).length === 0 && t('logged-in.error.list-empty'))}
          {dataList}
        </div>
      </div>
    </div>
  )
}