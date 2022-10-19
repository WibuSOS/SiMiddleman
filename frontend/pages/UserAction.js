import Card from 'react-bootstrap/Card';
import CreateRoom from './CreateRoom';
import JoinRoom from './JoinRoom';
import useTranslation from 'next-translate/useTranslation';

export default function UserAction ( props ) {
  const { t } = useTranslation('common');
  
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
      </div>
    </div>
  )
}