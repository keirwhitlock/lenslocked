GO_VERSION := 1.24.5

.PHONY: build

build:
	go build -o app .

run:
	go run main.go

deploy: build
	cd .terraform
	AWS_PROFILE=kwhitlockdev terraform plan -out tfplan.out
	AWS_PROFILE=kwhitlockdev terraform apply tfplan.out

psql:
	 docker compose exec -it db psql -U baloo -d lenslocked

restart:
	docker compose restart
