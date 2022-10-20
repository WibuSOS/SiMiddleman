import useTranslation from 'next-translate/useTranslation';
import { Button } from 'react-bootstrap';
import { signOut } from "next-auth/react";

export default function UserBanner () {
  const { t } = useTranslation('common');
  return (
    <div className='home-banner text-center'>
      <h2>{t('logged-in.banner.title')}</h2>
      <h3>{t('logged-in.banner.text')}</h3>
      <Button onClick={() => signOut()} className='btn-simiddleman'>{t('logged-in.user-action.title.3')}</Button>
    </div>
  )
}