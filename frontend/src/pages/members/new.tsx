import { formatRFC3339, parseISO } from "date-fns";
import { VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import {
  CreateMemberInput,
  Gender,
  useCreateMemberMutation,
} from "../../generated/graphql";

type Inputs = {
  number: number;
  name: string;
  kana: string;
  gender: Gender;
  dateOfBirth: string;
  phoneNumber: string;
  email: string;
  dateOfAdmission: string;
  memo: string;
};

export const NewMemberPage: VFC = () => {
  const navigate = useNavigate();
  const { register, handleSubmit } = useForm<Inputs>();
  const [createMember, { data, loading, error }] = useCreateMemberMutation();
  const onSubmit: SubmitHandler<Inputs> = (data) => {
    const input: CreateMemberInput = {
      ...data,
      dateOfBirth: formatRFC3339(parseISO(data.dateOfBirth)),
      dateOfAdmission: formatRFC3339(parseISO(data.dateOfAdmission)),
    };
    createMember({ variables: { input } }).then(() => navigate("/members"));
  };

  return (
    <>
      <h1>登録</h1>
      <form className="flex flex-col" onSubmit={handleSubmit(onSubmit)}>
        <label htmlFor="number">番号　</label>
        <input {...register("number", { required: true })} />
        <label htmlFor="name">名前</label>
        <input {...register("name", { required: true })} />
        <label htmlFor="kana">よみがな</label>
        <input {...register("kana", { required: true })} />
        <label htmlFor="gender">性別</label>
        <div className="flex flex-row items-center">
          女
          <input
            {...register("gender", { required: true })}
            type="radio"
            value="FEMALE"
          />
          男
          <input
            {...register("gender", { required: true })}
            type="radio"
            value="MALE"
          />
          その他
          <input
            {...register("gender", { required: true })}
            type="radio"
            value="OTHER"
          />
        </div>
        <label htmlFor="dateOfBirth">生年月日</label>
        <input type="date" {...register("dateOfBirth")} />
        <label htmlFor="phoneNumber">電話番号</label>
        <input type="tel" {...register("phoneNumber")} />
        <label htmlFor="email">メール</label>
        <input type="email" {...register("email")} />
        <label htmlFor="dateOfAdmission">入会日</label>
        <input type="date" {...register("dateOfAdmission")} />
        <label htmlFor="memo">メモ</label>
        <input type="text" {...register("memo")} />
        <input type="submit" />
      </form>
    </>
  );
};
