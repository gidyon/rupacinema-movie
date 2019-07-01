package service

import (
	"context"
	"database/sql"
	"time"
	// "encoding/json"
	"github.com/gidyon/rupacinema/movie/pkg/api"
	"github.com/gidyon/rupacinema/movie/pkg/logger"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

const (
	actionSet    = "set"
	actionVote   = "vote"
	actionDelete = "delete"
	actionList   = "list"
	moviesList   = "movies:list"
)

func getKey(movieID string) string {
	return "movies:" + movieID
}

func setMovieInCacheAndHandleErr(
	redisWorkerChan chan<- redisWorker,
	redisClient *redis.Client,
	movieItem *movie.Movie,
) {
	statusCMD := setMovieInCache(redisClient, movieItem)
	if statusCMD.Err() != nil {
		// Send the command that failed to redisWorker worker
		go sendRedisErrToChan(redisWorkerChan, statusCMD, actionSet)
	}
}

func setMovieInCache(
	redisClient *redis.Client,
	movieItem *movie.Movie,
) *redis.StatusCmd {

	// string id = 1;
	// string title = 2;
	// string price = 3;
	// string description = 4;
	// string trailer_url = 5;
	// string release_date = 6;
	// float ratings = 7;
	// int64 movie_duration_mins = 8;
	// repeated string photos = 9;
	// repeated string category = 10;
	// Audience audience_label = 11;
	// int32 all_votes = 12;
	// int32 current_votes = 13;

	movieMap := map[string]interface{}{
		"id":             movieItem.Id,
		"title":          movieItem.Title,
		"price":          movieItem.Price,
		"description":    movieItem.Description,
		"trailer_url":    movieItem.TrailerUrl,
		"release_date":   movieItem.ReleaseDate,
		"ratings":        movieItem.Ratings,
		"duration":       movieItem.MovieDurationMins,
		"photos":         strings.Join(movieItem.Photos, ";"),
		"category":       strings.Join(movieItem.Category, ";"),
		"audience_label": movieItem.AudienceLabel,
		"all_votes":      movieItem.AllVotes,
	}

	return redisClient.HMSet(getKey(movieItem.Id), movieMap)
}

func getMovieFromCache(
	redisClient *redis.Client,
	movieID string,
) (*movie.Movie, error) {

	strMapMapCmd := redisClient.HGetAll(getKey(movieID))

	movieItem, err := getMovieFromHGETALL(strMapMapCmd)
	if err != nil {
		return nil, err
	}

	return movieItem, nil
}

func getMovieFromHGETALL(strMapMapCmd *redis.StringStringMapCmd) (*movie.Movie, error) {
	if strMapMapCmd.Err() != nil {
		return nil, strMapMapCmd.Err()
	}

	movieMap := strMapMapCmd.Val()
	// movieMap := map[string]interface{}{
	// 	"id":             movieItem.Id,
	// 	"title":          movieItem.Title,
	// 	"price":          movieItem.Price,
	// 	"description":    movieItem.Description,
	// 	"trailer_url":    movieItem.TrailerUrl,
	// 	"release_date":   movieItem.ReleaseDate,
	// 	"ratings":        movieItem.Ratings,
	// 	"duration":       movieItem.MovieDurationMins,
	// 	"photos":         movieItem.Photos,
	// 	"category":       movieItem.Category,
	// 	"audience_label": movieItem.AudienceLabel,
	// 	"all_votes":      movieItem.AllVotes,
	// }

	ratings, err := strconv.ParseFloat(movieMap["ratings"], 32)
	if err != nil {
		return nil, errConvertingType(err, "String", "Float64")
	}

	duration, err := strconv.ParseInt(movieMap["duration"], 10, 64)
	if err != nil {
		return nil, errConvertingType(err, "String", "Int64")
	}

	audienceLabel, err := strconv.ParseInt(movieMap["audience_label"], 10, 32)
	if err != nil {
		return nil, errConvertingType(err, "String", "Int32")
	}

	allVotes, err := strconv.ParseInt(movieMap["all_votes"], 10, 32)
	if err != nil {
		return nil, errConvertingType(err, "String", "Int32")
	}

	movieItem := &movie.Movie{
		Id:                movieMap["id"],
		Title:             movieMap["title"],
		Price:             movieMap["price"],
		Description:       movieMap["description"],
		TrailerUrl:        movieMap["trailer_url"],
		ReleaseDate:       movieMap["release_date"],
		MovieDurationMins: int64(duration),
		Ratings:           float32(ratings),
		Photos:            strings.Split(movieMap["photos"], ";"),
		Category:          strings.Split(movieMap["category"], ";"),
		AudienceLabel:     movie.Audience(audienceLabel),
		AllVotes:          int32(allVotes),
	}

	return movieItem, nil
}

func rmMovieFromCacheAndHandleErr(
	redisWorkerChan chan<- redisWorker,
	redisClient *redis.Client,
	movieID string,
) {
	intCMD := rmMovieFromCache(redisClient, movieID)
	if err := intCMD.Err(); err != nil {
		// Send the command that failed to redisWorker worker
		go sendRedisErrToChan(redisWorkerChan, intCMD, actionDelete)
	}
}

func rmMovieFromCache(
	redisClient *redis.Client,
	movieID string,
) *redis.IntCmd {
	return redisClient.HDel(getKey(movieID))
}

func incrementVoteInCache(client *redis.Client, movieID string) *redis.IntCmd {
	pipeliner := client.Pipeline()

	intCMD := pipeliner.HIncrBy(getKey(movieID), "all_votes", 1)
	if intCMD.Err() != nil {
		return intCMD
	}

	return pipeliner.HIncrBy(getKey(movieID), "current_votes", 1)
}

func setMovieInCacheOneErr(
	redisWorkerChan chan<- redisWorker,
	redisClient *redis.Client,
	movieItem *movie.Movie,
	err error,
	op string,
) {
	switch {
	case err == redis.Nil:
		// Sets this movie in redis cache and handle error
		setMovieInCacheAndHandleErr(redisWorkerChan, redisClient, movieItem)
	default:
		logger.Log.Warn(
			"error occurred working with cache",
			zap.String("Operation", op),
			zap.Error(err),
		)
	}
}

func sendRedisErrToChan(redisWorkerChan chan<- redisWorker, cmd redis.Cmder, action string) {
	select {
	case <-time.After(time.Duration(time.Second * 3)):
	case redisWorkerChan <- redisWorker{
		statusCMD: cmd,
		err:       cmd.Err(),
		args:      cmd.Args(),
		action:    action,
	}:
	}
}

func sendFailedRedisCMDToWorker(redisWorkerChan chan<- redisWorker, redisWorkerStruct *redisWorker) {
	redisWorkerChan <- *redisWorkerStruct
}

func getListIDs(
	redisWorkerChan chan<- redisWorker, redisClient *redis.Client, start, stop int64,
) ([]string, error) {

	strSliceCMD := redisClient.LRange(moviesList, start, stop)

	// strSliceLen := len(strSliceCMD.Val())

	// switch {
	// case strSliceLen < pageSize:
	// 	strSliceCMD.Val() = strSliceCMD.Val()[:]
	// case strSliceLen < start:
	// 	strSliceCMD.Val() = strSliceCMD.Val()[:]

	// }

	if err := strSliceCMD.Err(); err != nil {
		go sendFailedRedisCMDToWorker(redisWorkerChan, &redisWorker{
			statusCMD: strSliceCMD,
			err:       err,
			args:      strSliceCMD.Args(),
			action:    "Get Govie Ids",
		})
		return nil, errRedisCmdFailed(err, "ListMovie: getting ids")
	}

	return strSliceCMD.Val(), nil
}

func loadMoviesIDToCache(
	ctx context.Context,
	redisWorkerChan chan<- redisWorker,
	redisClient *redis.Client,
	db *sql.DB,
) error {
	// Get all movies
	// Load their id in movies list

	// Prepare query
	query := `SELECT id FROM movies`
	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return errQueryFailed(err, "GetMovies (SELECT)")
	}
	defer rows.Close()

	for rows.Next() {
		movieID := ""

		err = rows.Scan(&movieID)
		if err != nil {
			return rows.Err()
		}

		// Load the movie ids in redis set cache
		intCMD := redisClient.LPush(moviesList, movieID)
		if err := intCMD.Err(); err != nil {
			logger.Log.Warn("error adding movie id to set", zap.Error(err))
			continue
		}
	}

	return nil
}

