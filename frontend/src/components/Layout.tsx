import { FC } from "react";
import { Outlet } from "react-router-dom";
import { SideBar } from ".";

export const Layout: FC = () => {
  return (
    <div className="flex min-h-screen flex-row">
      <div>
        <SideBar />
      </div>
      <div className="flex-grow">
        <Outlet />
      </div>
    </div>
  );
};
