import useTranslation from 'next-translate/useTranslation';
import { Button } from 'react-bootstrap';
import ModalUpdateProfile from './ModalUpdateProfile'

export default function UserProfile( props ) {
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
                    {props.error && <div>{t("load-fail")} {props.error.toString()}</div>}
                    {!props.data ? <div>{t("load.loading")}</div> : ((props.data?.data ?? []).length === 0 && <p className='text-xl p-8 text-center text-gray-100'>{t("list-empty")}</p>)}
                    <h2 className='nama-produk'>
                        {props.data?.data.nama}
                    </h2>
                    <div className="py-3">
                        <div className='d-flex flex-row justify-content-between'>
                            <div>
                                <h3>{t("email")}</h3>
                                <p>{props.data?.data.email}</p>
                            </div>
                            <div>
                                <h3>{t("handphone")}</h3>
                                <p>{props.data?.data.noHp}</p>
                            </div>
                        </div>
                        <div className='mt-4'>
                            <div className='col-lg-4 col-md-6 col-sm-8'>
                                <h3>{t("rekening")}</h3>
                                <p>{props.data?.data.noRek}</p>
                            </div>
                        </div>
                    </div>
                    <Button variant="simiddleman" onClick={props.openUpdateProfileModal}>{t("modal.btnUpdateProfile")}</Button>
                    <ModalUpdateProfile updateProfileModal={props.updateProfileModal} closeUpdateProfileModal={props.closeUpdateProfileModal} openUpdateProfileModal={props.openUpdateProfileModal} onChangeText={props.onChangeText} handleSubmitUpdateProfile={props.handleSubmitUpdateProfile} nama={props.nama} noHp={props.noHp} noRek={props.noRek}  />
                </div>
            </div>
        </div>
    )
}