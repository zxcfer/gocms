package entdbadapter

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/fastschema/fastschema/app"
	"github.com/fastschema/fastschema/pkg/utils"
)

type PredicateFN func(*sql.Selector) *sql.Predicate

// createEntPredicates creates ent sql predicates from the given predicates
func createEntPredicates(
	entAdapter EntAdapter,
	model *Model,
	predicates []*app.Predicate,
) (func(*sql.Selector) []*sql.Predicate, error) {
	var predicateFns = []PredicateFN{}

	for _, p := range predicates {
		if len(p.RelationFieldNames) > 0 {
			lastFieldPredicate := p.Clone()
			lastFieldPredicate.RelationFieldNames = []string{}
			relationPredicateFn, err := createRelationsPredicate(
				entAdapter,
				model,
				lastFieldPredicate,
				p.RelationFieldNames...,
			)

			if err != nil {
				return nil, err
			}

			predicateFns = append(predicateFns, relationPredicateFn)
			continue
		}

		if p.Field != "" {
			predicateFn, err := CreateFieldPredicate(p)
			if err != nil {
				return nil, err
			}

			predicateFns = append(predicateFns, predicateFn)
			continue
		}

		if p.And != nil {
			andPredicatesFn, err := createEntPredicates(entAdapter, model, p.And)
			if err != nil {
				return nil, err
			}

			predicateFns = append(predicateFns, func(s *sql.Selector) *sql.Predicate {
				return sql.And(andPredicatesFn(s)...)
			})
			continue
		}

		if p.Or != nil {
			orPredicatesFn, err := createEntPredicates(entAdapter, model, p.Or)
			if err != nil {
				return nil, err
			}

			predicateFns = append(predicateFns, func(s *sql.Selector) *sql.Predicate {
				return sql.Or(orPredicatesFn(s)...)
			})
			continue
		}
	}

	return func(s *sql.Selector) []*sql.Predicate {
		var entPredicates = []*sql.Predicate{}
		for _, predicateFn := range predicateFns {
			entPredicates = append(entPredicates, predicateFn(s))
		}

		return entPredicates
	}, nil
}

// createRelationsPredicate creates the relation predicate
func createRelationsPredicate(
	entAdapter EntAdapter,
	model *Model,
	lastFieldPredicate *app.Predicate,
	relationFieldNames ...string,
) (PredicateFN, error) {
	relationFieldName := relationFieldNames[0]
	relationFieldNames = relationFieldNames[1:]
	relationField, err := model.schema.Field(relationFieldName)

	if err != nil {
		return nil, err
	}

	relation := relationField.Relation

	targetModel, err := model.client.Model(relation.TargetSchemaName)
	if err != nil {
		return nil, err
	}

	entTargetModel, ok := targetModel.(*Model)
	if !ok {
		return nil, fmt.Errorf("model %s is not an ent model", targetModel.Schema().Name)
	}

	stepOption, err := entAdapter.NewEdgeStepOption(relation)
	if err != nil {
		return nil, fmt.Errorf("invalid edge step option '%s': %w", relationFieldName, err)
	}

	relationStep := sqlgraph.NewStep(
		sqlgraph.From(model.schema.Namespace, model.entIDColumn.Name),
		sqlgraph.To(entTargetModel.schema.Namespace, entTargetModel.entIDColumn.Name),
		stepOption,
	)

	var pred func(*sql.Selector)
	if len(relationFieldNames) > 0 {
		p, err := createRelationsPredicate(
			entAdapter,
			entTargetModel,
			lastFieldPredicate,
			relationFieldNames...,
		)

		if err != nil {
			return nil, err
		}

		pred = func(s2 *sql.Selector) {
			s2.Where(p(s2))
		}
	} else {
		predFn, err := createEntPredicates(entAdapter, model, []*app.Predicate{lastFieldPredicate})
		if err != nil {
			return nil, err
		}

		pred = func(s2 *sql.Selector) {
			s2.Where(sql.And(predFn(s2)...))
		}
	}

	return func(selector *sql.Selector) *sql.Predicate {
		s1 := selector.Clone().SetP(nil)
		sqlgraph.HasNeighborsWith(s1, relationStep, pred)
		return s1.P()
	}, nil
}

// CreateFieldPredicate convert a predicate to ent predicate
func CreateFieldPredicate(predicate *app.Predicate) (PredicateFN, error) {
	var columnWrap = func(field string, selectors ...*sql.Selector) string {
		if len(selectors) > 0 {
			return selectors[0].C(field)
		}

		return field
	}

	switch predicate.Operator {
	case app.OpEQ:
		return func(s *sql.Selector) *sql.Predicate {
			return sql.EQ(columnWrap(predicate.Field), predicate.Value)
		}, nil
	case app.OpNEQ:
		return func(s *sql.Selector) *sql.Predicate {
			return sql.NEQ(columnWrap(predicate.Field), predicate.Value)
		}, nil
	case app.OpGT:
		return func(s *sql.Selector) *sql.Predicate {
			return sql.GT(columnWrap(predicate.Field), predicate.Value)
		}, nil
	case app.OpGTE:
		return func(s *sql.Selector) *sql.Predicate {
			return sql.GTE(columnWrap(predicate.Field), predicate.Value)
		}, nil
	case app.OpLT:
		return func(s *sql.Selector) *sql.Predicate {
			return sql.LT(columnWrap(predicate.Field), predicate.Value)
		}, nil
	case app.OpLTE:
		return func(s *sql.Selector) *sql.Predicate {
			return sql.LTE(columnWrap(predicate.Field), predicate.Value)
		}, nil
	case app.OpLIKE:
		stringValue, ok := predicate.Value.(string)
		if !ok {
			return nil, fmt.Errorf(
				"value of field %s.%s = %v (%T) must be string",
				predicate.Field,
				app.OpLIKE,
				predicate.Value,
				predicate.Value,
			)
		}

		return func(s *sql.Selector) *sql.Predicate {
			return sql.Like(columnWrap(predicate.Field), stringValue)
		}, nil
	case app.OpIN, app.OpNIN:
		arrayValue, ok := predicate.Value.([]any)
		if !ok {
			return nil, fmt.Errorf(
				"value of field %s.%s = %v (%T) must be an array",
				predicate.Field,
				predicate.Operator,
				predicate.Value,
				predicate.Value,
			)
		}

		return func(s *sql.Selector) *sql.Predicate {
			op := utils.If(predicate.Operator == app.OpIN, sql.In, sql.NotIn)
			return op(columnWrap(predicate.Field), arrayValue...)
		}, nil
	case app.OpNULL:
		return func(s *sql.Selector) *sql.Predicate {
			op := utils.If(predicate.Value == true, sql.IsNull, sql.NotNull)
			return op(columnWrap(predicate.Field))
		}, nil
	default:
		return nil, fmt.Errorf("operator %s not supported", predicate.Operator)
	}
}
