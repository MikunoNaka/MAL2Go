/* MAL2Go - MyAnimeList V2 API wrapper for Go
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
  "strconv"
  "fmt"
  a "github.com/MikunoNaka/MAL2Go/anime"
  e "github.com/MikunoNaka/MAL2Go/errhandlers"
  u "github.com/MikunoNaka/MAL2Go/util"
)

const BASE_URL string = "https://api.myanimelist.net/v2"
const maxListLimit int = 1000

// Delete an anime from user's anime list
func (c Client)DeleteAnime(id int) (string, error) {
  endpoint := fmt.Sprintf("%s/anime/%d/my_list_status", BASE_URL, id)
  /* Returns 200 if anime successfully deleted
   * Alternatively returns 404 if anime not in user's anime list */
  return c.requestHandler(endpoint, "DELETE")
}

// Get authenticated user's anime list
// returns true as second value if there are more animes present
func (c Client) GetAnimeList(user, status, sort string, limit, offset int, fields []string) ([]a.Anime, bool, error){
  var userAnimeList []a.Anime
  // error handling for limit
  limitErr := e.LimitErrHandler(limit, maxListLimit)
  if limitErr != nil { 
    return userAnimeList, false, limitErr
  }

  // handle all the errors for the fields
  fields, err := e.FieldsErrHandler(fields)
  if err != nil {
    return userAnimeList, false, err
  }

  // append "list_status" field only used by this func.
  fields = append(fields, "list_status")

  // checks if valid sort is specified
  if !e.IsValidListSort(sort) {
    return userAnimeList, false, e.InvalidSortError
  }

  // checks if valid status is specified
  if status != "" && !e.IsValidListStatus(status) {
    return userAnimeList, false, e.InvalidStatusError
  }

  // get own list if user not specified
  if user == "" {
    user = "@me"
  }

  var endpoint string
  // if status is "" it returns all anime
  if status == "" {
    endpoint, _ = u.UrlGenerator(
      BASE_URL + "/users/" + user + "/animelist",
      []string{"sort", "limit", "offset", "fields"},
      [][]string{{sort}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields},
      true,
    )
  } else {
    // status gets included if specified
    endpoint, _ = u.UrlGenerator(
      BASE_URL + "/users/" + user + "/animelist",
      []string{"status", "sort", "limit", "offset", "fields"},
      [][]string{{status}, {sort}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields},
      true,
    )
  }

  // get data from API
  var animeListData animeListRaw
  data, err := c.requestHandler(endpoint, "GET")
  if err != nil {
    return userAnimeList, false, err
  }
  json.Unmarshal([]byte(data), &animeListData)

  nextPageExists := animeListData.Paging.NextPage != ""

  // set ListStatus for each element and add it to array
  for _, element := range animeListData.Data {
    a := element.Anime
    a.ListStatus = element.ListStatus

    userAnimeList = append(userAnimeList, a)
  }

  return userAnimeList, nextPageExists, nil
}

