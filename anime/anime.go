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

package anime

import (
	"encoding/json"
	"fmt"
	"strconv"
  e "github.com/MikunoNaka/MAL2Go/errhandlers"
  u "github.com/MikunoNaka/MAL2Go/util"
)

const BASE_URL string = "https://api.myanimelist.net/v2/anime"

// MAL Might change this
const maxAnimeLimit int = 500

// in MAL documentation this is named Get Anime List
func (c Client) SearchAnime(searchString string, limit, offset int, fields []string) ([]Anime, error) {
  var searchResults []Anime

  // error handling for limit
  limitErr := e.LimitErrHandler(limit, maxAnimeLimit)
  if limitErr != nil {
    return searchResults, limitErr
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
  for _, element := range animeSearchData.Data {
    searchResults = append(searchResults, element.Anime)
  } 

  return searchResults, nil
}

// Each anime has its own ID on MAL
func (c Client) GetAnimeById(animeId int, fields []string) (Anime, error) {
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
     * this might be an issue with MAL itself */
    [][]string{fields},
    true,
  )

  data := c.requestHandler(endpoint)
  json.Unmarshal([]byte(data), &anime)

  return anime, nil
}

// Ranking is a list of anime sorted by their rank
func (c Client) GetAnimeRanking(rankingType string, limit, offset int, fields []string) ([]rAnime, error) {
  var animeRanking []rAnime

  // error handling for limit
  limitErr := e.LimitErrHandler(limit, maxAnimeLimit)
  if limitErr != nil {
    return animeRanking, limitErr
  }

  // handle all the errors for the fields
  fields, err := e.FieldsErrHandler(fields)
  if err != nil {
    return animeRanking, err
  }

  // if ranking type is invalid
  if !e.IsValidRankingType(rankingType) {
    return animeRanking, e.InvalidRankingError
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
  for _, anime := range rankingData.Data {
    // set RankNum for anime
    a := anime.Anime
    a.RankNum = anime.Ranking.Rank
    // add anime to slice
    animeRanking = append(animeRanking, a)
  }

  return animeRanking, nil
}

// get list of animes from specified season
func (c Client) GetSeasonalAnime(year, season, sort string, limit, offset int, fields []string) (SeasonalAnime, error) {
  var seasonalAnime SeasonalAnime

  // error handling for limit
  limitErr := e.LimitErrHandler(limit, maxAnimeLimit)
  if limitErr != nil {
    return seasonalAnime, limitErr
  }

  // handle all the errors for the fields
  fields, err := e.FieldsErrHandler(fields)
  if err != nil {
    return seasonalAnime, err
  }

  // checks if valid season is specified
  if !e.IsValidSeason(season) {
    return seasonalAnime, e.InvalidSeasonError
  }

  // checks if valid sort is specified
  if !e.IsValidSeasonalSort(sort) {
    return seasonalAnime, e.InvalidSortError
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
    Season: seasonalAnimeData.Season,
  }

  return seasonalAnime, nil
}

// get anime suggestions for the user
func (c Client) GetSuggestedAnime(limit, offset int, fields []string) ([]Anime, error){
  var suggestedAnime []Anime

  // error handling for limit
  // limit for this is 100 unlike others in the current package
  limitErr := e.LimitErrHandler(limit, 100)
  if limitErr != nil {
    return suggestedAnime, limitErr
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
  for _, element := range suggestedAnimeData.Data {
    suggestedAnime = append(suggestedAnime, element.Anime)
  }

  return suggestedAnime, nil
}
