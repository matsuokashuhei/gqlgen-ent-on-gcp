import { Group, Home, Logout, Person, School } from "@mui/icons-material";
import {
  Box,
  List,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import { VFC } from "react";
import { Link } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";

export const SideBar: VFC = () => {
  const { signOut } = useAuth();

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
      }}
    >
      <Box>
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
              <ListItemText primary="SIGN OUT" />
            </ListItemButton>
          </Link>
        </List>
      </Box>
      <Box>
        <ListItemButton component="button" onClick={signOut}>
          <ListItemIcon>
            <Logout />
          </ListItemIcon>
          <ListItemText primary="SIGN OUT" />
        </ListItemButton>
      </Box>
    </Box>
  );
};
