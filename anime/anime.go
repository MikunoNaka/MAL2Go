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
	"strconv"
  e "github.com/MikunoNaka/mal2go/errhandlers"
  u "github.com/MikunoNaka/mal2go/util"
)

const BASE_URL string = "https://api.myanimelist.net/v2/anime"

// in MAL documentation this is named Get Anime List
func (c AnimeClient) SearchAnime(searchString string, limit, offset int, fields []string) (AnimeSearch, error) {
  var searchResults AnimeSearch

  // error handling for limit and offset
  limitsErr := e.LimitsErrHandler(limit, offset)
  if limitsErr != nil {
    return searchResults, limitsErr
  }

  // handle all the errors for the fields
  fields, err := e.FieldsErrHandler(fields)
  if err != nil {
    return searchResults, err
  }

  // generate endpoint url with custom params
  endpoint, _ := u.UrlGenerator(
    BASE_URL,
    []string{"q", "limit", "offset", "fields"},
    [][]string{{searchString}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields},
    true,
  )

  // gets data from API and stores it in a struct
  var animeSearchData AnimeSearchRaw
  data := c.requestHandler(endpoint)
  json.Unmarshal([]byte(data), &animeSearchData)

  // Adding all the animes to another list to get formatted results later
  var animes []Anime
  for _, element := range animeSearchData.Data { 
    animes = append(animes, element.Anime) 
  } 

  // finally generate AnimeList 
  searchResults = AnimeSearch {
    Animes: animes,
    Paging: animeSearchData.Paging,
  }

  return searchResults, nil
}

// Each anime has its own ID on MAL
func (c AnimeClient) GetAnimeById(animeId int, fields []string) (Anime, error) {
  var anime Anime

  // handle all the errors for the fields
  fields, err := e.FieldsErrHandler(fields)
  if err != nil {
    return anime, err
  }

  endpoint, _ := u.UrlGenerator(
    BASE_URL + "/" + strconv.Itoa(animeId),
    []string{"fields"},
    /* it seems to still return all fields from the API. 
     * this might be an issue with MAL itself
     * TODO: look into this */
    [][]string{fields},
    true,
  )

  data := c.requestHandler(endpoint)
  json.Unmarshal([]byte(data), &anime)

  return anime, nil
}

// Ranking is a list of anime sorted by their rank
func (c AnimeClient) GetAnimeRanking(rankingType string, limit, offset int, fields []string) (AnimeRanking, error) {
  var animeRanking AnimeRanking

  // error handling for limit and offset
  limitsErr := e.LimitsErrHandler(limit, offset)
  if limitsErr != nil {
    return animeRanking, limitsErr
  }

  // handle all the errors for the fields
  fields, err := e.FieldsErrHandler(fields)
  if err != nil {
    return animeRanking, err
  }

  // if ranking type is invalid
  if !e.IsValidRankingType(rankingType) {
    return animeRanking, errors.New(fmt.Sprintf("GetAnimeRanking: Invalid ranking type specified: \"%s\"", rankingType))
  }

  endpoint, _ := u.UrlGenerator(
    BASE_URL + "/ranking",
    []string{"ranking_type", "limit", "offset", "fields"},
    [][]string{{rankingType}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields},
    true,
  )

  // gets data from API and stores it in a struct
  var rankingData RawRanking
  data := c.requestHandler(endpoint)
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

// get list of animes from specified season
func (c AnimeClient) GetSeasonalAnime(year, season, sort string, limit, offset int, fields []string) (SeasonalAnime, error) {
  var seasonalAnime SeasonalAnime

  // error handling for limit and offset
  limitsErr := e.LimitsErrHandler(limit, offset)
  if limitsErr != nil {
    return seasonalAnime, limitsErr
  }

  // handle all the errors for the fields
  fields, err := e.FieldsErrHandler(fields)
  if err != nil {
    return seasonalAnime, err
  }

  // checks if valid season is specified
  if !e.IsValidSeason(season) {
    return seasonalAnime, errors.New(fmt.Sprintf("GetSeasonalAnime: Invalid season specified: \"%s\"", season))
  }

  // checks if valid sort is specified
  if !e.IsValidSort(sort) {
    return seasonalAnime, errors.New(fmt.Sprintf("GetSeasonalAnime: Invalid sort specified: \"%s\"", sort))
  }

  endpoint, _ := u.UrlGenerator(
    BASE_URL + fmt.Sprintf("/season/%s/%s", year, season),
    []string{"sort", "limit", "offset", "fields"},
    [][]string{{sort}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields},
    true,
  )

  // gets data from API and stores it in a struct
  var seasonalAnimeData SeasonalAnimeRaw
  data := c.requestHandler(endpoint)
  json.Unmarshal([]byte(data), &seasonalAnimeData)

  // Adding all the animes to another list to get formatted results later
  var animes []Anime
  for _, element := range seasonalAnimeData.Data {
    animes = append(animes, element.Anime)
  }

  // finally generate SeasonalAnime
  seasonalAnime = SeasonalAnime {
    Animes: animes,
    Paging: seasonalAnimeData.Paging,
    Season: seasonalAnimeData.Season,
  }

  return seasonalAnime, nil
}

// get anime suggestions for the user
func (c AnimeClient) GetSuggestedAnime(limit, offset int, fields []string) (SuggestedAnime, error){
  var suggestedAnime SuggestedAnime

  // error handling for limit and offset
  limitsErr := e.LimitsErrHandler(limit, offset)
  if limitsErr != nil {
    return suggestedAnime, limitsErr
  }

  // handle all the errors for the fields
  fields, err := e.FieldsErrHandler(fields)
  if err != nil {
    return suggestedAnime, err
  }

  endpoint, _ := u.UrlGenerator(
    BASE_URL + "/suggestions",
    []string{"limit", "offset", "fields"},
    [][]string{{strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields},
    true,
  )

  // gets data from API and stores it in a struct
  var suggestedAnimeData SuggestedAnimeRaw
  data := c.requestHandler(endpoint)
  json.Unmarshal([]byte(data), &suggestedAnimeData)

  // Adding all the animes to another list to get formatted results later
  var animes []Anime
  for _, element := range suggestedAnimeData.Data {
    animes = append(animes, element.Anime)
  }

  // finally generate RecommendedAnime struct
  suggestedAnime = SuggestedAnime {
    Animes: animes,
    Paging: suggestedAnimeData.Paging,
  }

  return suggestedAnime, nil
}
