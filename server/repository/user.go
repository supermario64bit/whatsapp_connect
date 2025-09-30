package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/supermario64bit/whatsapp_connect/db"
	"github.com/supermario64bit/whatsapp_connect/server/model"
)

const user_table_name string = "user"

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Find(filter *model.User) ([]*model.User, error)
	FindByID(id uint64) (*model.User, error)
	UpdateByID(updates *model.User, id uint64) (*model.User, error)
	DeleteByID(id uint64) error
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.New(),
	}
}

func (repo *userRepository) Create(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, fmt.Errorf("Cannot create organisation for nil reference")
	}

	if user.ID > 0 {
		return nil, fmt.Errorf("ID field should be empty")
	}

	if user.Name == "" || user.Handle == "" || user.Mobile == "" || user.Email == "" || user.Status == "" {
		return nil, fmt.Errorf("Name, Handle, Mobile Number. Email and Status field should not be empty")

	}

	if len(user.Mobile) != 10 {
		return nil, fmt.Errorf("Contact Number should be 10 digit")
	}

	colNames := []string{"name", "handle", "mobile_number", "email", "status"}
	values := [][]interface{}{
		{user.Name, user.Handle, user.Mobile, user.Email, user.Status},
	}

	qry, args := generateInsertQuery(user_table_name, colNames, values)

	var createdOrg model.User
	err := repo.db.QueryRow(qry, args...).Scan(&createdOrg.ID, &createdOrg.Name, &createdOrg.Handle, &createdOrg.Mobile, &createdOrg.Email, &createdOrg.Status, &createdOrg.CreatedAt, &createdOrg.UpdatedAt, &createdOrg.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &createdOrg, nil
}

func (repo *userRepository) Find(filter *model.User) ([]*model.User, error) {
	args := []interface{}{}
	whereParts := []string{}
	if filter != nil {
		if strings.TrimSpace(filter.Name) != "" {
			args = append(args, "%"+strings.TrimSpace(filter.Name)+"%")
			whereParts = append(whereParts, fmt.Sprintf("name LIKE $%d", len(args)))
		}

		if strings.TrimSpace(filter.Handle) != "" {
			args = append(args, "%"+strings.TrimSpace(filter.Handle)+"%")
			whereParts = append(whereParts, fmt.Sprintf("handle LIKE $%d", len(args)))

		}

		if strings.TrimSpace(filter.Mobile) != "" {
			args = append(args, "%"+strings.TrimSpace(filter.Mobile)+"%")
			whereParts = append(whereParts, fmt.Sprintf("mobile_number LIKE $%d", len(args)))

		}

		if strings.TrimSpace(filter.Email) != "" {
			args = append(args, "%"+strings.TrimSpace(filter.Email)+"%")
			whereParts = append(whereParts, fmt.Sprintf("email LIKE $%d", len(args)))
		}

		if strings.TrimSpace(filter.Status) != "" {
			args = append(args, "%"+strings.TrimSpace(filter.Status)+"%")
			whereParts = append(whereParts, fmt.Sprintf("status LIKE $%d", len(args)))
		}
	}

	whereClause := ""
	if len(whereClause) > 0 {
		whereClause = "WHERE " + strings.Join(whereParts, " AND ")
	}

	if whereClause != "" {
		whereClause = whereClause + " AND deleted_at IS NULL"
	} else {
		whereClause = whereClause + " WHERE deleted_at IS NULL"
	}

	qry := fmt.Sprintf("SELECT * FROM %s ", user_table_name) + whereClause
	rows, err := repo.db.Query(qry, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		var user model.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Handle,
			&user.Mobile,
			&user.Email,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (repo *userRepository) FindByID(id uint64) (*model.User, error) {
	qry := "SELECT * FROM " + user_table_name + " WHERE id = $1 AND deleted_at IS NULL LIMIT 1"

	var user model.User

	err := repo.db.QueryRow(qry, id).Scan(
		&user.ID,
		&user.Name,
		&user.Handle,
		&user.Mobile,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) UpdateByID(updates *model.User, id uint64) (*model.User, error) {
	_, err := repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	updatesParam := []string{}
	args := []interface{}{}
	argPos := 1
	if strings.TrimSpace(updates.Name) != "" {
		updatesParam = append(updatesParam, fmt.Sprintf("name = $%d", argPos))
		args = append(args, strings.TrimSpace(updates.Name))
		argPos++
	}

	if strings.TrimSpace(updates.Handle) != "" {
		updatesParam = append(updatesParam, fmt.Sprintf("handle = $%d", argPos))
		args = append(args, strings.TrimSpace(updates.Handle))
		argPos++
	}

	if strings.TrimSpace(updates.Mobile) != "" {
		updatesParam = append(updatesParam, fmt.Sprintf("mobile_number = $%d", argPos))
		args = append(args, strings.TrimSpace(updates.Mobile))
		argPos++
	}

	if strings.TrimSpace(updates.Email) != "" {
		updatesParam = append(updatesParam, fmt.Sprintf("email = $%d", argPos))
		args = append(args, strings.TrimSpace(updates.Email))
		argPos++
	}

	if strings.TrimSpace(updates.Status) != "" {
		updatesParam = append(updatesParam, fmt.Sprintf("status = $%d", argPos))
		args = append(args, strings.TrimSpace(updates.Status))
		argPos++
	}

	qry := "UPDATE " + user_table_name + " SET " + strings.Join(updatesParam, ", ") +
		fmt.Sprintf(" WHERE id = $%d AND deleted_at IS NULL RETURNING *", argPos)
	args = append(args, id)

	row := repo.db.QueryRow(qry, args...)
	var user model.User
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Handle,
		&user.Mobile,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) DeleteByID(id uint64) error {
	_, err := repo.FindByID(id)
	if err != nil {
		return err
	}

	qry := "UPDATE " + user_table_name + " SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL"
	_, err = repo.db.Exec(qry, time.Now(), id)
	return err
}