// func loadMoviesToCache(
// 	ctx context.Context,
// 	redisWorkerChan chan<- redisWorker,
// 	redisClient *redis.Client,
// 	db *sql.DB,
// ) error {
// 	// Get all movies
// 	// Load them in cache
// 	// Load their id in movies list

// 	// Prepare query
// 	query := `SELECT * FROM movies`
// 	// Execute query
// 	rows, err := db.QueryContext(ctx, query)
// 	if err != nil {
// 		return errQueryFailed(err, "GeTMovies (SELECT)")
// 	}

// 	for rows.Next() {
// 		movieItem := &movie.Movie{}
// 		audienceLabel := ""
// 		createTime := ""
// 		category := make([]byte, 0)
// 		photos := make([]byte, 0)

// 		// `id` varchar(50) NOT NULL,
// 		// `title` varchar(50) NOT NULL,
// 		// `price` varchar(15) NOT NULL DEFAULT 'NA',
// 		// `description` text NOT NULL,
// 		// `trailer_url` text,
// 		// `audience_label` enum('NR','GE','PG','PG-13','NC-17') NOT NULL DEFAULT 'NR',
// 		// `ratings` float NOT NULL DEFAULT '0',
// 		// `duration` int(11) DEFAULT 0,
// 		// `all_votes` int(11) DEFAULT 0,
// 		// `release_date` date DEFAULT NULL,
// 		// `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 		// `category` json DEFAULT NULL,
// 		// `photos` json DEFAULT NULL,

// 		err = rows.Scan(
// 			&movieItem.Id,
// 			&movieItem.Title,
// 			&movieItem.Price,
// 			&movieItem.Description,
// 			&movieItem.TrailerUrl,
// 			&audienceLabel,
// 			&movieItem.Ratings,
// 			&movieItem.MovieDurationMins,
// 			&movieItem.AllVotes,
// 			&movieItem.ReleaseDate,
// 			&createTime,
// 			&category,
// 			&photos,
// 		)

// 		if err != nil {
// 			return rows.Err()
// 		}

// 		err = json.Unmarshal(category, movieItem.Category)
// 		if err != nil {
// 			logger.Log.Warn("json unmarshal failed on Movie.Category", zap.Error(err))
// 			continue
// 		}

// 		err = json.Unmarshal(photos, movieItem.Photos)
// 		if err != nil {
// 			logger.Log.Warn("json unmarshal failed on Movie.Photos", zap.Error(err))
// 			continue
// 		}

// 		movieItem.AudienceLabel = movie.Audience(movie.Audience_value[audienceLabel])

// 		// Load the movie ids in cache
// 		redisClient.LPush(moviesList, movieItem.Id)

// 	}

// 	return nil
// }
