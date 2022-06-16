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

package user

import (
  "encoding/json"
  "errors"
)

const BASE_URL string = "https://api.myanimelist.net/v2/users"

// Get info of logged in user
func (c Client) GetSelfUserInfo() (UserInfo, error) {
  /* MAL only supports @me for this */
  endpoint := BASE_URL + "/@me?fields=anime_statistics"
  
  // get data from API
  var userData UserInfo
  data, err := c.requestHandler(endpoint)
  if err != nil {
    return userData, err
  }

  json.Unmarshal([]byte(data), &userData)

  return userData, nil
}
