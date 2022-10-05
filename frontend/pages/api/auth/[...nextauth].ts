import NextAuth from 'next-auth'
import CredentialProvider from 'next-auth/providers/credentials'
import { Alert } from 'react-bootstrap';
import Swal from 'sweetalert2'

export default NextAuth ({
    providers: [
        CredentialProvider({
            name: "credentials",
            credentials: {
              email: { 
                  label: "Email", 
                  type: "email", 
                  placeholder: "admin@localhost.com"},
              password: { 
                  label: "Password", 
                  type: "password"},
          },
            authorize : async (credentials, req) => {
                const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/login`, {
                  method: 'POST',
                  body: JSON.stringify(credentials),
                  headers: { "Content-Type": "application/json" }
                });
                const user = await res.json();
                if (res.ok && user) {
                  return user
                }
                return null;
            },
            
        }),
    ],
    pages: {
        signIn: "/Login",
    },
    callbacks: {
        jwt: ({ token, user }) => {
            // first time jwt callback is run, user object is available
            if (user) {
                token.id = user.id;
            }

            return token;
        },
        session: ( { session, token }) => {
            if (token) {
                session.id = token.id;
            }

            return session;
        },
    },
    secret: "test",
    jwt: {
        secret: "test",
    },
})