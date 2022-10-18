import useTranslation from 'next-translate/useTranslation';

export default function SimiddlemanSummaries () {
  const { t, lang } = useTranslation('common');

  return (
    <div className='simiddleman-summaries'>
      <div className='container'>
        <h2 className='text-center'>SiMiddleman+ {t("achievement.title")}</h2>
        <div className='row summaries-list text-center'>
          <div className='col'>
            <h3>100</h3>
            <span>{t("achievement.text.0")}</span>
          </div>
          <div className='col'>
            <h3>100</h3>
            <span>{t("achievement.text.1")}</span>
          </div>
          <div className='col'>
            <h3>50</h3>
            <span>{t("achievement.text.2")}</span>
          </div>
        </div>
      </div>
    </div>
  )
}