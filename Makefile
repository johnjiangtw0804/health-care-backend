#
# PostgreSQL Environment Variables
#
.EXPORT_ALL_VARIABLES:
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= jonathan
DB_PASSWORD ?= john0804
DB_NAME ?= health-care
DATABASE_URL ?= sslmode=disable host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME}

#
# postgres
#
stop-pg:
	@echo "stop postgres..."
	@docker stop health-care-pg | true > /dev/null

restart-pg: stop-pg
	@echo "restart postgres..."
	@docker run -d --rm --name health-care-pg \
				-p 5432:5432 -e POSTGRES_DB=health-care \
				-e POSTGRES_USER=jonathan -e POSTGRES_PASSWORD=john0804 \
				postgres:13.4-alpine | true

#
# Assume we are all mac users
# Before using the Go bindings, you must install the libpostal C library. Make sure you have the following prerequisites:
#
install-prerequisites:
	@brew install curl autoconf automake libtool pkg-config || true
run:restart-pg
	@echo "install ... "
	@go run main.go