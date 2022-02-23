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
	"io/ioutil"
	"log"
	"net/http"
)

// Handles HTTP request with your OAuth token as a Header
func (c Client) requestHandler(endpoint string) string {
  // generate request
  req, err := http.NewRequest("GET", endpoint, nil)
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

  return string(body)
}
