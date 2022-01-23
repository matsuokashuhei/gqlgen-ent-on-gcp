import { VFC } from "react";
import { useForm } from "react-hook-form";
import { useAuth } from "../../contexts/AuthContext";

export const SignIn: VFC = () => {
  const { signIn } = useAuth();

  type FormType = {
    email: string;
    password: string;
  };

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormType>();

  const handleSignIn = async ({ email, password }: FormType) => {
    console.log("email", email);
    console.log("password", password);
    try {
      await signIn(email, password);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <form>
      <input type="email" {...register("email", { required: true })} />
      <input type="password" {...register("password", { required: true })} />
      <button type="submit" onClick={handleSubmit(handleSignIn)}>
        Sign In
      </button>
    </form>
  );
};
