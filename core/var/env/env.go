package env

import (
	"os"
	"strconv"
	"strings"
)

// Port variables
var (
	PortTaskQueueService = strings.TrimSpace(os.Getenv("PORT_TASKQUEUE_SERVICE"))
)

// DB config variables
var (
	DBHost     = strings.TrimSpace(os.Getenv("DB_HOST"))
	DBPort, _  = strconv.Atoi(strings.TrimSpace(os.Getenv("DB_PORT")))
	DBUser     = strings.TrimSpace(os.Getenv("DB_USER"))
	DBPassword = strings.TrimSpace(os.Getenv("DB_PASSWORD"))

	DBNameSmsQueue = strings.TrimSpace(os.Getenv("DB_DBNAME_SMSQUEUE"))
)

// Utilze variables
var (
	JwtSecretKey = []byte(strings.TrimSpace(os.Getenv("JWT_SECRET_KEY")))
)
