/* mal2go - MyAnimeList V2 API wrapper for Go
 * Copyright (C) 2022  Vidhu Kant Sharma <vidhukant@protonmail.ch>

 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.

 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>. */

package anime

import (
	"encoding/json"
	"fmt"
  "github.com/MikunoNaka/mal2go/anime"
)

const BASE_URL string = "https://api.myanimelist.net/v2"

// Get authenticated user's anime list
func (c AnimeListClient) GetAnimeList(user, status, sort string/*, limit, offset int*/) {
  // get own list if user not specified
  if user == "" {
    user = "@me"
  }

  var userAnimeList anime.AnimeList
  endpoint := BASE_URL + "/users/0ZeroTsu/animelist?fields=list_status&limit=4"

  // get data from API
  var animeListData AnimeListRaw
  data := c.requestHandler(endpoint, "GET")
  json.Unmarshal([]byte(data), &animeListData)

  // set MyListStatus for each element and add it to array
  var animes []anime.Anime
  for _, element := range animeListData.Data {
    a := element.Anime
    a.MyListStatus = element.ListStatus

    animes = append(animes, a)
  }

  // finally create AnimeList
  userAnimeList = anime.AnimeList {
    Animes: animes,
    Paging: animeListData.Paging,
  }

  fmt.Println(userAnimeList)
}

