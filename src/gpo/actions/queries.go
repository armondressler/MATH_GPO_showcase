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

func GetGpoCompanies(gpo_id uuid.UUID) ([]*models.Company, error) {
	gpo := models.Gpo{}
	companies := []*models.Company{}
	err := models.DB.Find(&gpo, gpo_id)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("get companies of gpo with id %s : %s", gpo_id, err))
	}

	err = models.DB.RawQuery("select c.* from companies as c join companies_gpoes as cg on cg.company_id = c.id join gpoes as g on g.id = cg.gpo_id where g.id = ?;", gpo_id).All(&companies)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("get companies of gpo with id %s : %s", gpo_id, err))
	}
	return companies, err
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

func GetGPOContractors(gpo_id uuid.UUID) ([]*models.Contractor, error) {
	gpo := models.Gpo{}
	contractors := []*models.Contractor{}
	err := models.DB.Find(&gpo, gpo_id)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("get contractors of gpo with id %s : %s", gpo_id, err))
	}

	err = models.DB.RawQuery("select c.* from contractors as c join contractors_gpoes as cg on cg.contractor_id = c.id join gpoes as g on g.id = cg.gpo_id where g.id = ?;", gpo_id).All(&contractors)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("get contractors of gpo with id %s : %s", gpo_id, err))
	}
	return contractors, err
}

func GetAllGPOs() ([]*models.Gpo, error) {
	gpos := []*models.Gpo{}

	err := models.DB.RawQuery("select * from gpoes").All(&gpos)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("get gpoes: %s", err))
	}
	return gpos, err
}
