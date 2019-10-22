package models

type Shards struct {
    Total      int `json:"total"`
    Successful int `json:"successful"`
    Skipped    int `json:"skipped"`
    Failed     int `json:"failed"`
}

type Total struct { 
    Value    int    `json:"value"`
    Relation string `json:"relation"`
}

type Result struct {
    Index   string  `json:"_index"`
    Type    string  `json:"_type"`
    ID      string  `json:"_id"`
    Score   float64 `json:"_score"`
    Message Message `json:"_source"`
}

type Hits struct {
    Total    Total     `json:"total"`
    MaxScore float64   `json:"max_score"`
    Results  [] Result `json:"hits"`
}


type KibanaQueryResponse struct {
    ScrollId string `json:"_scroll_id"`
    Took     int    `json:"took"`
    TimedOut bool   `json:"timed_out"`
    Shards   Shards `json:"_shards"`
    Hits     Hits   `json:"hits"` 
}