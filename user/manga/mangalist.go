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
  "errors"
  e "github.com/MikunoNaka/MAL2Go/errhandlers"
  u "github.com/MikunoNaka/MAL2Go/util"
)

const BASE_URL string = "https://api.myanimelist.net/v2"
const maxListLimit int = 1000

// Delete a manga from user's manga list
func (c Client)DeleteManga(id int) string {
  endpoint := fmt.Sprintf("%s/manga/%d/my_list_status", BASE_URL, id)
  /* Returns 200 if manga successfully deleted
   * Alternatively returns 404 if manga not in user's manga list */
  return c.requestHandler(endpoint, "DELETE")
}

// Get authenticated user's manga list
func (c Client) GetMangaList(user, status, sort string, limit, offset int, fields []string) (MangaList, error){
  var userMangaList MangaList
  // error handling for limit
  limitErr := e.LimitErrHandler(limit, maxListLimit)
  if limitErr != nil { 
    return userMangaList, limitErr
  }

  // handle all the errors for the fields
  fields, err := e.FieldsErrHandler(fields)
  if err != nil {
    return userMangaList, err
  }

  // append "list_status" field only used by this func.
  fields = append(fields, "list_status")

  // checks if valid sort is specified
  if !e.IsValidMangaListSort(sort) {
    return userMangaList, errors.New(fmt.Sprintf("GetMangaList: Invalid sort specified: \"%s\"", sort))
  }

  // checks if valid status is specified
  if status != "" && !e.IsValidMangaListStatus(status) {
    return userMangaList, errors.New(fmt.Sprintf("GetMangaList: Invalid status specified: \"%s\"", status))
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
      []string{"sort", "limit", "offset", "fields"},
      [][]string{{sort}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields},
      true,
    )
  } else {
    // status gets included if specified
    endpoint, _ = u.UrlGenerator(
      BASE_URL + "/users/" + user + "/mangalist",
      []string{"status", "sort", "limit", "offset", "fields"},
      [][]string{{status}, {sort}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields},
      true,
    )
  }

  // get data from API
  var mangaListData MangaListRaw
  data := c.requestHandler(endpoint, "GET")
  json.Unmarshal([]byte(data), &mangaListData)

  // set MyListStatus for each element and add it to array
  var mangas []Manga
  for _, element := range mangaListData.Data {
    a := element.Manga
    a.ListStatus = element.ListStatus

    mangas = append(mangas, a)
  }

  // finally create AnimeList
  userMangaList = MangaList {
    Mangas: mangas,
    Paging: mangaListData.Paging,
  }

  return userMangaList, nil
}

