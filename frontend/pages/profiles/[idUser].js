import { useRouter } from "next/router";
import { useEffect, useState } from 'react';
import { getSession, signIn } from 'next-auth/react';
import jwt from "jsonwebtoken";
import UserProfile from './UserProfile';
import Swal from 'sweetalert2';

export default function Profile({ user }) {
    const [data, setData] = useState(null);
    const [error] = useState(null);
    const [nama, setNama] = useState("");
    const [noHp, setNoHp] = useState("");
    const [noRek, setNoRek] = useState("");
    const [updateProfileModal, setupdateProfileModal] = useState(false);
    const openUpdateProfileModal = () => setupdateProfileModal(true);
    const closeUpdateProfileModal = () => setupdateProfileModal(false);
    const router = useRouter();

    useEffect(() => {
        getUserDetails();
    }, [])
    const decoded = jwt.verify(user, process.env.NEXT_PUBLIC_JWT_SECRET);

    const getUserDetails = async () => {
        const idUser = decoded.ID;
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/user/${idUser}`, {
                method: 'GET',
                headers: {
                'Authorization': 'Bearer ' + user,
                }
            });
            const data = await res.json();
            setData(data);
            setNama(data?.data.nama);
            setNoHp(data?.data.noHp);
            setNoRek(data?.data.noRek);
        } catch (error) {
            console.error();
        }
    }

    const onChangeText = (e, type) => {
        if (type === "nama"){
            setNama(e.target.value);
        }
        if (type === "noHp"){
            setNoHp(e.target.value);
        }
        if (type === "noRek"){
            setNoRek(e.target.value);
        }
    }

    const handleSubmitUpdateProfile = async (e) => {
        closeUpdateProfileModal();
        e.preventDefault();
        const idUser = decoded.ID;
        const email = decoded.Email;
    
        const body = {
            Nama: nama,
            NoHp: noHp,
            Email: email,
            NoRek: noRek
        }
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${router.locale}/user/${idUser}`, {
                method: 'PUT',
                body: JSON.stringify(body),
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + user,
                }
            });
            const dataRes = await res.json();
    
            if (dataRes?.message === "success") {
                Swal.fire({ icon: 'success', title: 'Profile Berhasil diupdate', text: dataRes?.message, showConfirmButton: false, timer: 1500, })
                getUserDetails();
            } else {
                Swal.fire({ icon: 'error', title: 'Profile Gagal Diupdate', text: dataRes?.message, })
            }
        }
        catch (error) {
            console.log(error);
        }
    }

    return (
        <div className='content'>
            <UserProfile data={data} error={error} updateProfileModal={updateProfileModal} closeUpdateProfileModal={closeUpdateProfileModal} openUpdateProfileModal={openUpdateProfileModal} onChangeText={onChangeText} handleSubmitUpdateProfile={handleSubmitUpdateProfile} nama={nama} noHp={noHp} noRek={noRek} />
        </div >
    )
}

export async function getServerSideProps(ctx) {
    const session = await getSession(ctx)
    if (!session) {
        return {
        redirect: { permanent: false, destination: "/api/auth/signin" }
        }
    }
    const { user } = session;
    return {
        props: { user },
    }
}
