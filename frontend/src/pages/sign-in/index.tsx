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
import { useForm } from "react-hook-form";
import { useAuth } from "../../contexts/AuthContext";

export const SignIn: VFC = () => {
  const { signIn } = useAuth();

  type FormType = {
    email: string;
    password: string;
  };

  // const {
  //   register,
  //   handleSubmit,
  //   setValue,
  //   formState: { errors },
  // } = useForm<FormType>();

  // useEffect(() => {
  //   setValue("email", "matsuokashuheiii@gmail.com");
  //   setValue("password", "test1234");
  // }, [setValue]);

  const handleSignIn = async ({ email, password }: FormType) => {
    console.log("email", email);
    console.log("password", password);
    try {
      await signIn(email, password);
    } catch (error) {
      console.error(error);
    }
  };

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    const email = data.get("email")?.toString();
    const password = data.get("password")?.toString();
    if (email && password) {
      handleSignIn({ email, password });
    }
  };

  const theme = createTheme();

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
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
          <Box
            component="form"
            noValidate
            onSubmit={handleSubmit}
            sx={{ mt: 1 }}
          >
            <TextField
              margin="normal"
              required
              fullWidth
              id="email"
              label="Email"
              name="email"
              autoComplete="email"
              autoFocus
            />
            <TextField
              margin="normal"
              required
              fullWidth
              id="password"
              label="Password"
              name="password"
              type="password"
              autoComplete="current-password"
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Sign in
            </Button>
            <Grid container>
              <Grid item xs>
                <Link href="#" variant="body2">
                  Forgot password?
                </Link>
              </Grid>
              <Grid item xs>
                <Link href="/sign-up" variant="body2">
                  {"Don't have an account? Sign up"}
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
    // <div className="h-full">
    //   <div className="flex min-h-full items-center justify-center">
    //     <div className="w-full max-w-md space-y-8">
    //       <div>
    //         <h2 className="text-center font-extrabold text-gray-900">
    //           Sign up
    //         </h2>
    //       </div>
    //       <form className="mt-8 space-y-6">
    //         <div className="-space-y-px rounded-md shadow-sm">
    //           <div>
    //             <input
    //               id="email"
    //               type="email"
    //               {...register("email", { required: true })}
    //               className="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500"
    //               placeholder="Email"
    //             />
    //           </div>
    //           <div>
    //             <input
    //               id="password"
    //               type="password"
    //               {...register("password", { required: true })}
    //               className="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500"
    //               placeholder="Password"
    //             />
    //           </div>
    //         </div>
    //         <div className="flex items-center justify-between">
    //           <div className="flex items-center">
    //             <div className="text-sm">
    //               <a
    //                 href="/fogot-password"
    //                 className="font-medium text-indigo-600 hover:text-indigo-500"
    //               >
    //                 Forgot your password?
    //               </a>
    //             </div>
    //           </div>
    //         </div>
    //         <div>
    //           <button
    //             type="submit"
    //             onClick={handleSubmit(handleSignIn)}
    //             className="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
    //           >
    //             Sign in
    //           </button>
    //         </div>
    //       </form>
    //       <div className="flex items-center justify-between">
    //         <div className="flex items-center">
    //           <div className="text-sm">
    //             <a
    //               href="/sign-up"
    //               className="font-medium text-indigo-600 hover:text-indigo-500"
    //             >
    //               Sign up
    //             </a>
    //           </div>
    //         </div>
    //       </div>
    //     </div>
    //   </div>
    // </div>
  );
};
