import Swal from 'sweetalert2';
import ModalUpdateProduct from './ModalUpdateProduct';
import { useRouter } from 'next/router';

export default function UpdateProduct({ closeUpdateProductModal, updateProductModal, data, user, namaProduk, setNamaProduk, hargaProduk, setHargaProduk, deskripsiProduk, setDeskripsiProduk, kuantitasProduk, setKuantitasProduk, getRoomDetails }) {
    const router = useRouter();

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
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/updateproduct/${data?.data.product.ID}`, {
                method: 'PUT',
                body: JSON.stringify(body),
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + user,
                }
            });
            const dataRes = await res.json();
    
            if (dataRes?.message === "berhasil mengupdate data") {
                Swal.fire({ icon: 'success', title: 'Data Produk Berhasil Diupdate', showConfirmButton: false, timer: 1500, })
                getRoomDetails();
            } else {
                Swal.fire({ icon: 'error', title: 'Data Produk Gagal Diupdate', text: dataRes?.message, })
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