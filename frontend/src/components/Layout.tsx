import { FC } from "react";
import { SideBar } from ".";

export const Layout: FC = ({ children }) => {
  return (
    <div className="flex min-h-screen flex-row">
      <div>
        <SideBar />
      </div>
      <div className="flex-grow">{children}</div>
    </div>
  );
};
