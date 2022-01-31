package anime

import (
	"io/ioutil"
	"log"
	"net/http"
)

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
