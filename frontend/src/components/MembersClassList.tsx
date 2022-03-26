import { VFC } from "react";
import { GetMemberQuery } from "../generated/graphql";

type Props = {
  membersClasses: GetMemberQuery["member"]["membersClasses"];
};

export const MembersClassList: VFC<Props> = ({ membersClasses }) => {
  const renderMembersClasses = (
    membersClasses: GetMemberQuery["member"]["membersClasses"]
  ) => {
    return membersClasses.map((membersClass) => (
      <tr key={membersClass.id}>
        <td>
          {membersClass.class.name}
          {membersClass.class.level}
        </td>
        <td>{membersClass.class.instructor.name}</td>
        <td>{membersClass.dateOfAdmission}</td>
        <td>{membersClass.dateOfWithdrawal}</td>
      </tr>
    ));
  };

  return (
    <table>
      <thead>
        <tr>
          <th>クラス</th>
          <th>インストラクター</th>
          <th>入会日</th>
          <th>退会日</th>
        </tr>
      </thead>
      <tbody>{renderMembersClasses(membersClasses)}</tbody>
    </table>
  );
};
