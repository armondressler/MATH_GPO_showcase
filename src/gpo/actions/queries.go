package actions

import (
	"fmt"
	"gpo/models"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func GetEmployees(company_id uuid.UUID) ([]*models.User, error) {
	company := models.Company{}
	employees := []*models.User{}
	err := models.DB.Find(&company, company_id)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("get employees of company with id %s : %s", company_id, err))
	}

	err = models.DB.RawQuery("SELECT * FROM users WHERE company_id = ?", company_id).All(&employees)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("get employees of company with id %s : %s", company_id, err))
	}
	return employees, err
}

func GetGPOs(company_id uuid.UUID) ([]*models.Gpo, error) {
	company := models.Company{}
	gpos := []*models.Gpo{}
	err := models.DB.Find(&company, company_id)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("get gpos of company with id %s : %s", company_id, err))
	}

	err = models.DB.RawQuery("select g.* from gpoes as g join companies_gpoes as cg on cg.gpo_id = g.id join companies as c on c.id = cg.company_id where c.id = ?;", company_id).All(&gpos)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("get gpos of company with id %s : %s", company_id, err))
	}
	return gpos, err
}

func GetAllGPOs() ([]*models.Gpo, error) {
	gpos := []*models.Gpo{}

	err := models.DB.RawQuery("select * from gpoes").All(&gpos)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("get gpoes: %s", err))
	}
	return gpos, err
}
