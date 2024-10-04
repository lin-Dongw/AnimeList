// Code generated by ent, DO NOT EDIT.

package anime

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the anime type in the database.
	Label = "anime"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeSeasons holds the string denoting the seasons edge name in mutations.
	EdgeSeasons = "seasons"
	// EdgeTheaters holds the string denoting the theaters edge name in mutations.
	EdgeTheaters = "theaters"
	// Table holds the table name of the anime in the database.
	Table = "animes"
	// SeasonsTable is the table that holds the seasons relation/edge.
	SeasonsTable = "seasons"
	// SeasonsInverseTable is the table name for the Season entity.
	// It exists in this package in order to avoid circular dependency with the "season" package.
	SeasonsInverseTable = "seasons"
	// SeasonsColumn is the table column denoting the seasons relation/edge.
	SeasonsColumn = "anime_seasons"
	// TheatersTable is the table that holds the theaters relation/edge.
	TheatersTable = "theaters"
	// TheatersInverseTable is the table name for the Theater entity.
	// It exists in this package in order to avoid circular dependency with the "theater" package.
	TheatersInverseTable = "theaters"
	// TheatersColumn is the table column denoting the theaters relation/edge.
	TheatersColumn = "anime_theaters"
)

// Columns holds all SQL columns for anime fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldStatus,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Anime queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// BySeasonsCount orders the results by seasons count.
func BySeasonsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSeasonsStep(), opts...)
	}
}

// BySeasons orders the results by seasons terms.
func BySeasons(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSeasonsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTheatersCount orders the results by theaters count.
func ByTheatersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTheatersStep(), opts...)
	}
}

// ByTheaters orders the results by theaters terms.
func ByTheaters(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTheatersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSeasonsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SeasonsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SeasonsTable, SeasonsColumn),
	)
}
func newTheatersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TheatersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TheatersTable, TheatersColumn),
	)
}
