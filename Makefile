include Makefile.ansible

volumes:
	mkdir -p ./.docker/volumes/go/tls-certificates
	mkdir -p ./.docker/volumes/postgresql/data

env-docker-compose-development:
	rm -f docker-compose.yml
	ln -s ./docker/compose/development/docker-compose.yml docker-compose.yml

	cp .development.env .env

env-docker-compose-production:
	rm -f docker-compose.yml
	ln -s ./docker/compose/production/docker-compose.yml docker-compose.yml

	cp .production.env .env

env-up:
	docker-compose -f docker-compose.yml --env-file .env up -d

postgres-test-run:
	docker exec readytotouch_postgres_db psql -U yaaws_user -d yaaws -c "SELECT VERSION();"

logs:
	docker logs readytotouch_go_app

app:
	docker exec -it readytotouch_go_app sh

pg:
	docker exec -it readytotouch_postgres_db bash

pg-users:
	cat ./fixtures/postgres/users.sql | docker exec -i readytotouch_postgres_db psql -d yaaws -U yaaws_user

env-down:
	docker-compose -f docker-compose.yml --env-file .env down

env-down-with-clear:
	docker-compose -f docker-compose.yml --env-file .env down --remove-orphans -v # --rmi=all

app-build:
	docker exec readytotouch_go_app go build -o /bin/yaaws-server ./cmd/main.go

app-start:
	docker exec readytotouch_go_app yaaws-server

app-stop:
	docker exec readytotouch_go_app pkill yaaws-server || echo "yaaws-server already stopped"

app-restart: app-build app-stop app-start

test:
	docker exec readytotouch_go_app go test ./... -v -count=1

bench:
	docker exec readytotouch_go_app go test ./... -v -run=$$^ -bench=. -benchmem -benchtime=1000x

# make migrate-pgsql-create NAME=init
migrate-pgsql-create:
	goose -dir ./internal/storage/postgres/migrations -table schema_migrations create $(NAME) sql

migrate-pgsql-up:
	docker exec readytotouch_postgres_goose_migrations \
		goose --dir=/db/migrations --table=schema_migrations \
		up
migrate-pgsql-redo:
	docker exec readytotouch_postgres_goose_migrations \
		goose --dir=/db/migrations --table=schema_migrations \
		redo
migrate-pgsql-down:
	docker exec readytotouch_postgres_goose_migrations \
		goose --dir=/db/migrations --table=schema_migrations \
		down-to 20230815012800
migrate-pgsql-reset:
	docker exec readytotouch_postgres_goose_migrations \
		goose --dir=/db/migrations --table=schema_migrations \
		reset
migrate-pgsql-status:
	docker exec readytotouch_postgres_goose_migrations \
		goose --dir=/db/migrations --table=schema_migrations \
		status

generate-template:
	# go get -u github.com/valyala/quicktemplate/qtc
	qtc -dir=./internal/templates/v1 -skipLineComments
	qtc -dir=./internal/templates/dev -skipLineComments
	git add .

generate-sqlc:
	sqlc generate

esbuild-minify:
	npm --prefix=client i
	MINIFY=true npm run --prefix=client esbuild
	tree -h ./public/assets/js

esbuild:
	MINIFY=false npm run --prefix=client esbuild
	tree -h ./public/assets/js

