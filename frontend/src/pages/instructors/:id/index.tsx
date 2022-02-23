import { VFC } from "react";
import { useParams } from "react-router-dom";
import { Layout } from "../../../components";

export const InstructorPage: VFC = () => {
  const { instructorId } = useParams();

  return (
    <Layout>
      <h1>HHHHHHHHHHHHHHHHHHHHHH{instructorId}</h1>
    </Layout>
  );
};
