import { gql } from "@apollo/client";

gql`
  query getInstructors(
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
          createTime
          updateTime
        }
      }
    }
  }
`;

gql`
  query getInstructor($id: ID!) {
    instructor(id: $id) {
      id
      name
      syllabicCharacters
      biography
      phoneNumber
      email
      createTime
      updateTime
    }
  }
`;
