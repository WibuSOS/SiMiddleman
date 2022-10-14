import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import { useRouter } from 'next/router';
import sellerIcon from './assets/seller.png';
import buyerIcon from './assets/buyer.png'


export default function CardRoom(props) {
  const router = useRouter();
  console.log(props.decoded);
  const isSeller = (idPenjual) => {
    if (idPenjual === props.decoded) return (
      <>
        <img src={sellerIcon.src} className='seller-icon'></img>
        <p>Seller</p> 
      </>
    )
    else return (
      <>
        <img src={buyerIcon.src} className='seller-icon'></img>
        <p>Buyer</p> 
      </>
    )
  }
  return (
    <Card className='me-4 mb-4 room-card' style={{ width: '22rem' }}>
      <Card.Body>
        <Card.Title className='mb-5'>
        {isSeller(props.idPenjual)}
        </Card.Title>
        <Card.Text>
          <p className='card-nama-produk'>{props.namaProduk}</p>
          <p className='card-deskripsi-produk'>{props.deskripsiProduk}</p>
          <p className='card-harga-produk'>Rp{props.hargaProduk}</p>
        </Card.Text>
        <Button variant="simiddleman" className='w-100' onClick={() => {
          router.push(
            {
              pathname: '/rooms/[idRoom]',
              query: {
                id: `${props.idRoom}`,
                idRoom: `${props.kodeRuangan}`,
              },
            }, '/rooms/[idRoom]'
          )
        }}>Masuk Room</Button>
      </Card.Body>
    </Card>
  );
}
