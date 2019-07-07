package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/gidyon/rupacinema/movie/pkg/api"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
	"strings"
)

type updateMovieDS struct {
	res *empty.Empty
	err error
}

func (updateMovie *updateMovieDS) Update(
	ctx context.Context,
	sqlWorkerChan chan<- sqlWorker,
	redisWorkerChan chan<- redisWorker,
	updateReq *movie.UpdateMovieRequest,
	db *sql.DB,
	redisClient *redis.Client,
) {
	// Check if context is cancelled before proceeding
	if cancelled(ctx) {
		return
	}

	movieItem := updateReq.GetMovie()

	err := func() error {
		var err error
		// Validate movie input
		switch {
		case strings.Trim(movieItem.Id, " ") == "":
			err = errMissingCredential("Movie.Id")
		case strings.Trim(movieItem.Title, " ") == "":
			err = errMissingCredential("Movie Title")
		case strings.Trim(movieItem.Description, " ") == "":
			err = errMissingCredential("Movie Description")
		case strings.Trim(movieItem.ReleaseDate, " ") == "":
			err = errMissingCredential("Movie ReleaseDate")
		case movieItem.MovieDurationMins == 0:
			err = errMissingCredential("Movie MovieDurationMins")
		case len(movieItem.Photos) == 0:
			err = errMissingCredential("Movie Photos")
		case len(movieItem.Category) == 0:
			err = errMissingCredential("Movie Category")
		}

		return err
	}()

	if err != nil {
		updateMovie.err = err
		return
	}

	err = updateMovieToDB(ctx, sqlWorkerChan, db, movieItem)
	if err != nil {
		updateMovie.err = err
		return
	}

	updateMovie.res = &empty.Empty{}

	setMovieInCacheAndHandleErr(
		redisWorkerChan,
		redisClient,
		movieItem,
	)
}

func updateMovieToDB(
	ctx context.Context, sqlWorkerChan chan<- sqlWorker, db *sql.DB, movieItem *movie.Movie,
) error {
	// Check if context is cancelled
	if cancelled(ctx) {
		return ctx.Err()
	}

	category, err := json.Marshal(movieItem.Category)
	if err != nil {
		return errFromJSONMarshal(err, "Movie.Category")
	}

	photos, err := json.Marshal(movieItem.Photos)
	if err != nil {
		return errFromJSONMarshal(err, "Movie.Photos")
	}

	// `id` varchar(50) NOT NULL,
	// `title` varchar(50) NOT NULL,
	// `price` varchar(15) NOT NULL DEFAULT 'NA',
	// `description` text NOT NULL,
	// `trailer_url` text,
	// `audience_label` enum('NR','GE','PG','PG-13','NC-17') NOT NULL DEFAULT 'NR',
	// `ratings` float NOT NULL DEFAULT '0',
	// `duration` int(11) DEFAULT 0,
	// `all_votes` int(11) DEFAULT 0,
	// `release_date` date DEFAULT NULL,
	// `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	// `category` json DEFAULT NULL,
	// `photos` json DEFAULT NULL,

	// Prepare query
	query := `UPDATE movies SET title=?, price=?, description=?, trailer_url=?, audience_label=?, ratings=?, duration=?, release_date=?, category=?, photos=? WHERE id=?`

	// Execute query
	_, err = db.ExecContext(ctx, query,
		movieItem.Title,
		movieItem.Price,
		movieItem.Description,
		movieItem.TrailerUrl,
		movie.Audience_name[int32(movieItem.AudienceLabel)],
		movieItem.Ratings,
		movieItem.MovieDurationMins,
		movieItem.ReleaseDate,
		category,
		photos,
		movieItem.Id,
	)

	if err != nil {
		// Send the failed sql command to worker
		go sendSQLErrToChan(sqlWorkerChan, &sqlWorker{
			query: query,
			args: []interface{}{
				movieItem.Title,
				movieItem.Price,
				movieItem.Description,
				movieItem.TrailerUrl,
				movie.Audience_name[int32(movieItem.AudienceLabel)],
				movieItem.Ratings,
				movieItem.MovieDurationMins,
				movieItem.ReleaseDate,
				category,
				photos,
				movieItem.Id,
			},
			err: err,
		})
		return errQueryFailed(err, "UpdateMovie (UPDATE)")
	}

	return nil
}
