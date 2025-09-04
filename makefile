dev:
	air

gen-migration:
	migrate create -ext sql -dir db/migrations -seq $(name)