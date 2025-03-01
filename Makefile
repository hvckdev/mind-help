Arguments := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

create_migration:
	migrate create -ext sql -dir migrations -seq $(Arguments)