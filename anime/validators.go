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
      "favorite":
      return true
    }
  return false
}

// Checks if given rankingType is valid
func areValidFields(field string) bool {
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
      "avarage_episode_duration",
      "rating",
      "pictures",
      "background",
      "related_anime",
      "related_manga",
      "recommendations",
      "studios",
      "statistics":
      return true
    }
  return false
}

