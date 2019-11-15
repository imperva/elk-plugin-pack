package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"lsar/cmd/log-sender/message_builder"
	"lsar/package/models"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func get_message_json(message models.Message) string {

	json_message, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling JSON data:", err)
		return ""
	}

	return string(json_message)
}

func inet_worker(message_chan <-chan models.Message, host string, port int) {

	defer wg.Done()

	dest := host + ":" + strconv.Itoa(port)

	fmt.Printf("Connecting to %s\n", dest)

	conn, err := net.Dial("tcp", dest)

	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Problem connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
		return
	}

	defer conn.Close()

	//         conn.SetWriteDeadline( time.Now().Add( 1 * time.Second ) )
	var counter int = 0
	for {
		message, ok := <-message_chan
		if !ok {
			break
		}

		//            fmt.Printf( "[%d] inet_worker received id=%s\n", counter, message.Event.EventId )

		_, err := conn.Write([]byte("<14>" + get_message_json(message) + "\n"))
		//count, err := conn.Write( []byte( "<14>" + get_message_json( message ) + "\n" ) )
		if err != nil {
			fmt.Println("Error writing to stream.")
		} else {
			// fmt.Printf( "Wrote %d bytes to stream.\n", count )
		}

		counter++
	}

	fmt.Printf("inet_worker sent count=%d\n", counter)
}

func file_worker(message_chan <-chan models.Message, message_file string) {

	defer wg.Done()

	file, err := os.Create(message_file)
	if err != nil {
		fmt.Println("Error opening file ", message_file)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	var counter int = 1
	for {
		message, ok := <-message_chan
		if !ok {
			return
		}

		fmt.Printf("[%d] file_worker received id=%s\n", counter, message.Event.EventID)

		count, err := writer.WriteString("<14>" + get_message_json(message) + "\n")
		if err == nil {
			fmt.Printf("wrote %d bytes\n", count)
			writer.Flush()
		}
		counter++
	}
}

var wg sync.WaitGroup

func main() {

	target := strings.Split(os.Args[1], ":")

	fluentd_host := target[0]
	fluentd_port := 5514

	if len(target) > 1 {
		port, err := strconv.Atoi(target[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid target <host:port> %s specified", os.Args[1])
			os.Exit(22)
		}
		fluentd_port = port
	}

	message_count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid message count: %s", os.Args[2])
		os.Exit(22)
	}

	thread_count := 5

	if len(os.Args) > 3 {
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid thread count: %s", os.Args[3])
			os.Exit(22)
		}
		thread_count = count
	}

	fmt.Printf("Send %d messages to fluentd at %s:%d\n", message_count, fluentd_host, fluentd_port)

	start := time.Now()

	message_chan := make(chan models.Message, thread_count+10)

	wg.Add(1)
	go message_builder.Build(message_chan, message_count, &wg)

	wg.Add(thread_count)
	for count := 0; count < thread_count; count++ {
		go inet_worker(message_chan, fluentd_host, fluentd_port)
	}

	//go file_worker( message_chan, "m1.syslog" )
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("Message generator took %s", elapsed)
}
