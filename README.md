# MAL2Go
MyAnimeList V2 API wrapper for GoLang

## Installation
In a terminal, run
```
go get github.com/MikunoNaka/MAL2Go
```
Go needs to be installed and `$GOPATH` should be set up

## Package Structure
[anime](anime) package
contains all the functionality for pulling data about anime from the API.

[user](user) package
has the functionality for getting user data and updating information.

[user/animelist](user/anime) package
has the functionality to update the authenticated user's anime list status, etc.

[util](util) package has some code used 
by multiple packages that I think don't belong particularly to one single package

[errhandlers](errhandlers) package
contains the validators and error handlers mainly to be used by MAL2Go.

## Authentication
MyAnimeList V2 API uses OAuth to authenticate a user.

Some useful links about authenticating with MyAnimeList:
- A Client ID can be generated [here](https://myanimelist.net/apiconfig).
This is needed to create an auth token for use with your program.

- Official documentation for authenticating with MyAnimeList is [here](https://myanimelist.net/apiconfig/references/authorization)

- An awesome guide to simplify this process by [ZeroCrystal](https://myanimelist.net/blog.php?eid=835707).

- An easy but not the best (still really good for getting started quickly/testing)
way to generate a token is to use [my script](https://github.com/MikunoNaka/mal-authtoken-generator).

Each package has a Client struct that needs an `AuthToken: string` among other (optional)
values that don't do anything yet. 

**NOTE:** I'm new to OAuth and might make changes to the Client struct in each package.

## Licence
Licenced under GNU General Public Licence

GNU GPL License: [LICENSE](LICENSE)

Copyright (c) 2022 Vidhu Kant Sharma
