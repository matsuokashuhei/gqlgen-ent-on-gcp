import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
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

export type Class = Node & {
  __typename?: 'Class';
  createTime: Scalars['Time'];
  endDate?: Maybe<Scalars['Time']>;
  id: Scalars['ID'];
  instructor: Instructor;
  level: Scalars['String'];
  name: Scalars['String'];
  schedule: Schedule;
  startDate: Scalars['Time'];
  tuition: Scalars['Int'];
  updateTime: Scalars['Time'];
};

export type CreateClassInput = {
  endDate?: InputMaybe<Scalars['Time']>;
  instructorId: Scalars['ID'];
  level: Scalars['String'];
  name: Scalars['String'];
  scheduleId: Scalars['ID'];
  startDate: Scalars['Time'];
  tuition: Scalars['Int'];
};

export type CreateInstructorInput = {
  biography?: InputMaybe<Scalars['String']>;
  email?: InputMaybe<Scalars['String']>;
  kana: Scalars['String'];
  name: Scalars['String'];
  phoneNumber?: InputMaybe<Scalars['String']>;
};

export type CreateMemberInput = {
  dateOfAdmission: Scalars['Time'];
  dateOfBirth?: InputMaybe<Scalars['Time']>;
  dateOfWithdrawal?: InputMaybe<Scalars['Time']>;
  email?: InputMaybe<Scalars['String']>;
  gender: Gender;
  kana: Scalars['String'];
  memo?: InputMaybe<Scalars['String']>;
  name: Scalars['String'];
  number: Scalars['Int'];
  phoneNumber?: InputMaybe<Scalars['String']>;
};

export type CreateMembersClassInput = {
  classId: Scalars['ID'];
  dateOfAdmission: Scalars['Time'];
  dateOfWithdrawal?: InputMaybe<Scalars['Time']>;
  memberId: Scalars['ID'];
};

export type CreateRoomInput = {
  capacity: Scalars['Int'];
  name: Scalars['String'];
  studioID: Scalars['ID'];
};

export type CreateSchoolInput = {
  name: Scalars['String'];
};

export type CreateStudioInput = {
  location: Scalars['String'];
  name: Scalars['String'];
  schoolID: Scalars['ID'];
};

export type CreateUserInput = {
  email: Scalars['String'];
  password: Scalars['String'];
};

export type DeleteInstructorInput = {
  id: Scalars['ID'];
};

export type DeleteMemberInput = {
  id: Scalars['ID'];
};

export type DeleteMembersClassInput = {
  id: Scalars['ID'];
};

export enum Gender {
  Female = 'FEMALE',
  Male = 'MALE',
  Other = 'OTHER'
}

export type Instructor = Node & {
  __typename?: 'Instructor';
  biography?: Maybe<Scalars['String']>;
  createTime: Scalars['Time'];
  email?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  kana: Scalars['String'];
  name: Scalars['String'];
  phoneNumber?: Maybe<Scalars['String']>;
  updateTime: Scalars['Time'];
};

