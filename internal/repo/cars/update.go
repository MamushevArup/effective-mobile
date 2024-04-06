package cars

import (
	"context"
	"github.com/MamushevArup/test-effective-mob/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *Repo) Update(ctx context.Context, car *models.Car, owner *models.Owner) error {
	// Start building the query for cars
	sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	sql, err := r.updateCar(ctx, sb, car)
	if err != nil {
		r.lg.Errorf("failed to update car: %v", err)
		return err
	}

	// If owner information is provided, update the owner table as well
	err = r.updateOwner(ctx, *owner, car, sb, sql)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) updateOwner(ctx context.Context, owner models.Owner, car *models.Car, sb squirrel.StatementBuilderType, sql squirrel.UpdateBuilder) error {
	if owner.Name != "" || owner.Surname != "" || owner.Patronymic != "" {
		// First, retrieve the owner_id from the cars table
		ownerId, err := r.ownerId(ctx, *car)
		if err != nil {
			r.lg.Errorf("failed to retrieve owner_id: %v", err)
			return err
		}

		// Then, update the owner table
		sql = sb.Update("owner")

		sql = stringQueryBuild(owner.Name, "name", sql)
		sql = stringQueryBuild(owner.Surname, "surname", sql)
		sql = stringQueryBuild(owner.Patronymic, "patronymic", sql)

		sql = sql.Where(squirrel.Eq{"id": ownerId})

		err = r.executeQuery(ctx, sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) updateCar(ctx context.Context, sb squirrel.StatementBuilderType, car *models.Car) (squirrel.UpdateBuilder, error) {
	sql := sb.Update("cars")

	sql = stringQueryBuild(car.Mark, "mark", sql)
	sql = stringQueryBuild(car.Model, "model", sql)
	if car.Year != 0 {
		sql = sql.Set("year", car.Year)
	}

	sql = sql.Where(squirrel.Eq{"registration_number": car.RegNum})

	err := r.executeQuery(ctx, sql)
	return sql, err
}

func stringQueryBuild(field, column string, sql squirrel.UpdateBuilder) squirrel.UpdateBuilder {
	if field != "" {
		sql = sql.Set(column, field)
	}
	return sql

}

func (r *Repo) ownerId(ctx context.Context, car models.Car) (uuid.UUID, error) {
	var ownerId uuid.UUID
	err := r.db.QueryRow(ctx, `SELECT owner_id FROM cars WHERE registration_number = $1`, car.RegNum).Scan(&ownerId)
	return ownerId, err
}

func (r *Repo) executeQuery(ctx context.Context, sql squirrel.UpdateBuilder) error {
	queryStr, args, err := sql.ToSql()
	if err != nil {
		r.lg.Errorf("failed to build query: %v", err)
		return err
	}

	_, err = r.db.Exec(ctx, queryStr, args...)
	if err != nil {
		r.lg.Errorf("failed to execute query: %v", err)
		return err
	}
	return nil
}
