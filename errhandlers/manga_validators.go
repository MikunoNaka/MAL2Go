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

// Checks if given field is valid
func IsValidMangaField(field string) bool {
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
      "num_volumes",
      "num_chapters",
      "authors",
      "pictures",
      "background",
      "related_anime",
      "related_manga",
      "recommendations",
      "serialization": return true
    }
  return false
}

// Checks if given ranking type is valid
func IsValidMangaRankingType(rankingType string) bool {
  switch rankingType {
    case
      "all",
      "manga",
      "novels",
      "oneshots",
      "doujin",
      "manhwa",
      "manhua",
      "bypopularity",
      "favorite": return true
    }
  return false
}

// Checks if given sort is valid
func IsValidMangaListSort(sort string) bool {
  switch sort {
    case
      "list_score",
      "list_updated_at",
      "manga_title",
      "manga_start_date",
      "manga_id": return true
    }
  return false
}

// Checks if given list status is valid
func IsValidMangaListStatus(status string) bool {
  switch status {
    case
      "reading",
      "completed",
      "on_hold",
      "dropped",
      "plan_to_read": return true
    }
  return false
}
