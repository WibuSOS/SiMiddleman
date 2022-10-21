import { Button } from "react-bootstrap";
import RegisterForm from './register';
import useTranslation from 'next-translate/useTranslation';
import { useRouter } from "next/router";

export default function WelcomeBanner() {
  const { t } = useTranslation('common');
  const router = useRouter();
  const locale = router.locale === "id" ? router.locale : "";

  return (
    <div className='welcome-banner'>
      <div className='container'>
        <div className='row'>
          <div className='col-lg-6 banner-text-wrap'>
            <div className='banner-text'>
              <h2>{t("banner.title")}</h2><br/>
              <p>{t("banner.text")}</p><br/>
              <Button variant="white" onClick={() => router.push(`${locale}` + "/Login?callbackUrl=" + `${process.env.NEXT_PUBLIC_FRONTEND_URL}`)}>{t("banner.button-signin")} SiMiddleman+</Button><br/>
              <p>{t("banner.signup-text")}<RegisterForm /></p>
            </div>
          </div>
          <div className='col-lg-6 banner-image'></div>
        </div>
      </div>
    </div>
  )
}