# MAL2Go/anime
MAL2Go `anime` package has functionality related to getting data about anime.

To *update* anime status (score, status, etc) refer to [`user/anime`](../user/anime) package.

## Usage
Firstly, import this package and instanciate the client.
``` go
import (
  "github.com/MikunoNaka/MAL2Go/anime"
)
```

Now instanciate with
``` go
myClient := anime.Client {
  AuthToken: "Bearer " + yourTokenHere,
}
```

- ### Searching for an anime
``` go
searchString := "mushishi" // your search string here

// max amount of results to pull. Max is 500
limit := 10

// if the offset is 2 it will skip the first 2 results, then pull the next 10
offset := 0

// the API by default only returns some basic data
// you can specify some fields as a []string slice.
// it will return the default fields even if they aren't specified
fields := []string{
  "id", "title", "main_picture",
  "alternative_titles", "start_date",
  "end_date", "synopsis", "mean", "rank",
  "popularity", "num_list_users",
  "num_scoring_users", "nsfw", "created_at",
  "updated_at", "media_type", "status",
  "genres", "my_list_status", "num_episodes",
  "start_season", "broadcast", "source",
  "average_episode_duration", "rating",
  "pictures", "background", "related_anime",
  "related_manga", "recommendations",
  "studios", "statistics",
} // for all default fields fields := []string{} will also work

// finally making the API request
searchResults, err := myClient.SearchAnime(searchString, limit, offset, fields)

fmt.Println(searchResults.Animes) // print list of the search results

// results have page numbers
fmt.Println(searchResults.ListPaging.NextPage, searchResults.ListPaging.PrevPage)
```

- ### Getting an anime's info
Each anime on MyAnimeList has a unique ID, which you need to find it

Refer to [anime.structs.go](anime.structs.go) to find out all the keys the Anime struct has

``` go
animeId := 42351
fields := []string{} // pull every field

anime, err := myClient.GetAnimeById(animeId, fields)
if err != nil {
  fmt.Println(err)
}

fmt.Println(anime.Title, anime.MeanScore, anime.MyListStatus.Status)
```

- ### Get anime ranking
Ranking is a list of anime sorted by their rank

Possible ranking types are:
- `all`
- `airing`
- `upcoming`
- `tv`
- `ova`
- `movie`
- `special`
- `bypopularity`
- `favorite`

``` go
rankingType := "favorite"
limit, offset := 10, 0
fields := []string{"title", "media_type"}

ranking, err := myClient.GetAnimeRanking(rankingType, limit, offset, fields)

// loop over the array of "titles" returned by the API
for _, rankingAnime := range ranking.Titles {
  anime := rankingAnime.Anime
  rankNum := rankingAnime.RankNum

  fmt.Printf("Title: %s, Rank Number: %d", anime.Title, rankNum)
}

// ranking lists have page numbers
fmt.Println(ranking.Paging.NextPage, ranking.Paging.PrevPage)
```

- ### Get seasonal anime
Get a list of anime from a particular season/year

Possible seasons are:
- `winter`
- `spring`
- `summer`
- `fall`

Possible ways to sort are:
- `anime_score`
- `anime_num_list_users`

``` go
year := "2021"
season := "winter"
sort := "anime_score"

limit, offset := 10, 0

fields := []string{"title"}

seasonalAnime, err := myClient.GetSeasonalAnime(year, season, sort, limit, offset, fields)
if err != nil {
  fmt.Println(err)
}

for _, anime := range seasonalAnime.Animes {
  fmt.Println(anime.Title)
}

fmt.Println(seaonalAnime.Season)
fmt.Println(seasonalAnime.Paging.NextPage, seasonalAnime.Paging.PrevPage)
```

- ### Get suggested anime
Returns suggestions related to the authenticated user

``` go
limit, offset := 10, 0
fields := []string{"title"}

suggestedAnime, err := myClient.GetSuggestedAnime(limit, offset, fields)
if err != nil {
  fmt.Println(err)
}

for _, anime := range suggestedAnime.Animes {
  fmt.Println(anime.Titile)
}

fmt.Println(suggestedAnime.ListPaging.NextPage, suggestedAnime.ListPaging.PrevPage)
```

## Structure
- [anime.go](anime.go)
Contains all the exported functions for pulling data from the API.

- [anime.structs.go](anime.structs.go)
Contains all the structs representing an anime entry on MyAnimeList.

- [client.go](client.go)
The Client for accessing the API with this package.

- [general.structs.go](general.structs.go)
Some structs that can't particularly be grouped with another file.

- [ranking.structs.go](ranking.structs.go)
Representing anime ranking data both in the form returned by the API and
the formatted form to be returned by this package.

- [request_handler.go](request_handler.go)
Responsible for making HTTP requests

- [search.structs.go](search.structs.go)
Representing search results.

- [seasonal.structs.go](seasonal.structs.go)
Representing seasonal anime list.

- [suggestedanime.structs.go](suggestedanime.structs.go)
Representing suggested anime data.
