package service

import (
	"context"
	"database/sql"
	"github.com/gidyon/rupacinema/account/pkg/api"
	"github.com/gidyon/rupacinema/movie/pkg/api"
	"github.com/gidyon/rupacinema/notification/pkg/api"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
)

type movieAPIServer struct {
	ctx                       context.Context
	db                        *sql.DB
	redisClient               *redis.Client
	sqlWorkerChan             chan sqlWorker
	redisWorkerChan           chan redisWorker
	notificationServiceClient notification.NotificationServiceClient
	accountServiceClient      account.AccountAPIClient
}

type sqlWorker struct {
	query string
	args  []interface{}
	err   error
}

type redisWorker struct {
	statusCMD redis.Cmder
	err       error
	args      []interface{}
	action    string
}

// NewMovieAPI creates a new movie API server
func NewMovieAPI(
	ctx context.Context,
	redisClient *redis.Client,
	db *sql.DB,
	notificationServiceClient notification.NotificationServiceClient,
	accountServiceClient account.AccountAPIClient,
) (movie.MovieAPIServer, error) {
	movieSrv := &movieAPIServer{
		ctx:             ctx,
		db:              db,
		redisClient:     redisClient,
		sqlWorkerChan:   make(chan sqlWorker, 0),
		redisWorkerChan: make(chan redisWorker, 0),
		// Connection to remote services
		notificationServiceClient: notificationServiceClient,
		accountServiceClient:      accountServiceClient,
	}

	// Load the IDs of movies to cache
	err := loadMoviesIDToCache(
		ctx,
		movieSrv.redisWorkerChan,
		movieSrv.redisClient,
		movieSrv.db,
	)

	if err != nil {
		return nil, err
	}

	return movieSrv, nil
}

func (movieAPI *movieAPIServer) CreateMovie(
	ctx context.Context, createReq *movie.CreateMovieRequest,
) (*empty.Empty, error) {
	// Authenticate the request
	_, err := movieAPI.accountServiceClient.AuthenticateRequest(
		ctx, &empty.Empty{},
	)
	if err != nil {
		return nil, err
	}

	ctxCreate, cancel := context.WithCancel(ctx)
	defer cancel()

	createMovie := &createMovieDS{}

	createMovie.Create(
		ctxCreate,
		movieAPI.sqlWorkerChan,
		movieAPI.redisWorkerChan,
		createReq,
		movieAPI.db,
		movieAPI.redisClient,
		movieAPI.notificationServiceClient,
	)

	if cancelled(ctxCreate) {
		createMovie.err = contextError(ctx, "CreateMovie")
	}

	return createMovie.res, createMovie.err
}

func (movieAPI *movieAPIServer) UpdateMovie(
	ctx context.Context, updateReq *movie.UpdateMovieRequest,
) (*empty.Empty, error) {
	// Authenticate the request
	_, err := movieAPI.accountServiceClient.AuthenticateRequest(
		ctx, &empty.Empty{},
	)
	if err != nil {
		return nil, err
	}

	ctxUpdate, cancel := context.WithCancel(ctx)
	defer cancel()

	updateMovie := &updateMovieDS{}

	updateMovie.Update(
		ctxUpdate,
		movieAPI.sqlWorkerChan,
		movieAPI.redisWorkerChan,
		updateReq,
		movieAPI.db,
		movieAPI.redisClient,
	)

	if cancelled(ctxUpdate) {
		updateMovie.err = contextError(ctxUpdate, "UpdateMovie")
	}

	return updateMovie.res, updateMovie.err
}

func (movieAPI *movieAPIServer) DeleteMovie(
	ctx context.Context, delReq *movie.DeleteMovieRequest,
) (*empty.Empty, error) {
	// Authenticate the request
	_, err := movieAPI.accountServiceClient.AuthenticateRequest(
		ctx, &empty.Empty{},
	)
	if err != nil {
		return nil, err
	}

	ctxDel, cancel := context.WithCancel(ctx)
	defer cancel()

	deleteMovie := &deleteMovieDS{}

	deleteMovie.Delete(
		ctxDel,
		movieAPI.sqlWorkerChan,
		movieAPI.redisWorkerChan,
		delReq,
		movieAPI.db,
		movieAPI.redisClient,
	)

	if cancelled(ctxDel) {
		deleteMovie.err = contextError(ctxDel, "DeleteMovie")
	}

	return deleteMovie.res, deleteMovie.err
}

func (movieAPI *movieAPIServer) ListMovies(
	ctx context.Context, listReq *movie.ListMoviesRequest,
) (*movie.ListMoviesResponse, error) {
	return nil, nil
}

func (movieAPI *movieAPIServer) GetMovie(
	ctx context.Context, getReq *movie.GetMovieRequest,
) (*movie.Movie, error) {
	ctxGet, cancel := context.WithCancel(ctx)
	defer cancel()

	getMovie := &getMovieDS{}

	getMovie.Get(
		ctxGet,
		movieAPI.sqlWorkerChan,
		movieAPI.redisWorkerChan,
		getReq,
		movieAPI.db,
		movieAPI.redisClient,
	)

	if cancelled(ctxGet) {
		getMovie.err = contextError(ctx, "GetMovie")
	}

	return getMovie.res, getMovie.err
}

func (movieAPI *movieAPIServer) RequestMovieReplay(
	ctx context.Context, repReq *movie.RequestMovieReplayRequest,
) (*empty.Empty, error) {
	return nil, nil
}
