import NextAuth from 'next-auth'
import CredentialProvider from 'next-auth/providers/credentials'

export default NextAuth ({
  providers: [
    CredentialProvider({
      name: "credentials",
      credentials: {
        email: { 
          label: "Email", 
          type: "email",},
        password: { 
          label: "Password", 
          type: "password",},
      },
      authorize : async (credentials, req) => {
        const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/login`, {
          method: 'POST',
          body: JSON.stringify(credentials),
          headers: { "Content-Type": "application/json" }
        });
        const user = await res.json();
        if (res.ok && user) {
          return user.token;
        }
        return null;
      },
    }),
  ],
  pages: {
    signIn: "/Login",
  },
  callbacks: {
    async session({ session, token }) {
      session.user = token.user;
      return session;
    },
    async jwt ( {token, user} ) {
      if (user) {
        token.user = user;
      }
      return token;
    }
  },
  secret: process.env.NEXTAUTH_SECRET,
})