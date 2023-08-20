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

## How to run this project
1. You must have <a href="https://github.com/pressly/goose">Goose</a> and <a href="https://github.com/sqlc-dev/sqlc">SQLC</a> installed on your machine.
2. Run the migration by going to the sql/schema folder inside the project `cd sql/schema` and type `goose postgres postgres://postgres:@localhost:5432/rss-agg up` (change the postgres URL with your own).
3. Build the project `go build && ./go-rss-aggregator`

### After you have successfully installed and built the project in order setup the dummy data, you must do the following:
1. Open Postman or any app to access an API
2. Access `http:localhost:8080/users` with the POST method and insert any name in the body, for example: `{ "name": "Thomas Shelby" }`
3. Copy the `api_key` after the user was created.
4. Access `http://localhost:8080/feed` with POST method. Add the `api_key` to the Headers `Authorization: ApiKey {api_key}`. Insert these two data from the body (raw): 
```
{
    "name": "Lane's Blogs",
    "url": "https://wagslane.dev/index.xml"
}
and
{
    "name": "Boot Dev Blogs",
    "url": "https://blog.boot.dev/index.xml"
}
```
After that, make sure to copy the `id` from the created feed.
<br><br>
5. Now visit `http://localhost:8080/v1/feed-follows` with the <b>previous headers for authorization</b>, and add the following to the body request:
```
{
  feed_id: {your_feed_id}
}
```
You can add two of your recently created feeds.
<br><br>
6. Now you can rebuild the project and see the scrapper is running every minute. 
