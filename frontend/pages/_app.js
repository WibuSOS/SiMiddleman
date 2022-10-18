import { SessionProvider } from "next-auth/react"
import '../styles/globals.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import Navbar from './NavigationBar';
import Footer from "./Footer";
import Head from 'next/head';
import Link from "next/link";

function App({
  Component,
  pageProps: { session, ...pageProps },
}) {
  return (
    <>
      <Navbar />
      <SessionProvider session={session}>
        <Component {...pageProps} />
      </SessionProvider>
      <Footer />
    </>
  );
}

export default App
