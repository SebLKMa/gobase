module github.com/sebmaspd/rnd/ieq/db/postgres

go 1.15

replace (
	github.com/sebmaspd/rnd/ieq/models => ../../models
)

require (
    github.com/lib/pq v1.10.0 // indirect
	github.com/sebmaspd/rnd/ieq/models v0.0.0-00010101000000-000000000000
)