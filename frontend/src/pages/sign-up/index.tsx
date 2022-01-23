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
  const { signUp } = useAuth();

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
      await signUp(email, password);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <form>
      <input
        type="email"
        {...register("email", { required: true })}
        onChange={handleEmailChange}
        className="border"
      />
      <input
        type="password"
        {...register("password", { required: true })}
        onChange={handlePasswordChange}
        className="border"
      />
      <button
        type="submit"
        disabled={isSignUpButtonDisabled}
        onClick={handleSubmit(handleSignUp)}
        className="border"
      >
        Sign Up
      </button>
    </form>
  );
};
