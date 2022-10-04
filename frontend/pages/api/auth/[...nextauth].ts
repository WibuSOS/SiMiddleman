import NextAuth from 'next-auth'
import CredentialProvider from 'next-auth/providers/credentials'

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
            authorize: (credentials) => {
                // database lookup
                if (credentials.email === "admin@localhost.com" && credentials.password == "admin") {
                    return {
                        id: 1,
                        email: "admin@localhost.com"
                    };
                }
                // login failed
                return null;
            },
            
        }),
    ],
    callbacks: {
        jwt: () => {},
        session: () => {},
    },
    secret: "test",
    jwt: {
        secret: "test",
        encryption: true,
    }
})