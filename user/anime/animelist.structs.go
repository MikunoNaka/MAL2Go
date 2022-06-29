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
  "github.com/MikunoNaka/MAL2Go/v2/anime"
)

type animeListRaw struct {
  Data []struct {
    Anime      anime.Anime      `json:"node"`
    ListStatus anime.ListStatus `json:"list_status"`
  } `json:"data"`
  Paging struct {
    NextPage string `json:"next"`
    PrevPage string `json:"previous"`
  } `json:"paging"`
}

type UpdateAnimeData struct {
  Status         string
  IsRewatching   bool
  Score          int
  EpWatched      int
  Priority       int
  TimesRewatched int
  RewatchValue   int
  Tags           string
  Comments       string
}
