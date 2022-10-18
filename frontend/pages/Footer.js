import Link from "next/link"
import useTranslation from 'next-translate/useTranslation';

const Footer = () => {
  const { t, lang } = useTranslation('common');

  return (
    <div>
      <footer>
        <div className="container">
          <div className="footer-content row d-flex justify-content-between">
            <div className="col-lg-4 col-md-6 col-sm-6 text-white text-center"> 
              &#169; 2022, {t("footer.copyright")}
            </div>
            <div className="col-lg-4 col-md-6 col-sm-6 text-white" hidden>
              <div className="row d-flex justify-content-between link">
                <div className="col-sm-6 text-center">
                  <Link href="/">{t("footer.privacyPolicy")}</Link>
                </div>
                <div className="col-sm-6 text-center">
                  <Link href="/">{t("footer.termsCondition")}</Link>
                </div>
              </div>
            </div>
          </div>
        </div>
      </footer>
    </div>
  )
}

export default Footer