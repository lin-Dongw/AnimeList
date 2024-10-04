// Code generated by ent, DO NOT EDIT.

package season

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dongwlin/anime-list/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldName, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v int64) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldValue, v))
}

// Cover applies equality check predicate on the "cover" field. It's identical to CoverEQ.
func Cover(v string) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldCover, v))
}

// ReleasedAt applies equality check predicate on the "released_at" field. It's identical to ReleasedAtEQ.
func ReleasedAt(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldReleasedAt, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldDescription, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Season {
	return predicate.Season(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Season {
	return predicate.Season(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Season {
	return predicate.Season(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Season {
	return predicate.Season(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Season {
	return predicate.Season(sql.FieldContainsFold(FieldName, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v int64) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v int64) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...int64) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...int64) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v int64) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v int64) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v int64) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v int64) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldValue, v))
}

// CoverEQ applies the EQ predicate on the "cover" field.
func CoverEQ(v string) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldCover, v))
}

// CoverNEQ applies the NEQ predicate on the "cover" field.
func CoverNEQ(v string) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldCover, v))
}

// CoverIn applies the In predicate on the "cover" field.
func CoverIn(vs ...string) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldCover, vs...))
}

// CoverNotIn applies the NotIn predicate on the "cover" field.
func CoverNotIn(vs ...string) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldCover, vs...))
}

// CoverGT applies the GT predicate on the "cover" field.
func CoverGT(v string) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldCover, v))
}

// CoverGTE applies the GTE predicate on the "cover" field.
func CoverGTE(v string) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldCover, v))
}

// CoverLT applies the LT predicate on the "cover" field.
func CoverLT(v string) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldCover, v))
}

// CoverLTE applies the LTE predicate on the "cover" field.
func CoverLTE(v string) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldCover, v))
}

// CoverContains applies the Contains predicate on the "cover" field.
func CoverContains(v string) predicate.Season {
	return predicate.Season(sql.FieldContains(FieldCover, v))
}

// CoverHasPrefix applies the HasPrefix predicate on the "cover" field.
func CoverHasPrefix(v string) predicate.Season {
	return predicate.Season(sql.FieldHasPrefix(FieldCover, v))
}

// CoverHasSuffix applies the HasSuffix predicate on the "cover" field.
func CoverHasSuffix(v string) predicate.Season {
	return predicate.Season(sql.FieldHasSuffix(FieldCover, v))
}

// CoverEqualFold applies the EqualFold predicate on the "cover" field.
func CoverEqualFold(v string) predicate.Season {
	return predicate.Season(sql.FieldEqualFold(FieldCover, v))
}

// CoverContainsFold applies the ContainsFold predicate on the "cover" field.
func CoverContainsFold(v string) predicate.Season {
	return predicate.Season(sql.FieldContainsFold(FieldCover, v))
}

// ReleasedAtEQ applies the EQ predicate on the "released_at" field.
func ReleasedAtEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldReleasedAt, v))
}

// ReleasedAtNEQ applies the NEQ predicate on the "released_at" field.
func ReleasedAtNEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldReleasedAt, v))
}

// ReleasedAtIn applies the In predicate on the "released_at" field.
func ReleasedAtIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldReleasedAt, vs...))
}

// ReleasedAtNotIn applies the NotIn predicate on the "released_at" field.
func ReleasedAtNotIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldReleasedAt, vs...))
}

// ReleasedAtGT applies the GT predicate on the "released_at" field.
func ReleasedAtGT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldReleasedAt, v))
}

// ReleasedAtGTE applies the GTE predicate on the "released_at" field.
func ReleasedAtGTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldReleasedAt, v))
}

// ReleasedAtLT applies the LT predicate on the "released_at" field.
func ReleasedAtLT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldReleasedAt, v))
}

// ReleasedAtLTE applies the LTE predicate on the "released_at" field.
func ReleasedAtLTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldReleasedAt, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Season {
	return predicate.Season(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Season {
	return predicate.Season(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Season {
	return predicate.Season(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Season {
	return predicate.Season(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Season {
	return predicate.Season(sql.FieldContainsFold(FieldDescription, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Season {
	return predicate.Season(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Season {
	return predicate.Season(sql.FieldLTE(FieldCreatedAt, v))
}

// HasEpisodes applies the HasEdge predicate on the "episodes" edge.
func HasEpisodes() predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EpisodesTable, EpisodesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEpisodesWith applies the HasEdge predicate on the "episodes" edge with a given conditions (other predicates).
func HasEpisodesWith(preds ...predicate.Episode) predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		step := newEpisodesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnime applies the HasEdge predicate on the "anime" edge.
func HasAnime() predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AnimeTable, AnimeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnimeWith applies the HasEdge predicate on the "anime" edge with a given conditions (other predicates).
func HasAnimeWith(preds ...predicate.Anime) predicate.Season {
	return predicate.Season(func(s *sql.Selector) {
		step := newAnimeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Season) predicate.Season {
	return predicate.Season(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Season) predicate.Season {
	return predicate.Season(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Season) predicate.Season {
	return predicate.Season(sql.NotPredicates(p))
}
