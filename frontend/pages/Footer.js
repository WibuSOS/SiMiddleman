import useTranslation from 'next-translate/useTranslation';

const Footer = () => {
  const { t } = useTranslation('common');

  return (
    <div>
      <footer>
        <div className="container">
          <div className="footer-content row d-flex justify-content-between">
            <div className="col-lg-12 col-md-12 col-sm-12 text-white text-center"> 
              &#169; 2022, {t("footer.copyright")}
            </div>
          </div>
        </div>
      </footer>
    </div>
  )
}

export default Footer