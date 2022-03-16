import { gql } from "@apollo/client";

gql`
  mutation CreateInstructor($input: CreateInstructorInput!) {
    createInstructor(input: $input) {
      id
      name
      syllabicCharacters
      biography
      phoneNumber
      email
    }
  }
`;

gql`
  mutation UpdateInstructor($input: UpdateInstructorInput!) {
    updateInstructor(input: $input) {
      id
      name
      syllabicCharacters
      biography
      phoneNumber
      email
    }
  }
`;

gql`
  mutation DeleteInstructor($input: DeleteInstructorInput!) {
    deleteInstructor(input: $input) {
      id
      name
      syllabicCharacters
      biography
      phoneNumber
      email
    }
  }
`;

gql`
  mutation CreateClass($input: CreateClassInput!) {
    createClass(input: $input) {
      id
      name
      level
      startDate
      endDate
      instructor {
        id
      }
      schedule {
        id
      }
    }
  }
`;
