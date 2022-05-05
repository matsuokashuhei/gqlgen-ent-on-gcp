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

import { useNavigate } from "react-router-dom";

export const SideBar: VFC = () => {
  const navigate = useNavigate();
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
          <ListItemButton component="button" onClick={() => navigate("/")}>
            <ListItemIcon>
              <Home />
            </ListItemIcon>
            <ListItemText primary="HOME" />
          </ListItemButton>
          <ListItemButton
            component="button"
            onClick={() => navigate("/classes")}
          >
            <ListItemIcon>
              <School />
            </ListItemIcon>
            <ListItemText primary="CLASS" />
          </ListItemButton>
          <ListItemButton
            component="button"
            onClick={() => navigate("/members")}
          >
            <ListItemIcon>
              <Group />
            </ListItemIcon>
            <ListItemText primary="MEMBER" />
          </ListItemButton>
          <ListItemButton
            component="button"
            onClick={() => navigate("/instructors")}
          >
            <ListItemIcon>
              <Person />
            </ListItemIcon>
            <ListItemText primary="INSTRUCTORS" />
          </ListItemButton>
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
