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
          kana
          biography
          phoneNumber
          email
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
      kana
      biography
      phoneNumber
      email
    }
  }
`;

gql`
  query getClassesBySchool($id: ID!, $date: Time) {
    school(id: $id) {
      id
      name
      studios {
        id
        name
        location
        rooms {
          id
          name
          capacity
          schedules {
            id
            dayOfWeek
            startTime
            endTime
            class(date: $date) {
              id
              name
              level
              tuition
              startDate
              endDate
              instructor {
                id
                name
              }
            }
          }
        }
      }
    }
  }
`;

gql`
  query getClass($id: ID!) {
    class(id: $id) {
      id
      name
      level
      tuition
      startDate
      endDate
      instructor {
        id
        name
      }
      schedule {
        id
      }
    }
  }
`;

gql`
  query getScheduleAndInstructorsToRegisterNewClass($id: ID!) {
    schedule(id: $id) {
      id
      dayOfWeek
      startTime
      endTime
    }
    instructors(first: 100) {
      totalCount
      edges {
        node {
          id
          name
        }
      }
    }
  }
`;

gql`
  query getMembers($first: Int, $after: Cursor, $last: Int, $before: Cursor) {
    members(first: $first, after: $after, last: $last, before: $before) {
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
          number
          name
          kana
          gender
          dateOfAdmission
        }
      }
    }
  }
`;

gql`
  query getMember($id: ID!) {
    member(id: $id) {
      id
      number
      name
      kana
      gender
      dateOfBirth
      phoneNumber
      email
      dateOfAdmission
      dateOfWithdrawal
      memo
      membersClasses {
        id
        dateOfAdmission
        dateOfWithdrawal
        class {
          id
          name
          level
          instructor {
            id
            name
          }
        }
      }
    }
  }
`;
