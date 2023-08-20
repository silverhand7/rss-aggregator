# RSS Aggregator App with Golang 
Note: this documentation is created mainly for me if I want to run this project in the future. You can follow this document to run this project in your local too.

This project is a REST API to scrap the data from websites (different websites might still need a couple of adjustments to match the pattern of the data).

This project helps me to practice with: 
- Develop real go project from scratch
- Goroutine and concurrency
- Channel
- Context
- Struct, and struct method
- Routing
- and more

# How to run this project
1. You must have <a href="https://github.com/pressly/goose">Goose</a> and <a href="https://github.com/sqlc-dev/sqlc">SQLC</a> installed on your machine.
2. Run the migration by going to the sql/schema folder inside the project `cd sql/schema` and type `goose postgres postgres://postgres:@localhost:5432/rss-agg up` (change the postgres URL with your own).
3. Build the project `go build && ./go-rss-aggregator`

## After you have successfully installed and built the project in order setup the dummy data
