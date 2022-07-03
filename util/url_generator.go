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

package util

import (
  "errors"
  "net/url"
)

var URLNameValueError error = errors.New("URLNameValueError: Number of names and values passed to URLGenerator don't match.")

func UrlGenerator(baseUrl string, names []string, values [][]string, isPrimary bool) (string, error) {
  // length of names and values should be same
  if cap(names) != cap(values) {
    return "", URLNameValueError
  }

  var fields string

  for index, name := range(names) {
    var data string
    /* if the data is the first field in URL, 
     * it goes like ?key=value
     * else it is &nextkey=value */
    if isPrimary {
      data = "?" + name + "="
    } else {
      data = "&" + name + "="
    }

    // add values to data variable
    for i, j := range values[index] {
      if i > 0 {
        data = data + "," + url.QueryEscape(j)
      } else {
        data = data + url.QueryEscape(j)
      }
    }

    fields = fields + data

    // from now on all other fields will be secondary
    isPrimary = false
  }

  return baseUrl + fields, nil
}
