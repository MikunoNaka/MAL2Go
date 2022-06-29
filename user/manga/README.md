# MAL2Go/user/manga
MAL2Go `user/manga` package has functionality related to updating the user's manga list.

To *get* anime data, refer to the [`manga`](../../manga) package.

**There are multiple possible server responses and errors currently haven't been implemented yet.**

## Installation
In a terminal, run
``` fish
go get "github.com/MikunoNaka/MAL2Go/v2/user/manga"
```

## Usage
Firstly, import this package and instanciate the client.
``` go
import (
  "github.com/MikunoNaka/MAL2Go/v2/user/manga"
)
```

Now instanciate with
``` go
myClient := manga.Client {
  AuthToken: "Bearer " + yourTokenHere,
}
```

- ### Delete a manga from user's anime list
``` go
mangaId := 108407 // manga's ID

resp := myClient.DeleteManga(mangaId)

/* if manga is successfully deleted, resp is 200
 * if manga isn't in the list resp is 404 */
fmt.Println(resp)
```

- ### Get user's manga list
Possible statuses are:
- `reading`
- `completed`
- `on_hold`
- `dropped`
- `plan_to_watch`

Leaving blank (`""`) gets all the anime

Possible sorts are:
- `list_score`
- `list_updated_at`
- `manga_title`
- `manga_start_date`
- `manga_id` (beta)

Leaving user blank (`""`) or as `"@me"` returns the authenticated user's list

``` go
user := "0ZeroTsu" 
status := "reading"
sort := "list_score"

limit := 1000 // max is 1000
offset := 0

// fields := []string{} means get all the fields
fields := []string{"title"}

mangaList, err := myClient.GetMangaList(user, status, sort, limit, offset, fields)
if err != nil {
  fmt.Println(err)
}

// mangaList.Mangas is an array of the mangas in the list
for _, manga := range mangaList.Mangas {
  fmt.Println(manga.Title)
}

fmt.Println(mangaList.Paging.NextPage, mangaList.Paging.PrevPage)
```

- ### Set a manga's status
``` go
mangaId := 108407 // manga's ID
status := "dropped"
resp, _ := myClient.SetStatus(mangaId, status)
fmt.Println(resp.Error, resp.Message)
```

- ### Set read volumes
``` go
mangaId := 108407 // manga's ID
volumesRead := 10
resp, _ := myClient.SetVolumesRead(mangaId, volumesRead)
fmt.Println(resp.Error, resp.Message)
```

- ### Set read chapters
``` go
mangaId := 108407 // manga's ID
chaptersRead := 150
resp, _ := myClient.SetChaptersread(mangaId, chaptersRead)
fmt.Println(resp.Error, resp.Message)
```

- ### Set is rereading status
``` go
mangaId := 108407 // manga's ID
isRereading := true
_, _ := myClient.SetIsRereading(mangaId, isRereading)
```

- ### Set a manga's score
``` go
mangaId := 108407 // manga's ID
score := 1
_, _ := myClient.SetScore(mangaId, score)
```

- ### Set a manga's priority
Priority on MyAnimeList ranges from 0 to 2
``` go
mangaId := 108407 // manga's ID
priority := 2
_, _ := myClient.SetPriority(mangaId, priority)
```

- ### Set a manga's reread value
Reread value on MyAnimeList ranges from 0 to 5
``` go
mangaId := 108407 // manga's ID
rereadValue := 4
_, _ := myClient.SetRereadValue(mangaId, rereadValue)
```

- ### Set a manga's reread count
Number of times user has reread the manga. There is no limit
``` go
mangaId := 108407 // manga's ID
rereadCount := 69
_, _ := myClient.SetRereadCount(mangaId, rereadCount)
```

- ### Set a manga's tags
``` go
mangaId := 108407 // manga's ID
tags := "tags"
_, _ := myClient.UpdateTags(mangaId, tags)
```

- ### Set a manga's comments
``` go
mangaId := 108407 // manga's ID
comments := "I hate but love this"
_, _ := myClient.UpdateComments(mangaId, comments)
```

- ### Update all fields of a manga
WARNING: this function can overwrite any data and set it to null 
if you don't specify a value to it.

Refrain/use it carefully to avoid data loss.

``` go
updateData := manga.UpdateMangaData {
  Status: "dropped",
  IsRereading: true,
  Score: 1,
  VolumesRead: 10,
  ChaptersRead: 150,
  Priority: 2,
  TimesReread: 69,
  RereadValue: 4,
  Tags: "tags",
  Comments: "I hate but love this",
}

mangaId := 108407 // manga's ID

resp, err := myClient.UpdateManga(mangaId, updateData)
if err != nil {
  fmt.Println(err)
}

fmt.Println(resp.Error, resp.Message)
```

## Structure
- [mangalist.go](mangalist.go)
Contains the exported functions to do some basic functions with manga lists.

- [mangalist.structs.go](mangalist.structs.go)
Contains all the structs representing mangalist data pulled from MyAnimeList.

- [client.go](client.go)
The Client for accessing the API with this package.

- [request_handler.go](request_handler.go)
Responsible for making HTTP requests

- [update_mangalist.go](update_mangalist.go)
Contains all the exported functions to update a manga entry in user's mangalist.
