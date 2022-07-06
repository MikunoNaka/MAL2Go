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

package errhandlers

// Checks if given rankingType is valid
func IsValidRankingType(rankingType string) bool {
  switch rankingType {
    case
      "all",
      "airing",
      "upcoming",
      "tv",
      "ova",
      "movie",
      "special",
      "bypopularity",
      "favorite": return true
    }
  return false
}

// Checks if given field is valid
func IsValidField(field string) bool {
  switch field {
    case
      "id",
      "title",
      "main_picture",
      "alternative_titles",
      "start_date",
      "end_date",
      "synopsis",
      "mean",
      "rank",
      "popularity",
      "num_list_users",
      "num_scoring_users",
      "nsfw",
      "created_at",
      "updated_at",
      "media_type",
      "status",
      "genres",
      "my_list_status",
      "list_status",
      "num_episodes",
      "start_season",
      "broadcast",
      "source",
      "average_episode_duration",
      "rating",
      "pictures",
      "background",
      "related_anime",
      "related_manga",
      "recommendations",
      "studios",
      "statistics": return true
    }
  return false
}

// Checks if given season is valid
func IsValidSeason(season string) bool {
  switch season {
    case
      "winter",
      "spring",
      "summer",
      "fall": return true
    }
  return false
}

// Checks if given sort is valid
// For seasonal anime lists
func IsValidSeasonalSort(sort string) bool {
  switch sort {
    case
      "anime_score",
      "anime_num_list_users": return true
    }
  return false
}

// Checks if given sort is valid
// for user anime lists
func IsValidListSort(sort string) bool {
  switch sort {
    case
      "list_score",
      "list_updated_at",
      "anime_title",
      "anime_start_date",
      "anime_id": return true
    }
  return false
}

// Checks if given anime list status is valid
func IsValidListStatus(status string) bool {
  switch status {
    case
      "watching",
      "completed",
      "on_hold",
      "dropped",
      "plan_to_watch": return true
    }
  return false
}

// Checks if given anime score is valid
func IsValidScore(score int) bool {
  if score >= 0 && score <= 10 {
    return true
  } 
  return false
}

// Checks if given anime priority is valid
func IsValidPriority(priority int) bool {
  if priority >= 0 && priority <= 2 {
    return true
  } 
  return false
}

// Checks if given rewatch value is valid
func IsValidRewatchValue(r int) bool {
  if r >= 0 && r <= 5 {
    return true
  } 
  return false
}
