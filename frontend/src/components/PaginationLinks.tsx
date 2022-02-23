import { VFC } from "react";
import { PageInfo } from "../generated/graphql";
import { Link, useSearchParams } from "react-router-dom";
import { ChevronLeftIcon, ChevronRightIcon } from "@heroicons/react/solid";

type Props = {
  path: string;
  pageInfo: PageInfo;
};

export const PaginationLinks: VFC<Props> = ({ path, pageInfo }) => {
  const [searchParams] = useSearchParams();

  const renderPreviousPageButton = (pageInfo: PageInfo) => {
    if (searchParams.get("first") !== null) {
      if (searchParams.get("after") !== null) {
        return (
          <div>
            <Link to={`${path}?last=10&before=${pageInfo.startCursor}`}>
              <ChevronLeftIcon className="h-8 w-8 text-gray-500" />
            </Link>
          </div>
        );
      } else {
        return <></>;
      }
    }
    if (searchParams.get("last") !== null) {
      if (pageInfo.hasPreviousPage) {
        return (
          <div>
            <Link to={`${path}?last=10&before=${pageInfo.startCursor}`}>
              <ChevronLeftIcon className="h-8 w-8 text-gray-500" />
            </Link>
          </div>
        );
      } else {
        return <></>;
      }
    }
  };

  const renderNextPageButton = (pageInfo: PageInfo) => {
    if (searchParams.get("first") !== null) {
      if (pageInfo.hasNextPage) {
        return (
          <div>
            <Link to={`${path}?first=10&after=${pageInfo.endCursor}`}>
              <ChevronRightIcon className="h-8 w-8 text-gray-500" />
            </Link>
          </div>
        );
      } else {
        return <></>;
      }
    }
    if (searchParams.get("last") !== null) {
      if (searchParams.get("before") !== null) {
        return (
          <div>
            <Link to={`${path}?first=10&after=${pageInfo.endCursor}`}>
              <ChevronRightIcon className="h-8 w-8 text-gray-500" />
            </Link>
          </div>
        );
      } else {
        return <></>;
      }
    }
  };
  return (
    <div className="flex flex-row justify-between">
      <div>{renderPreviousPageButton(pageInfo)}</div>
      <div>{renderNextPageButton(pageInfo)}</div>
    </div>
  );
};