export type InstructorConnection = {
  __typename?: 'InstructorConnection';
  edges: Array<InstructorEdge>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type InstructorEdge = {
  __typename?: 'InstructorEdge';
  cursor: Scalars['Cursor'];
  node: Instructor;
};

export type InstructorOrder = {
  direction: OrderDirection;
  field?: InputMaybe<InstructorOrderField>;
};

export enum InstructorOrderField {
  Kana = 'KANA'
}

export type Member = Node & {
  __typename?: 'Member';
  dateOfAdmission: Scalars['Time'];
  dateOfBirth?: Maybe<Scalars['Time']>;
  dateOfWithdrawal?: Maybe<Scalars['Time']>;
  email?: Maybe<Scalars['String']>;
  gender: Gender;
  id: Scalars['ID'];
  kana: Scalars['String'];
  membersClasses: Array<MembersClass>;
  memo: Scalars['String'];
  name: Scalars['String'];
  number: Scalars['Int'];
  phoneNumber?: Maybe<Scalars['String']>;
};

export type MemberConnection = {
  __typename?: 'MemberConnection';
  edges: Array<MemberEdge>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type MemberEdge = {
  __typename?: 'MemberEdge';
  cursor: Scalars['Cursor'];
  node: Member;
};

export type MemberOrder = {
  direction: OrderDirection;
  field?: InputMaybe<MemberOrderField>;
};

export enum MemberOrderField {
  DateOfAdmission = 'DATE_OF_ADMISSION',
  Kana = 'KANA',
  Number = 'NUMBER'
}

export type MembersClass = Node & {
  __typename?: 'MembersClass';
  class: Class;
  dateOfAdmission: Scalars['Time'];
  dateOfWithdrawal?: Maybe<Scalars['Time']>;
  id: Scalars['ID'];
  member: Member;
};

export type Mutation = {
  __typename?: 'Mutation';
  createClass: Class;
  createInstructor: Instructor;
  createMember: Member;
  createMembersClass: MembersClass;
  createRoom: Room;
  createSchool: School;
  createStudio: Studio;
  createUser: User;
  deleteClass: Class;
  deleteInstructor: Instructor;
  deleteMember: Member;
  deleteMembersClass: MembersClass;
  deleteRoom: Room;
  deleteSchool: School;
  deleteStudio: Studio;
  signUp: User;
  updateClass: Class;
  updateInstructor: Instructor;
  updateMember: Member;
  updateMembersClass: MembersClass;
  updateRoom: Room;
  updateSchool: School;
  updateStudio: Studio;
  updateUser: User;
};


export type MutationCreateClassArgs = {
  input: CreateClassInput;
};


export type MutationCreateInstructorArgs = {
  input: CreateInstructorInput;
};


export type MutationCreateMemberArgs = {
  input: CreateMemberInput;
};


export type MutationCreateMembersClassArgs = {
  input: CreateMembersClassInput;
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


export type MutationDeleteClassArgs = {
  id: Scalars['ID'];
};


export type MutationDeleteInstructorArgs = {
  input: DeleteInstructorInput;
};


export type MutationDeleteMemberArgs = {
  input: DeleteMemberInput;
};


export type MutationDeleteMembersClassArgs = {
  input: DeleteMembersClassInput;
};


export type MutationDeleteRoomArgs = {
  id: Scalars['ID'];
};


export type MutationDeleteSchoolArgs = {
  id: Scalars['ID'];
};


export type MutationDeleteStudioArgs = {
  id: Scalars['ID'];
};


export type MutationSignUpArgs = {
  input: SignUpInput;
};


export type MutationUpdateClassArgs = {
  input: UpdateClassInput;
};


export type MutationUpdateInstructorArgs = {
  input: UpdateInstructorInput;
};


export type MutationUpdateMemberArgs = {
  input: UpdateMemberInput;
};


export type MutationUpdateMembersClassArgs = {
  input: UpdateMembersClassInput;
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
  id: Scalars['ID'];
};

export enum OrderDirection {
  Asc = 'ASC',
  Desc = 'DESC'
}

export type PageInfo = {
  __typename?: 'PageInfo';
  endCursor?: Maybe<Scalars['Cursor']>;
  hasNextPage: Scalars['Boolean'];
  hasPreviousPage: Scalars['Boolean'];
  startCursor?: Maybe<Scalars['Cursor']>;
};

export type Query = {
  __typename?: 'Query';
  class: Class;
  currentUser: User;
  instructor: Instructor;
  instructors: InstructorConnection;
  member: Member;
  members: MemberConnection;
  membersClass: MembersClass;
  node?: Maybe<Node>;
  nodes: Array<Node>;
  room: Room;
  rooms: Array<Room>;
  schedule: Schedule;
  school: School;
  schools: Array<School>;
  studio: Studio;
  user: User;
  users: Array<User>;
};


export type QueryClassArgs = {
  id: Scalars['ID'];
};


export type QueryInstructorArgs = {
  id: Scalars['ID'];
};


export type QueryInstructorsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<InstructorOrder>;
};


export type QueryMemberArgs = {
  id: Scalars['ID'];
};


export type QueryMembersArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<MemberOrder>;
};


export type QueryMembersClassArgs = {
  id: Scalars['ID'];
};


export type QueryNodeArgs = {
  id: Scalars['ID'];
};


export type QueryNodesArgs = {
  ids: Array<Scalars['ID']>;
};


export type QueryRoomArgs = {
  id: Scalars['ID'];
};


export type QueryScheduleArgs = {
  id: Scalars['ID'];
};


export type QuerySchoolArgs = {
  id: Scalars['ID'];
};


export type QueryStudioArgs = {
  id: Scalars['ID'];
};


export type QueryUserArgs = {
  id: Scalars['ID'];
};

export type Room = Node & {
  __typename?: 'Room';
  capacity: Scalars['Int'];
  createTime: Scalars['Time'];
  id: Scalars['ID'];
  name: Scalars['String'];
  schedules: Array<Schedule>;
  studio: Studio;
  updateTime: Scalars['Time'];
};

export type Schedule = Node & {
  __typename?: 'Schedule';
  class?: Maybe<Class>;
  createTime: Scalars['Time'];
  dayOfWeek: Scalars['Int'];
  endTime: Scalars['String'];
  id: Scalars['ID'];
  startTime: Scalars['String'];
  updateTime: Scalars['Time'];
};


export type ScheduleClassArgs = {
  date?: InputMaybe<Scalars['Time']>;
};

export enum ScheduleField {
  DayOfWeek = 'DAY_OF_WEEK',
  StartTime = 'START_TIME'
}

export type School = Node & {
  __typename?: 'School';
  createTime: Scalars['Time'];
  id: Scalars['ID'];
  name: Scalars['String'];
  studios: Array<Studio>;
  updateTime: Scalars['Time'];
};

export type SignUpInput = {
  email: Scalars['String'];
  password: Scalars['String'];
};

export type Studio = Node & {
  __typename?: 'Studio';
  createTime: Scalars['Time'];
  id: Scalars['ID'];
  location: Scalars['String'];
  name: Scalars['String'];
  rooms: Array<Room>;
  school: School;
  updateTime: Scalars['Time'];
};

export type UpdateClassInput = {
  endDate?: InputMaybe<Scalars['Time']>;
  id: Scalars['ID'];
  instructorId?: InputMaybe<Scalars['ID']>;
  level?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
  scheduleId?: InputMaybe<Scalars['ID']>;
  startDate?: InputMaybe<Scalars['Time']>;
  tuition?: InputMaybe<Scalars['Int']>;
};

export type UpdateInstructorInput = {
  biography?: InputMaybe<Scalars['String']>;
  email?: InputMaybe<Scalars['String']>;
  id: Scalars['ID'];
  kana?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
  phoneNumber?: InputMaybe<Scalars['String']>;
};

export type UpdateMemberInput = {
  dateOfAdmission?: InputMaybe<Scalars['Time']>;
  dateOfBirth?: InputMaybe<Scalars['Time']>;
  dateOfWithdrawal?: InputMaybe<Scalars['Time']>;
  email?: InputMaybe<Scalars['String']>;
  gender?: InputMaybe<Gender>;
  id: Scalars['ID'];
  kana?: InputMaybe<Scalars['String']>;
  memo?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
  number?: InputMaybe<Scalars['Int']>;
  phoneNumber?: InputMaybe<Scalars['String']>;
};

export type UpdateMembersClassInput = {
  classId?: InputMaybe<Scalars['ID']>;
  dateOfAdmission?: InputMaybe<Scalars['Time']>;
  dateOfWithdrawal?: InputMaybe<Scalars['Time']>;
  id: Scalars['ID'];
  memberId?: InputMaybe<Scalars['ID']>;
};

export type UpdateRoomInput = {
  capacity?: InputMaybe<Scalars['Int']>;
  id: Scalars['ID'];
  name?: InputMaybe<Scalars['String']>;
  studioID?: InputMaybe<Scalars['ID']>;
};

export type UpdateSchoolInput = {
  id: Scalars['ID'];
  name: Scalars['String'];
};

export type UpdateStudioInput = {
  id: Scalars['ID'];
  location?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
  schoolID?: InputMaybe<Scalars['ID']>;
};

export type UpdateUserInput = {
  id: Scalars['ID'];
  name: Scalars['String'];
};

export type User = Node & {
  __typename?: 'User';
  createTime: Scalars['Time'];
  firebaseUid: Scalars['String'];
  id: Scalars['ID'];
  name: Scalars['String'];
  updateTime: Scalars['Time'];
};

export type CreateInstructorMutationVariables = Exact<{
  input: CreateInstructorInput;
}>;


export type CreateInstructorMutation = { __typename?: 'Mutation', createInstructor: { __typename?: 'Instructor', id: string, name: string, kana: string, biography?: string | null, phoneNumber?: string | null, email?: string | null } };

export type UpdateInstructorMutationVariables = Exact<{
  input: UpdateInstructorInput;
}>;


export type UpdateInstructorMutation = { __typename?: 'Mutation', updateInstructor: { __typename?: 'Instructor', id: string, name: string, kana: string, biography?: string | null, phoneNumber?: string | null, email?: string | null } };

export type DeleteInstructorMutationVariables = Exact<{
  input: DeleteInstructorInput;
}>;


export type DeleteInstructorMutation = { __typename?: 'Mutation', deleteInstructor: { __typename?: 'Instructor', id: string, name: string, kana: string, biography?: string | null, phoneNumber?: string | null, email?: string | null } };

export type CreateClassMutationVariables = Exact<{
  input: CreateClassInput;
}>;


export type CreateClassMutation = { __typename?: 'Mutation', createClass: { __typename?: 'Class', id: string, name: string, level: string, startDate: any, endDate?: any | null, instructor: { __typename?: 'Instructor', id: string }, schedule: { __typename?: 'Schedule', id: string } } };

export type UpdateClassMutationVariables = Exact<{
  input: UpdateClassInput;
}>;


export type UpdateClassMutation = { __typename?: 'Mutation', updateClass: { __typename?: 'Class', id: string, name: string, level: string, startDate: any, endDate?: any | null, instructor: { __typename?: 'Instructor', id: string }, schedule: { __typename?: 'Schedule', id: string } } };

export type DeleteClassMutationVariables = Exact<{
  id: Scalars['ID'];
}>;


export type DeleteClassMutation = { __typename?: 'Mutation', deleteClass: { __typename?: 'Class', id: string, name: string, level: string, startDate: any, endDate?: any | null } };

export type CreateMemberMutationVariables = Exact<{
  input: CreateMemberInput;
}>;


export type CreateMemberMutation = { __typename?: 'Mutation', createMember: { __typename?: 'Member', id: string, number: number, name: string, kana: string, gender: Gender, dateOfBirth?: any | null, phoneNumber?: string | null, email?: string | null, dateOfAdmission: any, dateOfWithdrawal?: any | null, memo: string } };

export type UpdateMemberMutationVariables = Exact<{
  input: UpdateMemberInput;
}>;


export type UpdateMemberMutation = { __typename?: 'Mutation', updateMember: { __typename?: 'Member', id: string, number: number, name: string, kana: string, gender: Gender, dateOfBirth?: any | null, phoneNumber?: string | null, email?: string | null, dateOfAdmission: any, dateOfWithdrawal?: any | null, memo: string } };

export type CreateMembersClassMutationVariables = Exact<{
  input: CreateMembersClassInput;
}>;


export type CreateMembersClassMutation = { __typename?: 'Mutation', createMembersClass: { __typename?: 'MembersClass', id: string, dateOfAdmission: any, dateOfWithdrawal?: any | null, member: { __typename?: 'Member', id: string }, class: { __typename?: 'Class', id: string } } };

export type UpdateMembersClassMutationVariables = Exact<{
  input: UpdateMembersClassInput;
}>;


export type UpdateMembersClassMutation = { __typename?: 'Mutation', updateMembersClass: { __typename?: 'MembersClass', id: string, dateOfAdmission: any, dateOfWithdrawal?: any | null, member: { __typename?: 'Member', id: string }, class: { __typename?: 'Class', id: string, name: string, level: string, instructor: { __typename?: 'Instructor', id: string, name: string } } } };

export type GetInstructorsQueryVariables = Exact<{
  first?: InputMaybe<Scalars['Int']>;
  after?: InputMaybe<Scalars['Cursor']>;
  last?: InputMaybe<Scalars['Int']>;
  before?: InputMaybe<Scalars['Cursor']>;
}>;


export type GetInstructorsQuery = { __typename?: 'Query', instructors: { __typename?: 'InstructorConnection', totalCount: number, pageInfo: { __typename?: 'PageInfo', startCursor?: any | null, endCursor?: any | null, hasNextPage: boolean, hasPreviousPage: boolean }, edges: Array<{ __typename?: 'InstructorEdge', cursor: any, node: { __typename?: 'Instructor', id: string, name: string, kana: string, biography?: string | null, phoneNumber?: string | null, email?: string | null } }> } };

export type GetInstructorQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type GetInstructorQuery = { __typename?: 'Query', instructor: { __typename?: 'Instructor', id: string, name: string, kana: string, biography?: string | null, phoneNumber?: string | null, email?: string | null } };

export type GetClassesBySchoolQueryVariables = Exact<{
  id: Scalars['ID'];
  date?: InputMaybe<Scalars['Time']>;
}>;


export type GetClassesBySchoolQuery = { __typename?: 'Query', school: { __typename?: 'School', id: string, name: string, studios: Array<{ __typename?: 'Studio', id: string, name: string, location: string, rooms: Array<{ __typename?: 'Room', id: string, name: string, capacity: number, schedules: Array<{ __typename?: 'Schedule', id: string, dayOfWeek: number, startTime: string, endTime: string, class?: { __typename?: 'Class', id: string, name: string, level: string, tuition: number, startDate: any, endDate?: any | null, instructor: { __typename?: 'Instructor', id: string, name: string } } | null }> }> }> } };

export type GetClassQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type GetClassQuery = { __typename?: 'Query', class: { __typename?: 'Class', id: string, name: string, level: string, tuition: number, startDate: any, endDate?: any | null, instructor: { __typename?: 'Instructor', id: string, name: string }, schedule: { __typename?: 'Schedule', id: string } } };

export type GetScheduleAndInstructorsToRegisterNewClassQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type GetScheduleAndInstructorsToRegisterNewClassQuery = { __typename?: 'Query', schedule: { __typename?: 'Schedule', id: string, dayOfWeek: number, startTime: string, endTime: string }, instructors: { __typename?: 'InstructorConnection', totalCount: number, edges: Array<{ __typename?: 'InstructorEdge', node: { __typename?: 'Instructor', id: string, name: string } }> } };

export type GetMembersQueryVariables = Exact<{
  first?: InputMaybe<Scalars['Int']>;
  after?: InputMaybe<Scalars['Cursor']>;
  last?: InputMaybe<Scalars['Int']>;
  before?: InputMaybe<Scalars['Cursor']>;
}>;


export type GetMembersQuery = { __typename?: 'Query', members: { __typename?: 'MemberConnection', totalCount: number, pageInfo: { __typename?: 'PageInfo', startCursor?: any | null, endCursor?: any | null, hasNextPage: boolean, hasPreviousPage: boolean }, edges: Array<{ __typename?: 'MemberEdge', cursor: any, node: { __typename?: 'Member', id: string, number: number, name: string, kana: string, gender: Gender, dateOfAdmission: any } }> } };

export type GetMemberQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type GetMemberQuery = { __typename?: 'Query', member: { __typename?: 'Member', id: string, number: number, name: string, kana: string, gender: Gender, dateOfBirth?: any | null, phoneNumber?: string | null, email?: string | null, dateOfAdmission: any, dateOfWithdrawal?: any | null, memo: string, membersClasses: Array<{ __typename?: 'MembersClass', id: string, dateOfAdmission: any, dateOfWithdrawal?: any | null, member: { __typename?: 'Member', id: string }, class: { __typename?: 'Class', id: string, name: string, level: string, instructor: { __typename?: 'Instructor', id: string, name: string } } }> } };

export type GetMembersClassQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type GetMembersClassQuery = { __typename?: 'Query', membersClass: { __typename?: 'MembersClass', id: string, dateOfAdmission: any, dateOfWithdrawal?: any | null, member: { __typename?: 'Member', id: string }, class: { __typename?: 'Class', id: string, name: string, level: string, instructor: { __typename?: 'Instructor', name: string } } } };


export const CreateInstructorDocument = gql`
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
export type CreateInstructorMutationFn = Apollo.MutationFunction<CreateInstructorMutation, CreateInstructorMutationVariables>;

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
export function useCreateInstructorMutation(baseOptions?: Apollo.MutationHookOptions<CreateInstructorMutation, CreateInstructorMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateInstructorMutation, CreateInstructorMutationVariables>(CreateInstructorDocument, options);
      }
export type CreateInstructorMutationHookResult = ReturnType<typeof useCreateInstructorMutation>;
export type CreateInstructorMutationResult = Apollo.MutationResult<CreateInstructorMutation>;
export type CreateInstructorMutationOptions = Apollo.BaseMutationOptions<CreateInstructorMutation, CreateInstructorMutationVariables>;
export const UpdateInstructorDocument = gql`
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
export type UpdateInstructorMutationFn = Apollo.MutationFunction<UpdateInstructorMutation, UpdateInstructorMutationVariables>;

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
export function useUpdateInstructorMutation(baseOptions?: Apollo.MutationHookOptions<UpdateInstructorMutation, UpdateInstructorMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateInstructorMutation, UpdateInstructorMutationVariables>(UpdateInstructorDocument, options);
      }
export type UpdateInstructorMutationHookResult = ReturnType<typeof useUpdateInstructorMutation>;
export type UpdateInstructorMutationResult = Apollo.MutationResult<UpdateInstructorMutation>;
export type UpdateInstructorMutationOptions = Apollo.BaseMutationOptions<UpdateInstructorMutation, UpdateInstructorMutationVariables>;
export const DeleteInstructorDocument = gql`
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
export type DeleteInstructorMutationFn = Apollo.MutationFunction<DeleteInstructorMutation, DeleteInstructorMutationVariables>;

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
export function useDeleteInstructorMutation(baseOptions?: Apollo.MutationHookOptions<DeleteInstructorMutation, DeleteInstructorMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteInstructorMutation, DeleteInstructorMutationVariables>(DeleteInstructorDocument, options);
      }
export type DeleteInstructorMutationHookResult = ReturnType<typeof useDeleteInstructorMutation>;
export type DeleteInstructorMutationResult = Apollo.MutationResult<DeleteInstructorMutation>;
export type DeleteInstructorMutationOptions = Apollo.BaseMutationOptions<DeleteInstructorMutation, DeleteInstructorMutationVariables>;
export const CreateClassDocument = gql`
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
export type CreateClassMutationFn = Apollo.MutationFunction<CreateClassMutation, CreateClassMutationVariables>;

/**
 * __useCreateClassMutation__
 *
 * To run a mutation, you first call `useCreateClassMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateClassMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createClassMutation, { data, loading, error }] = useCreateClassMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateClassMutation(baseOptions?: Apollo.MutationHookOptions<CreateClassMutation, CreateClassMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateClassMutation, CreateClassMutationVariables>(CreateClassDocument, options);
      }
export type CreateClassMutationHookResult = ReturnType<typeof useCreateClassMutation>;
export type CreateClassMutationResult = Apollo.MutationResult<CreateClassMutation>;
export type CreateClassMutationOptions = Apollo.BaseMutationOptions<CreateClassMutation, CreateClassMutationVariables>;
export const UpdateClassDocument = gql`
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
export type UpdateClassMutationFn = Apollo.MutationFunction<UpdateClassMutation, UpdateClassMutationVariables>;

/**
 * __useUpdateClassMutation__
 *
 * To run a mutation, you first call `useUpdateClassMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateClassMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateClassMutation, { data, loading, error }] = useUpdateClassMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUpdateClassMutation(baseOptions?: Apollo.MutationHookOptions<UpdateClassMutation, UpdateClassMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateClassMutation, UpdateClassMutationVariables>(UpdateClassDocument, options);
      }
export type UpdateClassMutationHookResult = ReturnType<typeof useUpdateClassMutation>;
export type UpdateClassMutationResult = Apollo.MutationResult<UpdateClassMutation>;
export type UpdateClassMutationOptions = Apollo.BaseMutationOptions<UpdateClassMutation, UpdateClassMutationVariables>;
export const DeleteClassDocument = gql`
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
export type DeleteClassMutationFn = Apollo.MutationFunction<DeleteClassMutation, DeleteClassMutationVariables>;

/**
 * __useDeleteClassMutation__
 *
 * To run a mutation, you first call `useDeleteClassMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteClassMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteClassMutation, { data, loading, error }] = useDeleteClassMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteClassMutation(baseOptions?: Apollo.MutationHookOptions<DeleteClassMutation, DeleteClassMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteClassMutation, DeleteClassMutationVariables>(DeleteClassDocument, options);
      }
export type DeleteClassMutationHookResult = ReturnType<typeof useDeleteClassMutation>;
export type DeleteClassMutationResult = Apollo.MutationResult<DeleteClassMutation>;
export type DeleteClassMutationOptions = Apollo.BaseMutationOptions<DeleteClassMutation, DeleteClassMutationVariables>;
export const CreateMemberDocument = gql`
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
export type CreateMemberMutationFn = Apollo.MutationFunction<CreateMemberMutation, CreateMemberMutationVariables>;

/**
 * __useCreateMemberMutation__
 *
 * To run a mutation, you first call `useCreateMemberMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateMemberMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createMemberMutation, { data, loading, error }] = useCreateMemberMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateMemberMutation(baseOptions?: Apollo.MutationHookOptions<CreateMemberMutation, CreateMemberMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateMemberMutation, CreateMemberMutationVariables>(CreateMemberDocument, options);
      }
export type CreateMemberMutationHookResult = ReturnType<typeof useCreateMemberMutation>;
export type CreateMemberMutationResult = Apollo.MutationResult<CreateMemberMutation>;
export type CreateMemberMutationOptions = Apollo.BaseMutationOptions<CreateMemberMutation, CreateMemberMutationVariables>;
export const UpdateMemberDocument = gql`
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
export type UpdateMemberMutationFn = Apollo.MutationFunction<UpdateMemberMutation, UpdateMemberMutationVariables>;

/**
 * __useUpdateMemberMutation__
 *
 * To run a mutation, you first call `useUpdateMemberMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateMemberMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateMemberMutation, { data, loading, error }] = useUpdateMemberMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUpdateMemberMutation(baseOptions?: Apollo.MutationHookOptions<UpdateMemberMutation, UpdateMemberMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateMemberMutation, UpdateMemberMutationVariables>(UpdateMemberDocument, options);
      }
export type UpdateMemberMutationHookResult = ReturnType<typeof useUpdateMemberMutation>;
export type UpdateMemberMutationResult = Apollo.MutationResult<UpdateMemberMutation>;
export type UpdateMemberMutationOptions = Apollo.BaseMutationOptions<UpdateMemberMutation, UpdateMemberMutationVariables>;
export const CreateMembersClassDocument = gql`
    mutation CreateMembersClass($input: CreateMembersClassInput!) {
  createMembersClass(input: $input) {
    id
    member {
      id
    }
    class {
      id
    }
    dateOfAdmission
    dateOfWithdrawal
  }
}
    `;
export type CreateMembersClassMutationFn = Apollo.MutationFunction<CreateMembersClassMutation, CreateMembersClassMutationVariables>;

/**
 * __useCreateMembersClassMutation__
 *
 * To run a mutation, you first call `useCreateMembersClassMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateMembersClassMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createMembersClassMutation, { data, loading, error }] = useCreateMembersClassMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateMembersClassMutation(baseOptions?: Apollo.MutationHookOptions<CreateMembersClassMutation, CreateMembersClassMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateMembersClassMutation, CreateMembersClassMutationVariables>(CreateMembersClassDocument, options);
      }
export type CreateMembersClassMutationHookResult = ReturnType<typeof useCreateMembersClassMutation>;
export type CreateMembersClassMutationResult = Apollo.MutationResult<CreateMembersClassMutation>;
export type CreateMembersClassMutationOptions = Apollo.BaseMutationOptions<CreateMembersClassMutation, CreateMembersClassMutationVariables>;
export const UpdateMembersClassDocument = gql`
    mutation UpdateMembersClass($input: UpdateMembersClassInput!) {
  updateMembersClass(input: $input) {
    id
    member {
      id
    }
    class {
      id
      name
      level
      instructor {
        id
        name
      }
    }
    dateOfAdmission
    dateOfWithdrawal
  }
}
    `;
export type UpdateMembersClassMutationFn = Apollo.MutationFunction<UpdateMembersClassMutation, UpdateMembersClassMutationVariables>;

/**
 * __useUpdateMembersClassMutation__
 *
 * To run a mutation, you first call `useUpdateMembersClassMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateMembersClassMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateMembersClassMutation, { data, loading, error }] = useUpdateMembersClassMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUpdateMembersClassMutation(baseOptions?: Apollo.MutationHookOptions<UpdateMembersClassMutation, UpdateMembersClassMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateMembersClassMutation, UpdateMembersClassMutationVariables>(UpdateMembersClassDocument, options);
      }
export type UpdateMembersClassMutationHookResult = ReturnType<typeof useUpdateMembersClassMutation>;
export type UpdateMembersClassMutationResult = Apollo.MutationResult<UpdateMembersClassMutation>;
export type UpdateMembersClassMutationOptions = Apollo.BaseMutationOptions<UpdateMembersClassMutation, UpdateMembersClassMutationVariables>;
export const GetInstructorsDocument = gql`
    query getInstructors($first: Int, $after: Cursor, $last: Int, $before: Cursor) {
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
export function useGetInstructorsQuery(baseOptions?: Apollo.QueryHookOptions<GetInstructorsQuery, GetInstructorsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetInstructorsQuery, GetInstructorsQueryVariables>(GetInstructorsDocument, options);
      }
export function useGetInstructorsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetInstructorsQuery, GetInstructorsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetInstructorsQuery, GetInstructorsQueryVariables>(GetInstructorsDocument, options);
        }
export type GetInstructorsQueryHookResult = ReturnType<typeof useGetInstructorsQuery>;
export type GetInstructorsLazyQueryHookResult = ReturnType<typeof useGetInstructorsLazyQuery>;
export type GetInstructorsQueryResult = Apollo.QueryResult<GetInstructorsQuery, GetInstructorsQueryVariables>;
export const GetInstructorDocument = gql`
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
export function useGetInstructorQuery(baseOptions: Apollo.QueryHookOptions<GetInstructorQuery, GetInstructorQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetInstructorQuery, GetInstructorQueryVariables>(GetInstructorDocument, options);
      }
export function useGetInstructorLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetInstructorQuery, GetInstructorQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetInstructorQuery, GetInstructorQueryVariables>(GetInstructorDocument, options);
        }
export type GetInstructorQueryHookResult = ReturnType<typeof useGetInstructorQuery>;
export type GetInstructorLazyQueryHookResult = ReturnType<typeof useGetInstructorLazyQuery>;
export type GetInstructorQueryResult = Apollo.QueryResult<GetInstructorQuery, GetInstructorQueryVariables>;
export const GetClassesBySchoolDocument = gql`
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

/**
 * __useGetClassesBySchoolQuery__
 *
 * To run a query within a React component, call `useGetClassesBySchoolQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetClassesBySchoolQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetClassesBySchoolQuery({
 *   variables: {
 *      id: // value for 'id'
 *      date: // value for 'date'
 *   },
 * });
 */
export function useGetClassesBySchoolQuery(baseOptions: Apollo.QueryHookOptions<GetClassesBySchoolQuery, GetClassesBySchoolQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetClassesBySchoolQuery, GetClassesBySchoolQueryVariables>(GetClassesBySchoolDocument, options);
      }
export function useGetClassesBySchoolLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetClassesBySchoolQuery, GetClassesBySchoolQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetClassesBySchoolQuery, GetClassesBySchoolQueryVariables>(GetClassesBySchoolDocument, options);
        }
