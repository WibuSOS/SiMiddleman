import useTranslation from 'next-translate/useTranslation';
import { Button } from 'react-bootstrap';
import { signOut } from "next-auth/react";
import { useRouter } from "next/router";

export default function UserBanner ( props ) {
  const { t } = useTranslation('common');
  const router = useRouter();
  return (
    <div className='home-banner text-center'>
      <h2>{t('logged-in.banner.title')}</h2>
      <h3>{t('logged-in.banner.text')}</h3>
      <Button className='btn-simiddleman me-5' onClick={() => router.push({ pathname: '/profiles/[idUser]', query: {idUser: props.decoded?.ID}})}>{t('buttonProfile')}</Button>
      <Button onClick={() => signOut()} className='btn-simiddleman'>{t('logged-in.user-action.title.3')}</Button>
    </div>
  )
}