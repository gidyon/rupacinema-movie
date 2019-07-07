package grpc

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gidyon/rupacinema/account/pkg/api"
	"github.com/gidyon/rupacinema/movie/pkg/api"
	"github.com/gidyon/rupacinema/movie/pkg/config"
	"github.com/gidyon/rupacinema/notification/pkg/api"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gidyon/rupacinema/movie/internal/service"
)

// Creates the service
func createMovieAPIServer(
	ctx context.Context, cfg *config.Config,
) (movie.MovieAPIServer, error) {
	// Create a *sql.DB instance
	db, err := createMySQLConn(cfg)
	if err != nil {
		return nil, err
	}

	// Creates a redis client
	client := newRedisClient(cfg)

	// Remote services
	// Notification service
	notificationServiceConn, err := dialNotificationService(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to notification service: %v", err)
	}

	// Account service
	accountServiceConn, err := dialAccountService(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to account service: %v", err)
	}

	// Close the connection when context is cancelled
	go func() {
		<-ctx.Done()
		notificationServiceConn.Close()
		accountServiceConn.Close()
	}()

	return service.NewMovieAPI(
		ctx,
		client,
		db,
		notification.NewNotificationServiceClient(notificationServiceConn),
		account.NewAccountAPIClient(accountServiceConn),
	)
}

// Opens a connection to mysql database
func createMySQLConn(cfg *config.Config) (*sql.DB, error) {
	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBSchema,
		param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return db, nil
}

// Client uses a pool of connections to redis database.
func newRedisClient(cfg *config.Config) *redis.Client {
	redisURL := func(a, b string) string {
		if a == "" {
			return b
		}
		return a
	}
	return redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    redisURL(cfg.RedisURL, ":6379"),
	})
}

// creates a connection to the notification service
func dialNotificationService(
	ctx context.Context, cfg *config.Config,
) (*grpc.ClientConn, error) {

	creds, err := credentials.NewClientTLSFromFile(cfg.NotificationServiceCertPath, "localhost")
	if err != nil {
		return nil, err
	}

	return grpc.DialContext(
		ctx,
		cfg.NotificationServiceAddress+cfg.NotificationServicePort,
		grpc.WithTransportCredentials(creds),
	)
}

// creates a connection to the accounts service
func dialAccountService(
	ctx context.Context, cfg *config.Config,
) (*grpc.ClientConn, error) {

	creds, err := credentials.NewClientTLSFromFile(cfg.AccountServiceCertPath, "localhost")
	if err != nil {
		return nil, err
	}

	return grpc.DialContext(
		ctx,
		cfg.AccountServiceAddress+cfg.AccountServicePort,
		grpc.WithTransportCredentials(creds),
	)
}
