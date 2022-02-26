import { gql } from "@apollo/client";
import * as Apollo from "@apollo/client";
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Cursor: any;
  Time: any;
};

export type CreateInstructorInput = {
  biography?: InputMaybe<Scalars["String"]>;
  email?: InputMaybe<Scalars["String"]>;
  name: Scalars["String"];
  phoneNumber?: InputMaybe<Scalars["String"]>;
  syllabicCharacters: Scalars["String"];
};

export type CreateRoomInput = {
  capacity: Scalars["Int"];
  name: Scalars["String"];
  studioID: Scalars["ID"];
};

export type CreateSchoolInput = {
  name: Scalars["String"];
};

export type CreateStudioInput = {
  location: Scalars["String"];
  name: Scalars["String"];
  schoolID: Scalars["ID"];
};

export type CreateUserInput = {
  email: Scalars["String"];
  password: Scalars["String"];
};

export type DeleteInstructorInput = {
  id: Scalars["ID"];
};

export type Instructor = Node & {
  __typename?: "Instructor";
  biography?: Maybe<Scalars["String"]>;
  createTime: Scalars["Time"];
  email?: Maybe<Scalars["String"]>;
  id: Scalars["ID"];
  name: Scalars["String"];
  phoneNumber?: Maybe<Scalars["String"]>;
  syllabicCharacters: Scalars["String"];
  updateTime: Scalars["Time"];
};

export type InstructorConnection = {
  __typename?: "InstructorConnection";
  edges: Array<Maybe<InstructorEdge>>;
  pageInfo: PageInfo;
  totalCount: Scalars["Int"];
};

export type InstructorEdge = {
  __typename?: "InstructorEdge";
  cursor: Scalars["Cursor"];
  node: Instructor;
};

export type InstructorOrder = {
  direction: OrderDirection;
  field?: InputMaybe<InstructorOrderField>;
};

export enum InstructorOrderField {
  SyllabicCharacters = "SYLLABIC_CHARACTERS",
}

export type Mutation = {
  __typename?: "Mutation";
  createInstructor: Instructor;
  createRoom: Room;
  createSchool: School;
  createStudio: Studio;
  createUser: User;
  deleteInstructor: Instructor;
  deleteRoom: Room;
  deleteSchool: School;
  deleteStudio: Studio;
  signUp: User;
  updateInstructor: Instructor;
  updateRoom: Room;
  updateSchool: School;
  updateStudio: Studio;
  updateUser: User;
};

export type MutationCreateInstructorArgs = {
  input: CreateInstructorInput;
};

export type MutationCreateRoomArgs = {
  input: CreateRoomInput;
};

export type MutationCreateSchoolArgs = {
  input: CreateSchoolInput;
};

export type MutationCreateStudioArgs = {
  input: CreateStudioInput;
};

export type MutationCreateUserArgs = {
  input: CreateUserInput;
};

export type MutationDeleteInstructorArgs = {
  input: DeleteInstructorInput;
};

export type MutationDeleteRoomArgs = {
  id: Scalars["ID"];
};

export type MutationDeleteSchoolArgs = {
  id: Scalars["ID"];
};

export type MutationDeleteStudioArgs = {
  id: Scalars["ID"];
};

export type MutationSignUpArgs = {
  input: SignUpInput;
};

export type MutationUpdateInstructorArgs = {
  input: UpdateInstructorInput;
};

export type MutationUpdateRoomArgs = {
  input: UpdateRoomInput;
};

export type MutationUpdateSchoolArgs = {
  input: UpdateSchoolInput;
};

export type MutationUpdateStudioArgs = {
  input: UpdateStudioInput;
};

export type MutationUpdateUserArgs = {
  input: UpdateUserInput;
};

export type Node = {
  id: Scalars["ID"];
};

export enum OrderDirection {
  Asc = "ASC",
  Desc = "DESC",
}

export type PageInfo = {
  __typename?: "PageInfo";
  endCursor?: Maybe<Scalars["Cursor"]>;
  hasNextPage: Scalars["Boolean"];
  hasPreviousPage: Scalars["Boolean"];
  startCursor?: Maybe<Scalars["Cursor"]>;
};

