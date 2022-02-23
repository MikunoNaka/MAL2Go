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
	// "errors"
	// "fmt"
	"strconv"
  e "github.com/MikunoNaka/MAL2Go/errhandlers"
  u "github.com/MikunoNaka/MAL2Go/util"
)

const BASE_URL string = "https://api.myanimelist.net/v2/manga"

// MAL Might change this
const maxMangaLimit int = 100

// in MAL documentation this is named Get Manga List
func (c Client) SearchManga(searchString string, limit, offset int, fields []string) (MangaSearch, error) {
  var searchResults MangaSearch

  // error handling for limit
  limitErr := e.LimitErrHandler(limit, maxMangaLimit)
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
  fields, err := e.FieldsErrHandler(fields)
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

// // Ranking is a list of anime sorted by their rank
// func (c Client) GetAnimeRanking(rankingType string, limit, offset int, fields []string) (AnimeRanking, error) {
//   var animeRanking AnimeRanking
// 
//   // error handling for limit
//   limitErr := e.LimitErrHandler(limit, maxAnimeLimit)
//   if limitErr != nil {
//     return animeRanking, limitErr
//   }
// 
//   // handle all the errors for the fields
//   fields, err := e.FieldsErrHandler(fields)
//   if err != nil {
//     return animeRanking, err
//   }
// 
//   // if ranking type is invalid
//   if !e.IsValidRankingType(rankingType) {
//     return animeRanking, errors.New(fmt.Sprintf("GetAnimeRanking: Invalid ranking type specified: \"%s\"", rankingType))
//   }
// 
//   endpoint, _ := u.UrlGenerator(
//     BASE_URL + "/ranking",
//     []string{"ranking_type", "limit", "offset", "fields"},
//     [][]string{{rankingType}, {strconv.Itoa(limit)}, {strconv.Itoa(offset)}, fields},
//     true,
//   )
// 
//   // gets data from API and stores it in a struct
//   var rankingData RawRanking
//   data := c.requestHandler(endpoint)
//   json.Unmarshal([]byte(data), &rankingData)
// 
//   // Adding all the animes in ranking list to a slice
//   var animeRankingTitles []AnimeRankingTitle
//   for _, element := range rankingData.Data {
//     animeRankingTitles = append(
//       animeRankingTitles,
//       AnimeRankingTitle {
//         Anime:   element.Anime, 
//         RankNum: element.Ranking.Rank,
//       },
//     )
//   }
// 
//   // Finally, create the AnimeRanking object
//   animeRanking = AnimeRanking {
//     Titles: animeRankingTitles,
//     Paging: ListPaging {
//       NextPage: rankingData.Paging.NextPage,
//       PrevPage: rankingData.Paging.PrevPage,
//     },
//   }
// 
//   return animeRanking, nil
// }
