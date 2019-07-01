package service

import (
	"context"
	"database/sql"
	"github.com/gidyon/rupacinema/movie/pkg/api"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
)

type deleteMovieDS struct {
	res *empty.Empty
	err error
}

func (deleteMovie *deleteMovieDS) Delete(
	ctx context.Context,
	sqlWorkerChan chan<- sqlWorker,
	redisWorkerChan chan<- redisWorker,
	delReq *movie.DeleteMovieRequest,
	db *sql.DB,
	redisClient *redis.Client,
) {
	// Check if context is cancelled before proceeding
	if cancelled(ctx) {
		return
	}

	// Check that the movie Id is provided, error if it isn't
	movieID := delReq.GetMovieId()
	if movieID == "" {
		deleteMovie.err = errMissingCredential("Movie ID")
		return
	}

	// Prepare query
	query := "DELETE * FROM movies WHERE id=?"

	// Execute query
	_, err := db.ExecContext(ctx, query, movieID)
	if err != nil {
		go sendSQLErrToChan(sqlWorkerChan, &sqlWorker{
			query: query,
			args:  []interface{}{movieID},
			err:   err,
		})
		deleteMovie.err = errQueryFailed(err, "Delete Movie")
		return
	}

	deleteMovie.res = &empty.Empty{}

	rmMovieFromCacheAndHandleErr(redisWorkerChan, redisClient, movieID)
}
