import { useRouter } from "next/router";
import { useEffect, useState } from 'react';
import { getSession, signIn } from 'next-auth/react';
import jwt from "jsonwebtoken";
import UserProfile from './UserProfile'

export default function Profile({ user }) {
    const [data, setData] = useState(null);
    const [error] = useState(null);
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
        } catch (error) {
            console.error();
        }
    }

    return (
        <div className='content'>
            <UserProfile data={data} error={error} />
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
