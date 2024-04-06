package cars

import (
	"context"
	"errors"
)

func (r *Repo) Delete(ctx context.Context, regNum string) error {
	query := `DELETE FROM cars WHERE registration_number = $1`

	exec, err := r.db.Exec(ctx, query, regNum)
	if err != nil {
		r.lg.Errorf("failed to execute query: %v", err)
		return err
	}

	if !exec.Delete() {
		r.lg.Errorf("failed to delete car: %v", err)
		return errors.New("failed to delete car")
	}
	return nil
}
