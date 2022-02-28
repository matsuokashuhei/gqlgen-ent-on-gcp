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

gql`
  query getSchedules($schoolID: ID!) {
    school(id: $schoolID) {
      id
      name
      createTime
      updateTime
      studios {
        id
        name
        location
        createTime
        updateTime
        rooms {
          id
          name
          capacity
          createTime
          updateTime
          schedules {
            id
            dayOfWeek
            startTime
            endTime
            createTime
            updateTime
          }
        }
      }
    }
  }
`;
