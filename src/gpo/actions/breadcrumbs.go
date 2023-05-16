package actions

import (
	"fmt"
	"gpo/models"

	"github.com/gobuffalo/buffalo"
)

func addBreadcrumbs(c buffalo.Context) error {
	breadcrumbs := [][]string{}
	breadcrumbs = append(breadcrumbs, []string{"Home", "/"})
	if company := c.Value("company"); company != nil {
		breadcrumbs = append(breadcrumbs, []string{"Unternehmen", "/companies"})
		co, ok := company.(*models.Company)
		if !ok {
			return c.Error(503, fmt.Errorf("creating breadcrumbs, no company found in context"))
		}
		breadcrumbs = append(breadcrumbs, []string{co.Name, fmt.Sprintf("/companies/%s", co.ID.String())})
		if gpo := c.Value("gpo"); gpo != nil {
			gpo, ok := gpo.(*models.Gpo)
			if !ok {
				return c.Error(503, fmt.Errorf("creating breadcrumbs, no gpo found in context"))
			}
			breadcrumbs = append(breadcrumbs, []string{gpo.Name, fmt.Sprintf("/companies/%s/gpoes/%s", co.ID.String(), gpo.ID.String())})
		}
	}

	c.Set("breadcrumbs", breadcrumbs)
	return nil
}
