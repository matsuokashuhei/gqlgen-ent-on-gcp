import { Typography } from "@mui/material";
import { VFC } from "react";

// export const Copyright: VFC = (props: SxProps<Theme>) => {
export const Copyright: VFC = (props: any) => {
  return (
    <Typography
      variant="body2"
      color="text.secondary"
      align="center"
      {...props}
    >
      {"Copyright © Shuhei Matsuoka "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
};