export type GetClassesBySchoolQueryHookResult = ReturnType<typeof useGetClassesBySchoolQuery>;
export type GetClassesBySchoolLazyQueryHookResult = ReturnType<typeof useGetClassesBySchoolLazyQuery>;
export type GetClassesBySchoolQueryResult = Apollo.QueryResult<GetClassesBySchoolQuery, GetClassesBySchoolQueryVariables>;
export const GetClassDocument = gql`
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

/**
 * __useGetClassQuery__
 *
 * To run a query within a React component, call `useGetClassQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetClassQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetClassQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useGetClassQuery(baseOptions: Apollo.QueryHookOptions<GetClassQuery, GetClassQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetClassQuery, GetClassQueryVariables>(GetClassDocument, options);
      }
export function useGetClassLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetClassQuery, GetClassQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetClassQuery, GetClassQueryVariables>(GetClassDocument, options);
        }
export type GetClassQueryHookResult = ReturnType<typeof useGetClassQuery>;
export type GetClassLazyQueryHookResult = ReturnType<typeof useGetClassLazyQuery>;
export type GetClassQueryResult = Apollo.QueryResult<GetClassQuery, GetClassQueryVariables>;
export const GetScheduleAndInstructorsToRegisterNewClassDocument = gql`
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

/**
 * __useGetScheduleAndInstructorsToRegisterNewClassQuery__
 *
 * To run a query within a React component, call `useGetScheduleAndInstructorsToRegisterNewClassQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetScheduleAndInstructorsToRegisterNewClassQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetScheduleAndInstructorsToRegisterNewClassQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useGetScheduleAndInstructorsToRegisterNewClassQuery(baseOptions: Apollo.QueryHookOptions<GetScheduleAndInstructorsToRegisterNewClassQuery, GetScheduleAndInstructorsToRegisterNewClassQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetScheduleAndInstructorsToRegisterNewClassQuery, GetScheduleAndInstructorsToRegisterNewClassQueryVariables>(GetScheduleAndInstructorsToRegisterNewClassDocument, options);
      }
