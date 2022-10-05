import Button from 'react-bootstrap/Button';
import { signOut, signIn, useSession } from "next-auth/react";
import CreateRoom from './CreateRoom';
import RegisterForm from './register';

function Home() {    
  const { data: session } = useSession();

  if (session) {
    return (
      <>
        <Button onClick={() => signOut()}>Sign out</Button>
        <CreateRoom/>
      </>
    )
  }
  else {
    return (
      <div className='container mx-10 my-7'>
          <Button variant="primary" onClick={() => signIn()}>
              Login
          </Button>
          <RegisterForm/>
      </div>
    );
  }
}

export default Home;