export type Query = {
  __typename?: "Query";
  currentUser: User;
  instructor: Instructor;
  instructors?: Maybe<InstructorConnection>;
  node?: Maybe<Node>;
  nodes: Array<Maybe<Node>>;
  room: Room;
  rooms: Array<Maybe<Room>>;
  school: School;
  schools: Array<Maybe<School>>;
  studio: Studio;
  user: User;
  users: Array<Maybe<User>>;
};

export type QueryInstructorArgs = {
  id: Scalars["ID"];
};

export type QueryInstructorsArgs = {
  after?: InputMaybe<Scalars["Cursor"]>;
  before?: InputMaybe<Scalars["Cursor"]>;
  first?: InputMaybe<Scalars["Int"]>;
  last?: InputMaybe<Scalars["Int"]>;
  orderBy?: InputMaybe<InstructorOrder>;
};

export type QueryNodeArgs = {
  id: Scalars["ID"];
};

export type QueryNodesArgs = {
  ids: Array<Scalars["ID"]>;
};

export type QueryRoomArgs = {
  id: Scalars["ID"];
};

export type QuerySchoolArgs = {
  id: Scalars["ID"];
};

export type QueryStudioArgs = {
  id: Scalars["ID"];
};

export type QueryUserArgs = {
  id: Scalars["ID"];
};

export type Room = Node & {
  __typename?: "Room";
  capacity: Scalars["Int"];
  createTime: Scalars["Time"];
  id: Scalars["ID"];
  name: Scalars["String"];
  schedules: Array<Maybe<Schedule>>;
  studio: Studio;
  updateTime: Scalars["Time"];
};

export type Schedule = Node & {
  __typename?: "Schedule";
  createdAt: Scalars["Time"];
  dayOfWeek: Scalars["Int"];
  endTime: Scalars["String"];
  id: Scalars["ID"];
  startTime: Scalars["String"];
  updatedAt: Scalars["Time"];
};

export type School = Node & {
  __typename?: "School";
  createTime: Scalars["Time"];
  id: Scalars["ID"];
  name: Scalars["String"];
  studios: Array<Maybe<Studio>>;
  updateTime: Scalars["Time"];
};

export type SignUpInput = {
  email: Scalars["String"];
  password: Scalars["String"];
};

export type Studio = Node & {
  __typename?: "Studio";
  createTime: Scalars["Time"];
  id: Scalars["ID"];
  location: Scalars["String"];
  name: Scalars["String"];
  rooms: Array<Maybe<Room>>;
  school: School;
  updateTime: Scalars["Time"];
};

export type UpdateInstructorInput = {
  biography?: InputMaybe<Scalars["String"]>;
  email?: InputMaybe<Scalars["String"]>;
  id: Scalars["ID"];
  name?: InputMaybe<Scalars["String"]>;
  phoneNumber?: InputMaybe<Scalars["String"]>;
  syllabicCharacters?: InputMaybe<Scalars["String"]>;
};

export type UpdateRoomInput = {
  capacity?: InputMaybe<Scalars["Int"]>;
  id: Scalars["ID"];
  name?: InputMaybe<Scalars["String"]>;
  studioID?: InputMaybe<Scalars["ID"]>;
};

export type UpdateSchoolInput = {
  id: Scalars["ID"];
  name: Scalars["String"];
};

export type UpdateStudioInput = {
  id: Scalars["ID"];
  location?: InputMaybe<Scalars["String"]>;
  name?: InputMaybe<Scalars["String"]>;
  schoolID?: InputMaybe<Scalars["ID"]>;
};

export type UpdateUserInput = {
  id: Scalars["ID"];
  name: Scalars["String"];
};

export type User = Node & {
  __typename?: "User";
  createTime: Scalars["Time"];
  firebaseUid: Scalars["String"];
  id: Scalars["ID"];
  name: Scalars["String"];
  updateTime: Scalars["Time"];
};

export type CreateInstructorMutationVariables = Exact<{
  input: CreateInstructorInput;
}>;

