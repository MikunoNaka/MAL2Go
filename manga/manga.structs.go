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
  "github.com/MikunoNaka/MAL2Go/util"
)

type ListPaging util.ListPaging

type Author struct {
  // TODO: add stuff
}

type ListStatus struct {
  util.DefaultListStatus
  VolumesRead  int  `json:"num_volumes_read"`
  ChaptersRead int  `json:"num_chapters_read"`
  IsRereading  bool `json:"is_rereading"`
  TimesReread  int  `json:"num_times_reread"`
  RereadValue  int  `json:"reread_value"`
}

type Manga struct {
  Id            int             `json:"id"`
  Title         string          `json:"title"`
  MainPicture   util.Picture    `json:"main_picture"`
  AltTitles     util.AltTitles  `json:"alternative_titles"`
  StartDate     string          `json:"start_date"`
  EndDate       string          `json:"end_date"`
  Synopsis      string          `json:"synopsis"`
  MeanScore     float32         `json:"mean"`
  Rank          int             `json:"rank"`
  Popularity    int             `json:"popularity"` 
  NumListUsers  int             `json:"num_list_users"`
  NsfwStatus    string         `json:"nsfw"`
  Genres        []util.Genre   `json:"genres"`
  CreatedAt     string         `json:"created_at"`
  UpdatedAt     string         `json:"updated_at"`
  MediaType     string         `json:"media_type"`
  Status        string         `json:"status"`
  MyListStattus ListStatus     `json:"my_list_status"`
  NumVolumes    int            `json:"num_volumes"`
  NumChapters   int            `json:"num_chapters"`
  Authors       []Author       `json:"authors"`
  Pictures      []util.Picture `json:"pictures"`
  Background    string         `json:"background"`
  /* TODO add these fields:
   * related_anime, related_manga, recommendations, serialization */
}
