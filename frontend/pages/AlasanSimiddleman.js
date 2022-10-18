import useTranslation from 'next-translate/useTranslation';

export default function AlasanSimiddleman() {
  const { t, lang } = useTranslation('common');

  return (
    <div className='alasan-simiddleman'>
      <div className='container'>
        <h2 className='text-center'>{t("simiddleman.title")} SiMiddleman+ ?</h2><br/>
        <div className='row'>
          <div className='col-lg-6 alasan-text-wrap'>
            <div className='alasan-text'>
              <h3><strong>{t("simiddleman.sub-title.0")}.</strong></h3><br/>
              <p>{t("simiddleman.text.0")}</p>
            </div>
          </div>
          <div className='col-lg-6 alasan-image alasan-img-1'></div>

          <div className='pt-5'></div>

          <div className='col-lg-6 alasan-image alasan-img-2'></div>
          <div className='col-lg-6 alasan-text-wrap'>
            <div className='alasan-text'>
              <h3><strong>{t("simiddleman.sub-title.1")}.</strong></h3><br/>
              <p>{t("simiddleman.text.1")}</p>
            </div>
          </div>

          <div className='pt-5'></div>

          <div className='col-lg-6 alasan-text-wrap'>
            <div className='alasan-text'>
              <h3><strong>{t("simiddleman.sub-title.2")}.</strong></h3><br/>
              <p>{t("simiddleman.text.2")}</p>
            </div>
          </div>
          <div className='col-lg-6 alasan-image alasan-img-3'></div>
        </div>
      </div>
    </div>
  )
}