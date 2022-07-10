module github.com/sebmaspd/rnd/ieq/frontend

go 1.16

replace (
	github.com/sebmaspd/rnd/ieq/configs => ../configs
	github.com/sebmaspd/rnd/ieq/db/postgres => ../db/postgres
	github.com/sebmaspd/rnd/ieq/formulas => ../formulas
	github.com/sebmaspd/rnd/ieq/interfaces => ../interfaces
	github.com/sebmaspd/rnd/ieq/models => ../models
	github.com/sebmaspd/rnd/ieq/ratings => ../ratings
	github.com/sebmaspd/rnd/ieq/sensors/awair => ../sensors/awair
	github.com/sebmaspd/rnd/ieq/sensors/uhoo => ../sensors/uhoo
	github.com/sebmaspd/rnd/ieq/tasks => ../tasks
	github.com/sebmaspd/rnd/ieq/utils => ../utils
	github.com/sebmaspd/rnd/ieq/utils/skiptree => ../utils/skiptree
)

require (
	github.com/christophwitzko/go-curl v0.0.0-20171216141518-4203158d6acb // indirect
	github.com/sebmaspd/rnd/ieq/configs v0.0.0-00010101000000-000000000000 // indirect
	github.com/sebmaspd/rnd/ieq/db/postgres v0.0.0-00010101000000-000000000000
	github.com/sebmaspd/rnd/ieq/formulas v0.0.0-00010101000000-000000000000 // indirect
	github.com/sebmaspd/rnd/ieq/interfaces v0.0.0-00010101000000-000000000000
	github.com/sebmaspd/rnd/ieq/models v0.0.0-00010101000000-000000000000
	github.com/sebmaspd/rnd/ieq/ratings v0.0.0-00010101000000-000000000000
	github.com/sebmaspd/rnd/ieq/sensors/awair v0.0.0-00010101000000-000000000000 // indirect
	github.com/sebmaspd/rnd/ieq/sensors/uhoo v0.0.0-00010101000000-000000000000 // indirect
	github.com/sebmaspd/rnd/ieq/tasks v0.0.0-00010101000000-000000000000 // indirect
	github.com/sebmaspd/rnd/ieq/utils v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
