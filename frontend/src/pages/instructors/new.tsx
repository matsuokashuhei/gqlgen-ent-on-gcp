import { VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { Layout } from "../../components";
import {
  CreateInstructorInput,
  useCreateInstructorMutation,
} from "../../generated/graphql";

type Inputs = {
  name: string;
  syllabicCharacters: string;
  biography: string;
  phoneNumber: string;
  email: string;
};

export const NewInstructorPage: VFC = () => {
  const navigate = useNavigate();
  const { register, handleSubmit } = useForm<Inputs>();
  const [createInstructor, { data, loading, error }] =
    useCreateInstructorMutation();
  const onSubmit: SubmitHandler<Inputs> = (data) => {
    const { name, syllabicCharacters, biography, phoneNumber, email } = data;
    const input: CreateInstructorInput = {
      name,
      syllabicCharacters,
      biography,
      phoneNumber,
      email,
    };
    createInstructor({ variables: { input } }).then((instructor) =>
      navigate("/instructors")
    );
  };

  return (
    <Layout>
      <form className="flex flex-col">
        <input {...register("name", { required: true })} />
        <input {...register("syllabicCharacters", { required: true })} />
        <input {...register("biography")} />
        <input {...register("phoneNumber")} />
        <input {...register("email")} />
      </form>
      <button type="submit" onClick={handleSubmit(onSubmit)}>
        登録
      </button>
    </Layout>
  );
};
