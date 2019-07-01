package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/gidyon/rupacinema/movie/pkg/api"
	"github.com/go-redis/redis"
	"strings"
)

type getMovieDS struct {
	res *movie.Movie
	err error
}

func (getMovie *getMovieDS) Get(
	ctx context.Context,
	sqlWorkerChan chan<- sqlWorker,
	redisWorkerChan chan<- redisWorker,
	getReq *movie.GetMovieRequest,
	db *sql.DB,
	redisClient *redis.Client,
) {
	// Check if context is cancelled before proceeding
	if cancelled(ctx) {
		return
	}

	movieID := getReq.GetMovieId()
	if strings.Trim(movieID, " ") == "" {
		getMovie.err = errMissingCredential("Movie Id")
		return
	}

	// Get movie from redis cache
	movieItem, err := getMovieFromCache(redisClient, movieID)
	if err != nil {
		var err0 error
		// Get movie from mysql database
		movieItem, err0 = getMovieFromDB(ctx, sqlWorkerChan, db, movieID)
		if err0 != nil {
			getMovie.err = err0
			return
		}

		getMovie.res = movieItem

		setMovieInCacheOneErr(redisWorkerChan, redisClient, movieItem, err, "GetMovie")

		return
	}

	getMovie.res = movieItem
}

func getMovieFromDB(
	ctx context.Context, sqlWorkerChan chan<- sqlWorker, db *sql.DB, movieID string,
) (*movie.Movie, error) {
	// Check if context is cancelled before proceeding
	if cancelled(ctx) {
		return nil, ctx.Err()
	}

	movieItem := &movie.Movie{}

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

	audienceLabel := ""
	createTime := ""
	category := make([]byte, 0)
	photos := make([]byte, 0)

	// Prepare query
	query := `SELECT * FROM movies WHERE id=?`

	// Execute query
	row := db.QueryRowContext(ctx, query, movieID)
	err := row.Scan(
		&movieItem.Id,
		&movieItem.Title,
		&movieItem.Price,
		&movieItem.Description,
		&movieItem.TrailerUrl,
		&audienceLabel,
		&movieItem.Ratings,
		&movieItem.MovieDurationMins,
		&movieItem.AllVotes,
		&movieItem.ReleaseDate,
		&createTime,
		&category,
		&photos,
	)

	if err != nil {
		go sendSQLErrToChan(sqlWorkerChan, &sqlWorker{
			query: query,
			args:  []interface{}{movieID},
			err:   err,
		})
		return nil, errQueryFailed(err, "GetMovie (SELECT)")
	}

	err = json.Unmarshal(category, movieItem.Category)
	if err != nil {
		return nil, errFromJSONUnMarshal(err, "Movie.Category")
	}

	err = json.Unmarshal(photos, movieItem.Photos)
	if err != nil {
		return nil, errFromJSONUnMarshal(err, "Movie.Photos")
	}

	movieItem.AudienceLabel = movie.Audience(movie.Audience_value[audienceLabel])

	return movieItem, nil
}
