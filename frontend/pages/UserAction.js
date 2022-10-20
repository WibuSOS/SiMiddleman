import Card from 'react-bootstrap/Card';
import CreateRoom from './CreateRoom';
import JoinRoom from './JoinRoom';
import useTranslation from 'next-translate/useTranslation';
import { useRouter } from "next/router";
import { Button } from 'react-bootstrap';

export default function UserAction ( props ) {
  const { t } = useTranslation('common');
  const router = useRouter();
  
  return (
    <div className='user-action-wrapper'>
      <div className='row d-flex justify-content-around p-2'>
        <Card className='user-action col-lg-4 col-md-5 col-sm-12'>
          <Card.Body className='d-flex flex-column justify-content-around'>
            <Card.Title className='mb-5'>{t('logged-in.user-action.title.0')}</Card.Title>
            <Card.Text>{t('logged-in.user-action.text.0')}</Card.Text>
            <CreateRoom idPenjual={props.decoded?.ID} sessionToken={props.user} />
          </Card.Body>
        </Card>
        <Card className='user-action col-lg-4 col-md-5 col-sm-12'>
          <Card.Body className='d-flex flex-column justify-content-around'>
            <Card.Title className='mb-5'>{t('logged-in.user-action.title.1')}</Card.Title>
            <Card.Text>{t('logged-in.user-action.text.1')}</Card.Text>
            <JoinRoom idPembeli={props.decoded?.ID} sessionToken={props.user} />
          </Card.Body>
        </Card>
        <Card className='user-action col-lg-4 col-md-5 col-sm-12'>
          <Card.Body className='d-flex flex-column justify-content-around'>
            <Card.Title className='mb-5'>{t('logged-in.user-action.title.2')}</Card.Title>
            <Card.Text>{t('logged-in.user-action.text.2')}</Card.Text>
            <Button className='w-100 btn-simiddleman' onClick={() => router.push({ pathname: '/profiles/[idUser]', query: {idUser: props.decoded?.ID}})}>{t('buttonProfile')}</Button>
          </Card.Body>
        </Card>
      </div>
    </div>
  )
}