package core

import (
    "encoding/json"
    "fmt"
    "log"
	"net/http"
	
	"lsar/package/models"
)

func QueryKibanaByIndex( kibana_host string, kibana_port int, kibana_index string ) int {

    base_url := fmt.Sprintf( "http://%s:%d/", kibana_host, kibana_port )

    client := &http.Client{}

    log.Printf( "Query Kibana for index %s, host=%s port=%d\n", kibana_index, kibana_host, kibana_port )

    var url = base_url + fmt.Sprintf( "%s/_search/?scroll=1m", kibana_index )

    var total_count = 0
    var scroll_id = ""

    for {

        log.Printf( "Query Kibana GET URL=%s\n", url )

        request, err := http.NewRequest( "GET", url, nil )
        if err != nil {
            log.Fatal( "Failed to create HTTP request: ", err )
            return 0
        }

        response, err := client.Do( request )
        if err != nil {
            log.Fatal( "Failed to send HTTP request: ", err )
            return 0
        }

        defer response.Body.Close()

        id, return_count := processResponse( response )
        if return_count == 0 { 
            break
        }
        
        scroll_id = id

        url = base_url + fmt.Sprintf( "_search/scroll?scroll=1m&scroll_id=%s", scroll_id )
        total_count = total_count + return_count
    }

    if len( scroll_id ) > 0 {
        url = base_url + fmt.Sprintf( "_search/scroll/%s", scroll_id )

        log.Printf( "Query Kibana DELETE URL=%s\n", url )
        request, err := http.NewRequest( "DELETE", url, nil )
        if err == nil {
            _, err := client.Do( request )
            if err != nil {
                log.Printf( "Failed to send HTTP DELETE: %v", err )
            }
        }
    }

    return total_count
}

func processResponse( response *http.Response ) (string, int) {

    var query_response models.KibanaQueryResponse

    if err := json.NewDecoder( response.Body ).Decode( &query_response ) ; err != nil {
        log.Println( err )
        return "", 0
    }

    total_hits := query_response.Hits.Total.Value
    scroll_id  := query_response.ScrollId

    log.Printf( "Response total_hits=%d timed_out=%t\n", total_hits, query_response.TimedOut )
    log.Printf( "ScrollId=[%s]\n", scroll_id )

    return_count := len( query_response.Hits.Results )
    log.Printf( "Return count %d\n", return_count )

    for index, result := range query_response.Hits.Results {

        message := result.Message
		event   := message.Event
		
		LogRequest(message.Timestamp)
        log.Printf( "[%d] _id=%s, event-id=%s, timestamp=%s\n", index,  result.ID, event.EventID, message.Timestamp )
    }
    
    return scroll_id, return_count
}
