import { useEffect, VFC } from "react";
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
    setValue,
    formState: { errors },
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
                  className="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500"
                  placeholder="Email"
                />
              </div>
              <div>
                <input
                  id="password"
                  type="password"
                  {...register("password", { required: true })}
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
                onClick={handleSubmit(handleSignIn)}
                className="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
              >
                Sign in
              </button>
            </div>
          </form>
          <div className="flex items-center justify-between">
            <div className="flex items-center">
              <div className="text-sm">
                <a
                  href="/sign-up"
                  className="font-medium text-indigo-600 hover:text-indigo-500"
                >
                  Sign up
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
