package cars

import (
	"context"
	"errors"
	"fmt"
	"github.com/MamushevArup/test-effective-mob/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repo) Insert(ctx context.Context, car *models.Car, owner models.Owner) error {
	fmt.Println(owner)
	insertCar := `INSERT INTO cars VALUES ($1, $2, $3, $4, $5)`
	insertOwner := `INSERT INTO owner VALUES ($1, $2, $3, $4)`

	tx, err := r.db.Begin(ctx)
	if err != nil {
		r.lg.Errorf("failed to begin transaction: %v", err)
		return err
	}
	r.lg.Info("transaction started")

	exec, err := tx.Exec(ctx, insertOwner, owner.Id, owner.Name, owner.Surname, owner.Patronymic)

	err = r.insertCheck("failed to insert owner", err, exec)
	if err != nil {
		return err
	}
	exec, err = tx.Exec(ctx, insertCar, car.RegNum, car.Mark, car.Model, car.Year, owner.Id)

	err = r.insertCheck("failed to insert car", err, exec)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.lg.Errorf("failed to commit transaction: %v", err)
		return err
	}

	return nil
}

func (r *Repo) insertCheck(message string, err error, exec pgconn.CommandTag) error {
	if err != nil {
		r.lg.Errorf("%v: %v", message, err)
		return err
	}
	if !exec.Insert() {
		r.lg.Errorf("%v: %v", message, err)
		return errors.New(message)
	}
	return nil
}
