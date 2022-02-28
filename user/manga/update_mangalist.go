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

package manga

import (
  e "github.com/MikunoNaka/MAL2Go/errhandlers"
  "errors"
  "fmt"
	"net/url"
	"strconv"
)

// generate the endpoint url with the manga id
func endpointGenerator(id int) string {
  return fmt.Sprintf(BASE_URL + "/manga/%d/my_list_status", id)
}

// update just an manga's status
func (c Client)SetStatus(id int, status string) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // checks if specified list status is valid
  if !e.IsValidMangaListStatus(status) {
    return serverResponse{}, errors.New(fmt.Sprintf("SetStatus: Invalid list status: \"%s\"", status))
  }

  // data to be sent to the server
  params := url.Values{}
  params.Set("status", status)

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

// update just a manga's num of volumes read
func (c Client)SetVolumesRead(id int, volumes int) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("num_volumes_read", strconv.Itoa(volumes))

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

// update just a manga's num of chapters read
func (c Client)SetChaptersRead(id int, chapters int) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("num_chapters_read", strconv.Itoa(chapters))

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

// update just a manga's rereading status
func (c Client)SetIsRereading(id int, isRereading bool) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("is_rereading", strconv.FormatBool(isRereading))

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

// update just the manga's score
func (c Client)SetScore(id int, score int) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // checks if specified score is valid
  if !e.IsValidScore(score) {
    return serverResponse{}, errors.New(fmt.Sprintf("SetScore: Invalid score: %d doesn't lie within 0-10", score))
  }

  // data to be sent to the server
  params := url.Values{}
  params.Set("score", strconv.Itoa(score))

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

// update just a manga's priority
func (c Client)SetPriority(id int, priority int) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // checks if specified priority is valid
  if !e.IsValidPriority(priority) {
    return serverResponse{}, errors.New(fmt.Sprintf("SetPriority: Invalid priority: %d doesn't lie within 0-2", priority))
  }

  // data to be sent to the server
  params := url.Values{}
  params.Set("priority", strconv.Itoa(priority))

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

// update just a manga's reread value
func (c Client)SetRereadValue(id int, rereadValue int) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // checks if specified reread value is valid
  if !e.IsValidRewatchValue(rereadValue) {
    return serverResponse{}, errors.New(fmt.Sprintf("SetRereadValue: Invalid rewatch value: %d doesn't lie within 0-5", rereadValue))
  }

  // data to be sent to the server
  params := url.Values{}
  params.Set("reread_value", strconv.Itoa(rereadValue))

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

// update just a manga's reread count
func (c Client)SetRereadCount(id int, rereadCount int) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("num_times_reread", strconv.Itoa(rereadCount))

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

// update just a manga's tags
func (c Client)UpdateTags(id int, tags string) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("tags", tags)

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

// update just a manga's comments
func (c Client)UpdateComments(id int, comments string) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // data to be sent to the server
  params := url.Values{}
  params.Set("comments", comments)

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

/* NOTE: This will overwrite everything
 * i won't use it.. but it's pretty flexible
 * so this will stay here */
// Update/Add a manga to user's manga list
func (c Client)UpdateManga(id int, data UpdateMangaData) (serverResponse, error) {
  endpoint := endpointGenerator(id)

  // checks if specified list status is valid
  if !e.IsValidMangaListStatus(data.Status) {
    return serverResponse{}, errors.New(fmt.Sprintf("UpdateManga: Invalid list status: \"%s\"", data.Status))
  }

  // checks if specified score is valid
  if !e.IsValidScore(data.Score) {
    return serverResponse{}, errors.New(fmt.Sprintf("UpdateManga: Invalid score: %d doesn't lie within 0-10", data.Score))
  }

  // checks if specified priority is valid
  if !e.IsValidPriority(data.Priority) {
    return serverResponse{}, errors.New(fmt.Sprintf("UpdateManga: Invalid priority: %d doesn't lie within 0-2", data.Priority))
  }

  // checks if specified rewatch value is valid
  if !e.IsValidRewatchValue(data.RereadValue) {
    return serverResponse{}, errors.New(fmt.Sprintf("UpdateManga: Invalid reread value: %d doesn't lie within 0-5", data.RereadValue))
  }

  params := url.Values{}

  /* NOTE: THIS WILL OVERWRITE ANY DATA THAT 
   * IS NOT SPECIFIED AND SET IT TO NULL */
  params.Set("status",            data.Status)
  params.Set("is_rereading",      strconv.FormatBool(data.IsRereading))
  params.Set("score",             strconv.Itoa(data.Score))
  params.Set("num_chapters_read", strconv.Itoa(data.ChaptersRead))
  params.Set("priority",          strconv.Itoa(data.Priority))
  params.Set("num_times_reread",  strconv.Itoa(data.TimesReread))
  params.Set("reread_value",      strconv.Itoa(data.RereadValue))
  params.Set("tags",              data.Tags)
  params.Set("comments",          data.Comments)

  // make API request
  return c.putRequestHandler(endpoint, params), nil
}

