import { getSession } from "next-auth/react";
import jwt from "jsonwebtoken";
import { useEffect, useState } from 'react';
import WelcomeBanner from './WelcomeBanner';
import AlasanSimiddleman from './AlasanSimiddleman';
import SimiddlemanSummaries from './SimiddlemanSummaries';
import UserAction from './UserAction';
import ShowRoomList from './ShowRoomList';
import UserBanner from './UserBanner';

function Home({ user }) {
  const [data, setData] = useState(null);
  
  useEffect(() => {
    if (user) GetAllRoom();
  }, [user])
  
  if (!user) return (
    <div className='content'>
      <WelcomeBanner />
      <AlasanSimiddleman />
      <SimiddlemanSummaries />
    </div>
  )
  
  const decoded = jwt.verify(user, process.env.NEXT_PUBLIC_JWT_SECRET);
  const GetAllRoom = async () => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/rooms/${decoded.ID}`, {
        method: 'GET',
        headers: {
          'Authorization': 'Bearer ' + user
        },
      });
      const data = await res.json();
      setData(data);
    } catch (error) {
      console.error();
    }
  }

  return (
    <div className='content'>
      <UserBanner decoded={decoded}/>
      <UserAction decoded={decoded} user={user}/>
      <ShowRoomList data={data} decoded={decoded}/>
    </div>
  )
}

export async function getServerSideProps(ctx) {
  const session = await getSession(ctx)
  if (!session) {
    return {
      props: {}
    }
  }
  const { user } = session;
  return {
    props: { user },
  }
}

export default Home;
