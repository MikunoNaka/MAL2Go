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

type AnimePicture struct {
  Medium string `json:"large"`
  Large  string `json:"medium"`
}

type StatusStatistics struct {
  Watching    string `json:"watching"` 
  Completed   string `json:"completed"`
  OnHold      string `json:"on_hold"`
  Dropped     string `json:"dropped"`
  PlanToWatch string `json:"plan_to_watch"`
}

type AnimeStatistics struct {
  Status       StatusStatistics `json:"status"`
  NumListUsers int              `json:"num_list_users"`
}

type Genre struct {
  Id   int    `json:"id"`
  Name string `json:"name"`
}

type ListStatus struct {
  Status       string `json:"status"`
  Score        int    `json:"score"`
  EpWatched    int    `json:"num_episodes_watched"`
  IsRewatching bool   `json:"is_rewatching"`
  UpdatedAt    string `json:"updated_at"`
}

type Season struct {
  Year int    `json:"year"`
  Name string `json:"season"`
}

type Broadcast struct {
  Day  string `json:"day_of_the_week"`
  Time string `json:"start_time"`
}

type Related struct {
  Anime                 Anime  `json:"node"`
  RelationType          string `json:"relation_type"`
  RelationTypeFormatted string `json:"relation_type_formatted"`
}

type Studio struct {
  Id   int    `json:"id"`
  Name string `json:"name"`
}

type Recommendation struct {
  Anime Anime `json:"node"`
  Num   int   `json:"num_recommendations"`
}

type AltTitles struct {
  Synonyms []string `json:"synonyms"`
  En       string   `json:"en"`
  Ja       string   `json:"ja"`
}

type Anime struct {
  Id                int              `json:"id"`
  Title             string           `json:"title"`
  MainPicture       AnimePicture     `json:"main_picture"`
  AltTitles         AltTitles        `json:"alternative_titles"`
  StartDate         string           `json:"start_date"`
  EndDate           string           `json:"end_date"`
  Synopsis          string           `json:"synopsis"`
  MeanScore         float32          `json:"mean"`
  Rank              int              `json:"rank"`
  Popularity        int              `json:"popularity"`
  NumListUsers      int              `json:"num_list_users"`
  NumScoringUsers   int              `json:"num_scoring_users"`
  NsfwStatus        string           `json:"nsfw"` // find out what values are there
  CreatedAt         string           `json:"created_at"`
  UpdatedAt         string           `json:"updated_at"`
  MediaType         string           `json:"media_type"`
  Status            string           `json:"status"`
  Genres            []Genre          `json:"genres"`
  MyListStatus      ListStatus       `json:"my_list_status"`
  NumEpisodes       int              `json:"num_episodes"`
  StartSeason       Season           `json:"start_season"`
  Broadcast         Broadcast        `json:"broadcast"`
  Source            string           `json:"source"`
  DurationSeconds   int              `json:"average_episode_duration"`
  // Rating as in R, PG13, etc
  Rating            string           `json:"rating"`
  Pictures          []AnimePicture   `json:"pictures"`
  Background        string           `json:"background"`
  RelatedAnime      []Related        `json:"related_anime"`
  Recommendations   []Recommendation `json:"recommendations"`
  Studios           []Studio         `json:"studios"`
  Statistics        AnimeStatistics  `json:"statistics"`
}
