import { Menu, Notifications } from "@mui/icons-material";
import { AppBar, IconButton, Toolbar, Typography } from "@mui/material";
import { VFC } from "react";

export const Header: VFC = () => {
  return (
    <AppBar position="static">
      <Toolbar variant="regular"></Toolbar>
    </AppBar>
  );
};
