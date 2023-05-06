package actions

import (
	"gpo/models"

	"github.com/gobuffalo/nulls"
	"github.com/pkg/errors"
)

func generate_data() error {
	tx := models.DB

	gpoes := []*models.Gpo{}
	gpo_octa := &models.Gpo{
		Name:        "Einkaufsgemeinschaft Octamed AG",
		Description: nulls.NewString("Einkaufskooperation für Gesundheitsdienstleister im Raum Zentral- und Ostschweiz"),
	}

	gpo_srk := &models.Gpo{
		Name:        "SRK EK GmbH",
		Description: nulls.NewString("Verbund für den gemeinschaftlichen Einkauf des Schweizerischen roten Kreuzes"),
	}

	gpoes = append(gpoes, gpo_octa, gpo_srk)
	for _, gpo := range gpoes {
		q := tx.Where("name = ?", gpo.Name)
		exists, err := q.Exists("gpoes")
		if err != nil {
			return errors.WithStack(err)
		}
		if exists {
			continue
		}
		if err != nil {
			return errors.WithStack(err)
		}
		err = tx.Save(gpo)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	companies := []*models.Company{}
	co_lu := &models.Company{
		Name:        "Kantonsspital Luzern",
		Description: nulls.NewString("Account für Projektarbeit"),
		Street:      nulls.NewString("Feldstrasse"),
		Streetnr:    nulls.NewString("3"),
		City:        nulls.NewString("Luzern"),
		Postalnr:    nulls.NewString("6004"),
		Gpoes:       []models.Gpo{*gpo_octa}}

	co_triemli := &models.Company{
		Name:        "Triemlispital",
		Description: nulls.NewString("Zentraler Gesundheitsdienstleister"),
		Street:      nulls.NewString("Zürcherstrasse"),
		Streetnr:    nulls.NewString("1"),
		City:        nulls.NewString("Zürich"),
		Postalnr:    nulls.NewString("8003"),
		Gpoes:       []models.Gpo{*gpo_octa}}

	co_biel := &models.Company{
		Name:     "Spital Biel/Bienne",
		Street:   nulls.NewString("Bielerweg"),
		Streetnr: nulls.NewString("22"),
		City:     nulls.NewString("Biel/Bienne"),
		Postalnr: nulls.NewString("2501"),
		Gpoes:    []models.Gpo{*gpo_octa, *gpo_srk}}

	co_zug := &models.Company{
		Name:        "Kantonsspital Zug",
		Description: nulls.NewString("We can fix it, if you can afford it."),
		Street:      nulls.NewString("Baarerstrasse"),
		Streetnr:    nulls.NewString("6"),
		City:        nulls.NewString("Baar"),
		Postalnr:    nulls.NewString("6320"),
		Gpoes:       []models.Gpo{*gpo_srk}}

	co_unizh := &models.Company{
		Name:        "Unispital Zürich",
		Description: nulls.NewString("Unser Motto: Trial and Error"),
		Street:      nulls.NewString("Seeweg"),
		Streetnr:    nulls.NewString("1"),
		City:        nulls.NewString("Zürich"),
		Postalnr:    nulls.NewString("8003"),
		Gpoes:       []models.Gpo{}}

	companies = append(companies, co_lu, co_triemli, co_biel, co_zug, co_unizh)
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
		FirstName:  "Armon",
		LastName:   "Dressler",
		ProviderID: "114725001636885388051",
		Provider:   "google",
		CompanyID:  co_lu.ID,
	}
	u_rolf := &models.User{
		Email:      "rolf.feldmann@proton.ch",
		FirstName:  "Rolf",
		LastName:   "Feldmann",
		ProviderID: "114725001636885388052",
		Provider:   "google",
		CompanyID:  co_lu.ID,
	}
	u_martin := &models.User{
		Email:      "martin.mueller@proton.ch",
		FirstName:  "Martin",
		LastName:   "Müller",
		ProviderID: "114725001636885388053",
		Provider:   "google",
		CompanyID:  co_biel.ID,
	}
	u_felix := &models.User{
		Email:      "felix.rotmann@gmail.com",
		FirstName:  "Felix",
		LastName:   "Rotmann",
		ProviderID: "114725001636885388054",
		Provider:   "google",
		CompanyID:  co_biel.ID,
	}
	u_robert := &models.User{
		Email:      "robert.heller@proton.ch",
		FirstName:  "Robert",
		LastName:   "Heller",
		ProviderID: "114725001636885388055",
		Provider:   "google",
		CompanyID:  co_lu.ID,
	}
	u_lorenz := &models.User{
		Email:      "lorenz.baum@gmx.ch",
		FirstName:  "Lorenz",
		LastName:   "Baum",
		ProviderID: "114725001636885388056",
		Provider:   "google",
		CompanyID:  co_triemli.ID,
	}
	u_bernhard := &models.User{
		Email:      "b.felder@gmx.net",
		FirstName:  "Bernhard",
		LastName:   "Felder",
		ProviderID: "114725001636885388056",
		Provider:   "google",
		CompanyID:  co_triemli.ID,
	}
	u_ismail := &models.User{
		Email:      "ismail.erden@gmail.com",
		FirstName:  "Ismail",
		LastName:   "Erden",
		ProviderID: "114725001636885388056",
		Provider:   "google",
		CompanyID:  co_triemli.ID,
	}
	users = append(users, u_armon, u_felix, u_lorenz, u_martin, u_robert, u_rolf, u_bernhard, u_ismail)
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