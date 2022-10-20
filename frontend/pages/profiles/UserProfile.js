import useTranslation from 'next-translate/useTranslation';

export default function UserProfile({ data, error }) {
    const { t, lang } = useTranslation('userProfile');

    return (
        <div className='content'>
            <div className="detail-produk-header">
                <div className='container'>
                    <h2>{t("header")}</h2>
                </div>
            </div>
            <div className='container pt-5'>
                <div className='detail-produk'>
                    {error && <div>{t("load-fail")} {error.toString()}</div>}
                    {!data ? <div>{t("load.loading")}</div> : ((data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>{t("list-empty")}</p>)}
                    <h2 className='nama-produk'>
                        {data?.data.nama}
                    </h2>
                    <div className="py-3">
                        <div className='d-flex flex-row justify-content-between'>
                            <div>
                                <h3>{t("email")}</h3>
                                <p>{data?.data.email}</p>
                            </div>
                            <div>
                                <h3>{t("handphone")}</h3>
                                <p>{data?.data.noHp}</p>
                            </div>
                        </div>
                        <div className='mt-4'>
                            <div className='col-lg-4 col-md-6 col-sm-8'>
                                <h3>{t("rekening")}</h3>
                                <p>{data?.data.noRek}</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}