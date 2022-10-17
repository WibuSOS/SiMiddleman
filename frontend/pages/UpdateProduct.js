import { Form, Modal, Button } from 'react-bootstrap';
import Swal from 'sweetalert2';
import logo from './assets/logo.png';

export default function UpdateProduct({ closeUpdateProductModal, updateProductModal, data, user, namaProduk, setNamaProduk, hargaProduk, setHargaProduk, deskripsiProduk, setDeskripsiProduk, kuantitasProduk, setKuantitasProduk, getRoomDetails }) {
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
        <Modal show={updateProductModal} onHide={closeUpdateProductModal}
            aria-labelledby="contained-modal-title-vcenter"
            data-testid="modalRegister"
            centered>
            <Modal.Header closeButton data-testid="modalHeader">
                <div className="avatar">
                    <img src={logo.src} alt="logo SiMiddleman+" data-testid="logo" />
                </div>
                <Modal.Title className="ms-auto" data-testid="title">Update Product</Modal.Title>
            </Modal.Header>

            <Modal.Body data-testid="modalBody">
                <Form onSubmit={handleSubmitUpdateProduct} id="createRoomForm">
                    <Form.Group className="mb-3">
                        <Form.Control
                        type="text"
                        placeholder="Nama Produk"
                        data-testid="namaProduk"
                        name="namaProduk"
                        value={namaProduk}
                        onChange={(e) => onChangeText(e , "namaProduk")}
                        autoFocus
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                        type="number"
                        placeholder="Harga Produk"
                        data-testid="hargaProduk"
                        name="hargaProduk"
                        value={hargaProduk}
                        onChange={(e) => onChangeText(e , "hargaProduk")}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                        type="number"
                        placeholder="Kuantitas Produk"
                        data-testid="kuantitasProduk"
                        name="kuantitasProduk"
                        value={kuantitasProduk}
                        onChange={(e) => onChangeText(e , "kuantitasProduk")}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Control
                        placeholder="Deskripsi Produk"
                        as="textarea"
                        rows={5}
                        data-testid="deskripsiProduk"
                        name="deskripsiProduk"
                        value={deskripsiProduk}
                        onChange={(e) => onChangeText(e , "deskripsiProduk")}
                        />
                    </Form.Group>
                    <Button variant='merah'
                        className='w-100'
                        type='submit'
                        form='createRoomForm'
                        data-testid="buttonCreateRoom">Update Product</Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}