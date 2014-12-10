package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Request struct {
	gets map[string]string
}

type Response struct {
	body map[string]interface{}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req := Request{}

		req.gets = make(map[string]string)

		for k, v := range r.URL.Query() {
			req.gets[k] = v[0]
		}

		// get and combine result
		response := Response{}
		response.body = make(map[string]interface{})

		c := make(chan struct {
			name string
			f    interface{}
		})

		var wg sync.WaitGroup

		for name, url := range req.gets {
			wg.Add(1)
			go func(name, url string) {
				defer wg.Done()
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
				} else {
					result, _ := ioutil.ReadAll(resp.Body)
					defer resp.Body.Close()
					var f interface{}

					json.Unmarshal(result, &f)

					r := struct {
						name string
						f    interface{}
					}{
						name,
						f,
					}

					c <- r
				}
			}(name, url)
		}

		go func() {
			wg.Wait()
			close(c)
		}()

	Loop:
		for {
			select {
			case r, more := <-c:
				response.body[r.name] = r.f
				if more == false {
					break Loop
				}
			case <-time.After(time.Second * 1):
				fmt.Println("timeout")
				break Loop
			}
		}

		// write the result
		res, _ := json.Marshal(response.body)
		// fmt.Println(response)
		// fmt.Println(response.body)
		fmt.Fprint(w, string(res[:]))

	})

	fmt.Println("Serving")
	http.ListenAndServe(":8080", nil)
}
