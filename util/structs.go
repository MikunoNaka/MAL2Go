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

package util

/* NOTE: MAL still seems to send some fields
 * even if they aren't requested.
 * those include Title, Picture, Id, etc */
// default fields to use when none are specified
var DefaultFields []string = []string{
  "id", "title", "main_picture",
  "alternative_titles", "start_date",
  "end_date", "synopsis", "mean", "rank",
  "popularity", "num_list_users",
  "num_scoring_users", "nsfw", "created_at",
  "updated_at", "media_type", "status",
  "genres", "my_list_status", "num_episodes",
  "start_season", "broadcast", "source",
  "average_episode_duration", "rating",
  "pictures", "background", "related_anime",
  "related_manga", "recommendations",
  "studios", "statistics",
}

// contains previous/next page for anime list
// we don't actually need this.
// TODO: for compatibility's sake, keep this but also define methods
// to get the prev. and next page's elements automatically
type ListPaging struct {
  NextPage string `json:"next"`
  PrevPage string `json:"previous"`
}

/* these structs are used 
 * both by anime and manga package */
type Picture struct {
  Medium string `json:"medium"`
  Large  string `json:"large"`
}

type StatusStatistics struct {
  Watching    string `json:"watching"` 
  Completed   string `json:"completed"`
  OnHold      string `json:"on_hold"`
  Dropped     string `json:"dropped"`
  PlanToWatch string `json:"plan_to_watch"`
}

type Genre struct {
  Id   int    `json:"id"`
  Name string `json:"name"`
}

type DefaultListStatus struct {
  Status       string `json:"status"`
  Score        int    `json:"score"`
  StartDate    string `json:"start_date"`
  FinishDate   string `json:"finish_date"`
  Priority     int    `json:"priority"`
  Tags         string `json:"tags"`
  Comments     string `json:"comments"`
  UpdatedAt    string `json:"updated_at"`
}

type AltTitles struct {
  Synonyms []string `json:"synonyms"`
  En       string   `json:"en"`
  Ja       string   `json:"ja"`
}
