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
  "github.com/MikunoNaka/MAL2Go/v4/util"
  a "github.com/MikunoNaka/MAL2Go/v4/anime"
)

type Magazine struct {
  Id   int    `json:"id"`
  Name string `json:"name"`
}

type MangaSerialization struct {
  Magazine Magazine `json:"node"`
  Role     string   `json:"role"`
}

type Recommendation struct {
  Manga Manga `json:"node"`
  Num   int   `json:"num_recommendations"`
}

type Author struct {
  Id        int    `json:"id"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
}

type MangaAuthor struct {
  Author
  Role string `json:"role"`
}

type Related struct {
  Manga                 Manga  `json:"node"`
  RelationType          string `json:"relation_type"`
  RelationTypeFormatted string `json:"relation_type_formatted"`
}

type ListStatus struct {
  Status       string `json:"status"`
  Score        int    `json:"score"`
  StartDate    string `json:"start_date"`
  FinishDate   string `json:"finish_date"`
  Priority     int    `json:"priority"`
  Tags         string `json:"tags"`
  Comments     string `json:"comments"`
  UpdatedAt    string `json:"updated_at"`

  // thsese fields are exclusive to manga
  VolumesRead  int  `json:"num_volumes_read"`
  ChaptersRead int  `json:"num_chapters_read"`
  IsRereading  bool `json:"is_rereading"`
  TimesReread  int  `json:"num_times_reread"`
  RereadValue  int  `json:"reread_value"`
}

type Manga struct {
  Id              int                  `json:"id"`
  Title           string               `json:"title"`
  MainPicture     util.Picture         `json:"main_picture"`
  AltTitles       util.AltTitles       `json:"alternative_titles"`
  StartDate       string               `json:"start_date"`
  EndDate         string               `json:"end_date"`
  Synopsis        string               `json:"synopsis"`
  MeanScore       float32              `json:"mean"`
  Rank            int                  `json:"rank"`
  Popularity      int                  `json:"popularity"`
  NumListUsers    int                  `json:"num_list_users"`
  NsfwStatus      string               `json:"nsfw"`
  Genres          []util.Genre         `json:"genres"`
  CreatedAt       string               `json:"created_at"`
  UpdatedAt       string               `json:"updated_at"`
  MediaType       string               `json:"media_type"`
  Status          string               `json:"status"`
  /* MyListStatus refers to the authenticated user's info
   * while ListStatus can be used for other users */
  MyListStatus    ListStatus           `json:"my_list_status"`
  ListStatus      ListStatus           `json:"list_status"`
  NumVolumes      int                  `json:"num_volumes"`
  NumChapters     int                  `json:"num_chapters"`
  Authors         []MangaAuthor        `json:"authors"`
  Pictures        []util.Picture       `json:"pictures"`
  Background      string               `json:"background"`
  // NOTE: RelatedAnime doesn't seem to work, not sure if that's the API's fault
  RelatedAnime    []a.Related          `json:"related_anime"`
  RelatedManga    []Related            `json:"related_manga"`
  Recommendations []Recommendation     `json:"recommendations"`
  Serialization   []MangaSerialization `json:"serialization"`
}
