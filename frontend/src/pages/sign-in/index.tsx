import { ThemeProvider } from "@emotion/react";
import { LockOutlined } from "@mui/icons-material";
import {
  Avatar,
  Box,
  Button,
  Container,
  createTheme,
  CssBaseline,
  Grid,
  Link,
  TextField,
  Typography,
} from "@mui/material";
import { FormEvent, useEffect, VFC } from "react";
import { Controller, useForm } from "react-hook-form";
import { useAuth } from "../../contexts/AuthContext";
import { Copyright } from "../../components";

export const SignIn: VFC = () => {
  const { signIn } = useAuth();

  type FormType = {
    email: string;
    password: string;
  };

  const {
    // register,
    handleSubmit,
    setValue,
    // formState: { errors },
    reset,
    control,
  } = useForm<FormType>();

  useEffect(() => {
    setValue("email", "matsuokashuheiii@gmail.com");
    setValue("password", "test1234");
  }, [setValue]);

  const handleSignIn = async ({ email, password }: FormType) => {
    console.log("email", email);
    console.log("password", password);
    try {
      await signIn(email, password);
    } catch (error) {
      console.error(error);
    }
  };

  const theme = createTheme();

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="sm">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
            <LockOutlined />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <Box component="form" noValidate sx={{ mt: 2 }}>
            <Controller
              name="email"
              control={control}
              render={({ field: { onChange, value } }) => (
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  id="email"
                  label="email"
                  name="email"
                  autoComplete="email"
                  autoFocus
                  onChange={onChange}
                  value={value}
                />
              )}
            />
            <Controller
              name="password"
              control={control}
              render={({ field: { onChange, value } }) => (
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  id="password"
                  label={"password"}
                  name="password"
                  type="password"
                  autoComplete="current-password"
                  onChange={onChange}
                  value={value}
                />
              )}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 4, mb: 2 }}
              onClick={handleSubmit(handleSignIn)}
            >
              Sign in
            </Button>
            <Grid container justifyContent="space-between">
              <Grid item>
                <Link href="#" variant="body2">
                  Forgot password?
                </Link>
              </Grid>
              <Grid item>
                <Link href="/sign-up" variant="body2">
                  {"Don't have an account? Sign up"}
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
        <Copyright />
      </Container>
    </ThemeProvider>
  );
};
