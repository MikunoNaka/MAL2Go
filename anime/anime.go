package anime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
  "errors"
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

// Each anime has its own ID on MAL
func GetAnimeById(token string, animeId int) Anime {
  endpoint := fmt.Sprintf("https://api.myanimelist.net/v2/anime/%d?fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration,rating,pictures,background,related_anime,related_manga,recommendations,studios,statistics", animeId)

  var anime Anime
  data := requestHandler(token, endpoint)
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

// Ranking is a list of anime sorted by their rank
func GetAnimeRanking(token string, rankingType string) (AnimeRanking, error) {
  var animeRanking AnimeRanking
  if !isValidRankingType(rankingType) {
    return animeRanking, errors.New(fmt.Sprintf("GetAnimeRanking: Invalid Ranking Type Given (\"%s\")", rankingType))
  }

  endpoint := "https://api.myanimelist.net/v2/anime/ranking?ranking_type=all&limit=4"

  // gets data from API and stores it in a struct
  var rankingData RawRanking
  data := requestHandler(token, endpoint)
  json.Unmarshal([]byte(data), &rankingData)

  // Adding all the animes in ranking list to a slice
  var animeRankingTitles []AnimeRankingTitle
  for _, element := range rankingData.Data {
    animeRankingTitles = append(
      animeRankingTitles,
      AnimeRankingTitle {
        Anime:   element.Anime, 
        RankNum: element.Ranking.Rank,
      },
    )
  }

  // Finally, create the AnimeRanking object
  animeRanking = AnimeRanking {
    Titles: animeRankingTitles,
    Paging: ListPaging {
      NextPage: rankingData.Paging.NextPage,
      PrevPage: rankingData.Paging.PrevPage,
    },
  }

  return animeRanking, nil
}
