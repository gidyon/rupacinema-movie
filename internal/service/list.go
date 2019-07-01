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

	// Get the movie obj
	pipeline := redisClient.Pipeline()

	for _, movieID := range movieIDs {
		pipeline.HGetAll(getKey(movieID))
	}

	fn := func(pipeline redis.Pipeliner) error {
		return nil
	}

	sliceCMD, err := redisClient.Pipelined(fn)
	if err != nil {
		listMovies.err = errRedisCmdFailed(err, "List Movies")
		return
	}

	listMovies.res = &movie.ListMoviesResponse{
		NextPageToken: pageToken + 1,
		Movies:        make([]*movie.Movie, 0, len(sliceCMD)),
	}

	for _, cmd := range sliceCMD {
		strMapMapCmd, ok := cmd.(*redis.StringStringMapCmd)
		if !ok {
			listMovies.err = errFailedTypeConversion("redis.Cmder", "redis.StringStringMapCmd")
			return
		}
		movieItem, err := getMovieFromHGETALL(strMapMapCmd)
		if err != nil {
			go sendRedisErrToChan(redisWorkerChan, strMapMapCmd, actionList)
			continue
		}

		listMovies.res.Movies = append(listMovies.res.Movies, movieItem)
	}
}
