import { VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import {
  CreateInstructorInput,
  useCreateInstructorMutation,
} from "../../generated/graphql";

type Inputs = {
  name: string;
  kana: string;
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
    const { name, kana, biography, phoneNumber, email } = data;
    const input: CreateInstructorInput = {
      name,
      kana,
      biography,
      phoneNumber,
      email,
    };
    createInstructor({ variables: { input } }).then((instructor) =>
      navigate("/instructors")
    );
  };

  return (
    <>
      <form className="flex flex-col">
        <label htmlFor="name">名前</label>
        <input {...register("name", { required: true })} />
        <label htmlFor="syllabicCharacters">よみがな</label>
        <input {...register("kana", { required: true })} />
        <label htmlFor="biography">紹介文</label>
        <input {...register("biography")} />
        <label htmlFor="phoneNumber">電話番号</label>
        <input {...register("phoneNumber")} />
        <label htmlFor="email">メール</label>
        <input {...register("email")} />
      </form>
      <button type="submit" onClick={handleSubmit(onSubmit)}>
        登録
      </button>
    </>
  );
};
