package anime

import (
	"io/ioutil"
	"log"
	"net/http"
  "errors"
)

// Handles HTTP request with your OAuth token as a Header
// TODO: Verify that this function is safe to use
func requestHandler(token string, endpoint string) string {
  client := &http.Client{}

  // generate request
  req, err := http.NewRequest("GET", endpoint, nil)
  if err != nil {
      log.Fatal(err)
  }
  req.Header.Add("Authorization", token)

  // do request
  res, err := client.Do(req)
  if err != nil {
      log.Fatal(err)
  }
  defer res.Body.Close()

  // read body
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
      log.Fatal(err)
  }

  return string(body)
}

func urlGenerator(baseUrl string, names []string, values [][]string, isPrimary bool) (string, error) {
  // length of names and values should be same
  if cap(names) != cap(values) {
    return "", errors.New("urlGenerator: Error: Length of names and values don't match.")
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
        data = data + "," + j
      } else {
        data = data + j 
      }
    }

    fields = fields + data

    // from now on all other fields will be secondary
    isPrimary = false
  }

  return baseUrl + fields, nil
}
