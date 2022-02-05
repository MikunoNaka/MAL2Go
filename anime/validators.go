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

// Checks if given rankingType is valid
func isValidRankingType(rankingType string) bool {
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

// Checks if given rankingType is valid
func isValidField(field string) bool {
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
func isValidSeason(season string) bool {
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
func isValidSort(sort string) bool {
  switch sort {
    case
      "anime_score",
      "anime_num_list_users": return true
    }
  return false
}
