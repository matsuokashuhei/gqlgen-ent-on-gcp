import { FC } from "react";
import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";

export const Public: FC<{}> = ({ children }) => {
  const { currentUser } = useAuth();
  const location = useLocation();
  if (currentUser) {
    return <Navigate to="/" state={{ from: location }} replace />;
  } else {
    return <>{children}</>;
  }
};
