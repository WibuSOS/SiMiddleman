import Head from 'next/head';
import { useState, useEffect } from 'react';
import Modal from 'react-modal';

const customStyles = {
  content: {
    top: '50%',
    left: '50%',
    right: 'auto',
    bottom: 'auto',
    marginRight: '-50%',
    transform: 'translate(-50%, -50%)',
    width: '50%'
  },
};

function ShowRegisterModal(props){
  return(
    <Modal
      {...props}>
        <div className={"w-100"}>
            <div className={"modal-header"}>
                <h4 className={"modal-title"}>Register</h4>
                <button className={"btn btn-close"} onClick={props.closeModal}></button>
            </div>
            <div className={"w-100 d-flex flex-row justify-content-between my-1"}>
              <form className={"mx-1"} onSubmit={() => console.log("Submit")}>
                <input className='mx-1' name="name" type="text" placeholder='Nama' />
                <br />
                <input className='mx-1' name="email" type="email" placeholder='Email' />
                <br />
                <input className='mx-1' name="password" type="password" placeholder='Password' />
                <br />
                <input className='mx-1' name="confirm-password" type="password" placeholder='Confirm Password' />
                <br />
                <button className='btn btn-danger'>Submit</button>
              </form>
            </div>    
        </div>
    </Modal>
  )
}

export default function Home() {
  const [registerModal, setRegisterModal] = useState(false);
  useEffect(() => {
    getData()
  }, []);

  const getData = async () => {
    
  }

  const openRegisterModal = () =>{
    setRegisterModal(true)
  }

  const closeRegisterModal = () =>{
    setRegisterModal(false)
  }

  return (
    <div className='container mx-10 my-7'>
      <button className='btn btn-primary' onClick={() => openRegisterModal()}>
        Register
      </button>
      
      <ShowRegisterModal
        isOpen={registerModal}
        onRequestClose={closeRegisterModal}
        contentLabel="Register Modal" 
        closeModal={closeRegisterModal}
        style={customStyles} />
    </div>
  )
}