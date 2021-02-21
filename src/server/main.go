package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type counters struct {
	sync.Mutex
	view  int
	click int
}

type content_time struct {
	sync.Mutex 
	content string 
	time time.Time
	counter_at_time counters 
}
var (
	c = counters{}
	ct = []content_time{}
	content = []string{"sports", "entertainment", "business", "education"}
	prev_time = time.Time 
	rate_limit = 5
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to EQ Works 😎")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	c.Lock()
	c.view++
	c.Unlock()

	err := processRequest(r) // ?? 
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}
	data := content[rand.Intn(len(content))] 

	// simulate random click call
	if rand.Intn(100) < 50 {
		processClick(data)
	}


	
	store_counter()

	fmt.Fprint(w, ct)
}
func store_counter() {
	data := content[rand.Intn(len(content))] 
	
	now_time := time.Now()

	var contnet_time_obj = new (content_time)
	contnet_time_obj.content = data 
	contnet_time_obj.time  = now_time
	contnet_time_obj.counter_at_time = c
	fmt.Println("added counter")
	ct = append(ct, *contnet_time_obj)
}
func processRequest(r *http.Request) error {

	time.Sleep(time.Duration(rand.Int31n(50)) * time.Millisecond)
	return nil
}

func processClick(data string) error {
	c.Lock()
	c.click++
	c.Unlock()

	return nil
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	// for rate limiting
	if !isAllowed() {
		w.WriteHeader(429)
		return
	}
}

/* 
	rate limiter for stats handler
*/
func isAllowed() bool {
	return true
}

func uploadCounters() error {
	return nil
}

func main() {
	
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/stats/", statsHandler)
	fmt.Println("Server is running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

    go store_counter()
}
