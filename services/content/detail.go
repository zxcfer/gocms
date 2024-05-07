package contentservice

import (
	"strings"

	"github.com/fastschema/fastschema/app"
	"github.com/fastschema/fastschema/pkg/errors"
	"github.com/fastschema/fastschema/pkg/utils"
	"github.com/fastschema/fastschema/schema"
)

func (cs *ContentService) Detail(c app.Context, _ *any) (*schema.Entity, error) {
	id := c.ArgInt("id")
	schemaName := c.Arg("schema")
	model, err := cs.DB().Model(schemaName)
	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	columns := []string{}
	if fields := c.Arg("select", ""); fields != "" {
		columns = strings.Split(fields, ",")
	}

	entity, err := model.Query(app.EQ("id", id)).
		Select(columns...).
		First(c.Context())
	if err != nil {
		e := utils.If(app.IsNotFound(err), errors.NotFound, errors.InternalServerError)
		return nil, e(err.Error())
	}

	if schemaName == "user" {
		entity.Delete("password")
	}

	return entity, nil
}
