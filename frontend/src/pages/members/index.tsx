import { useEffect, VFC } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { PaginationLinks } from "../../components";
import {
  GetMembersQuery,
  useGetMembersLazyQuery,
} from "../../generated/graphql";

export const MembersPage: VFC = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const [getMembers, { data, loading, error }] = useGetMembersLazyQuery();

  useEffect(() => {
    const after = searchParams.get("after");
    const before = searchParams.get("before");
    if (!after && !before) {
      getMembers({ variables: { first: 10 } });
    } else if (after) {
      getMembers({ variables: { first: 10, after } });
    } else {
      getMembers({ variables: { last: 10, before } });
    }
  }, [searchParams, getMembers]);

  const rennderMembers = (edges: GetMembersQuery["members"]["edges"]) => {
    return edges.map(({ cursor, node }) => {
      return (
        <tr key={cursor} onClick={() => navigate(`/members/${node.id}`)}>
          <td>{node.number}</td>
          <td>{node.name}</td>
          <td>{node.gender}</td>
          <td>{node.dateOfAdmission}</td>
        </tr>
      );
    });
  };

  if (!data) return <></>;

  const {
    members: { edges, pageInfo },
  } = data;

  return (
    <>
      <h1>会員</h1>
      <button onClick={() => navigate("/members/new")}>登録</button>
      <table>
        <thead>
          <tr>
            <th>番号</th>
            <th>名前</th>
            <th>性別</th>
            <th>入会日</th>
          </tr>
        </thead>
        <tbody>{rennderMembers(edges)}</tbody>
      </table>
      <PaginationLinks path="/members" pageInfo={pageInfo} />
    </>
  );
};
