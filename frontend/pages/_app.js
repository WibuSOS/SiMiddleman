import { SessionProvider } from "next-auth/react"
import '../styles/globals.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import Navbar from './NavigationBar';
import Footer from "./Footer";
import Head from 'next/head';
import Link from "next/link";

// add bootstrap css 
import 'bootstrap/dist/css/bootstrap.css'

function App({
  Component,
  pageProps: { session, ...pageProps },
}) {
  return (
    <>
      <Head>
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"/>
      </Head>
      <Navbar/>
      <SessionProvider session={pageProps.session}>
        <Component {... pageProps} />
      </SessionProvider>
      <Footer/>
    </>
  );
}

export default App