export function useGetScheduleAndInstructorsToRegisterNewClassLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetScheduleAndInstructorsToRegisterNewClassQuery, GetScheduleAndInstructorsToRegisterNewClassQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetScheduleAndInstructorsToRegisterNewClassQuery, GetScheduleAndInstructorsToRegisterNewClassQueryVariables>(GetScheduleAndInstructorsToRegisterNewClassDocument, options);
        }
export type GetScheduleAndInstructorsToRegisterNewClassQueryHookResult = ReturnType<typeof useGetScheduleAndInstructorsToRegisterNewClassQuery>;
export type GetScheduleAndInstructorsToRegisterNewClassLazyQueryHookResult = ReturnType<typeof useGetScheduleAndInstructorsToRegisterNewClassLazyQuery>;
export type GetScheduleAndInstructorsToRegisterNewClassQueryResult = Apollo.QueryResult<GetScheduleAndInstructorsToRegisterNewClassQuery, GetScheduleAndInstructorsToRegisterNewClassQueryVariables>;
export const GetMembersDocument = gql`
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

/**
 * __useGetMembersQuery__
 *
 * To run a query within a React component, call `useGetMembersQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetMembersQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetMembersQuery({
 *   variables: {
 *      first: // value for 'first'
 *      after: // value for 'after'
 *      last: // value for 'last'
 *      before: // value for 'before'
 *   },
 * });
 */
