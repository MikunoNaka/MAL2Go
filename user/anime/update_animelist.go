/* MAL2Go - MyAnimeList V2 API wrapper for Go
 * Copyright (C) 2022  Vidhu Kant Sharma <vidhukant@protonmail.ch>

 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.

 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>. */

package anime

import (
  e "github.com/MikunoNaka/MAL2Go/v2/errhandlers"
  "fmt"
  "net/url"
  "strconv"
)

// generate the endpoint url with the anime id
func endpointGenerator(id int) string {
  return fmt.Sprintf("%s/anime/%d/my_list_status", BASE_URL, id)
}

// update just an anime's status
func (c Client)SetStatus(id int, status string) (UpdateResponse, error) {
  endpoint := endpointGenerator(id)

  // checks if specified list status is valid
  if !e.IsValidListStatus(status) {
    return UpdateResponse{}, e.InvalidStatusError
  }

  // data to be sent to the server
  params := url.Values{}
  params.Set("status", status)

  // make API request
  return c.putRequestHandler(endpoint, params)
}

// update just an anime's num of episodes watched
func (c Client)SetWatchedEpisodes(id int, episodesWatched int) (UpdateResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("num_watched_episodes", strconv.Itoa(episodesWatched))

  // make API request
  return c.putRequestHandler(endpoint, params)
}

// update just an anime's rewatching status
func (c Client)SetIsRewatching(id int, isRewatching bool) (UpdateResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("is_rewatching", strconv.FormatBool(isRewatching))

  // make API request
  return c.putRequestHandler(endpoint, params)
}

// update just the anime's score
func (c Client)SetScore(id int, score int) (UpdateResponse, error) {
  endpoint := endpointGenerator(id)

  // checks if specified score is valid
  if !e.IsValidScore(score) {
    return UpdateResponse{}, e.InvalidScoreError
  }

  // data to be sent to the server
  params := url.Values{}
  params.Set("score", strconv.Itoa(score))

  // make API request
  return c.putRequestHandler(endpoint, params)
}

// update just an anime's priority
func (c Client)SetPriority(id int, priority int) (UpdateResponse, error) {
  endpoint := endpointGenerator(id)

  // checks if specified priority is valid
  if !e.IsValidPriority(priority) {
    return UpdateResponse{}, e.InvalidPriorityError
  }

  // data to be sent to the server
  params := url.Values{}
  params.Set("priority", strconv.Itoa(priority))

  // make API request
  return c.putRequestHandler(endpoint, params)
}

// update just an anime's rewatch value
func (c Client)SetRewatchValue(id int, rewatchValue int) (UpdateResponse, error) {
  endpoint := endpointGenerator(id)

  // checks if specified rewatch value is valid
  if !e.IsValidRewatchValue(rewatchValue) {
    return UpdateResponse{}, e.InvalidRewatchValueError
  }

  // data to be sent to the server
  params := url.Values{}
  params.Set("rewatch_value", strconv.Itoa(rewatchValue))

  // make API request
  return c.putRequestHandler(endpoint, params)
}

// update just an anime's rewatch count
func (c Client)SetRewatchCount(id int, rewatchCount int) (UpdateResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("num_times_rewatched", strconv.Itoa(rewatchCount))

  // make API request
  return c.putRequestHandler(endpoint, params)
}

// update just an anime's tags
func (c Client)UpdateTags(id int, tags string) (UpdateResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("tags", tags)

  // make API request
  return c.putRequestHandler(endpoint, params)
}

// update just an anime's comments
func (c Client)UpdateComments(id int, comments string) (UpdateResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("comments", comments)

  // make API request
  return c.putRequestHandler(endpoint, params)
}

/* This will overwrite everything
 * i won't use it.. but it's pretty flexible
 * so this will stay here */
// Update/Add an anime to user's anime list
func (c Client)UpdateAnime(id int, data UpdateAnimeData) (UpdateResponse, error) {
  endpoint := endpointGenerator(id)

  // checks if specified list status is valid
  if !e.IsValidListStatus(data.Status) {
    return UpdateResponse{}, e.InvalidStatusError
  }

  // checks if specified score is valid
  if !e.IsValidScore(data.Score) {
    return UpdateResponse{}, e.InvalidScoreError
  }

  // checks if specified priority is valid
  if !e.IsValidPriority(data.Priority) {
    return UpdateResponse{}, e.InvalidPriorityError
  }

  // checks if specified rewatch value is valid
  if !e.IsValidRewatchValue(data.RewatchValue) {
    return UpdateResponse{}, e.InvalidRereadValueError
  }

  params := url.Values{}

  /* NOTE: THIS WILL OVERWRITE ANY DATA THAT 
   * IS NOT SPECIFIED AND SET IT TO NULL */
  params.Set("status",               data.Status)
  params.Set("is_rewatching",        strconv.FormatBool(data.IsRewatching))
  params.Set("score",                strconv.Itoa(data.Score))
  params.Set("num_watched_episodes", strconv.Itoa(data.EpWatched))
  params.Set("priority",             strconv.Itoa(data.Priority))
  params.Set("num_times_rewatched",  strconv.Itoa(data.TimesRewatched))
  params.Set("rewatch_value",        strconv.Itoa(data.RewatchValue))
  params.Set("tags",                 data.Tags)
  params.Set("comments",             data.Comments)

  // make API request
  return c.putRequestHandler(endpoint, params)
}

