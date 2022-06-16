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
  "encoding/json"
  "strconv"
  e "github.com/MikunoNaka/MAL2Go/errhandlers"
  u "github.com/MikunoNaka/MAL2Go/util"
)

const BASE_URL string = "https://api.myanimelist.net/v2/manga"

// in MAL documentation this is named Get Manga List
func (c Client) SearchManga(searchString string, limit, offset int, fields []string) ([]Manga, error) {
  var searchResults []Manga

  // error handling for limit
  limitErr := e.LimitErrHandler(limit, 100)
  if limitErr != nil {
    return searchResults, limitErr
  }

  // handle all the errors for the fields
  fields, err := e.MangaFieldsErrHandler(fields)
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
  var mangaSearchData MangaSearchRaw
  data, err := c.requestHandler(endpoint)
  if err != nil {
    return searchResults, err
  }
  json.Unmarshal([]byte(data), &mangaSearchData)

  for _, element := range mangaSearchData.Data {
    searchResults = append(searchResults, element.Manga)
  } 

  return searchResults, nil
}
 
// Each manga has its own ID on MAL
func (c Client) GetMangaById(mangaId int, fields []string) (Manga, error) {
  var manga Manga

  // handle all the errors for the fields
  fields, err := e.MangaFieldsErrHandler(fields)
  if err != nil {
    return manga, err
  }

  endpoint, _ := u.UrlGenerator(
    BASE_URL + "/" + strconv.Itoa(mangaId),
    []string{"fields"},
    [][]string{fields},
    true,
  )

  data, err := c.requestHandler(endpoint)
  if err != nil {
    return manga, err
  }
  json.Unmarshal([]byte(data), &manga)

  return manga, nil
}

// Ranking is a list of manga sorted by their rank
func (c Client) GetMangaRanking(rankingType string, limit, offset int, fields []string) ([]rManga, error) {
  var mangaRanking []rManga

  // error handling for limit
  limitErr := e.LimitErrHandler(limit, 500)
  if limitErr != nil {
    return mangaRanking, limitErr
  }

  // handle all the errors for the fields
  fields, err := e.MangaFieldsErrHandler(fields)
  if err != nil {
    return mangaRanking, err
  }

  // if ranking type is invalid
  if !e.IsValidMangaRankingType(rankingType) {
    return mangaRanking, e.InvalidRankingError
  }

  endpoint, _ := u.UrlGenerator(
    BASE_URL + "/ranking",
    []string{"ranking_type", "limit", "offset", "fields"},
    [][]string{{rankingType}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields},
    true,
  )

  // gets data from API and stores it in a struct
  var rankingData RawRanking
  data, err := c.requestHandler(endpoint)
  if err != nil {
    return mangaRanking, err
  }
  json.Unmarshal([]byte(data), &rankingData)

  // Adding all the mangas in ranking list to a slice
  for _, manga := range rankingData.Data {
    // set RankNum for manga
    m := manga.Manga
    m.RankNum = manga.Ranking.Rank

    // add manga to list
    mangaRanking = append(mangaRanking, m)
  }

  return mangaRanking, nil
}