# make design DESIGN="~/go/src/github.com/readytotouch/readytotouch.github.io"
design:
	$(eval DESIGN := ~/go/src/github.com/readytotouch/readytotouch.github.io)
	rm -rf ./public/assets/images/pages ./public/assets/fonts ./public/assets/js ./public/design
	mkdir -p ./public/assets/images/pages ./public/assets/fonts ./public/assets/js ./public/design

	cp -r $(DESIGN)/public/assets/images/pages/* ./public/assets/images/pages
	cp -r $(DESIGN)/public/assets/fonts/* ./public/assets/fonts
	cp -r $(DESIGN)/public/assets/js/* ./public/assets/js
	cp -r $(DESIGN)/public/*.html ./public/design

	git restore ./public/assets/images/pages/online-new/cz_flag.svg

	git add .

companies-and-connections:
	firefox ./public/chatgpt-design/companies-and-connections.html

generate-protos:
	# https://grpc.io/docs/languages/go/quickstart/
	# sudo apt install -y protobuf-compiler
	# protoc --version
	protoc -I . protos/auth/*.proto --go_out=./internal/

generate-organizers:
	go run ./cmd/organizers/generate/main.go

review-logos:
	go run ./cmd/organizers/review-logos/main.go

# POSTGRES_PASSWORD=$(echo "$RANDOM$RANDOM" | sha256sum | head -c 32; echo;) JWT_SECRET_KEY=$(echo "$RANDOM$RANDOM" | sha256sum | head -c 32; echo;) make generate-production-environment-file
generate-production-environment-file:
	touch .production.env

	grep -qF 'ENVIRONMENT=' .production.env || echo 'ENVIRONMENT="production"' >> .production.env

	# Database
	grep -qF 'POSTGRES_USER=' .production.env || echo 'POSTGRES_USER="u8user"' >> .production.env
	grep -qF 'POSTGRES_PASSWORD=' .production.env || echo 'POSTGRES_PASSWORD="$(POSTGRES_PASSWORD)"' >> .production.env
	grep -qF 'POSTGRES_DB=' .production.env || echo 'POSTGRES_DB="yaaws"' >> .production.env
	grep -qF 'POSTGRES_DSN=' .production.env || echo 'POSTGRES_DSN="postgresql://u8user:$(POSTGRES_PASSWORD)@postgres:5432/yaaws?sslmode=disable"' >> .production.env
	grep -qF 'HOSTS=' .production.env || echo 'HOSTS="readytotouch.com,dev.readytotouch.com,www.readytotouch.com"' >> .production.env

	grep -qF 'JWT_SECRET_KEY=' .production.env || echo 'JWT_SECRET_KEY="$(JWT_SECRET_KEY)"' >> .production.env

	grep -qF 'GITHUB_CLIENT_ID=' .production.env || echo 'GITHUB_CLIENT_ID="8dce25b763367e846763"' >> .production.env
	grep -qF 'GITHUB_CLIENT_SECRET=' .production.env || echo 'GITHUB_CLIENT_SECRET=""' >> .production.env
	grep -qF 'GITHUB_REDIRECT_URL=' .production.env || echo 'GITHUB_REDIRECT_URL="https://readytotouch.com/auth/github/callback"' >> .production.env
	grep -qF 'GITLAB_CLIENT_ID=' .production.env || echo 'GITLAB_CLIENT_ID="1f8bc1174d17998654c82400ff7a230c87d4e633327c17c2414f315f62b80d28"' >> .production.env
	grep -qF 'GITLAB_CLIENT_SECRET=' .production.env || echo 'GITLAB_CLIENT_SECRET=""' >> .production.env
	grep -qF 'GITLAB_REDIRECT_URL=' .production.env || echo 'GITLAB_REDIRECT_URL="https://readytotouch.com/auth/gitlab/callback"' >> .production.env
	grep -qF 'BITBUCKET_CLIENT_ID=' .production.env || echo 'BITBUCKET_CLIENT_ID="PY4qXGrqgvCS34DuqT"' >> .production.env
	grep -qF 'BITBUCKET_CLIENT_SECRET=' .production.env || echo 'BITBUCKET_CLIENT_SECRET=""' >> .production.env
	grep -qF 'BITBUCKET_REDIRECT_URL=' .production.env || echo 'BITBUCKET_REDIRECT_URL="https://readytotouch.com/auth/bitbucket/callback"' >> .production.env
	grep -qF 'LINKEDIN_OAUTH2_TOKEN=' .production.env || echo 'LINKEDIN_OAUTH2_TOKEN=""' >> .production.env

	cat .production.env

ssh:
	ssh -t root@70.34.247.27 "cd /var/go/readytotouch/; bash --login"

ssh-copy-tls-certificates:
	scp -r root@70.34.247.27:/var/go/readytotouch/.docker/volumes/go/tls-certificates ./.docker/volumes/go

unsafe-browser-extension:
	mkdir -p ./public/extensions/readytotouch-unsafe-browser-extension/icons
	cp ./public/favicon-16x16.png ./public/extensions/readytotouch-unsafe-browser-extension/icons
	cd ./public/extensions \
		&& rm -rf readytotouch-unsafe-browser-extension.zip \
		&& zip -r readytotouch-unsafe-browser-extension.zip readytotouch-unsafe-browser-extension \
		&& chmod 777 readytotouch-unsafe-browser-extension.zip
