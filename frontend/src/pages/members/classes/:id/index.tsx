import { useEffect, useState, VFC } from "react";
import { useForm } from "react-hook-form";
import { useParams } from "react-router-dom";
import { useGetMembersClassLazyQuery } from "../../../../generated/graphql";
import { Dialog } from "../../../../components/modals/Dialog";
import { Timetable } from "../../../../components/Timetable";
import { format, formatISO, parseISO } from "date-fns";

type Inputs = {
  id: number;
  memberId: number;
  classId: number;
  dateOfAdmission: string;
  dateOfWithdrawal: string;
};

export const MembersClassPage: VFC = () => {
  // const { memberId, classId } = useParams();

  // const [getMembersClass, { data, loading, error }] =
  //   useGetMembersClassLazyQuery();

  // const [isOpenTimetable, setIsOpenTimetable] = useState<boolean>(false);
  // const [className, setClassName] = useState<string>("");
  // const { register, handleSubmit, setValue } = useForm<Inputs>();

  // useEffect(() => {
  //   if (!memberId) return;
  //   if (!classId) return;

  //   getMembersClass({ variables: { memberId, classId } });
  // }, [memberId, classId, getMembersClass]);

  // useEffect(() => {
  //   if (!data) return;
  //   const { membersClass } = data;

  //   if (!membersClass) return;

  //   setValue("id", parseInt(membersClass.id));
  //   setValue("memberId", parseInt(membersClass.member.id));
  //   setValue("classId", parseInt(membersClass.class.id));
  //   setClassName(
  //     `${membersClass.class.name}${membersClass.class.level} ${membersClass.class.instructor.name}`
  //   );
  //   setValue("classId", parseInt(membersClass.class.id));
  //   setValue(
  //     "dateOfAdmission",
  //     format(parseISO(membersClass.dateOfAdmission), "yyyy-MM-dd")
  //   );
  // }, [data, setValue]);

  // return (
  //   <>
  //     <form className="flex flex-col">
  //       <input type="hidden" {...register("id", { required: true })} />
  //       <input type="hidden" {...register("memberId", { required: true })} />
  //       <label htmlFor="classNasme">クラス</label>
  //       <input type="hidden" {...register("classId", { required: true })} />
  //       <input
  //         id="className"
  //         type="text"
  //         value={className}
  //         onClick={() => setIsOpenTimetable(true)}
  //         readOnly={true}
  //       />
  //       <label htmlFor="dateOfAdmission">入会日</label>
  //       <input
  //         type="date"
  //         {...register("dateOfAdmission", { required: true })}
  //       />
  //       <label htmlFor="dateOfWithdrawal">退会日</label>
  //       <input
  //         type="date"
  //         {...register("dateOfWithdrawal", { required: true })}
  //       />
  //     </form>
  //     <Dialog
  //       isOpen={isOpenTimetable}
  //       onClose={() => setIsOpenTimetable(false)}
  //     >
  //       <Timetable
  //         date={format(new Date(), "yyyyMMdd")}
  //         onClickClass={(clazz) => {
  //           setValue("classId", parseInt(clazz.id));
  //           setClassName(
  //             `${clazz.name}${clazz.level} ${clazz.instructor.name}`
  //           );
  //           setIsOpenTimetable(false);
  //         }}
  //       />
  //     </Dialog>
  //   </>
  // );
  return <></>;
};
