import { useEffect, FC, VFC } from "react";
import { Layout } from "../../components";
import {
  useGetInstructorsLazyQuery,
  GetInstructorsQuery,
  PageInfo,
} from "../../generated/graphql";
import { Link, useSearchParams } from "react-router-dom";

export const InstructorsPage: VFC = () => {
  const [searchParams] = useSearchParams();
  const [GetInstructors, { data, loading, error }] =
    useGetInstructorsLazyQuery();
  //   const { data, loading, error } = useGetInstructorsQuery({
  //     variables: { first: 10 },
  //   });

  useEffect(() => {
    // const first = parseInt(searchParams.get("first") || "10");
    const after = searchParams.get("after");
    const before = searchParams.get("before");
    if (!after && !before) {
      GetInstructors({ variables: { first: 10 } });
    } else if (after) {
      GetInstructors({ variables: { first: 10, after: after } });
    } else {
      GetInstructors({ variables: { last: 10, before: before } });
    }
    // const last = parseInt(searchParams.get("last") || "10");
  }, [GetInstructors, searchParams]);

  const renderInstructors = (data?: GetInstructorsQuery) => {
    return data?.instructors?.edges.map((edge) => (
      <tr key={edge?.cursor}>
        <td>{edge?.node?.id}</td>
        <td>{edge?.node?.name}</td>
      </tr>
    ));
  };

  const renderPagination = (data?: GetInstructorsQuery) => {
    const renderPreviousPageButton = (pageInfo: PageInfo) => {
      if (pageInfo.hasPreviousPage) {
        return (
          <div>
            <Link to={`/instructors?last=10&before=${pageInfo.startCursor}`}>
              {pageInfo.startCursor}
            </Link>
          </div>
        );
      } else {
        return <></>;
      }
    };
    const renderNextPageButton = (pageInfo: PageInfo) => {
      if (pageInfo.hasNextPage) {
        return (
          <div>
            <Link to={`/instructors?first=10&after=${pageInfo.endCursor}`}>
              {pageInfo.endCursor}
            </Link>
          </div>
        );
      } else {
        return <></>;
      }
    };
    const pageInfo = data?.instructors?.pageInfo;
    if (!pageInfo) return;

    return (
      <div className="flex flex-row justify-between">
        <div>{renderPreviousPageButton(pageInfo)}</div>
        <div>{renderNextPageButton(pageInfo)}</div>
      </div>
    );
  };

  if (loading) {
    return <></>;
  }

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
        {renderPagination(data)}
      </div>
    </Layout>
  );
};
