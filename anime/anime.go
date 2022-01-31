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

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

const BASE_URL string = "https://api.myanimelist.net/v2/anime"

// Each anime has its own ID on MAL
func GetAnimeById(token string, animeId int, fields []string) (Anime, error) {
  var anime Anime

  // Check if given fields are valid
  for _, j := range(fields) {
    if !isValidField(j) {
      return anime, errors.New(fmt.Sprintf("GetAnimeById: Invalid field specified: \"%s\"", j))
    }
  }

  // default fields to use when none are specified
  defaultFields := []string{
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

  if cap(fields) == 0 {
    fields = defaultFields
    log.Println("GetAnimeById: WARN: No fields specified, using all default fields to get data")
  }

  endpoint, _ := urlGenerator(
    BASE_URL + "/" + strconv.Itoa(animeId),
    []string{"fields"},
    /* it seems to still return all fields from the API. 
     * this might be an issue with MAL itself
     * TODO: look into this */
    [][]string{fields},
    true,
  )

  data := requestHandler(token, endpoint)
  json.Unmarshal([]byte(data), &anime)

  return anime, nil
}

// Ranking is a list of anime sorted by their rank
func GetAnimeRanking(token string, rankingType string, limit int, offset int) (AnimeRanking, error) {
  var animeRanking AnimeRanking

  // if limit exceeds what MAL supports
  if limit > 500 {
    return animeRanking, errors.New(fmt.Sprintf("GetAnimeRanking: Limit too high(%d). Max limit is 500", limit))
  } else if offset > 499 {
    return animeRanking, errors.New(fmt.Sprintf("GetAnimeRanking: Offset too high(%d). Max offset for mal2go is 499", offset))
  }

  // if ranking type is invalid
  if !isValidRankingType(rankingType) {
    return animeRanking, errors.New(fmt.Sprintf("GetAnimeRanking: Invalid ranking type specified: \"%s\"", rankingType))
  }

  endpoint, _ := urlGenerator(
    BASE_URL + "/ranking",
    []string{"ranking_type", "limit", "offset"},
    [][]string{{rankingType}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}},
    true,
  )

  // gets data from API and stores it in a struct
  var rankingData RawRanking
  data := requestHandler(token, endpoint)
  json.Unmarshal([]byte(data), &rankingData)

  // Adding all the animes in ranking list to a slice
  var animeRankingTitles []AnimeRankingTitle
  for _, element := range rankingData.Data {
    animeRankingTitles = append(
      animeRankingTitles,
      AnimeRankingTitle {
        Anime:   element.Anime, 
        RankNum: element.Ranking.Rank,
      },
    )
  }

  // Finally, create the AnimeRanking object
  animeRanking = AnimeRanking {
    Titles: animeRankingTitles,
    Paging: ListPaging {
      NextPage: rankingData.Paging.NextPage,
      PrevPage: rankingData.Paging.PrevPage,
    },
  }

  return animeRanking, nil
}
