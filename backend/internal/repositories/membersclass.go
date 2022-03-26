package repositories

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
)

type MembersClassRepository struct {
	client *ent.Client
}

func (r *MembersClassRepository) Find(ctx context.Context, id int) (*ent.MembersClass, error) {
	return r.client.MembersClass.Get(ctx, id)
}

func (r *MembersClassRepository) Create(ctx context.Context, input model.CreateMembersClassInput) (*ent.MembersClass, error) {
	return r.client.MembersClass.Create().
		SetMemberID(input.MemberID).
		SetClassID(input.ClassID).
		SetDateOfAdmission(input.DateOfAdmission).
		SetNillableDateOfWithdrawal(input.DateOfWithdrawal).
		Save(ctx)
}

func (r *MembersClassRepository) Update(ctx context.Context, input model.UpdateMembersClassInput) (*ent.MembersClass, error) {
	updation := r.client.MembersClass.UpdateOneID(input.ID)
	if input.MemberID != nil {
		updation.SetMemberID(*input.MemberID)
	}
	if input.ClassID != nil {
		updation.SetClassID(*input.ClassID)
	}
	if input.DateOfAdmission != nil {
		updation.SetDateOfAdmission(*input.DateOfAdmission)
	}
	if input.DateOfWithdrawal != nil {
		updation.SetDateOfWithdrawal(*input.DateOfWithdrawal)
	}
	return updation.Save(ctx)
}

func (r *MembersClassRepository) Delete(ctx context.Context, input model.DeleteMembersClassInput) (*ent.MembersClass, error) {
	membersclass, err := r.Find(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if err := r.client.MembersClass.DeleteOne(membersclass).Exec(ctx); err != nil {
		return nil, err
	}
	return membersclass, nil
}

func NewMembersClassRepository(client *ent.Client) *MembersClassRepository {
	return &MembersClassRepository{client: client}
}
