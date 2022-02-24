import { useEffect, useState, VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useParams } from "react-router-dom";
import { Layout } from "../../../components";
import {
  Instructor,
  UpdateInstructorInput,
  useGetInstructorLazyQuery,
  useUpdateInstructorMutation,
} from "../../../generated/graphql";

type Inputs = {
  id: number;
  name: string;
  syllabicCharacters: string;
  biography: string;
  phoneNumber: string;
  email: string;
};
export const InstructorPage: VFC = () => {
  const { id } = useParams();
  const [editable, setEditable] = useState<Boolean>(false);
  const [getInstructor, { data, loading, error }] = useGetInstructorLazyQuery();
  const [updateInstructor, { data: data2 }] = useUpdateInstructorMutation();
  //   const [instructor, setInstructor] = useState<Instructor | null>(null);

  const { register, handleSubmit } = useForm<Inputs>();

  useEffect(() => {
    if (id) {
      getInstructor({ variables: { id } });
    }
  }, [id, getInstructor]);

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    const input: UpdateInstructorInput = {
      id: data.id.toString(),
      name: data.name,
      syllabicCharacters: data.syllabicCharacters,
      biography: data.biography,
      phoneNumber: data.phoneNumber,
      email: data.email,
    };
    updateInstructor({ variables: { input } })
      .then((instrutor) =>
        getInstructor({ variables: { id: data.id.toString() } })
      )
      .then(() => setEditable(false));
  };

  const renderInstructor = (instructor: Instructor) => {
    if (editable) {
      return (
        <form className="flex flex-col">
          <input
            type="hidden"
            {...register("id", { required: true })}
            defaultValue={instructor.id}
          />
          <input
            {...register("name", { required: true })}
            defaultValue={instructor.name}
          />
          <input
            {...register("syllabicCharacters", { required: true })}
            defaultValue={instructor.syllabicCharacters}
          />
          <input
            {...register("biography")}
            defaultValue={instructor.biography ?? ""}
          />
          <input
            {...register("phoneNumber")}
            defaultValue={instructor.phoneNumber ?? ""}
          />
          <input {...register("email")} defaultValue={instructor.email ?? ""} />
        </form>
      );
    } else {
      return (
        <div className="flex flex-col">
          <div>{instructor.id}</div>
          <div>{instructor.name}</div>
          <div>{instructor.syllabicCharacters}</div>
          <div>{instructor.biography}</div>
          <div>{instructor.phoneNumber}</div>
          <div>{instructor.email}</div>
          <div>{instructor.createTime}</div>
          <div>{instructor.updateTime}</div>
        </div>
      );
    }
  };

  const renderEditButton = () => {
    if (editable) {
      return (
        <button type="submit" onClick={handleSubmit(onSubmit)}>
          保存
        </button>
      );
    } else {
      return <button onClick={() => setEditable(true)}>編集</button>;
    }
  };

  const instructor = data?.instructor;
  if (!instructor) return <></>;

  return (
    <Layout>
      {renderInstructor(instructor)}
      {renderEditButton()}
    </Layout>
  );
};
