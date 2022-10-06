import { SessionProvider } from "next-auth/react"
import '../styles/globals.css'
import 'bootstrap/dist/css/bootstrap.min.css';

// add bootstrap css 
import 'bootstrap/dist/css/bootstrap.css'

function App({
  Component,
  pageProps: { session, ...pageProps },
}) {
  return (
    <SessionProvider session={session}>
      <Component {... pageProps} />
    </SessionProvider>
  );
}

export default App
