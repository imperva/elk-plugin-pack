package main

import (
    "strings"
    "strconv"
    "time"
    "log"
    "os"

    "lsar/cmd/kibana-querier/core"
)

func main() {

    target := strings.Split( os.Args[ 1 ], ":" )

    kibana_host := target[ 0 ]
    kibana_port := 9200

    if len( target ) > 1 {
        port, err := strconv.Atoi( target[ 1 ] )
        if err != nil {
            log.Printf("%s | Invalid target <host:port> %s specified", os.Stderr, os.Args[ 1 ] )
            os.Exit( 22 )
        }
        kibana_port = port
    }

    now := time.Now()
    kibana_index := "lsar-" + now.Format( "2006.01.02" )

    if len( os.Args ) > 2 {
        kibana_index = os.Args[ 2 ]
    }

    core.InitStats()

    total_count := core.QueryKibanaByIndex( kibana_host, kibana_port, kibana_index )

    log.Printf( "Kibana query returned total_count=%d\n", total_count )
    core.ProcessStats()
}
