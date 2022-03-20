package repositories

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/ent/member"
	"github.com/matsuokashuhei/landin/graph/model"
)

type MemberRepository struct {
	client *ent.Client
}

func (r *MemberRepository) Find(ctx context.Context, id int) (*ent.Member, error) {
	return r.client.Member.Get(ctx, id)
}

func (r *MemberRepository) Create(ctx context.Context, input model.CreateMemberInput) (*ent.Member, error) {
	return r.client.Member.Create().
		SetName(input.Name).
		SetKana(input.Kana).
		SetGender(member.Gender(input.Gender)).
		SetNillableDateOfBirth(input.DateOfBirth).
		SetNillablePhoneNumber(input.PhoneNumber).
		SetNillableEmail(input.Email).
		SetDateOfAdmission(input.DateOfAdmission).
		SetNillableMemo(input.Memo).
		Save(ctx)
}

func (r *MemberRepository) Update(ctx context.Context, input model.UpdateMemberInput) (*ent.Member, error) {
	updation := r.client.Member.UpdateOneID(input.ID)
	if input.Name != nil {
		updation.SetName(*input.Name)
	}
	if input.Kana != nil {
		updation.SetKana(*input.Kana)
	}
	if input.Gender != nil {
		updation.SetGender(member.Gender(*input.Gender))
	}
	if input.DateOfBirth != nil {
		updation.SetDateOfBirth(*input.DateOfBirth)
	}
	if input.PhoneNumber != nil {
		updation.SetPhoneNumber(*input.PhoneNumber)
	}
	if input.Email != nil {
		updation.SetEmail(*input.Email)
	}
	if input.DateOfAdmission != nil {
		updation.SetDateOfAdmission(*input.DateOfAdmission)
	}
	if input.DateOfWithdrawal != nil {
		updation.SetDateOfWithdrawal(*input.DateOfWithdrawal)
	}
	if input.Memo != nil {
		updation.SetMemo(*input.Memo)
	}
	return updation.Save(ctx)
}

func (r *MemberRepository) Delete(ctx context.Context, input model.DeleteMemberInput) (*ent.Member, error) {
	member, err := r.Find(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if err := r.client.Member.DeleteOne(member).Exec(ctx); err != nil {
		return nil, err
	}
	return member, nil
}

func NewMemberRepsitory(client *ent.Client) *MemberRepository {
	return &MemberRepository{client: client}
}