export type CreateInstructorMutation = {
  __typename?: "Mutation";
  createInstructor: {
    __typename?: "Instructor";
    id: string;
    name: string;
    syllabicCharacters: string;
    biography?: string | null;
    phoneNumber?: string | null;
    email?: string | null;
  };
};

export type UpdateInstructorMutationVariables = Exact<{
  input: UpdateInstructorInput;
}>;

export type UpdateInstructorMutation = {
  __typename?: "Mutation";
  updateInstructor: {
    __typename?: "Instructor";
    id: string;
    name: string;
    syllabicCharacters: string;
    biography?: string | null;
    phoneNumber?: string | null;
    email?: string | null;
  };
};

export type DeleteInstructorMutationVariables = Exact<{
  input: DeleteInstructorInput;
}>;

export type DeleteInstructorMutation = {
  __typename?: "Mutation";
  deleteInstructor: {
    __typename?: "Instructor";
    id: string;
    name: string;
    syllabicCharacters: string;
    biography?: string | null;
    phoneNumber?: string | null;
    email?: string | null;
  };
};

export type GetInstructorsQueryVariables = Exact<{
  first?: InputMaybe<Scalars["Int"]>;
  after?: InputMaybe<Scalars["Cursor"]>;
  last?: InputMaybe<Scalars["Int"]>;
  before?: InputMaybe<Scalars["Cursor"]>;
}>;

export type GetInstructorsQuery = {
  __typename?: "Query";
  instructors?: {
    __typename?: "InstructorConnection";
    totalCount: number;
    pageInfo: {
      __typename?: "PageInfo";
      startCursor?: any | null;
      endCursor?: any | null;
      hasNextPage: boolean;
      hasPreviousPage: boolean;
    };
    edges: Array<{
      __typename?: "InstructorEdge";
      cursor: any;
      node: {
        __typename?: "Instructor";
        id: string;
        name: string;
        syllabicCharacters: string;
        biography?: string | null;
        phoneNumber?: string | null;
        email?: string | null;
      };
    } | null>;
  } | null;
};

export type GetInstructorQueryVariables = Exact<{
  id: Scalars["ID"];
}>;

export type GetInstructorQuery = {
  __typename?: "Query";
  instructor: {
    __typename?: "Instructor";
    id: string;
    name: string;
    syllabicCharacters: string;
    biography?: string | null;
    phoneNumber?: string | null;
    email?: string | null;
    createTime: any;
    updateTime: any;
  };
};

export const CreateInstructorDocument = gql`
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
export type CreateInstructorMutationFn = Apollo.MutationFunction<
  CreateInstructorMutation,
  CreateInstructorMutationVariables
>;

/**
 * __useCreateInstructorMutation__
 *
 * To run a mutation, you first call `useCreateInstructorMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateInstructorMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createInstructorMutation, { data, loading, error }] = useCreateInstructorMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateInstructorMutation(
  baseOptions?: Apollo.MutationHookOptions<
    CreateInstructorMutation,
    CreateInstructorMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useMutation<
    CreateInstructorMutation,
    CreateInstructorMutationVariables
  >(CreateInstructorDocument, options);
}
export type CreateInstructorMutationHookResult = ReturnType<
  typeof useCreateInstructorMutation
>;
export type CreateInstructorMutationResult =
  Apollo.MutationResult<CreateInstructorMutation>;
export type CreateInstructorMutationOptions = Apollo.BaseMutationOptions<
  CreateInstructorMutation,
  CreateInstructorMutationVariables
>;
export const UpdateInstructorDocument = gql`
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
export type UpdateInstructorMutationFn = Apollo.MutationFunction<
  UpdateInstructorMutation,
  UpdateInstructorMutationVariables
>;

/**
 * __useUpdateInstructorMutation__
 *
 * To run a mutation, you first call `useUpdateInstructorMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateInstructorMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateInstructorMutation, { data, loading, error }] = useUpdateInstructorMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUpdateInstructorMutation(
  baseOptions?: Apollo.MutationHookOptions<
    UpdateInstructorMutation,
    UpdateInstructorMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useMutation<
    UpdateInstructorMutation,
    UpdateInstructorMutationVariables
  >(UpdateInstructorDocument, options);
}
export type UpdateInstructorMutationHookResult = ReturnType<
  typeof useUpdateInstructorMutation
