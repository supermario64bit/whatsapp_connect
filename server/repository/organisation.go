package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/supermario64bit/whatsapp_connect/db"
	"github.com/supermario64bit/whatsapp_connect/server/model"
)

const table_name string = "organisations"

type organisationRepository struct {
	db *sql.DB
}

type OrganisationRepository interface {
	Create(org *model.Organisation) (*model.Organisation, error)
	Find(filter *model.Organisation) ([]*model.Organisation, error)
	FindByID(id uint64) (*model.Organisation, error)
	UpdateByID(updates *model.Organisation, id uint64) (*model.Organisation, error)
	DeleteByID(id uint64) error
}

func NewOrganisationRepository() OrganisationRepository {
	return &organisationRepository{
		db: db.New(),
	}
}

func (repo *organisationRepository) Create(org *model.Organisation) (*model.Organisation, error) {
	if org == nil {
		return nil, fmt.Errorf("Cannot create organisation for nil reference")
	}

	if org.ID > 0 {
		return nil, fmt.Errorf("ID field should be empty")
	}

	if org.Name == "" || org.ContactNumber == "" || org.Email == "" || org.Status == "" {
		return nil, fmt.Errorf("Name, Contact Number. Email and Status field should not be empty")

	}

	if len(org.ContactNumber) != 10 {
		return nil, fmt.Errorf("Contact Number should be 10 digit")
	}

	colNames := []string{"name", "contact_number", "email", "status"}
	values := [][]interface{}{
		{org.Name, org.ContactNumber, org.Email, org.Status},
	}

	qry, args := generateInsertQuery(table_name, colNames, values)

	var createdOrg model.Organisation
	err := repo.db.QueryRow(qry, args...).Scan(&org.ID, &org.Name, &org.ContactNumber, &org.Email, &org.Status, &org.CreatedAt, &org.UpdatedAt, &org.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &createdOrg, nil
}

func (repo *organisationRepository) Find(filter *model.Organisation) ([]*model.Organisation, error) {
	args := []interface{}{}
	whereParts := []string{}
	if filter != nil {
		if strings.TrimSpace(filter.Name) != "" {
			args = append(args, "%"+strings.TrimSpace(filter.Name)+"%")
			whereParts = append(whereParts, fmt.Sprintf("name LIKE $%d", len(args)))
		}

		if strings.TrimSpace(filter.ContactNumber) != "" {
			args = append(args, "%"+strings.TrimSpace(filter.ContactNumber)+"%")
			whereParts = append(whereParts, fmt.Sprintf("contact_number LIKE $%d", len(args)))

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
	whereClause = whereClause + " AND deleted_at IS NULL"

	qry := fmt.Sprintf("SELECT * FROM %s ", table_name) + whereClause
	rows, err := repo.db.Query(qry, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orgs []*model.Organisation

	for rows.Next() {
		var org model.Organisation

		err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.ContactNumber,
			&org.Email,
			&org.Status,
			&org.CreatedAt,
			&org.UpdatedAt,
			&org.DeletedAt,
		)

		if err != nil {
			return nil, err
		}

		orgs = append(orgs, &org)
	}

	return orgs, nil
}

func (repo *organisationRepository) FindByID(id uint64) (*model.Organisation, error) {
	qry := "SELECT * FROM " + table_name + " WHERE id = $1 AND deleted_at IS NULL LIMIT 1"

	var org model.Organisation

	err := repo.db.QueryRow(qry, id).Scan(
		&org.ID,
		&org.Name,
		&org.ContactNumber,
		&org.Email,
		&org.Status,
		&org.CreatedAt,
		&org.UpdatedAt,
		&org.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &org, nil
}

func (repo *organisationRepository) UpdateByID(updates *model.Organisation, id uint64) (*model.Organisation, error) {
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

	if strings.TrimSpace(updates.ContactNumber) != "" {
		updatesParam = append(updatesParam, fmt.Sprintf("name = $%d", argPos))
		args = append(args, strings.TrimSpace(updates.Name))
		argPos++
	}

	if strings.TrimSpace(updates.Email) != "" {
		updatesParam = append(updatesParam, fmt.Sprintf("name = $%d", argPos))
		args = append(args, strings.TrimSpace(updates.Name))
		argPos++
	}

	if strings.TrimSpace(updates.Status) != "" {
		updatesParam = append(updatesParam, fmt.Sprintf("name = $%d", argPos))
		args = append(args, strings.TrimSpace(updates.Name))
		argPos++
	}

	qry := "UPDATE " + table_name + " SET " + strings.Join(updatesParam, ", ") +
		fmt.Sprintf(" WHERE id = $%d AND deleted_at IS NULL RETURNING *", argPos)
	args = append(args, id)

	row := repo.db.QueryRow(qry, args...)
	var org model.Organisation
	err = row.Scan(
		&org.ID,
		&org.Name,
		&org.ContactNumber,
		&org.Email,
		&org.Status,
		&org.CreatedAt,
		&org.UpdatedAt,
		&org.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &org, nil
}

func (repo *organisationRepository) DeleteByID(id uint64) error {
	_, err := repo.FindByID(id)
	if err != nil {
		return err
	}

	qry := "UPDATE " + table_name + " SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL"
	_, err = repo.db.Exec(qry, time.Now(), id)
	return err
}
