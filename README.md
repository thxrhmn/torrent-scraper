# TORRENT-SCRAPER

## GETTING STARTED

### Clone this repository
```
git clone https://github.com/thxrhmn/torrent-scraper.git
```

### Install dependencies
```
go get
```

```
cp .example_env .env
```

### Run the project
```
go run main.go
```

## API Documentation

### Endpoint: /api/v1/btdig

### Description
This endpoint is used to retrieve a list of torrent from the BTDIG.

### HTTP Method
- GET

### Query Parameters
- `keyword` (required): The keyword to be scrapped.
- `startpage` (optional): The page to start scraping from (default: 1).
- `endpage` (optional): The page to end scraping at (default: 2).
- `order` (optional): (default: 1 | "relevance")


### Example Usage
- Search for torrents with the title "Udemy":
```
GET /api/v1/btdig?keyword=Udemy
```

- Search for torrents with title and pagination
```
GET /api/v1/btdig?keyword=Udemy&startpage=1&endpage=3
```

- Searhc for torrents with title, pagination, and order

```
GET /api/v1/btdig?keyword=Udemy&startpage=1&endpage=3&order=age
```