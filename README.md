# Gator

Gator is a command-line RSS feed aggregator written in Go. It lets users register, add RSS feeds, follow feeds, periodically scrape posts from those feeds, and browse saved posts from the terminal.

## Requirements

Before running Gator, install:

- Go
- PostgreSQL

For development and database setup, you will also want:

- Goose for migrations
- SQLC for generated database code

## Install

Install the CLI with:

```bash
go install github.com/amanuelmr/gator@latest
```

After installation, run the compiled binary directly:

```bash
gator
```

`go run .` is useful during development, but `gator` is the production command after `go install`.

## Configuration

Create a config file in your home directory:

```bash
~/.gatorconfig.json
```

Example:

```json
{
  "db_url": "postgres://amanuel:@localhost:5432/gator?sslmode=disable"
}
```

Replace the connection string with your local PostgreSQL connection string. The `current_user_name` field is managed by the app.

## Database Setup

Create a local PostgreSQL database named `gator`, then run migrations from the schema directory:

```bash
cd sql/schema
goose postgres "postgres://amanuel:@localhost:5432/gator" up
```

If you change SQL queries during development, regenerate the database package from the project root:

```bash
sqlc generate
```

## Commands

Register and log in:

```bash
gator register alice
gator login alice
```

Add and list feeds:

```bash
gator addfeed "Hacker News RSS" "https://news.ycombinator.com/rss"
gator feeds
```

Follow and unfollow feeds:

```bash
gator follow "https://news.ycombinator.com/rss"
gator following
gator unfollow "https://news.ycombinator.com/rss"
```

Start the feed scraper:

```bash
gator agg 1m
```

The `agg` command runs continuously and fetches feeds every duration you provide, such as `10s`, `1m`, or `1h`. Stop it with `Ctrl+C`.

Browse saved posts:

```bash
gator browse
gator browse 10
```

`browse` defaults to 2 posts when no limit is provided.

Reset all users and their feeds:

```bash
gator reset
```

This is mainly useful while developing or testing locally.
