import { ThemeProvider } from "@emotion/react";
import { Box, createTheme, CssBaseline } from "@mui/material";
import { FC } from "react";
import { Outlet } from "react-router-dom";
import { Footer, Header, SideBar } from ".";

const theme = createTheme();

export const Layout: FC = () => {
  return (
    <ThemeProvider theme={theme}>
      <Box sx={{ display: "flex", flexDirection: "column", height: "100vh" }}>
        <CssBaseline />
        <Header />
        <Box sx={{ display: "flex", flexDirection: "row", flexGrow: "1" }}>
          <SideBar />
          <Box sx={{ color: "inherit" }}>
            <Outlet />
          </Box>
        </Box>
        <Footer />
      </Box>
    </ThemeProvider>
  );
};
