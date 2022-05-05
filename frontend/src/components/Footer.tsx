import { Box, Typography } from "@mui/material";
import { VFC } from "react";

export const Footer: VFC = () => {
  return (
    <Box sx={{ m: 1 }}>
      <Typography variant="body2" color="text.secondary" align="center">
        {"Copyright © Shuhei Matsuoka "}
        {new Date().getFullYear()}
        {"."}
      </Typography>
    </Box>
  );
};
