# TORRENT-SCRAPER

## FEATURES
- Filter
- Pagination
- Save to csv

## GETTING STARTED

### Clone this repository
```shell
git clone https://github.com/thxrhmn/torrent-scraper.git
```

### Install dependencies
```shell
go get
```

```shell
cp example_env .env
```

### Run the project
```shell
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
- `order` (optional): Order by relevance | age | size | files (default: relevance)


### Example Usage
- Search for torrents with the title "Udemy":
```
GET /api/v1/btdig?keyword=Udemy
```

- Search for torrents with title and pagination
```
GET /api/v1/btdig?keyword=Udemy&startpage=1&endpage=3
```

- Search for torrents with title, pagination, and order

```
GET /api/v1/btdig?keyword=Udemy&startpage=1&endpage=3&order=age
```

### Example response
```json
{
    "status": 200,
    "data": [
        {
            "Title": "[UdemyTuts] - Cisco Switched Network Implementation",
            "Date": "found 1 year ago",
            "Files": "55 Files",
            "Size": "443.87 MB",
            "Link": "https://btdig.com/b7387d3db3f2a2d4f6fb7c32efb24d88106fb09b/udemy",
            "MagnetURL": "magnet:?xt=urn:btih:b7387d3db3f2a2d4f6fb7c32efb24d88106fb09b&dn=%5BUdemyTuts%5D+-+Cisco+Switched+Network+Implementation&tr=udp://tracker.openbittorrent.com:80&tr=udp://tracker.opentrackr.org:1337/announce"
        },
        {
            "Title": "[UdemyTuts] - Data Science- Machine Learning and Statistical Modeling in R",
            "Date": "found 3 years ago",
            "Files": "228 Files",
            "Size": "1.23 GB",
            "Link": "https://btdig.com/712b248d849f3c36614b9698217e5c8aae6d46a5/udemy",
            "MagnetURL": "magnet:?xt=urn:btih:712b248d849f3c36614b9698217e5c8aae6d46a5&dn=%5BUdemyTuts%5D+-+Data+Science-+Machine+Learning+and+Statistical+Modeling+in+R&tr=udp://tracker.openbittorrent.com:80&tr=udp://tracker.opentrackr.org:1337/announce"
        },
    ]
}
```