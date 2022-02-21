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
  "fmt"
  "github.com/MikunoNaka/MAL2Go/util"
)

// if fields aren't specified
func FieldsErrHandler(fields []string) ([]string, error) {
  if cap(fields) == 0 {
    // uses all the default fields if none specified
    return util.DefaultFields, nil
  }

  // checks if each given field is valid
  for _, j := range(fields) {
    if !IsValidField(j) {
      return []string{}, errors.New(fmt.Sprintf("InvalidFieldError: Invalid field specified: \"%s\"", j))
    }
  }

  // everything's fine!
  return fields, nil
}

// if limit or error specified are above the limit
func LimitErrHandler(limit, maxLimit int) error {
  if limit > maxLimit {
    return errors.New(fmt.Sprintf("InvalidLimitError: Limit specified too high (%d > %d).", limit, maxLimit))
  } 
  // return nil if no error
  return nil
}
