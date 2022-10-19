import Button from 'react-bootstrap/Button';

function Room(RoomDetails) {
  return (
    <div className='container pt-5'>
      <Button type='submit'>Close</Button>
      <div className="d-flex justify-content-between">
        <div className='pt-5'>
          <h2>Detail Produk</h2>
          <h4></h4>
        </div>
        <div className='pt-5'>
          <Button type='submit'>Edit Produk</Button>
        </div>
      </div>
      <div className='pt-5'>
        <h5> NAMA PRODUK : </h5>
        <h5> Razer </h5>
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
        <h5> Lorem Ipsum is simply dummy text of the printing and typesetting industry.</h5>
      </div>
      <div className="row pt-5">
        <Button type='submit'>Checkout</Button>
      </div>
    </div>
  )
}

export async function getServerSideProps({ query }) {
  const roomId = query.kodeRuangan;
  const userId = query.idPenjual;
  let roomDetails = null;
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/joinroom/${roomId}/${userId}`, {
      method: 'GET',
    });
    roomDetails = await res.json();
  } catch (error) {
    console.error();
  }

  return {
    props: {
      roomDetails
    },
  }
}

export default Room;
