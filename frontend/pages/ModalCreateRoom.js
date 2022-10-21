import { Form, Modal, Button } from 'react-bootstrap';
import logo from './assets/logo.png';
import Swal from 'sweetalert2';
import useTranslation from 'next-translate/useTranslation';
import { useRouter } from "next/router";

export default function ModalCreateRoom({ idPenjual, sessionToken, closeCreateRoomModal, createRoomModal, GetAllRoom }) {
  const { t, lang } = useTranslation('createRoom');
  const router = useRouter();

  const handleSubmitCreateRoom = async (e) => {
    closeCreateRoomModal();
    e.preventDefault();

    const formData = new FormData(e.currentTarget);
    const body = {
      id: idPenjual,
      product: {
        nama: formData.get("namaProduk"),
        deskripsi: formData.get("deskripsiProduk"),
        harga: parseInt(formData.get("hargaProduk")),
        kuantitas: parseInt(formData.get("kuantitasProduk")),
      }
    }
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/rooms`, {
        method: 'POST',
        body: JSON.stringify(body),
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ' + sessionToken,
        }
      });
      const data = await res.json();

      if (data?.message === "Success Create Room" || data?.message === "Berhasil membuat ruangan") {
        Swal.fire({ icon: 'success', title: t("successCreate"), text: data.message, showConfirmButton: false, timer: 1500, })
        GetAllRoom();
      } else {
        Swal.fire({ icon: 'error', title: t("failCreate"), text: data.message, })
      }
    }
    catch (error) {
      console.log(error);
    }
  }
  return (
    <Modal show={createRoomModal} onHide={closeCreateRoomModal}
      aria-labelledby="contained-modal-title-vcenter"
      data-testid="ModalCreateRoom"
      centered>
      <Modal.Header data-testid="ModalHeader" closeButton>
        <div className="avatar" data-testid="avatar">
          <img src={logo.src} alt="logo SiMiddleman+" data-testid="logo" />
        </div>
        <Modal.Title className="ms-auto" data-testid="title">{t("modalTitle")}</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form onSubmit={handleSubmitCreateRoom} id="createRoomForm">
          <Form.Group className="mb-3">
            <Form.Control
              type="text"
              placeholder={t("placeholder.0")}
              data-testid="namaProduk"
              name="namaProduk"
              autoFocus
            />
          </Form.Group>
          <Form.Group className="mb-3">
            <Form.Control
              type="number"
              placeholder={t("placeholder.1")}
              data-testid="hargaProduk"
              name="hargaProduk"
              min={1}
              autoFocus
            />
          </Form.Group>
          <Form.Group className="mb-3">
            <Form.Control
              type="number"
              placeholder={t("placeholder.2")}
              data-testid="kuantitasProduk"
              name="kuantitasProduk"
              min={1}
              autoFocus
            />
          </Form.Group>
          <Form.Group className="mb-3">
            <Form.Control
              placeholder={t("placeholder.3")}
              as="textarea"
              rows={5}
              data-testid="deskripsiProduk"
              name="deskripsiProduk"
              autoFocus
            />
          </Form.Group>
          <Button variant='merah'
            className='w-100'
            type='submit'
            form='createRoomForm'
            data-testid="buttonCreateRoom">{t("createRoomBtnModal")}</Button>
        </Form>
      </Modal.Body>
    </Modal>
  )
}