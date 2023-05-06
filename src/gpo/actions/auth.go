package actions

import (
	"fmt"
	"gpo/models"
	"net/http"
	"os"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/pkg/errors"
)

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/google/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(http.StatusUnauthorized, err)
	}
	tx := c.Value("tx").(*pop.Connection)
	q := tx.Where("provider = ? and provider_id = ?", user.Provider, user.UserID)
	exists, err := q.Exists("users")
	if err != nil {
		return errors.WithStack(err)
	}
	u := &models.User{}
	if exists {
		q.First(u)
		if err != nil {
			return errors.WithStack(err)
		}
		c.Logger().Info("found existing user ", user.Email, " with id ", user.UserID)
	} else {
		c.Logger().Info("generating new user ", user.Email, " with id ", user.UserID)
		u.Email = user.Email
		u.ProviderID = user.UserID
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
		u.Provider = user.Provider
		err = tx.Save(u)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	c.Session().Set("current_user_id", u.ID)
	err = c.Session().Save()
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Redirect(302, "/")
}

func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	err := c.Session().Save()
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Redirect(302, "/")
}

func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				c.Session().Delete("current_user_id")
				return next(c)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

func SetCurrentCompany(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if user := c.Value("current_user"); user != nil {
			company := &models.Company{}
			cid := user.(*models.User).CompanyID
			err := models.DB.Find(company, cid)
			if err != nil {
				return fmt.Errorf("could not determine company of user %s", user.(*models.User).Email)
			}
			c.Set("current_company", company)
		}
		return next(c)
	}
}

func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			return c.Redirect(302, "/login/show")
		}
		return next(c)
	}
}
