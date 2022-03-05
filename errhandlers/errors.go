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

package errhandlers

import (
  "errors"
)

var InvalidFieldError        error = errors.New("InvalidFieldError: Invalid field specified.")
var InvalidLimitError        error = errors.New("InvalidLimitError: Limit specified too high.")
var InvalidLimitError500     error = errors.New("InvalidLimitError: Limit specified too high. (max is 500)")
var InvalidLimitError1000    error = errors.New("InvalidLimitError: Limit specified too high. (max is 1000)")

var InvalidRankingError      error = errors.New("InvalidRankingError: Invalid ranking type specified.")
var InvalidSeasonError       error = errors.New("InvalidSeasonError: Invalid season specifield.")
var InvalidSortError         error = errors.New("InvalidSortError: Invalid sort type specifield.")
var InvalidStatusError       error = errors.New("InvalidStatusError: Invalid status specified.")

var URLNameValueError        error = errors.New("URLNameValueError: Number of names and values passed to URLGenerator don't match.")

var InvalidScoreError        error = errors.New("InvalidScoreError: Score should lie between 0-10.")
var InvalidPriorityError     error = errors.New("InvalidPriorityError: Priority should lie between 0-2.")
var InvalidRewatchValueError error = errors.New("InvalidRewatchValueError: Rewatch value should lie between 0-5.")
var InvalidRereadValueError  error = errors.New("InvalidRereadValueError: Reread value should lie between 0-5.")
