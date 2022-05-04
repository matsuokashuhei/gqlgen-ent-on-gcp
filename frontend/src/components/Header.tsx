import { AppBar, Toolbar } from "@mui/material";
import { VFC } from "react";

export const Header: VFC = () => {
  return (
    <AppBar position="static">
      <Toolbar variant="regular"></Toolbar>
    </AppBar>
  );
};
