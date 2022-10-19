import Button from 'react-bootstrap/Button';
import RegisterForm from './register';
import useTranslation from 'next-translate/useTranslation';
import LoginForm from './LoginForm';

export default function Login() {
  const { t } = useTranslation('common');
  return (   
    <div className="content" data-testid="main">
      <div className="container">
        <div className="row d-flex justify-content-center align-items-center">
          <div className="col-lg-8 col-sm-12 login-illustration"></div>
          <div className="col-lg-4 col-sm-12 mt-3">
            <div className="login-form-wrapper text-center">
              <h2 data-testid="Title">{t('login-page.title')}</h2>
              <LoginForm/>
              <div className='d-flex justify-content-between align-items-center'>
                <Button type='submit' className="w-100 mb-3" variant='simiddleman' form='loginForm' data-testid="masukButton">{t('login-page.btn-login')}</Button>
              </div>
              <p className='or' data-testid="OR">{t('login-page.or')}</p>
              <RegisterForm />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}