package repository

import (
	"fmt"
	"log/slog"
	"strings"

	domain "github.com/egor/watcher/pkg/model"
	"github.com/jmoiron/sqlx"
)

type DomainTargetPostgres struct {
	db *sqlx.DB
}

func NewDomainTargetPostgres(db *sqlx.DB) *DomainTargetPostgres {
	return &DomainTargetPostgres{db: db}
}

func (r *DomainTargetPostgres) Create(userId int, target domain.Target) (int, error) {

	var targetId int
	CreateTargetQuery := fmt.Sprintf("INSERT INTO %s(url, title, user_id) VALUES ($1, $2, $3) RETURNING id", targetsTable)
	row := r.db.QueryRow(CreateTargetQuery, target.URL, target.Title, userId)
	if err := row.Scan(&targetId); err != nil {
		return 0, err
	}

	return targetId, nil

}
func (r *DomainTargetPostgres) GetAll(userId int) ([]domain.Target, error) {
	var targets []domain.Target

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", targetsTable)

	err := r.db.Select(&targets, query, userId)

	return targets, err

}
func (r *DomainTargetPostgres) GetById(userId, targetId int) (domain.Target, error) {
	var target domain.Target

	query := fmt.Sprintf("SELECT id, url, user_id FROM %s WHERE id = $1 AND user_id = $2", targetsTable)

	err := r.db.Get(&target, query, targetId, userId)

	return target, err
}
func (r *DomainTargetPostgres) Update(userId, targetId int, input domain.UpdateTargetInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.URL != nil {
		setValues = append(setValues, fmt.Sprintf("url = $%d", argId))
		args = append(args, *input.URL)
		argId++
	}
	if input.Interval != nil {
		setValues = append(setValues, fmt.Sprintf("interval = $%d", argId))
		args = append(args, *input.Interval)
		argId++
	}

	if len(setValues) == 0 {
		return nil
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND user_id=$%d",
		targetsTable, setQuery, argId, argId+1)
	args = append(args, targetId, userId)

	slog.Debug("database update trace", "query", query, "args", args)

	_, err := r.db.Exec(query, args...)

	return err

}
func (r *DomainTargetPostgres) Delete(userId, targetId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", targetsTable)
	_, err := r.db.Exec(query, targetId, userId)

	return err

}

func (r *DomainTargetPostgres) GetAllForWorker() ([]domain.Target, error) {
	var targets []domain.Target
	query := fmt.Sprintf("SELECT * FROM %s", targetsTable)
	err := r.db.Select(&targets, query)
	return targets, err
}

func (r *DomainTargetPostgres) UpdateStatus(id int, status bool) error {
	query := fmt.Sprintf("UPDATE %s SET status = $1, last_check = NOW() WHERE id = $2", targetsTable)
	_, err := r.db.Exec(query, status, id)
	return err
}