>;
export type UpdateInstructorMutationResult =
  Apollo.MutationResult<UpdateInstructorMutation>;
export type UpdateInstructorMutationOptions = Apollo.BaseMutationOptions<
  UpdateInstructorMutation,
  UpdateInstructorMutationVariables
>;
export const DeleteInstructorDocument = gql`
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
export type DeleteInstructorMutationFn = Apollo.MutationFunction<
  DeleteInstructorMutation,
  DeleteInstructorMutationVariables
>;

/**
 * __useDeleteInstructorMutation__
 *
 * To run a mutation, you first call `useDeleteInstructorMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteInstructorMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteInstructorMutation, { data, loading, error }] = useDeleteInstructorMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useDeleteInstructorMutation(
  baseOptions?: Apollo.MutationHookOptions<
    DeleteInstructorMutation,
    DeleteInstructorMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useMutation<
    DeleteInstructorMutation,
    DeleteInstructorMutationVariables
  >(DeleteInstructorDocument, options);
}
export type DeleteInstructorMutationHookResult = ReturnType<
  typeof useDeleteInstructorMutation
>;
export type DeleteInstructorMutationResult =
  Apollo.MutationResult<DeleteInstructorMutation>;
export type DeleteInstructorMutationOptions = Apollo.BaseMutationOptions<
  DeleteInstructorMutation,
  DeleteInstructorMutationVariables
>;
export const GetInstructorsDocument = gql`
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
        }
      }
    }
  }
`;

/**
 * __useGetInstructorsQuery__
 *
 * To run a query within a React component, call `useGetInstructorsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetInstructorsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetInstructorsQuery({
 *   variables: {
 *      first: // value for 'first'
 *      after: // value for 'after'
 *      last: // value for 'last'
 *      before: // value for 'before'
 *   },
 * });
 */
export function useGetInstructorsQuery(
  baseOptions?: Apollo.QueryHookOptions<
    GetInstructorsQuery,
    GetInstructorsQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useQuery<GetInstructorsQuery, GetInstructorsQueryVariables>(
    GetInstructorsDocument,
    options
  );
}
export function useGetInstructorsLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    GetInstructorsQuery,
    GetInstructorsQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useLazyQuery<GetInstructorsQuery, GetInstructorsQueryVariables>(
    GetInstructorsDocument,
    options
  );
}
export type GetInstructorsQueryHookResult = ReturnType<
  typeof useGetInstructorsQuery
>;
export type GetInstructorsLazyQueryHookResult = ReturnType<
  typeof useGetInstructorsLazyQuery
>;
export type GetInstructorsQueryResult = Apollo.QueryResult<
  GetInstructorsQuery,
  GetInstructorsQueryVariables
>;
export const GetInstructorDocument = gql`
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

/**
 * __useGetInstructorQuery__
 *
 * To run a query within a React component, call `useGetInstructorQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetInstructorQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetInstructorQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useGetInstructorQuery(
  baseOptions: Apollo.QueryHookOptions<
    GetInstructorQuery,
    GetInstructorQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useQuery<GetInstructorQuery, GetInstructorQueryVariables>(
    GetInstructorDocument,
    options
  );
}
export function useGetInstructorLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    GetInstructorQuery,
    GetInstructorQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useLazyQuery<GetInstructorQuery, GetInstructorQueryVariables>(
    GetInstructorDocument,
    options
  );
}
export type GetInstructorQueryHookResult = ReturnType<
  typeof useGetInstructorQuery
>;
export type GetInstructorLazyQueryHookResult = ReturnType<
  typeof useGetInstructorLazyQuery
>;
export type GetInstructorQueryResult = Apollo.QueryResult<
  GetInstructorQuery,
  GetInstructorQueryVariables
>;
