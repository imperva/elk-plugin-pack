package message_builder

import (
    "math/rand"
	"strconv"
	"sync"
    "time"
    "fmt"
    "lsar/package/models"

    "github.com/brianvoe/gofakeit"
)

func Build( message_chan chan<- models.Message, message_count int, wg *sync.WaitGroup ) {

    defer wg.Done()

    fmt.Printf( "message_builder message_count=%d\n", message_count )

    rand_source := rand.NewSource( time.Now().UnixNano() )
    random := rand.New( rand_source )

    for count := 0 ; count < message_count ; count++ {

        message := generate_message( count, random.Intn( 1000 ) )
        message_chan <- message
    }

    fmt.Printf( "message_builder sent count=%d\n", message_count )
    close( message_chan )
}

func generate_message( which int, random_int int ) models.Message {

	//    fmt.Printf( "random=%d\n", random_int )
	now := time.Now()

	header := models.Header{ "Imperva Inc.",
						"SecureSphere",
						"Audit",
						"Database",
						"1.6",
						"June 6, 2019",
					}

	event_id := now.Format( "20060102150405" ) + fmt.Sprintf( "%05d", which )

	event_type := "Query"

	user_group         := "Default MsSql group"

	source_application := ""

	os_user_chain  := ""
	raw_query      := "\"CREATE TABLE Hotdog\""
	parsed_query   := "create table hotdog" // or "N/A (login)"
	bind_variables := ""
	service_type   := "MsSql"

	mxip               := gofakeit.IPv4Address() 
	gwip               := gofakeit.IPv4Address() 
	agent_name         := ""
	db_schema_pair     := "(master,)"
	db_schema_name     := ""
	operation_name     := "select" // or "Login"
	database_name      := "master"
	table_group_name   := ""

	object_type        := "table"
	objects_list       := "hotdog_s"

	query_group        := "(hotdog,select)(hotdog,delete)(hotdog,insert)(hotdog,create)(hotdog,drop)"


	event := models.Event{ 
					gofakeit.Username(),
					event_id,
					event_type,
					user_group,
					strconv.FormatBool( gofakeit.Bool() ),
					gofakeit.Username(),
					source_application,
					gofakeit.Username(),
					os_user_chain,
					raw_query,
					parsed_query,
					bind_variables,
					gofakeit.Sentence(30)
					gofakeit.Number(1, 99999),
					service_type,
					mxip,
					gwip,
					agent_name,
					db_schema_pair,
					db_schema_name,
					strconv.FormatBool( gofakeit.Bool() ),
					operation_name,
					database_name,
					table_group_name,
					strconv.FormatBool( gofakeit.Bool() ),
					strconv.FormatBool( gofakeit.Bool() ),
					object_type,
					objects_list,
					gofakeit.Number(1, 99999),
					gofakeit.Number(1, 2000),
					strconv.FormatBool( gofakeit.Bool() ),
					gofakeit.Sentence(30),
					query_group,
				}

	message := models.Message{ header,
						now.Format( "2006-01-02 15:04:05 +0000" ),
						gofakeit.WeekDay(),
						fmt.Sprintf("%d", gofakeit.Hour()),
						"13.5.0.10_0",
						mxip,
						gwip,
						gofakeit.IPv4Address(),
						strconv.Itoa(gofakeit.Number(1000, 9999)),
						gofakeit.DomainName(),
						gofakeit.IPv4Address(),
						strconv.Itoa(gofakeit.Number(1000, 9999)),
						"TCP",
						"Database connections",
						"Windows",
						"MsSQL DB Service (Demo)",
						"Default MsSql Application",
						event,
						now,
					}

	return message
}