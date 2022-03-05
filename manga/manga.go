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
	"fmt"
	"strconv"
  e "github.com/MikunoNaka/MAL2Go/errhandlers"
  u "github.com/MikunoNaka/MAL2Go/util"
)

const BASE_URL string = "https://api.myanimelist.net/v2/manga"

// MAL Might change this
const maxMangaLimit int = 500

// in MAL documentation this is named Get Manga List
func (c Client) SearchManga(searchString string, limit, offset int, fields []string) (MangaSearch, error) {
  var searchResults MangaSearch

  // error handling for limit
  limitErr := e.LimitErrHandler(limit, maxMangaLimit)
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
  data := c.requestHandler(endpoint)
  json.Unmarshal([]byte(data), &mangaSearchData)

  // Adding all the mangas to another list to get formatted results later
  var mangas []Manga
  for _, element := range mangaSearchData.Data { 
    mangas = append(mangas, element.Manga) 
  } 

  // finally generate AnimeList 
  searchResults = MangaSearch {
    Mangas: mangas,
    Paging: mangaSearchData.Paging,
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

  data := c.requestHandler(endpoint)
  json.Unmarshal([]byte(data), &manga)

  return manga, nil
}

// Ranking is a list of manga sorted by their rank
func (c Client) GetMangaRanking(rankingType string, limit, offset int, fields []string) (MangaRanking, error) {
  var mangaRanking MangaRanking

  // error handling for limit
  limitErr := e.LimitErrHandler(limit, maxMangaLimit)
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
  data := c.requestHandler(endpoint)
  json.Unmarshal([]byte(data), &rankingData)

  // Adding all the mangas in ranking list to a slice
  var mangas []rManga

  for _, manga := range rankingData.Data {
    // set RankNum for manga
    newManga := manga.Manga
    newManga.RankNum = manga.Ranking.Rank

    // add newManga to list
    mangas = append(mangas, newManga)
  }

  // Finally, create the AnimeRanking object
  mangaRanking = MangaRanking {
    Mangas: mangas,
    Paging: ListPaging {
      NextPage: rankingData.Paging.NextPage,
      PrevPage: rankingData.Paging.PrevPage,
    },
  }

  return mangaRanking, nil
}
