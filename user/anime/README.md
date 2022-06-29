# MAL2Go/user/anime
MAL2Go `user/anime` package has functionality related to updating the user's anime list.

To *get* anime data, refer to the [`anime`](../../anime) package.

**There are multiple possible server responses and errors currently haven't been implemented yet.**

## Installation
In a terminal, run
``` fish
go get "github.com/MikunoNaka/MAL2Go/v2/user/anime"
```

## Usage
Firstly, import this package and instanciate the client.
``` go
import (
  "github.com/MikunoNaka/MAL2Go/v2/user/anime"
)
```

Now instanciate with
``` go
myClient := anime.Client {
  AuthToken: "Bearer " + yourTokenHere,
}
```

- ### Delete an anime from user's anime list
``` go
animeId := 457 // anime's ID

resp := myClient.DeleteAnime(animeId)

/* if anime is successfully deleted, resp is 200
 * if anime isn't in the list resp is 404 */
fmt.Println(resp)
```

- ### Get user's anime list
Possible statuses are:
- `watching`
- `completed`
- `on_hold`
- `dropped`
- `plan_to_watch`

Leaving blank (`""`) gets all the anime

Possible sorts are:
- `list_score`
- `list_updated_at`
- `anime_title`
- `anime_start_date`
- `anime_id` (beta)

Leaving user blank (`""`) or as `"@me"` returns the authenticated user's list

``` go
user := "0ZeroTsu" 
status := "watching"
sort := "list_score"

limit := 1000 // max is 1000
offset := 0

// fields := []string{} means get all the fields
fields := []string{"title"}

animeList, err := myClient.GetAnimeList(user, status, sort, limit, offset, fields)
if err != nil {
  fmt.Println(err)
}

// animeList.Animes is an array of the animes in the list
for _, anime := range animeList.Animes {
  fmt.Println(anime.Title)
}

fmt.Println(animeList.Paging.NextPage, animeList.Paging.PrevPage)
```

- ### Set an anime's status
``` go
animeId := 457 // anime's ID
status := "dropped"
resp, _ := myClient.SetStatus(animeId, status)
fmt.Println(resp.Error, resp.Message)
```

- ### Set watched episodes
``` go
animeId := 457 // anime's ID
epWatched := 22
resp, _ := myClient.SetWatchedEpisodes(animeId, epWatched)
fmt.Println(resp.Error, resp.Message)
```

- ### Set is rewatching status
``` go
animeId := 457 // anime's ID
isRewatching := true
_, _ := myClient.SetIsRewatching(animeId, isRewatching)
```

- ### Set an anime's score
``` go
animeId := 457 // anime's ID
score := 10
_, _ := myClient.SetScore(animeId, score)
```

- ### Set an anime's priority
Priority on MyAnimeList ranges from 0 to 2
``` go
animeId := 457 // anime's ID
priority := 2
_, _ := myClient.SetPriority(animeId, priority)
```

- ### Set an anime's rewatch value
Rewatch value on MyAnimeList ranges from 0 to 5
``` go
animeId := 457 // anime's ID
rewatchValue := 4
_, _ := myClient.SetRewatchValue(animeId, rewatchValue)
```

- ### Set an anime's rewatch count
Number of times user has rewatched the anime. There is no limit
``` go
animeId := 457 // anime's ID
rewatchCount := 69
_, _ := myClient.SetRewatchCount(animeId, rewatchCount)
```

- ### Set an anime's tags
``` go
animeId := 457 // anime's ID
tags := "tags"
_, _ := myClient.UpdateTags(animeId, tags)
```

- ### Set an anime's comments
``` go
animeId := 457 // anime's ID
comments := "I love this"
_, _ := myClient.UpdateComments(animeId, comments)
```

- ### Update all fields of an anime
WARNING: this function can overwrite any data and set it to null 
if you don't specify a value to it.

Refrain/use it carefully to avoid data loss.

``` go
updateData := anime.UpdateAnimeData {
  Status: "watching",
  IsRewatching: true,
  Score: 10,
  EpWatched: 22,
  Priority: 2,
  TimesRewatched: 69,
  RewatchValue: 4,
  Tags: "tags",
  Comments: "I love this",
}

animeId := 457 // anime's ID

resp, err := myClient.UpdateAnime(animeId, updateData)
if err != nil {
  fmt.Println(err)
}

fmt.Println(resp.Error, resp.Message)
```

## Structure
- [animelist.go](animelist.go)
Contains the exported functions to do some basic functions with anime lists.

- [animelist.structs.go](animelist.structs.go)
Contains all the structs representing animelist data pulled from MyAnimeList.

- [client.go](client.go)
The Client for accessing the API with this package.

- [request_handler.go](request_handler.go)
Responsible for making HTTP requests

- [update_animelist.go](update_animelist.go)
Contains all the exported functions to update an anime entry in user's animelist.
