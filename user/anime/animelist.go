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
  "fmt"
)

const BASE_URL string = "https://api.myanimelist.net/v2"

// Get authenticated user's anime list
func (c AnimeListClient) GetAnimeList() {
  endpoint := BASE_URL + "/users/@me/animelist?fields=list_status&limit=4"

  data := c.requestHandler(endpoint, "GET")
  fmt.Println(data)
}

