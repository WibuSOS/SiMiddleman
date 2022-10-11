import { useRef, useState } from 'react';
import { Form, Modal, Button } from 'react-bootstrap';
import logo from './assets/logo.png';
import Swal from 'sweetalert2';

export default function ModalShowRoomCode(props) {
  return (
    <Modal show={props.showRoomCodeModal} onHide={props.closeShowRoomCodeModal}
      aria-labelledby="contained-modal-title-vcenter" centered>
      <Modal.Header data-testid="ModalHeader" closeButton>
        <div className="avatar" data-testid="avatar">
          <img src={logo.src} alt="logo SiMiddleman+" data-testid="logo" />
        </div>
        <Modal.Title className="ms-auto mt-4" data-testid="title">Room Code</Modal.Title>
      </Modal.Header>
      <Modal.Body className='mx-auto text-center'>
        <div className='roomCode' data-testid="roomCode">
          <p><strong>{props.roomCode}</strong></p>
        </div>
        <Button data-testid="buttonSalin" onClick={() => {
          navigator.clipboard.writeText(props.roomCode);
          Swal.fire({
            icon: 'success',
            title: 'Kode ruangan berhasil disalin',
            showConfirmButton: false,
            timer: 1500,
        })
          }}>Salin Kode Ruangan</Button>
      </Modal.Body>
    </Modal>
  )
}