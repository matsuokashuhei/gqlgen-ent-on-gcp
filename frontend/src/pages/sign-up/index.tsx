import { gql, useMutation } from "@apollo/client";
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
  ThemeProvider,
  Typography,
} from "@mui/material";
import { VFC } from "react";
import { Controller, useForm } from "react-hook-form";
import { Copyright } from "../../components";
import { useAuth } from "../../contexts/AuthContext";

// type State = {
//   email: string;
//   password: string;
//   isSignUpButtonDisabled: boolean;
// };

// const initialState: State = {
//   email: "",
//   password: "",
//   isSignUpButtonDisabled: true,
// };

// type Action =
//   | { type: "setEmail"; payload: string }
//   | { type: "setPassword"; payload: string }
//   | { type: "setIsSignUpButtonDisabled"; payload: boolean };
// //   | { type: "signupSucceeded"; payload: string }
// //   | { type: "signUpFailed"; payload: string };

const theme = createTheme();

// const reducer = (state: State, action: Action): State => {
//   switch (action.type) {
//     case "setEmail":
//       return {
//         ...state,
//         email: action.payload,
//       };
//     case "setPassword":
//       return {
//         ...state,
//         password: action.payload,
//       };
//     case "setIsSignUpButtonDisabled":
//       return {
//         ...state,
//         isSignUpButtonDisabled: action.payload,
//       };
//     default:
//       return state;
//   }
// };

const SIGN_UP = gql`
  mutation SignUp($input: SignUpInput!) {
    signUp(input: $input) {
      id
      firebaseUid
    }
  }
`;

// https://github.com/mui/material-ui/blob/v5.6.4/docs/data/material/getting-started/templates/sign-up/SignUp.tsx
export const SignUp: VFC = () => {
  type FormType = {
    email: string;
    password: string;
  };
  const {
    // register,
    handleSubmit,
    // formState: { errors },
    control,
  } = useForm<FormType>();
  // const [state, dispatch] = useReducer(reducer, initialState);
  const { signIn } = useAuth();
  const [signUp, { loading, error }] = useMutation(SIGN_UP);

  // const { email, password, isSignUpButtonDisabled } = state;

  // useEffect(() => {
  //   console.log(email.length > 0 && password.length);
  //   dispatch({
  //     type: "setIsSignUpButtonDisabled",
  //     payload: email.length === 0 || password.length === 0,
  //   });
  // }, [email, password]);

  const handleSignUp = async ({ email, password }: FormType) => {
    try {
      await signUp({
        variables: { input: { email: email, password: password } },
      });
      await signIn(email, password);
    } catch (error) {
      console.error(error);
    }
  };

  if (loading) return <>Submitting...</>;
  if (error) return <>{error.message}</>;

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
            Sign up
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
                  label="Email"
                  name="email"
                  autoComplete="email"
                  autoFocus
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
                  name="password"
                  label="Password"
                  type="password"
                  id="password"
                  autoComplete="new-password"
                />
              )}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              onClick={handleSubmit(handleSignUp)}
              sx={{ mt: 4, mb: 2 }}
            >
              Sign up
            </Button>
            <Grid container justifyContent="flex-end">
              <Grid item>
                <Link href="/sign-in" variant="body2">
                  Already have an account? Sign in
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
