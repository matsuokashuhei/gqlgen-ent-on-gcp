import { format, formatRFC3339, parseISO } from "date-fns";
import { useCallback, useEffect, useMemo, useState, VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useLocation, useNavigate, useParams } from "react-router-dom";
import {
  Gender,
  UpdateMemberInput,
  useGetMemberLazyQuery,
  useUpdateMemberMutation,
} from "../../../generated/graphql";

type Inputs = {
  id: number;
  number: number;
  name: string;
  kana: string;
  gender: Gender;
  dateOfBirth: string;
  phoneNumber: string;
  email: string;
  dateOfAdmission: string;
  dateOfWithdrawal: string;
  memo: string;
};

export const MemberPage: VFC = () => {
  const { id } = useParams();
  const { register, handleSubmit, setValue } = useForm<Inputs>();
  const [getMember, { data, loading, error }] = useGetMemberLazyQuery();
  const [updateMember] = useUpdateMemberMutation();
  const [editable, setEditable] = useState<boolean>(false);

  useEffect(() => {
    if (!id) return;

    getMember({ variables: { id } });
  }, [id, getMember]);

  const setValueFromdata = useCallback(() => {
    if (!data) return;
    const { member } = data;
    setValue("id", parseInt(member.id));
    setValue("number", member.number);
    setValue("name", member.name);
    setValue("kana", member.kana);
    setValue("gender", member.gender);
    setValue("dateOfBirth", format(parseISO(member.dateOfBirth), "yyyy-MM-dd"));
    setValue("phoneNumber", member.phoneNumber || "");
    setValue("email", member.email || "");
    setValue(
      "dateOfAdmission",
      format(parseISO(member.dateOfAdmission), "yyyy-MM-dd")
    );
    setValue(
      "dateOfWithdrawal",
      member.dateOfWithdrawal
        ? format(parseISO(member.dateOfWithdrawal), "yyyy-MM-dd")
        : ""
    );
    setValue("memo", member.memo);
    setGender(member.gender);
  }, [data, setValue]);

  useEffect(() => {
    setValueFromdata();
  }, [setValueFromdata]);

  const [gender, setGender] = useState<Gender | null>(null);

  const renderCancelButton = () => {
    if (editable) {
      return (
        <button
          type="button"
          onClick={() => {
            setEditable(!editable);
            setValueFromdata();
          }}
        >
          キャンセル
        </button>
      );
    } else {
      return <div></div>;
    }
  };

  const renderEditButton = () => {
    if (editable) {
      return (
        <button type="submit" onClick={handleSubmit(onSubmit)}>
          変更
        </button>
      );
    } else {
      return (
        <button type="button" onClick={() => setEditable(true)}>
          編集
        </button>
      );
    }
  };

  const renderButtons = () => {
    return (
      <div className="flex flex-row justify-between">
        {renderCancelButton()}
        {renderEditButton()}
      </div>
    );
  };

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    const input: UpdateMemberInput = {
      ...data,
      id: data.id.toString(),
      dateOfBirth: data.dateOfBirth
        ? formatRFC3339(parseISO(data.dateOfBirth))
        : null,
      dateOfAdmission: formatRFC3339(parseISO(data.dateOfAdmission)),
      dateOfWithdrawal: data.dateOfWithdrawal
        ? formatRFC3339(parseISO(data.dateOfWithdrawal))
        : null,
    };
    updateMember({ variables: { input } })
      .then((member) => getMember({ variables: { id: data.id.toString() } }))
      .finally(() => setEditable(false));
  };

  if (!data) return <></>;
  const { member } = data;
  if (!member) return <></>;

  return (
    <>
      <h1>登録</h1>
      <form className="flex flex-col">
        <label htmlFor="number">番号　</label>
        <input
          {...register("number", { required: true, disabled: !editable })}
        />
        <label htmlFor="name">名前</label>
        <input {...register("name", { required: true, disabled: !editable })} />
        <label htmlFor="kana">よみがな</label>
        <input {...register("kana", { required: true, disabled: !editable })} />
        <label htmlFor="gender">性別</label>
        <div className="flex flex-row items-center">
          女
          <input
            {...register("gender", {
              required: true,
              disabled: !editable,
            })}
            type="radio"
            value={Gender.Female.toString()}
            onChange={() => setGender(Gender.Female)}
            checked={gender === Gender.Female}
          />
          男
          <input
            {...register("gender", {
              required: true,
              disabled: !editable,
            })}
            type="radio"
            value={Gender.Male.toString()}
            onChange={() => setGender(Gender.Male)}
            checked={gender === Gender.Male}
          />
          その他
          <input
            {...register("gender", {
              required: true,
              disabled: !editable,
            })}
            type="radio"
            value={Gender.Other.toString()}
            onChange={() => setGender(Gender.Other)}
            checked={gender === Gender.Other}
          />
        </div>
        <label htmlFor="dateOfBirth">生年月日</label>
        <input
          type="date"
          {...register("dateOfBirth", { disabled: !editable })}
        />
        <label htmlFor="phoneNumber">電話番号</label>
        <input
          type="tel"
          {...register("phoneNumber", { disabled: !editable })}
        />
        <label htmlFor="email">メール</label>
        <input type="email" {...register("email", { disabled: !editable })} />
        <label htmlFor="dateOfAdmission">入会日</label>
        <input
          type="date"
          {...register("dateOfAdmission", { disabled: !editable })}
        />
        <label htmlFor="dateOfWithdrawal">退会日</label>
        <input
          type="date"
          {...register("dateOfWithdrawal", { disabled: !editable })}
        />
        <label htmlFor="memo">メモ</label>
        <input type="text" {...register("memo", { disabled: !editable })} />
      </form>
      {renderButtons()}
    </>
  );
};
