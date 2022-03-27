import { gql, useMutation } from "@apollo/client";
import { ChangeEventHandler, useEffect, useReducer, VFC } from "react";
import { useForm } from "react-hook-form";
import { useAuth } from "../../contexts/AuthContext";

type State = {
  email: string;
  password: string;
  isSignUpButtonDisabled: boolean;
};

const initialState: State = {
  email: "",
  password: "",
  isSignUpButtonDisabled: true,
};

type Action =
  | { type: "setEmail"; payload: string }
  | { type: "setPassword"; payload: string }
  | { type: "setIsSignUpButtonDisabled"; payload: boolean };
//   | { type: "signupSucceeded"; payload: string }
//   | { type: "signUpFailed"; payload: string };

const reducer = (state: State, action: Action): State => {
  switch (action.type) {
    case "setEmail":
      return {
        ...state,
        email: action.payload,
      };
    case "setPassword":
      return {
        ...state,
        password: action.payload,
      };
    case "setIsSignUpButtonDisabled":
      return {
        ...state,
        isSignUpButtonDisabled: action.payload,
      };
    default:
      return state;
  }
};

const SIGN_UP = gql`
  mutation SignUp($input: SignUpInput!) {
    signUp(input: $input) {
      id
      firebaseUid
    }
  }
`;

export const SignUp: VFC = () => {
  type FormType = {
    email: string;
    password: string;
  };
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormType>();
  const [state, dispatch] = useReducer(reducer, initialState);
  const { signIn } = useAuth();
  const [signUp, { loading, error }] = useMutation(SIGN_UP);

  const { email, password, isSignUpButtonDisabled } = state;

  useEffect(() => {
    console.log(email.length > 0 && password.length);
    dispatch({
      type: "setIsSignUpButtonDisabled",
      payload: email.length === 0 || password.length === 0,
    });
  }, [email, password]);

  const handleEmailChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    dispatch({
      type: "setEmail",
      payload: event.target.value,
    });
  };

  const handlePasswordChange: ChangeEventHandler<HTMLInputElement> = (
    event
  ) => {
    dispatch({
      type: "setPassword",
      payload: event.target.value,
    });
  };

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
    <div className="h-full">
      <div className="flex min-h-full items-center justify-center">
        <div className="w-full max-w-md space-y-8">
          <div>
            <h2 className="text-center font-extrabold text-gray-900">
              Sign up
            </h2>
          </div>
          <form className="mt-8 space-y-6">
            <div className="-space-y-px rounded-md shadow-sm">
              <div>
                <input
                  id="email"
                  type="email"
                  {...register("email", { required: true })}
                  onChange={handleEmailChange}
                  className="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500"
                  placeholder="Email"
                />
              </div>
              <div>
                <input
                  id="password"
                  type="password"
                  {...register("password", { required: true })}
                  onChange={handlePasswordChange}
                  className="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500"
                  placeholder="Password"
                />
              </div>
            </div>
            <div className="flex items-center justify-between">
              <div className="flex items-center">
                <div className="text-sm">
                  <a
                    href="/fogot-password"
                    className="font-medium text-indigo-600 hover:text-indigo-500"
                  >
                    Forgot your password?
                  </a>
                </div>
              </div>
            </div>
            <div>
              <button
                type="submit"
                disabled={isSignUpButtonDisabled}
                onClick={handleSubmit(handleSignUp)}
                className="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
              >
                Sign Up
              </button>
            </div>
          </form>
          <div className="flex items-center justify-between">
            <div className="flex items-center">
              <div className="text-sm">
                <a
                  href="/sign-in"
                  className="font-medium text-indigo-600 hover:text-indigo-500"
                >
                  Sign in
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
