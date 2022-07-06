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
  "github.com/MikunoNaka/MAL2Go/v4/manga"
)

type mangaListRaw struct {
  Data []struct {
    Manga      manga.Manga      `json:"node"`
    ListStatus manga.ListStatus `json:"list_status"`
  } `json:"data"`
  Paging struct {
    NextPage string `json:"next"`
    PrevPage string `json:"previous"`
  } `json:"paging"`
}

type UpdateMangaData struct {
  Status         string
  IsRereading    bool
  Score          int
  VolumesRead    int
  ChaptersRead   int
  Priority       int
  TimesReread    int
  RereadValue    int
  Tags           string
  Comments       string
}
