package actions

import "github.com/markbates/buffalo"

type AdminPage struct {
	Title    string
	Subtitle string
	Page     string
}

// makeAdminPage returns an AdminPage struct struct which populates
// the admin template.  All pages with the admin templates should include
// this by calling c.Set("adminpage", makeAdminPage("Title", "Subtitle", "Page Name"))
func makeAdminPage(title, subtitle, page string) AdminPage {
	return AdminPage{
		Title:    title,
		Subtitle: subtitle,
		Page:     page,
	}
}

// AdminHandler serves the admin root
func AdminHandler(c buffalo.Context) error {
	c.Set("adminpage", makeAdminPage("Dashboard", "get stuff done!", "Index"))
	return c.Render(200, r.HTML("index.html", "main.html"))
}
