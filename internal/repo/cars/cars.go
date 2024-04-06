package cars

import (
	"context"
	"github.com/MamushevArup/test-effective-mob/internal/models"
	"github.com/Masterminds/squirrel"
)

func (r *Repo) Cars(ctx context.Context, car models.Car, limit, offset int) ([]models.Car, error) {
	sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	sql := sb.Select("*").From("cars")

	if car.Mark != "" {
		sql = sql.Where(squirrel.Eq{"mark": car.Mark})
	}
	if car.Model != "" {
		sql = sql.Where(squirrel.Eq{"model": car.Model})
	}
	if car.Year != 0 {
		sql = sql.Where(squirrel.Eq{"year": car.Year})
	}

	sql = sql.Limit(uint64(limit)).Offset(uint64(offset))

	queryStr, args, err := sql.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var c models.Car
		err = rows.Scan(&car.RegNum, &c.Mark, &c.Model, &c.Year)
		if err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}

	return cars, nil
}
