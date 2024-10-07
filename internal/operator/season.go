package operator

import (
	"context"
	"github.com/dongwlin/anime-list/internal/ent/episode"
	"time"

	"github.com/dongwlin/anime-list/internal/ent"
	"github.com/dongwlin/anime-list/internal/ent/anime"
	"github.com/dongwlin/anime-list/internal/ent/season"
)

type CreateSeasonParams struct {
	AnimeID     int
	Name        string
	Value       int64
	Cover       string
	ReleasedAt  time.Time
	Description string
	Status      int
}

type ListSeasonParams struct {
	AnimeID int
	Offset  int
	Limit   int
}

type SearchSeasonParams struct {
	Query   string
	AnimeID int
	Offset  int
	Limit   int
}

type UpdateSeasonParams struct {
	ID          int
	Name        string
	Value       int64
	Cover       string
	ReleasedAt  time.Time
	Description string
	Status      int
}

type SeasonOperator interface {
	Create(ctx context.Context, params CreateSeasonParams) (*ent.Season, error)
	List(ctx context.Context, params ListSeasonParams) (int, []*ent.Season, error)
	Search(ctx context.Context, params SearchSeasonParams) (int, []*ent.Season, error)
	Get(ctx context.Context, id int) (*ent.Season, error)
	Update(ctx context.Context, params UpdateSeasonParams) (*ent.Season, error)
	Delete(ctx context.Context, id int) error
}

type seasonOperator struct {
	db *ent.Client
}

// Create implements SeasonOperator.
func (so *seasonOperator) Create(ctx context.Context, params CreateSeasonParams) (*ent.Season, error) {
	s, err := so.db.Season.
		Create().
		SetName(params.Name).
		SetValue(params.Value).
		SetCover(params.Cover).
		SetReleasedAt(params.ReleasedAt).
		SetDescription(params.Description).
		SetStatus(params.Status).
		SetAnimeID(params.AnimeID).
		Save(ctx)
	return s, err
}

// List implements SeasonOperator.
func (so *seasonOperator) List(ctx context.Context, params ListSeasonParams) (int, []*ent.Season, error) {
	query := so.db.Season.Query().
		Where(season.HasAnimeWith(anime.ID(params.AnimeID)))

	count, err := query.Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	seasons, err := query.
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}

	return count, seasons, nil
}

// Search implements SeasonOperator.
func (so *seasonOperator) Search(ctx context.Context, params SearchSeasonParams) (int, []*ent.Season, error) {
	query := so.db.Season.Query().
		Where(season.HasAnimeWith(anime.ID(params.AnimeID))).
		Where(season.NameContains(params.Query))

	count, err := query.Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	seasons, err := query.
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}

	return count, seasons, nil

}

// Get implements SeasonOperator.
func (so *seasonOperator) Get(ctx context.Context, id int) (*ent.Season, error) {
	s, err := so.db.Season.
		Get(ctx, id)
	return s, err
}

// Update implements SeasonOperator.
func (so *seasonOperator) Update(ctx context.Context, params UpdateSeasonParams) (*ent.Season, error) {
	s, err := so.db.Season.
		UpdateOneID(params.ID).
		SetName(params.Name).
		SetValue(params.Value).
		SetCover(params.Cover).
		SetReleasedAt(params.ReleasedAt).
		SetDescription(params.Description).
		SetStatus(params.Status).
		Save(ctx)
	return s, err
}

// Delete implements SeasonOperator.
func (so *seasonOperator) Delete(ctx context.Context, id int) error {
	tx, err := so.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.Episode.
		Delete().
		Where(episode.HasSeasonWith(season.ID(id))).
		Exec(ctx)
	if err != nil {
		return err
	}

	err = tx.Season.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func NewSeasonOperator(db *ent.Client) SeasonOperator {
	return &seasonOperator{
		db: db,
	}
}
