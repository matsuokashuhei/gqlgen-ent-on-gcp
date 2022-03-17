import { formatRFC3339, parseISO } from "date-fns";
import { useEffect, VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import { Layout } from "../../../../components";
import {
  CreateClassInput,
  useCreateClassMutation,
  useGetInstructorsLazyQuery,
} from "../../../../generated/graphql";

type Inputs = {
  name: string;
  level: string;
  tuition: number;
  startDate: string;
  instructorId: number;
  scheduleId: number;
};

export const NewClassPage: VFC = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { register, handleSubmit } = useForm<Inputs>();
  const [getInstructors, { data, loading, error }] =
    useGetInstructorsLazyQuery();
  const [createClass] = useCreateClassMutation();

  useEffect(() => {
    getInstructors({ variables: { first: 100 } });
  }, [getInstructors]);

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    console.log("data", data);
    console.log(formatRFC3339(parseISO(data.startDate)));
    const input: CreateClassInput = {
      name: data.name,
      level: data.level,
      tuition: data.tuition,
      startDate: formatRFC3339(parseISO(data.startDate)),
      instructorId: data.instructorId.toString(),
      scheduleId: data.scheduleId.toString(),
    };
    createClass({ variables: { input } }).then((data) => navigate("/classes"));
  };

  return (
    <Layout>
      <h1>クラス</h1>
      <form className="flex flex-col">
        <input
          type="hidden"
          {...register("scheduleId", { required: true })}
          defaultValue={id}
        />
        <label htmlFor="name">名前</label>
        <input type="text" {...register("name", { required: true })} />
        <label htmlFor="level">レベル</label>
        <input type="text" {...register("level", { required: true })} />
        <label htmlFor="tuition">料金</label>
        <input type="number" {...register("tuition", { required: true })} />
        <label htmlFor="startDate">開始日</label>
        <input type="date" {...register("startDate", { required: true })} />
        <label htmlFor="instructorId">インストラクター</label>
        <select {...register("instructorId")}>
          <option disabled selected>
            --
          </option>
          {data?.instructors.edges.map(({ node }) => (
            <option value={node.id}>{node.name}</option>
          ))}
        </select>
      </form>
      <button type="submit" onClick={handleSubmit(onSubmit)}>
        登録
      </button>
    </Layout>
  );
};
