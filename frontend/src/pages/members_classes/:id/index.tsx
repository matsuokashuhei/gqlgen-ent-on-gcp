import { format, parseISO, formatRFC3339 } from "date-fns";
import { useEffect, useState, VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useParams } from "react-router-dom";
import { Dialog } from "../../../components/modals/Dialog";
import { Timetable } from "../../../components/Timetable";
import {
  UpdateMembersClassInput,
  useGetMembersClassLazyQuery,
  useUpdateMembersClassMutation,
} from "../../../generated/graphql";

type Inputs = {
  id: number;
  memberId: number;
  classId: number;
  dateOfAdmission: string;
  dateOfWithdrawal: string;
};

export const MembersClassPage: VFC = () => {
  const { membersClassId } = useParams();

  const [getMembersClass, { data, loading, error }] =
    useGetMembersClassLazyQuery();
  const [updateMembersClass] = useUpdateMembersClassMutation();

  const [editable, setEditable] = useState(false);
  const [isOpenTimetable, setIsOpenTimetable] = useState(false);
  const [className, setClassName] = useState("");
  const { register, handleSubmit, setValue } = useForm<Inputs>();

  useEffect(() => {
    if (!membersClassId) return;

    getMembersClass({ variables: { id: membersClassId } });
  }, [membersClassId, getMembersClass]);

  useEffect(() => {
    if (!data) return;
    const { membersClass } = data;

    if (!membersClass) return;

    setValue("id", parseInt(membersClass.id));
    setValue("memberId", parseInt(membersClass.member.id));
    setValue("classId", parseInt(membersClass.class.id));
    setClassName(
      `${membersClass.class.name}${membersClass.class.level} ${membersClass.class.instructor.name}`
    );
    setValue("classId", parseInt(membersClass.class.id));
    setValue(
      "dateOfAdmission",
      format(parseISO(membersClass.dateOfAdmission), "yyyy-MM-dd")
    );
  }, [data, setValue]);

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    const input: UpdateMembersClassInput = {
      id: data.id.toString(),
      memberId: data.memberId.toString(),
      classId: data.memberId.toString(),
      dateOfAdmission: data.dateOfAdmission
        ? formatRFC3339(parseISO(data.dateOfAdmission))
        : null,
      dateOfWithdrawal: data.dateOfWithdrawal
        ? formatRFC3339(parseISO(data.dateOfWithdrawal))
        : null,
    };
    updateMembersClass({ variables: { input } });
  };

  return (
    <>
      <form className="flex flex-col">
        <input type="hidden" {...register("id", { required: true })} />
        <input type="hidden" {...register("memberId", { required: true })} />
        <label htmlFor="classNasme">クラス</label>
        <input type="hidden" {...register("classId", { required: true })} />
        <input
          id="className"
          type="text"
          value={className}
          onClick={() => editable && setIsOpenTimetable(true)}
          readOnly={true}
        />
        <label htmlFor="dateOfAdmission">入会日</label>
        <input
          type="date"
          {...register("dateOfAdmission", {
            required: true,
            disabled: !editable,
          })}
        />
        <label htmlFor="dateOfWithdrawal">退会日</label>
        <input
          type="date"
          {...register("dateOfWithdrawal", {
            disabled: !editable,
          })}
        />
      </form>
      <div className="flex w-full flex-row justify-between">
        {editable ? (
          <>
            <button type="button" onClick={() => setEditable(false)}>
              キャンセル
            </button>
            <button type="submit" onClick={handleSubmit(onSubmit)}>
              更新
            </button>
          </>
        ) : (
          <>
            <div></div>
            <button type="button" onClick={() => setEditable(true)}>
              編集
            </button>
          </>
        )}
      </div>
      <Dialog
        isOpen={isOpenTimetable}
        onClose={() => setIsOpenTimetable(false)}
      >
        <Timetable
          date={format(new Date(), "yyyyMMdd")}
          onClickClass={(clazz) => {
            setValue("classId", parseInt(clazz.id));
            setClassName(
              `${clazz.name}${clazz.level} ${clazz.instructor.name}`
            );
            setIsOpenTimetable(false);
          }}
        />
      </Dialog>
    </>
  );
};
