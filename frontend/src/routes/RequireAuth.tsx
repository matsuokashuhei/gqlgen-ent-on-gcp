import { FC } from "react";
import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";

export const RequiredAuth: FC<{}> = ({ children }) => {
  const { currentUser } = useAuth();
  const location = useLocation();
  if (currentUser) {
    return <>{children}</>;
  }
  return <Navigate to="/sign-in" state={{ from: location }} replace />;
};
