import { format, formatRFC3339, parseISO } from "date-fns";
import { useState, VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import {
  CreateMembersClassInput,
  GetMemberQuery,
  useCreateMembersClassMutation,
} from "../../generated/graphql";
import { Dialog } from "../modals/Dialog";
import { Timetable } from "../Timetable";

type Props = {
  member: GetMemberQuery["member"];
  onSubmitted: () => void;
};

type Inputs = {
  memberId: string;
  classId: string;
  dateOfAdmission: string;
  dateOfWithdrawal?: string;
};

export const MembersClassForm: VFC<Props> = ({ member, onSubmitted }) => {
  const [createMembersClass] = useCreateMembersClassMutation();
  const { register, setValue, handleSubmit } = useForm<Inputs>();
  const [isOpenTimetable, setIsOpenTimetable] = useState<boolean>(false);
  const [className, setClassName] = useState<string>("");

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    console.log(123);
    const input: CreateMembersClassInput = {
      ...data,
      dateOfAdmission: formatRFC3339(parseISO(data.dateOfAdmission)),
      dateOfWithdrawal: data.dateOfWithdrawal
        ? formatRFC3339(parseISO(data.dateOfWithdrawal))
        : null,
    };
    createMembersClass({ variables: { input } }).then(() => onSubmitted());
  };

  return (
    <>
      <form className="flex flex-col">
        <input
          type="hidden"
          {...register("memberId", { required: true, value: member.id })}
        />
        <label
          htmlFor="classId"
          onClick={() => {
            setIsOpenTimetable(true);
          }}
        >
          クラス
        </label>
        <input type="hidden" {...register("classId", { required: true })} />
        <div>{className}</div>
        <label htmlFor="dateOfAdmission">入会日</label>
        <input
          type="date"
          {...register("dateOfAdmission", { required: true })}
        />
        <label htmlFor="dateOfWithdrawal">退会日</label>
        <input type="date" {...register("dateOfWithdrawal")} />
        <button type="submit" onClick={handleSubmit(onSubmit)}>
          登録
        </button>
      </form>
      <Dialog
        isOpen={isOpenTimetable}
        onClose={() => setIsOpenTimetable(false)}
      >
        {/* <Timetable
          date={format(new Date(), "yyyyMMdd")}
          onClickClass={(clazz) => {
            setValue("classId", clazz.id);
            setClassName(
              `${clazz.name}${clazz.level} ${clazz.instructor.name}`
            );
            setIsOpenTimetable(false);
          }}
        /> */}
      </Dialog>
    </>
  );
};
