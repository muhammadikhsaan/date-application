package static

import (
	"os"
	"strings"
)

var (
	PORT         = strings.TrimSpace(os.Getenv("PORT"))
	MODE         = strings.TrimSpace(os.Getenv("MODE"))
	ENVIRONTMENT = strings.TrimSpace(os.Getenv("ENVIRONTMENT"))

	CLIENT_SCHEMA = strings.TrimSpace(os.Getenv("CLIENT_SCHEMA"))
	CLENT_URL     = strings.TrimSpace(os.Getenv("CLENT_URL"))

	REQUEST_TIMEOUT = strings.TrimSpace(os.Getenv("REQUEST_TIMEOUT"))

	JWT_ISSUER      = strings.TrimSpace(os.Getenv("JWT_ISSUER"))
	CREDENTIAL_PATH = strings.TrimSpace(os.Getenv("CREDENTIAL_PATH"))

	COOKIE_SECURE = strings.TrimSpace(os.Getenv("COOKIE_SECURE")) == "true"

	DATABASE_HOST     = strings.TrimSpace(os.Getenv("DATABASE_HOST"))
	DATABASE_PORT     = strings.TrimSpace(os.Getenv("DATABASE_PORT"))
	DATABASE_USER     = strings.TrimSpace(os.Getenv("DATABASE_USER"))
	DATABASE_PASSWORD = strings.TrimSpace(os.Getenv("DATABASE_PASSWORD"))
	DATABASE_NAME     = strings.TrimSpace(os.Getenv("DATABASE_NAME"))

	HASH_SALT = strings.TrimSpace(os.Getenv("HASH_SALT"))
	HASH_COST = strings.TrimSpace(os.Getenv("HASH_COST"))
)