export function useGetMembersQuery(baseOptions?: Apollo.QueryHookOptions<GetMembersQuery, GetMembersQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetMembersQuery, GetMembersQueryVariables>(GetMembersDocument, options);
      }
export function useGetMembersLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetMembersQuery, GetMembersQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetMembersQuery, GetMembersQueryVariables>(GetMembersDocument, options);
        }
export type GetMembersQueryHookResult = ReturnType<typeof useGetMembersQuery>;
export type GetMembersLazyQueryHookResult = ReturnType<typeof useGetMembersLazyQuery>;
export type GetMembersQueryResult = Apollo.QueryResult<GetMembersQuery, GetMembersQueryVariables>;
export const GetMemberDocument = gql`
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
      member {
        id
      }
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

/**
 * __useGetMemberQuery__
 *
 * To run a query within a React component, call `useGetMemberQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetMemberQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetMemberQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useGetMemberQuery(baseOptions: Apollo.QueryHookOptions<GetMemberQuery, GetMemberQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetMemberQuery, GetMemberQueryVariables>(GetMemberDocument, options);
      }
export function useGetMemberLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetMemberQuery, GetMemberQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetMemberQuery, GetMemberQueryVariables>(GetMemberDocument, options);
        }
export type GetMemberQueryHookResult = ReturnType<typeof useGetMemberQuery>;
export type GetMemberLazyQueryHookResult = ReturnType<typeof useGetMemberLazyQuery>;
export type GetMemberQueryResult = Apollo.QueryResult<GetMemberQuery, GetMemberQueryVariables>;
export const GetMembersClassDocument = gql`
    query getMembersClass($id: ID!) {
  membersClass(id: $id) {
    id
    member {
      id
    }
    class {
      id
      name
      level
      instructor {
        name
      }
    }
    dateOfAdmission
    dateOfWithdrawal
  }
}
    `;

/**
 * __useGetMembersClassQuery__
 *
 * To run a query within a React component, call `useGetMembersClassQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetMembersClassQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetMembersClassQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useGetMembersClassQuery(baseOptions: Apollo.QueryHookOptions<GetMembersClassQuery, GetMembersClassQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetMembersClassQuery, GetMembersClassQueryVariables>(GetMembersClassDocument, options);
      }
export function useGetMembersClassLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetMembersClassQuery, GetMembersClassQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetMembersClassQuery, GetMembersClassQueryVariables>(GetMembersClassDocument, options);
        }
export type GetMembersClassQueryHookResult = ReturnType<typeof useGetMembersClassQuery>;
export type GetMembersClassLazyQueryHookResult = ReturnType<typeof useGetMembersClassLazyQuery>;
export type GetMembersClassQueryResult = Apollo.QueryResult<GetMembersClassQuery, GetMembersClassQueryVariables>;