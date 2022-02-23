import { useEffect, VFC } from "react";
import { Link, useNavigate, useSearchParams } from "react-router-dom";
import { Layout, PaginationLinks } from "../../components";
import {
  GetInstructorsQuery,
  useGetInstructorsLazyQuery,
  useGetInstructorsQuery,
} from "../../generated/graphql";

export const InstructorsPage: VFC = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const [GetInstructors, { data, loading, error }] =
    useGetInstructorsLazyQuery();
  //   const { data, loading, error } = useGetInstructorsQuery({
  //     variables: { first: 10 },
  //   });
  // const {data, loading, error } = useGetInstructorsQuery()

  useEffect(() => {
    const after = searchParams.get("after");
    const before = searchParams.get("before");
    if (!after && !before) {
      GetInstructors({ variables: { first: 10 } });
    } else if (after) {
      GetInstructors({ variables: { first: 10, after: after } });
    } else {
      GetInstructors({ variables: { last: 10, before: before } });
    }
  }, [GetInstructors, searchParams]);

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
          <tbody>{!loading && renderInstructors(data)}</tbody>
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
