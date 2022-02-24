import { filter, compact } from "lodash-es";
import { useEffect, useMemo, VFC } from "react";
import { Link, useNavigate, useSearchParams } from "react-router-dom";
import { Layout, PaginationLinks } from "../../components";
import {
  GetInstructorsQuery,
  Instructor,
  InstructorEdge,
  useGetInstructorsLazyQuery,
  useGetInstructorsQuery,
} from "../../generated/graphql";

export const InstructorsPage: VFC = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const [getInstructors, { data, loading, error }] =
    useGetInstructorsLazyQuery();
  //   const { data, loading, error } = useGetInstructorsQuery({
  //     variables: { first: 10 },
  //   });
  // const {data, loading, error } = useGetInstructorsQuery()

  useEffect(() => {
    const after = searchParams.get("after");
    const before = searchParams.get("before");
    if (!after && !before) {
      getInstructors({ variables: { first: 10 } });
    } else if (after) {
      getInstructors({ variables: { first: 10, after: after } });
    } else {
      getInstructors({ variables: { last: 10, before: before } });
    }
  }, [getInstructors, searchParams]);

  const renderInstructors = (data?: GetInstructorsQuery) => {
    return data?.instructors?.edges.map((edge) => (
      <tr
        key={edge?.cursor}
        onClick={() => navigate(`/instructors/${edge?.node.id}`)}
      >
        <td>{edge?.node?.id}</td>
        <td>{edge?.node?.name}</td>
      </tr>
    ));
  };
  // const renderInstructors = (instructors: Instructor[]) => {
  //   return instructors.map((instructor) => (
  //     <tr
  //       key={instructor.id}
  //       onClick={() => navigate(`/instructors/${instructor.id}`)}
  //     >
  //       <td>{instructor.id}</td>
  //       <td>{instructor.name}</td>
  //     </tr>
  //   ));
  // };

  // const instructors: Instructor[] = compact(
  //   data?.instructors?.edges?.map((edge) => edge?.node) ?? []
  // );
  // const instructors: Instructor | undefined =
  //   return data?.instructors?.edges.map((edge) => edge?.node);

  return (
    <Layout>
      <div>
        <h1>インストラクター</h1>
        <table>
          <thead>
            <tr>
              <td>ID</td>
              <td>Name</td>
            </tr>
          </thead>
          <tbody>{renderInstructors(data)}</tbody>
        </table>
        {data?.instructors?.pageInfo && (
          <PaginationLinks
            path="/instructors"
            pageInfo={data?.instructors?.pageInfo}
          />
        )}
      </div>
    </Layout>
  );
};
