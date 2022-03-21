import { useEffect, VFC } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { PaginationLinks } from "../../components";
import {
  GetInstructorsQuery,
  useGetInstructorsLazyQuery,
} from "../../generated/graphql";

export const InstructorsPage: VFC = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const [getInstructors, { data, loading, error }] =
    useGetInstructorsLazyQuery();

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

  const renderInstructors = (
    edges: GetInstructorsQuery["instructors"]["edges"]
  ) => {
    return edges.map(({ cursor, node }) => (
      <tr key={cursor} onClick={() => navigate(`/instructors/${node.id}`)}>
        <td>{node.id}</td>
        <td>{node.name}</td>
      </tr>
    ));
  };

  if (!data) return <></>;

  const {
    instructors: { edges, pageInfo },
  } = data;

  return (
    <>
      <h1>インストラクター</h1>
      <button onClick={() => navigate("/instructors/new")}>登録</button>
      <table>
        <thead>
          <tr>
            <td>ID</td>
            <td>Name</td>
          </tr>
        </thead>
        <tbody>{renderInstructors(edges)}</tbody>
      </table>
      <PaginationLinks path="/instructors" pageInfo={pageInfo} />
    </>
  );
};
