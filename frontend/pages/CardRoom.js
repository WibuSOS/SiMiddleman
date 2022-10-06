import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';

export default function CardRoom() {
  return (
    <Card className='mt-5' style={{ width: '18rem' }}>
      <Card.Body>
        <Card.Title>Card Title</Card.Title>
        <Card.Text>
          Some quick example text to build on the card title and make up the
          bulk of the card's content.
        </Card.Text>
        <Button variant="primary">Go somewhere</Button>
      </Card.Body>
    </Card>
  );
}
