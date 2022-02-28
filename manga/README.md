# MAL2Go/manga
MAL2Go `manga` package has functionality related to getting data about manga.

To *update* manga status (score, status, etc) refer to [`user/manga`](../user/manga) package.

## Installation
In a terminal, run
``` fish
go get "github.com/MikunoNaka/MAL2Go/manga"
```

## Usage
Firstly, import this package and instanciate the client.
``` go
import (
  "github.com/MikunoNaka/MAL2Go/manga"
)
```

Now instanciate with
``` go
myClient := manga.Client {
  AuthToken: "Bearer " + yourTokenHere,
}
```

- ### Searching for a manga
``` go
searchString := "kanojo okarishimasu" // your search string here

// max amount of results to pull. Max is 500
limit := 10

// if the offset is 2 it will skip the first 2 results, then pull the next 10
offset := 0

// the API by default only returns some basic data
// you can specify some fields as a []string slice.
// it will return the default fields even if they aren't specified
var DefaultMangaFields []string = []string{
  "id", "title", "main_picture",
  "alternative_titles", "start_date", "end_date",
  "synopsis", "mean", "rank",
  "popularity", "num_list_users", "num_scoring_users",
  "nsfw", "created_at", "media_type",
  "status", "genres", "my_list_status",
  "num_volumes", "num_chapters", "authors",
  "pictures", "background", "related_anime",
  "related_manga", "recommendations", "serialization",
} // for all default fields fields := []string{} will also work

// finally making the API request
searchResults, err := myClient.SearchManga(searchString, limit, offset, fields)

// searchResults.Mangas is a list of all the Mangas returned by the API as search results
// print list of the search results
for _, manga := range searchResults.Mangas {
  fmt.Println(manga.Title)
}

// results have page numbers
fmt.Println(searchResults.Paging.NextPage, searchResults.Paging.PrevPage)
```

- ### Getting a manga's info
Each manga on MyAnimeList has a unique ID, which you need to find it

``` go
mangaId := 108407
fields := []string{} // pull every field

manga, err := myClient.GetMangaById(mangaId, fields)
if err != nil {
  fmt.Println(err)
}

fmt.Println(manga.Title, manga.MeanScore, manga.ListStatus.Status)
```

- ### Get manga ranking
Ranking is a list of mangas sorted by their rank

Possible ranking types are:
- `all`
- `manga`
- `novels`
- `oneshots`
- `doujin`
- `manhwa`
- `manhua`
- `bypopularity`
- `favorite`

``` go
rankingType := "favorite"
limit, offset := 10, 0
fields := []string{"title"}

ranking, err := myClient.GetMangaRanking(rankingType, limit, offset, fields)

// loop over the array mangas returned by the API
for _, manga := range ranking.Mangas {
  fmt.Printf("Title: %s, Rank Number: %d", manga.Title, manga.RankNum)
}

// ranking lists have page numbers
fmt.Println(ranking.Paging.NextPage, ranking.Paging.PrevPage)
```

## Structure
- [manga.go](anime.go)
Contains all the exported functions for pulling data from the API.

- [manga.structs.go](anime.structs.go)
Contains all the structs representing a manga entry on MyAnimeList.

- [ranking.structs.go](ranking.structs.go)
Representing anime ranking data both in the form returned 
by the API and the formatted form to be returned by this package.

- [search.structs.go](search.structs.go)
Representing search results.

- [request_handler.go](request_handler.go)
Responsible for making HTTP requests.

- [client.go](client.go)
The Client for accessing the API with this package.
