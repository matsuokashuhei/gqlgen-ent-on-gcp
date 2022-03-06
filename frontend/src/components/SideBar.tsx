import {
  CalendarIcon,
  HomeIcon,
  UserIcon,
  UsersIcon,
} from "@heroicons/react/solid";
import { VFC } from "react";
import { Link } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";

export const SideBar: VFC = () => {
  const { signOut } = useAuth();

  const renderSignOutButton = () => <button onClick={signOut}>Sign Out</button>;

  return (
    <div className="text-md relative flex h-full w-36 flex-col bg-slate-500 text-slate-50">
      <div className="flex items-center gap-1">
        <HomeIcon className="h-5" />
        <div>ホーム</div>
      </div>
      <div className="flex items-center gap-1">
        <CalendarIcon className="h-5" />
        <div>
          <Link to="/classes">クラス</Link>
        </div>
      </div>
      <div className="flex items-center gap-1">
        <UsersIcon className="h-5" />
        <div>会員</div>
      </div>
      <div className="flex items-center gap-1">
        <UserIcon className="h-5" />
        <div>
          <Link to="/instructors?first=10">インストラクター</Link>
        </div>
      </div>
      <div className="absolute bottom-0 left-0">{renderSignOutButton()}</div>
    </div>
  );
};
