import { VFC } from "react";
import { useAuth } from "../contexts/AuthContext";

export const Home: VFC = () => {
  const { signOut } = useAuth();

  const handleSignOut = () => {
    signOut();
  };

  return (
    <>
      <h1>Home</h1>
      <button onClick={handleSignOut}>Sign Out</button>
    </>
  );
};
