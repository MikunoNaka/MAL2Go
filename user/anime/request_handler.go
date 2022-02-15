/* mal2go - MyAnimeList V2 API wrapper for Go
 * Copyright (C) 2022  Vidhu Kant Sharma <vidhukant@protonmail.ch>

 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.

 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>. */

package anime

import (
  "strings"
	"encoding/json"
  "net/url"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type serverResponse struct {
  Message string
  Error string
}

// Handles HTTP request with your OAuth token as a Header
func (c AnimeListClient) requestHandler(endpoint, method string) string {
  // generate request
  req, err := http.NewRequest(method, endpoint, nil)
  if err != nil {
      log.Fatal(err)
  }
  req.Header.Add("Authorization", c.AuthToken)

  // do request
  res, err := c.HttpClient.Do(req)
  if err != nil {
      log.Fatal(err)
  }
  defer res.Body.Close()

  // read body
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
      log.Fatal(err)
  }

  // for DeleteAnime, its endpoint returns null data
  if method == "DELETE" {
    return strconv.Itoa(res.StatusCode)
  }

  return string(body)
}

// for PUT requests (used by UpdateAnime)
func (c AnimeListClient) putRequestHandler(endpoint string, updateData UpdateAnimeData) serverResponse {
  // TODO: make this do other stuff
  params := url.Values{}

  /* NOTE: THIS WILL OVERWRITE ANY DATA THAT 
   * IS NOT SPECIFIED AND SET IT TO NULL */
  params.Set("status",               updateData.Status)
  params.Set("is_rewatching",        strconv.FormatBool(updateData.IsRewatching))
  params.Set("score",                strconv.Itoa(updateData.Score))
  params.Set("num_watched_episodes", strconv.Itoa(updateData.EpWatched))
  params.Set("priority",             strconv.Itoa(updateData.Priority))
  params.Set("num_times_rewatched",  strconv.Itoa(updateData.TimesRewatched))
  params.Set("rewatch_value",        strconv.Itoa(updateData.RewatchValue))
  params.Set("tags",                 updateData.Tags)
  params.Set("comments",             updateData.Comments)

  paramsEncoded := params.Encode()

  // generate request
  req, err := http.NewRequest(http.MethodPut, endpoint, strings.NewReader(paramsEncoded))
  if err != nil {
      log.Fatal(err)
  }
  req.Header.Add("Authorization", c.AuthToken)
  // this makes the sending-data-to-server magic work
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Content-Length", strconv.Itoa(len(paramsEncoded)))

  // do request
  res, err := c.HttpClient.Do(req)
  if err != nil {
      log.Fatal(err)
  }
  defer res.Body.Close()

  // read body
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
      log.Fatal(err)
  }

  // server response, ie message / error
  var resp serverResponse
  json.Unmarshal(body, &resp)

  return resp
}
