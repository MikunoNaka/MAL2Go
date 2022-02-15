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
  "strconv"
  "fmt"
  "errors"
  a "github.com/MikunoNaka/mal2go/anime"
  e "github.com/MikunoNaka/mal2go/errhandlers"
)

const BASE_URL string = "https://api.myanimelist.net/v2"
const maxListLimit int = 1000

// Delete an anime from user's anime list
func (c AnimeListClient)DeleteAnime(id int) string {
  endpoint := fmt.Sprintf("%s/anime/%d/my_list_status", BASE_URL, id)
  /* Returns 200 if anime successfully deleted
   * Alternatively returns 404 if anime not in user's anime list */
  return c.requestHandler(endpoint, "DELETE")
}

// Update/Add an anime to user's anime list
func (c AnimeListClient)UpdateAnime(id int, data UpdateAnimeData) (serverResponse, error) {
  endpoint := fmt.Sprintf("%s/anime/%d/my_list_status", BASE_URL, id)

  // checks if specified list status is valid
  if !e.IsValidListStatus(data.Status) {
    return serverResponse{}, errors.New(fmt.Sprintf("UpdateAnime: Invalid list status: \"%s\"", data.Status))
  }

  // checks if specified score is valid
  if !e.IsValidScore(data.Score) {
    return serverResponse{}, errors.New(fmt.Sprintf("UpdateAnime: Invalid score: %d doesn't lie within 0-10", data.Score))
  }

  // checks if specified priority is valid
  if !e.IsValidPriority(data.Priority) {
    return serverResponse{}, errors.New(fmt.Sprintf("UpdateAnime: Invalid priority: %d doesn't lie within 0-2", data.Priority))
  }

  // checks if specified rewatch value is valid
  if !e.IsValidRewatchValue(data.RewatchValue) {
    return serverResponse{}, errors.New(fmt.Sprintf("UpdateAnime: Invalid rewatch value: %d doesn't lie within 0-5", data.RewatchValue))
  }

  // make API request
  return c.putRequestHandler(endpoint, data), nil
}

// Get authenticated user's anime list
func (c AnimeListClient) GetAnimeList(user, status, sort string, limit, offset int) (a.AnimeList, error){
  var userAnimeList a.AnimeList
  // error handling for limit and offset
  limitsErr := e.LimitsErrHandler(limit, offset, maxListLimit)
  if limitsErr != nil { return userAnimeList, limitsErr
  }

  // checks if valid sort is specified
  if !e.IsValidListSort(sort) {
    return userAnimeList, errors.New(fmt.Sprintf("GetAnimeList: Invalid sort specified: \"%s\"", sort))
  }

  // checks if valid status is specified
  if status != "" && !e.IsValidListStatus(status) {
    return userAnimeList, errors.New(fmt.Sprintf("GetAnimeList: Invalid status specified: \"%s\"", status))
  }

  // get own list if user not specified
  if user == "" {
    user = "@me"
  }

  // if status is "" it returns all anime
  var endpoint string
  if status == "" {
    endpoint = BASE_URL + "/users/" + user + "/animelist?sort=" + sort + "&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset)
  } else {
    endpoint = BASE_URL + "/users/" + user + "/animelist?status=" + status + "&sort=" + sort + "&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset)
  }

  // get data from API
  var animeListData AnimeListRaw
  data := c.requestHandler(endpoint, "GET")
  json.Unmarshal([]byte(data), &animeListData)

  // set MyListStatus for each element and add it to array
  var animes []a.Anime
  for _, element := range animeListData.Data {
    a := element.Anime
    a.MyListStatus = element.ListStatus

    animes = append(animes, a)
  }

  // finally create AnimeList
  userAnimeList = a.AnimeList {
    Animes: animes,
    Paging: animeListData.Paging,
  }

  return userAnimeList, nil
}

