schema using goose https://github.com/pressly/goose

up schema 
``goose postgres postgres://user:password@localhost:5432/go-inventory up``

down schema
``goose postgres postgres://user:password@localhost:5432/go-inventory down``