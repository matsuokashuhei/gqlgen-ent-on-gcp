import { FC } from "react";
import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";

export const Private: FC<{}> = ({ children }) => {
  const { currentUser } = useAuth();
  const location = useLocation();
  if (!currentUser) {
    return <Navigate to="/sign-in" state={{ from: location }} replace />;
  } else {
    return <>{children}</>;
  }
};
