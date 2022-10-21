import Swal from 'sweetalert2';
import ModalUpdateProduct from './ModalUpdateProduct';
import { useRouter } from 'next/router';
import useTranslation from 'next-translate/useTranslation';

export default function UpdateProduct({ closeUpdateProductModal, updateProductModal, data, user, namaProduk, setNamaProduk, hargaProduk, setHargaProduk, deskripsiProduk, setDeskripsiProduk, kuantitasProduk, setKuantitasProduk, getRoomDetails }) {
    const router = useRouter();
    const { t, lang } = useTranslation('detailProduct');

    const handleSubmitUpdateProduct = async (e) => {
        closeUpdateProductModal();
        e.preventDefault();
    
        const body = {
            Nama: namaProduk,
            Deskripsi: deskripsiProduk,
            Harga: parseInt(hargaProduk),
            Kuantitas: parseInt(kuantitasProduk),
        }
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/rooms/details/updateproduct/${data?.data.product.ID}`, {
                method: 'PUT',
                body: JSON.stringify(body),
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + user,
                }
            });
            const dataRes = await res.json();
    
            if (dataRes?.message === "Success update data product" || dataRes?.message === "Berhasil update data produk") {
                Swal.fire({ icon: 'success', title: t("updateProductModal.successUpdate"), text: dataRes?.message, showConfirmButton: false, timer: 1500, })
                getRoomDetails();
            } else {
                Swal.fire({ icon: 'error', title: t("updateProductModal.failUpdate"), text: dataRes?.message, })
            }
        }
        catch (error) {
            console.log(error);
        }
    }

    const onChangeText = (e, type) => {
        if (type === "namaProduk"){
            setNamaProduk(e.target.value);
        }
        if (type === "kuantitasProduk"){
            setKuantitasProduk(e.target.value);
        }
        if (type === "deskripsiProduk"){
            setDeskripsiProduk(e.target.value);
        }
        if (type === "hargaProduk"){
            setHargaProduk(e.target.value);
        }
    }

    return (
        <ModalUpdateProduct closeUpdateProductModal={closeUpdateProductModal} updateProductModal={updateProductModal} handleSubmitUpdateProduct={handleSubmitUpdateProduct} namaProduk={namaProduk} hargaProduk={hargaProduk} kuantitasProduk={kuantitasProduk} deskripsiProduk={deskripsiProduk} onChangeText={onChangeText} />
    )
}