# MAL2Go/manga
MAL2Go `manga` package has functionality related to getting data about anime.

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
```

- ### Getting a manga's info
``` go
```

- ### Get manga ranking
``` go
```

## Structure
- [manga.go](anime.go)
Contains all the exported functions for pulling data from the API.

- [manga.structs.go](anime.structs.go)
Contains all the structs representing a manga entry on MyAnimeList.

- [client.go](client.go)
The Client for accessing the API with this package.
