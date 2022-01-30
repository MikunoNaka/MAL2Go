package anime

// contains previous/next page for anime list
type ListPaging struct {
  NextPage string `json:"next"`
  PrevPage string `json:"previous"` // might need checking
}

// this is how the API returns data (looks horrible)
type RawRanking struct {
  Data []struct {
    Anime Anime `json:"node"`
    Ranking struct {
      Rank int `json:"rank"`
    } `json:"ranking"`
  } `json:"data"`

  Paging ListPaging `json:"paging"`
}

// each anime has a ranking number
type AnimeRankingTitle struct {
  Anime   Anime
  RankNum int
}

// this is how mal2go returns data
type AnimeRanking struct {
  Titles []AnimeRankingTitle
  Paging ListPaging
}
