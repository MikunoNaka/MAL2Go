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

package manga

import (
  "encoding/json"
  "strconv"
  "fmt"
  e "github.com/MikunoNaka/MAL2Go/v4/errhandlers"
  u "github.com/MikunoNaka/MAL2Go/v4/util"
  m "github.com/MikunoNaka/MAL2Go/v4/manga"
)

const BASE_URL string = "https://api.myanimelist.net/v2"
const maxListLimit int = 1000

// Delete a manga from user's manga list
func (c Client)DeleteManga(id int) (string, error) {
  endpoint := fmt.Sprintf("%s/manga/%d/my_list_status", BASE_URL, id)
  /* Returns 200 if manga successfully deleted
   * Alternatively returns 404 if manga not in user's manga list */
  return c.requestHandler(endpoint, "DELETE")
}

// Get authenticated user's manga list
// returns true as second value if there are more mangas present
func (c Client) GetMangaList(user, status, sort string, limit, offset int, nsfw bool, fields []string) ([]m.Manga, bool, error){
  var userMangaList []m.Manga
  // error handling for limit
  limitErr := e.LimitErrHandler(limit, maxListLimit)
  if limitErr != nil { 
    return userMangaList, false, limitErr
  }

  // handle all the errors for the fields
  fields, err := e.MangaFieldsErrHandler(fields)
  if err != nil {
    return userMangaList, false, err
  }

  // append "list_status" field only used by this func.
  fields = append(fields, "list_status")

  // checks if valid sort is specified
  if !e.IsValidMangaListSort(sort) {
    return userMangaList, false, e.InvalidSortError
  }

  // checks if valid status is specified
  if status != "" && !e.IsValidMangaListStatus(status) {
    return userMangaList, false, e.InvalidStatusError
  }

  // get own list if user not specified
  if user == "" {
    user = "@me"
  }

  var endpoint string
  // if status is "" it returns all anime
  if status == "" {
    endpoint, _ = u.UrlGenerator(
      BASE_URL + "/users/" + user + "/mangalist",
      []string{"sort", "limit", "offset", "fields", "nsfw"},
      [][]string{{sort}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields, {strconv.FormatBool(nsfw)}},
      true,
    )
  } else {
    // status gets included if specified
    endpoint, _ = u.UrlGenerator(
      BASE_URL + "/users/" + user + "/mangalist",
      []string{"status", "sort", "limit", "offset", "fields", "nsfw"},
      [][]string{{status}, {sort}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields, {strconv.FormatBool(nsfw)}},
      true,
    )
  }

  // get data from API
  var mangaListData mangaListRaw
  data, err := c.requestHandler(endpoint, "GET")
  if err != nil {
    return userMangaList, false, err
  }
  json.Unmarshal([]byte(data), &mangaListData)

  nextPageExists := mangaListData.Paging.NextPage != ""

  // set MyListStatus for each element and add it to slice
  for _, element := range mangaListData.Data {
    m := element.Manga
    m.ListStatus = element.ListStatus

    userMangaList = append(userMangaList, m)
  }

  return userMangaList, nextPageExists, nil
}

