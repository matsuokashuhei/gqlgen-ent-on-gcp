import { gql } from "@apollo/client";

gql`
  query GetInstructors(
    $first: Int
    $after: Cursor
    $last: Int
    $before: Cursor
  ) {
    instructors(first: $first, after: $after, last: $last, before: $before) {
      totalCount
      pageInfo {
        startCursor
        endCursor
        hasNextPage
        hasPreviousPage
      }
      edges {
        cursor
        node {
          id
          name
          syllabicCharacters
          biography
          phoneNumber
          email
        }
      }
    }
  }
`;
