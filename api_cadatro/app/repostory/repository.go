package repostory

import (
	"context"
	"fmt"
)

const (
	register = `
	INSERT INTO 
	    merchant(
			email, name, phone
	)values (
         $1, $2, $3
	)
`
)

type valuesDB struct {
	Email string
	Name  string
	Phone string
}

type toolsDatabase struct {
	db DB
}

func (t *toolsDatabase) Create(ctx context.Context, v valuesDB) error {
	_, err := t.db.Exec(ctx, register,v.Email, v.Name, v.Phone)
	if err != nil {
		return fmt.Errorf("error on create on database %v", err)
	}
	return nil
}
