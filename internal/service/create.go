package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gidyon/rupacinema/movie/pkg/api"
	"github.com/gidyon/rupacinema/notification/pkg/api"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"strings"
	"time"
)

type createMovieDS struct {
	res *empty.Empty
	err error
}

func (createMovie *createMovieDS) Create(
	ctx context.Context,
	sqlWorkerChan chan<- sqlWorker,
	redisWorkerChan chan<- redisWorker,
	createReq *movie.CreateMovieRequest,
	db *sql.DB,
	redisClient *redis.Client,
	notificationServiceClient notification.NotificationServiceClient,
) {
	// Check if context is cancelled before proceeding
	if cancelled(ctx) {
		return
	}

	movieItem := createReq.GetMovie()

	err := func() error {
		var err error
		// Validate movie fields
		switch {
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
		createMovie.err = err
		return
	}

	// Create a unique Id for the movie
	movieItem.Id = uuid.New().String()

	// Check that the movie title doesn't exist
	exist, err := checkMovieExists(ctx, db, movieItem.Title)
	if err != nil {
		createMovie.err = err
		return
	}

	if exist {
		createMovie.err = errMovieExists(movieItem.Title)
		return
	}

	err = insertMovieToDB(ctx, sqlWorkerChan, db, movieItem)
	if err != nil {
		createMovie.err = err
		return
	}

	createMovie.res = &empty.Empty{}

	// Notify movies subscribers
	notificationServiceClient.MultiTrigger(
		ctx,
		&notification.Notification{
			NotificationId: uuid.New().String(),
			Priority:       notification.Priority_LOW,
			SendMethod:     notification.SendMethod_EMAIL_AND_SMS,
			CreateTime:     &timestamp.Timestamp{Nanos: int32(time.Now().Nanosecond())},
			EmailNotification: &notification.EmailNotification{
				Subject:         "New Movie Released",
				BodyContentType: "text/html",
				Body: fmt.Sprintf(
					"%s movie is now available at Rupa Mall Cinema. Check our website for more info and getting tickets for the show. Rupa Love :)",
					movieItem.Title,
				),
			},
			SmsNotification: &notification.SMSNotification{
				Message: fmt.Sprintf(
					"%s movie is now available at Rupa Mall Cinema. Check our website for more info and getting tickets for the show. Rupa Love :)",
					movieItem.Title,
				),
			},
			BulkChannel: "new-releases",
			Bulk:        true,
			Save:        false,
		},
	)

	setMovieInCacheAndHandleErr(
		redisWorkerChan,
		redisClient,
		movieItem,
	)
}

func insertMovieToDB(
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
	query := `INSERT INTO movies (id, title, price, description, trailer_url, audience_label, ratings, duration, all_votes, release_date, create_time, category, photos) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, DATE(?), ?, ?, ?)`

	// Execute query
	_, err = db.ExecContext(ctx, query,
		movieItem.Id,
		movieItem.Title,
		movieItem.Price,
		movieItem.Description,
		movieItem.TrailerUrl,
		movie.Audience_name[int32(movieItem.AudienceLabel)],
		movieItem.Ratings,
		movieItem.MovieDurationMins,
		movieItem.AllVotes,
		movieItem.ReleaseDate,
		`NOW()`,
		category,
		photos,
	)

	if err != nil {
		// Send the failed sql to channel worker
		go sendSQLErrToChan(sqlWorkerChan, &sqlWorker{
			query: query,
			args: []interface{}{
				movieItem.Id,
				movieItem.Title,
				movieItem.Price,
				movieItem.Description,
				movieItem.TrailerUrl,
				movie.Audience_name[int32(movieItem.AudienceLabel)],
				movieItem.Ratings,
				movieItem.MovieDurationMins,
				movieItem.AllVotes,
				movieItem.ReleaseDate,
				`NOW()`,
				category,
				photos,
			},
			err: err,
		})
		return errQueryFailed(err, "CreateMovie (INSERT)")
	}

	return nil
}

func sendSQLErrToChan(sqlWorkerChan chan<- sqlWorker, workerInfo *sqlWorker) {
	select {
	case <-time.After(time.Second * 3):
	case sqlWorkerChan <- *workerInfo:
	}
}

func checkMovieExists(
	ctx context.Context, db *sql.DB, movieName string,
) (bool, error) {
	// Check if context is cancelled before proceeding
	if cancelled(ctx) {
		return false, ctx.Err()
	}

	title := ""

	// Prepare query
	query := `SELECT title FROM movies WHERE id=?`

	// Execute query
	row := db.QueryRowContext(ctx, query, movieName)
	err := row.Scan(&title)

	switch err {
	case nil:
	case sql.ErrNoRows:
		return false, nil
	default:
		return false, errQueryFailed(err, "GetMovie (SELECT)")
	}

	return true, nil
}
