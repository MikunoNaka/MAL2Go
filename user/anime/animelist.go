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
  "errors"
  a "github.com/MikunoNaka/MAL2Go/anime"
  e "github.com/MikunoNaka/MAL2Go/errhandlers"
)

const BASE_URL string = "https://api.myanimelist.net/v2"
const maxListLimit int = 1000

// Delete an anime from user's anime list
func (c Client)DeleteAnime(id int) string {
  endpoint := fmt.Sprintf("%s/anime/%d/my_list_status", BASE_URL, id)
  /* Returns 200 if anime successfully deleted
   * Alternatively returns 404 if anime not in user's anime list */
  return c.requestHandler(endpoint, "DELETE")
}

// Get authenticated user's anime list
func (c Client) GetAnimeList(user, status, sort string, limit, offset int) (a.AnimeList, error){
  var userAnimeList a.AnimeList
  // error handling for limit
  limitErr := e.LimitErrHandler(limit, maxListLimit)
  if limitErr != nil { 
    return userAnimeList, limitErr
  }

  // checks if valid sort is specified
  if !e.IsValidListSort(sort) {
    return userAnimeList, errors.New(fmt.Sprintf("GetAnimeList: Invalid sort specified: \"%s\"", sort))
  }

  // checks if valid status is specified
  if status != "" && !e.IsValidListStatus(status) {
    return userAnimeList, errors.New(fmt.Sprintf("GetAnimeList: Invalid status specified: \"%s\"", status))
  }

  // get own list if user not specified
  if user == "" {
    user = "@me"
  }

  // if status is "" it returns all anime
  var endpoint string
  if status == "" {
    endpoint = BASE_URL + "/users/" + user + "/animelist?sort=" + sort + "&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset)
  } else {
    endpoint = BASE_URL + "/users/" + user + "/animelist?status=" + status + "&sort=" + sort + "&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset)
  }

  // get data from API
  var animeListData AnimeListRaw
  data := c.requestHandler(endpoint, "GET")
  json.Unmarshal([]byte(data), &animeListData)

  // set MyListStatus for each element and add it to array
  var animes []a.Anime
  for _, element := range animeListData.Data {
    a := element.Anime
    a.ListStatus = element.ListStatus

    animes = append(animes, a)
  }

  // finally create AnimeList
  userAnimeList = a.AnimeList {
    Animes: animes,
    Paging: animeListData.Paging,
  }

  return userAnimeList, nil
}

