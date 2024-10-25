source local.env

export MIGRATION_DSN="host=$DB_HOST port=$LOCAL_PG_INNER_PORT dbname=$LOCAL_PG_DATABASE_NAME user=$LOCAL_PG_USER password=$LOCAL_PG_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${LOCAL_MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v && echo "SUCCESS"