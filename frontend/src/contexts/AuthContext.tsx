import { createContext, FC, useContext, useEffect, useState } from "react";
import { initializeApp } from "firebase/app";
import {
  getAuth,
  User,
  signInWithEmailAndPassword,
  createUserWithEmailAndPassword,
  UserCredential,
} from "firebase/auth";
import { config } from "../firebase";

initializeApp(config);

type AuthContextType = {
  currentUser: User | null;
  signUp: (emall: string, password: string) => Promise<UserCredential>;
  signIn: (email: string, password: string) => Promise<UserCredential>;
  signOut: () => Promise<void>;
};

const AuthContext = createContext<AuthContextType | null>(null);

export const AuthProvider: FC = ({ children }) => {
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  //   const [loading, setLoading] = useState(true);
  const auth = getAuth();

  const signUp = (email: string, password: string) => {
    return createUserWithEmailAndPassword(auth, email, password);
  };

  const signIn = (email: string, password: string) => {
    return signInWithEmailAndPassword(auth, email, password);
  };

  const signOut = () => {
    return auth.signOut();
  };

  useEffect(() => {
    auth.onAuthStateChanged((user) => setCurrentUser(user));
  }, [auth]);

  return (
    <AuthContext.Provider
      value={{ currentUser, signUp, signIn, signOut }}
      children={children}
    />
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (context === null) {
    throw new Error();
  }
  return context;
};
