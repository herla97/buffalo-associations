package actions

import (
	"coke/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

// BooksRead default implementation.
func BooksRead(c buffalo.Context) error {
	// Books Model
	books := &models.Books{}

	// Get the DB connection from the context.
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	// Add Order for date.
	q := tx.PaginateFromParams(c.Params()).Order("created_at asc")

	// Retrieve all Users from the DB. Select all except password.
	if err := q.Eager().All(books); err != nil {
		return errors.WithStack(err)
	}
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.Auto(c, books))
}
