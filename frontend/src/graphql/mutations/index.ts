import { gql } from "@apollo/client";

gql`
  mutation CreateInstructor($input: CreateInstructorInput!) {
    createInstructor(input: $input) {
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
  mutation UpdateInstructor($input: UpdateInstructorInput!) {
    updateInstructor(input: $input) {
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
  mutation DeleteInstructor($input: DeleteInstructorInput!) {
    deleteInstructor(input: $input) {
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

gql`
  mutation UpdateClass($input: UpdateClassInput!) {
    updateClass(input: $input) {
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

gql`
  mutation DeleteClass($id: ID!) {
    deleteClass(id: $id) {
      id
      name
      level
      startDate
      endDate
    }
  }
`;

gql`
  mutation CreateMember($input: CreateMemberInput!) {
    createMember(input: $input) {
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
    }
  }
`;

gql`
  mutation UpdateMember($input: UpdateMemberInput!) {
    updateMember(input: $input) {
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
    }
  }
`;
