package contentservice

import (
	"github.com/fastschema/fastschema/app"
	"github.com/fastschema/fastschema/pkg/errors"
	"github.com/fastschema/fastschema/pkg/utils"
)

func (cs *ContentService) Delete(c app.Context, _ *any) (any, error) {
	model, err := cs.DB().Model(c.Arg("schema"))
	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	id := c.ArgInt("id")

	_, err = model.Query(app.EQ("id", id)).Only()

	if err != nil {
		e := utils.If(app.IsNotFound(err), errors.NotFound, errors.InternalServerError)
		return nil, e(err.Error())
	}

	if _, err := model.Mutation().Where(app.EQ("id", id)).Delete(); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	return nil, nil
}
