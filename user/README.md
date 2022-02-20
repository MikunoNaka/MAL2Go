# MAL2Go/user
MAL2Go `user` package has functionality related to getting data about the authenticated user.

## Usage
Firstly, import this package and instanciate the client.
``` go
import (
  "github.com/MikunoNaka/MAL2Go/user"
)
```

Now instanciate with
``` go
myClient := user.Client {
  AuthToken: "Bearer " + yourTokenHere,
}
```

- ### Get authenticated user's info
``` go
userData := myClient.GetSelfUserInfo()

fmt.Println(userData.Name, userData.Picture)
```

The `User` datatype has the following fields:
- `Id` `int`
- `Name` `string`
- `Picture` `string`
- `Gender` `string`
- `Birthday` `string`
- `Location` `string`
- `JoinedAt` `string`
- `TimeZone` `string`
- `IsSupporter` `bool`

## Structure
- [user.go](user.go)
Contains all the exported functions for pulling data from the API.

- [user.structs.go](user.structs.go)
Represents the user data returned by the API

- [client.go](client.go)
The Client for accessing the API with this package.

- [request_handler.go](request_handler.go)
Responsible for making HTTP requests
