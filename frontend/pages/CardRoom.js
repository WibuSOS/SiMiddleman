import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import { useRouter } from 'next/router';

export default function CardRoom(props) {
  const router = useRouter();
  return (
    <Card className='mt-5' style={{ width: '22rem' }}>
      <Card.Body>
        <Card.Title className='mb-5'>Kode Ruangan: {props.kodeRuangan}</Card.Title>
        <Card.Text>
          <b>Nama: </b>{props.namaProduk}
          <br />
          <b>Deskripsi: </b>{props.deskripsiProduk}
          <br />
          <b>Harga:</b> Rp{props.hargaProduk}
        </Card.Text>
        <Button variant="primary" className='w-100' onClick={() => {
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
