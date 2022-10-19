import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import { useRouter } from 'next/router';
import sellerIcon from './assets/seller.png';
import buyerIcon from './assets/buyer.png'
import useTranslation from 'next-translate/useTranslation';

export default function CardRoom(props) {
  const router = useRouter();
  const { t, lang } = useTranslation('cardRoom');
  const isSeller = (idPenjual) => {
    if (idPenjual === props.decoded) return (
      <>
        <img src={sellerIcon.src} className='seller-icon'></img>
        <p>{t("Seller")}</p> 
      </>
    )
    else return (
      <>
        <img src={buyerIcon.src} className='seller-icon'></img>
        <p>{t("Buyer")}</p> 
      </>
    )
  }
  return (
    <Card className='me-4 mb-4 room-card' style={{ width: '22rem' }}>
      <Card.Body className='d-flex flex-column justify-content-between'>
        <Card.Title className='mb-1'>
          {isSeller(props.idPenjual)}
        </Card.Title>
        <Card.Body className='d-flex flex-column justify-content-between'>
          <div>
            <p className='card-nama-produk'>{props.namaProduk}</p>
            <p className='card-deskripsi-produk'>{props.deskripsiProduk}</p>
          </div>
          <div>
            <p className='card-harga-produk'>Rp{props.hargaProduk}</p>
          </div>
        </Card.Body>
        <Button variant="simiddleman" className='w-100' onClick={() => {
          router.push(
            {
              pathname: '/rooms/[idRoom]',
              query: {
                id: `${props.idRoom}`,
                idRoom: `${props.kodeRuangan}`,
              }
            }
          )
        }}>{t("masukButton")}</Button>
      </Card.Body>
    </Card>
  );
}
