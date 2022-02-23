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

// Anime but with an extra RankNum field
type rAnime struct {
  Anime
  RankNum int
}

// this is how the API returns data (looks horrible)
type RawRanking struct {
  Data []struct {
    Anime rAnime `json:"node"`
    Ranking struct {
      Rank int `json:"rank"`
    } `json:"ranking"`
  } `json:"data"`

  Paging ListPaging `json:"paging"`
}

// this is how mal2go returns data
type AnimeRanking struct {
  Animes []rAnime
  Paging ListPaging
}
