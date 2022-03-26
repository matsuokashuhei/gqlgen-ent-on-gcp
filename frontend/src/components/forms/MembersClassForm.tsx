import { useEffect, useRef, useState, VFC } from "react";
import { useForm } from "react-hook-form";
import {
  GetMemberQuery,
  useGetClassesBySchoolLazyQuery,
} from "../../generated/graphql";
import { Dialog } from "../modals/Dialog";
import { Timetable } from "../Timetable";

type Props = {
  member: GetMemberQuery["member"];
};

type Inputs = {
  memberId: number;
  classId: number;
  dateOfAdmission: string;
  dateOfWithdrawal?: string;
};

export const MembersClassForm: VFC<Props> = ({ member }) => {
  const [getClassesBySchool, { data, loading, error }] =
    useGetClassesBySchoolLazyQuery();
  const { register, handleSubmit } = useForm<Inputs>();
  const [showDialog, setShowDialog] = useState<boolean>(false);
  let refDiv = useRef(null);

  useEffect(() => {
    getClassesBySchool({ variables: { id: "1" } });
  }, [getClassesBySchool]);

  return (
    <>
      <form className="flex flex-col">
        <input type="hidden" {...register("memberId", { required: true })} />
        <label
          htmlFor="classId"
          onClick={() => {
            setShowDialog(true);
          }}
        >
          クラス
        </label>
        <input type="hidden" {...register("classId", { required: true })} />
        <label htmlFor="dateOfAdmission">入会日</label>
        <input
          type="date"
          {...register("dateOfAdmission", { required: true })}
        />
        <label htmlFor="dateOfAdmission">退会日</label>
        <input type="date" {...register("dateOfWithdrawal")} />
      </form>
      <Dialog isOpen={showDialog} onClose={() => setShowDialog(false)}>
        <Timetable date={"20220325"} onClickClass={() => console.log(1)} />
      </Dialog>
      <button type="submit">登録</button>
    </>
  );
};
