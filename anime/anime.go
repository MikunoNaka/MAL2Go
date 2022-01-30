package anime

import (
	"encoding/json"
	"fmt"
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

func GetAnimeById(token string, animeId int) Anime {
  endpoint := fmt.Sprintf("https://api.myanimelist.net/v2/anime/%d?fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration,rating,pictures,background,related_anime,related_manga,recommendations,studios,statistics", animeId)

  data := requestHandler(token, endpoint)
  var anime Anime
  json.Unmarshal([]byte(data), &anime)

  return anime
}

// Checks if given rankingType is valid
func isValidRankingType(rankingType string) bool {
    switch rankingType {
    case
        "all",
        "airing",
        "upcoming",
        "tv",
        "ova",
        "movie",
        "special",
        "bypopularity",
        "favorite":
        return true
    }
    return false
}

func GetAnimeRanking(token string, rankingType string) {
  if !isValidRankingType(rankingType) {
    log.Fatal(fmt.Sprintf("GetAnimeRanking: Invalid Ranking Type Given (\"%s\")", rankingType))
  }
}
