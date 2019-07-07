package service

import (
	"context"
	"github.com/gidyon/rupacinema/movie/pkg/api"
	"github.com/go-redis/redis"
)

const pageSize = 20

type listMoviesDS struct {
	res *movie.ListMoviesResponse
	err error
}

func (listMovies *listMoviesDS) List(
	ctx context.Context,
	redisWorkerChan chan<- redisWorker,
	listReq *movie.ListMoviesRequest,
	redisClient *redis.Client,
) {
	// Check if context is cancelled before proceeding
	if cancelled(ctx) {
		return
	}

	pageToken := listReq.GetPageToken()
	if pageToken < 0 {
		listMovies.err = errIncorrectVal("Page token")
		return
	}

	// Movie ids stored in redis list
	movieIDs, err := getListIDs(
		redisWorkerChan,
		redisClient,
		int64(pageToken*pageSize),
		int64(pageToken*pageSize+pageSize),
	)
	if err != nil {
		listMovies.err = err
		return
	}

	// Prepare a pipeline
	pipeline := redisClient.Pipeline()

	// Channel to receive results
	strMapMapCmdChan := make(chan *redis.StringStringMapCmd, pageSize)

	for _, movieID := range movieIDs {
		strMapMapCmdChan <- pipeline.HGetAll(getKey(movieID))
	}

	close(strMapMapCmdChan)

	// Execeute the pipeline
	sliceCMD, err := pipeline.Exec()
	if err != nil {
		listMovies.err = errRedisCmdFailed(err, "List Movies")
		return
	}

	listMovies.res = &movie.ListMoviesResponse{
		NextPageToken: pageToken + 1,
		Movies:        make([]*movie.Movie, 0, len(sliceCMD)),
	}

	for strMapMapCmd := range strMapMapCmdChan {
		movieItem, err := getMovieFromHGETALL(strMapMapCmd)
		if err != nil {
			go sendRedisErrToChan(redisWorkerChan, strMapMapCmd, actionList)
			continue
		}

		listMovies.res.Movies = append(listMovies.res.Movies, movieItem)
	}
}
