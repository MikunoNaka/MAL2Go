/* MAL2Go - MyAnimeList V2 API wrapper for Go
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

package manga

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
  "strconv"
  "strings"
  "errors"
  "github.com/MikunoNaka/MAL2Go/v3/util"
)

type UpdateResponse struct {
  Status         string `json:"status"`
  Score          int    `json:"score"`
  VolumesRead    int    `json:"num_volumes_read"`
  ChaptersRead   int    `json:"num_chapters_read"`
  IsRereading    bool   `json:"is_rereading"`
  Priority       string `json:"priority"`
  TimesReread    string `json:"num_times_reread"`
  RereadValue    string `json:"reread_value"`
  Tags           string `json:"tags"`
  Comments       string `json:"string"`
}

// Handles HTTP request with your OAuth token as a Header
func (c Client) requestHandler(endpoint, method string) (string, error) {
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

  // error handling (if API returns error)
  var errMsg util.APIError
  json.Unmarshal(body, &errMsg)
  if errMsg.Err != "" {
    return string(body), errors.New(errMsg.Err + " " + errMsg.Msg)
  }

  // for DeleteAnime, its endpoint returns null data
  if method == "DELETE" {
    return strconv.Itoa(res.StatusCode), nil
  }

  return string(body), nil
}

// for PUT requests (used for updating anime)
func (c Client) putRequestHandler(endpoint string, params url.Values) (UpdateResponse, error) {
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

  // error handling (if API returns error)
  var errMsg util.APIError
  json.Unmarshal(body, &errMsg)
  if errMsg.Err != "" {
    return UpdateResponse{}, errors.New(errMsg.Err + " " + errMsg.Msg)
  }

  var resp UpdateResponse
  json.Unmarshal(body, &resp)

  return resp, nil
}
