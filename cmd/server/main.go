package main

import (
	"bufio"
	"context"
	"flag"
	"github.com/Sirupsen/logrus"
	http_server "github.com/gidyon/rupacinema/movie/internal/protocol/http"
	"os"
	"strconv"

	"github.com/gidyon/rupacinema/movie/pkg/config"
)

var (
	defaultLogLevel      = 0
	defaultLogTimeFormat = "2006-01-02T15:04:05Z07:00"
)

func main() {
	var (
		cfg      = &config.Config{}
		useFlags bool
	)

	flag.BoolVar(
		&useFlags,
		"uflag", false,
		"Whether to pass config in flags",
	)

	// gRPC section
	flag.StringVar(
		&cfg.GRPCPort,
		"grpc-port", ":5500",
		"gRPC port to bind",
	)

	// DB section
	flag.StringVar(
		&cfg.DBHost,
		"db-host", "mysqldb",
		"Database host",
	)
	flag.StringVar(
		&cfg.DBUser,
		"db-user", "root",
		"Database user",
	)
	flag.StringVar(
		&cfg.DBPassword,
		"db-password", "hakty11",
		"Database password",
	)
	flag.StringVar(
		&cfg.DBSchema,
		"db-schema", "rupa-movie",
		"Database Schema to use",
	)

	// Redis section
	flag.StringVar(
		&cfg.RedisHost,
		"redis-host", "redisdb",
		"Redis host",
	)
	flag.StringVar(
		&cfg.RedisUser,
		"redis-user", "rupacinema",
		"Redis user",
	)
	flag.StringVar(
		&cfg.RedisPort,
		"redis-port", ":6379",
		"Redis port",
	)
	flag.StringVar(
		&cfg.RedisPassword,
		"redis-password", "hakty11",
		"Redis password",
	)

	// Logging section
	flag.IntVar(
		&cfg.LogLevel,
		"log-level", defaultLogLevel,
		"Global log level",
	)
	flag.StringVar(
		&cfg.LogTimeFormat,
		"log-time-format", defaultLogTimeFormat,
		"Print time format for logger e.g 2006-01-02T15:04:05Z07:00",
	)

	// TLS Certificate and Private key paths for service
	flag.StringVar(
		&cfg.TLSCertPath,
		"tls-cert", "certs/cert.pem",
		"Path to TLS certificate for the service",
	)
	flag.StringVar(
		&cfg.TLSKeyPath,
		"tls-key", "certs/key.pem",
		"Path to Private key for the service",
	)

	// External Services
	// Notification Service
	flag.StringVar(
		&cfg.NotificationServiceAddress,
		"notification-host", "localhost",
		"Address of the notification service",
	)
	flag.StringVar(
		&cfg.NotificationServicePort,
		"notification-port", ":5540",
		"Port where the notification service is running",
	)
	flag.StringVar(
		&cfg.NotificationServiceCertPath,
		"notification-cert", "certs/cert.pem",
		"Path to TLS certificate for notification service",
	)
	// Account Service
	flag.StringVar(
		&cfg.AccountServiceAddress,
		"account-host", "localhost",
		"Address of the account service",
	)
	flag.StringVar(
		&cfg.AccountServicePort,
		"account-port", ":5500",
		"Port where the account service is running",
	)
	flag.StringVar(
		&cfg.AccountServiceCertPath,
		"account-cert", "certs/cert.pem",
		"Path to TLS certificate for account service",
	)

	flag.Parse()

	if !useFlags {
		// Get from environmnent variables
		cfg = &config.Config{
			// GRPC section
			GRPCPort: os.Getenv("GRPC_PORT"),
			// Mysql section
			DBHost:     os.Getenv("MYSQL_HOST"),
			DBUser:     os.Getenv("MYSQL_USER"),
			DBPassword: os.Getenv("MYSQL_PASSWORD"),
			DBSchema:   os.Getenv("MYSQL_DATABASE"),
			// Redis section
			RedisURL:      os.Getenv("REDIS_URL"),
			RedisHost:     os.Getenv("REDIS_HOST"),
			RedisPort:     os.Getenv("REDIS_PORT"),
			RedisUser:     os.Getenv("REDIS_USER"),
			RedisPassword: os.Getenv("REDIS_PASSWORD"),
			// TLS certificate and private key paths
			TLSCertPath: os.Getenv("TLS_CERT_PATH"),
			TLSKeyPath:  os.Getenv("TLS_KEY_PATH"),
			// External services section
			// Notification service
			NotificationServiceAddress:  os.Getenv("NOTIFICATION_ADDRESS"),
			NotificationServicePort:     os.Getenv("NOTIFICATION_PORT"),
			NotificationServiceCertPath: os.Getenv("NOTIFICATION_CERT_PATH"),
			// Account service
			AccountServiceAddress:  os.Getenv("ACCOUNT_ADDRESS"),
			AccountServicePort:     os.Getenv("ACCOUNT_PORT"),
			AccountServiceCertPath: os.Getenv("ACCOUNT_CERT_PATH"),
		}
		logLevel := os.Getenv("LOG_LEVEL")
		logTimeFormat := os.Getenv("LOG_TIME_FORMAT")

		// Log Level
		if logLevel == "" {
			cfg.LogLevel = defaultLogLevel
		} else {
			logLevelInt64, err := strconv.ParseInt(logLevel, 10, 64)
			if err != nil {
				panic(err)
			}
			cfg.LogLevel = int(logLevelInt64)
		}

		// Log Time Format
		if logTimeFormat == "" {
			cfg.LogTimeFormat = defaultLogTimeFormat
		} else {
			cfg.LogTimeFormat = logTimeFormat
		}
	}

	cfg.RedisURL = cfg.RedisHost + cfg.RedisPort

	ctx, cancel := context.WithCancel(context.Background())

	s := bufio.NewScanner(os.Stdin)
	defer cancel()

	logrus.Infof(
		"Type %q or %q or %q or %q to stop the service",
		"kill", "KILL", "quit", "QUIT",
	)

	// Shutdown when user press q or Q
	go func() {
		for s.Scan() {
			if s.Text() == "kill" || s.Text() == "KILL" || s.Text() == "quit" || s.Text() == "QUIT" {
				cancel()
				return
			}
		}
	}()

	if err := http_server.Serve(ctx, cfg); err != nil {
		cancel()
		logrus.Fatalf("%v\n", err)
	}
}
