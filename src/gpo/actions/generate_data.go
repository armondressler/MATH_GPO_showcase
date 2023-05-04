package actions

import (
	"gpo/models"

	"github.com/gobuffalo/nulls"
	"github.com/pkg/errors"
)

func generate_data() error {
	tx := models.DB
	companies := []*models.Company{}
	co_lu := &models.Company{
		Name:     "Kantonsspital Luzern",
		Street:   nulls.NewString("Feldstrasse"),
		Streetnr: nulls.NewString("3"),
		City:     nulls.NewString("Luzern"),
		Postalnr: nulls.NewString("6004")}

	co_triemli := &models.Company{
		Name:     "Triemlispital",
		Street:   nulls.NewString("Zürcherstrasse"),
		Streetnr: nulls.NewString("1"),
		City:     nulls.NewString("Zürich"),
		Postalnr: nulls.NewString("8003")}

	co_biel := &models.Company{
		Name:     "Spital Biel/Bienne",
		Street:   nulls.NewString("Bielerweg"),
		Streetnr: nulls.NewString("22"),
		City:     nulls.NewString("Biel/Bienne"),
		Postalnr: nulls.NewString("2501")}
	companies = append(companies, co_lu, co_triemli, co_biel)
	for _, co := range companies {
		q := tx.Where("name = ? and street = ? and streetnr = ? and city = ?", co.Name, co.Street, co.Streetnr, co.City)
		exists, err := q.Exists("companies")
		if err != nil {
			return errors.WithStack(err)
		}
		if exists {
			continue
		}
		if err != nil {
			return errors.WithStack(err)
		}
		err = tx.Save(co)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	users := []*models.User{}
	u_armon := &models.User{
		Email:      "armon.dressler@gmail.com",
		ProviderID: "114725001636885388051",
		Provider:   "google",
		CompanyID:  co_lu.ID,
	}
	u_rolf := &models.User{
		Email:      "rolf.feldmann@proton.ch",
		ProviderID: "114725001636885388052",
		Provider:   "google",
		CompanyID:  co_lu.ID,
	}
	u_martin := &models.User{
		Email:      "martin.mueller@proton.ch",
		ProviderID: "114725001636885388053",
		Provider:   "google",
		CompanyID:  co_biel.ID,
	}
	u_felix := &models.User{
		Email:      "felix.rotmann@gmail.com",
		ProviderID: "114725001636885388054",
		Provider:   "google",
		CompanyID:  co_biel.ID,
	}
	u_robert := &models.User{
		Email:      "robert.heller@proton.ch",
		ProviderID: "114725001636885388055",
		Provider:   "google",
		CompanyID:  co_lu.ID,
	}
	u_lorenz := &models.User{
		Email:      "lorenz.baum@gmx.ch",
		ProviderID: "114725001636885388056",
		Provider:   "google",
		CompanyID:  co_triemli.ID,
	}
	users = append(users, u_armon, u_felix, u_lorenz, u_martin, u_robert, u_rolf)
	for _, user := range users {
		q := tx.Where("email = ?", user.Email)
		exists, err := q.Exists("users")
		if err != nil {
			return errors.WithStack(err)
		}
		if exists {
			continue
		}

		err = tx.Save(user)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
