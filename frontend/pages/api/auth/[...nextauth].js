import NextAuth from 'next-auth'
import CredentialProvider from 'next-auth/providers/credentials'
import jwt from "jsonwebtoken"

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
  // secret: "simiddleman",
  // jwt: {
  //   encode: async ({ secret, token }) => {
  //     return jwt.sign(token as any, secret, { algorithm: "HS256" });
  //   },
  //   decode: async ( {secret, token} ) => {
  //     return jwt.verify(token as string, secret, {
  //       algorithms: ["HS256"],
  //     }) as any;
  //   },
  // },
})