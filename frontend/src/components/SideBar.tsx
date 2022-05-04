import { Group, Home, Person, School } from "@mui/icons-material";
import {
  List,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import { VFC } from "react";
import { Link, useNavigate } from "react-router-dom";

export const SideBar: VFC = () => {
  const navigate = useNavigate();
  /*
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
        <div>
          <Link to="/members">会員</Link>
        </div>
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
  */
  return (
    <List component="nav">
      <Link to="/">
        <ListItemButton>
          <ListItemIcon>
            <Home />
          </ListItemIcon>
          <ListItemText primary="HOME" />
        </ListItemButton>
      </Link>
      <Link to="/classes">
        <ListItemButton>
          <ListItemIcon>
            <School />
          </ListItemIcon>
          <ListItemText primary="CLASS" />
        </ListItemButton>
      </Link>
      <Link to="/members">
        <ListItemButton>
          <ListItemIcon>
            <Group />
          </ListItemIcon>
          <ListItemText primary="MEMBER" />
        </ListItemButton>
      </Link>
      <Link to="/instructors">
        <ListItemButton>
          <ListItemIcon>
            <Person />
          </ListItemIcon>
          <ListItemText primary="INSTRUCTOR" />
        </ListItemButton>
      </Link>
    </List>
  );
};
