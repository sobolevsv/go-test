package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var msg = []byte(`{"events":[{"action_type":"view","attributes":{"app_version":"2.0.0","namespace":"bx","platform":"site","screen_resolution_x":1920,"screen_resolution_y":1080,"user_agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"},"composer":{"slice":{},"cell":{"index":1,"id":"https://www.ozon.ru/highlight/28470/?advert=zqGXx4HkhwPSmK5TQ5_sJ5h1ftgZs_r92r70qxCuDWJBVOWjx_7autvI_6BOXUFQ_4KX1UB5uVfsL6dccZ8flthoVkPhm4EQ7XJyDfQSIZQ&hs=1","link":"https://www.ozon.ru/highlight/28470/?advert=zqGXx4HkhwPSmK5TQ5_sJ5h1ftgZs_r92r70qxCuDWJBVOWjx_7autvI_6BOXUFQ_4KX1UB5uVfsL6dccZ8flthoVkPhm4EQ7XJyDfQSIZQ&hs=1","type":"banner","advId":"zqGXx4HkhwPSmK5TQ5_sJ5h1ftgZs_r92r70qxCuDWJBVOWjx_7autvI_6BOXUFQ_4KX1UB5uVfsL6dccZ8flthoVkPhm4EQ7XJyDfQSIZQ"}},"number":2,"page":{"category_id":0,"current":"home","current_url":"https://www.ozon.ru/","page_view_id":"2f3d6779-4fd0-48e8-b556-5b0d7b41e58b","previous":"home","previous_page_view_id":"9fc796bb-45f6-4d18-acfb-aea4811db95f","referral_url":""},"timestamp":"2019-10-10 00:12:29 +0300","user":{"ab_group":501,"client_id":0,"full_name":"","google_id":"GA1.2.231942499.1559849202","is_premium":false,"latitude":55.755787,"long_cookie":"cbdb46bccd6117999997b45c2c0fc67f","longitude":37.617634,"region_id":2,"session_id":"hauariikbkdl01ulqqxs1koa","yandex_id":"156258079712708837"},"uuid":"2ea7e610-16d0-42d6-935f-5ed34e9290b1","version":"2","widget":{"type":"rtb","name":"rtb.advBanner","id":231005,"revision_id":174535,"config_revision_id":0,"index":20}},{"action_type":"view","attributes":{"app_version":"2.0.0","namespace":"bx","platform":"site","screen_resolution_x":1920,"screen_resolution_y":1080,"user_agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"},"composer":{"slice":{},"cell":{"index":1,"type":"banner","link":"/highlight/28580/"}},"number":3,"page":{"category_id":0,"current":"home","current_url":"https://www.ozon.ru/","page_view_id":"2f3d6779-4fd0-48e8-b556-5b0d7b41e58b","previous":"home","previous_page_view_id":"9fc796bb-45f6-4d18-acfb-aea4811db95f","referral_url":""},"timestamp":"2019-10-10 00:12:29 +0300","user":{"ab_group":501,"client_id":0,"full_name":"","google_id":"GA1.2.231942499.1559849202","is_premium":false,"latitude":55.755787,"long_cookie":"cbdb46bccd6117999997b45c2c0fc67f","longitude":37.617634,"region_id":2,"session_id":"hauariikbkdl01ulqqxs1koa","yandex_id":"156258079712708837"},"uuid":"ebda1fe9-a85a-4063-b736-8f476d8df42c","version":"2","widget":{"type":"cms","name":"cms.banner","id":242068,"revision_id":191887,"config_revision_id":0,"index":23}},{"action_type":"view","attributes":{"app_version":"2.0.0","namespace":"bx","platform":"site","screen_resolution_x":1920,"screen_resolution_y":1080,"user_agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"},"composer":{"slice":{},"cell":{"index":1,"type":"banner","link":"/highlight/28581/"}},"number":4,"page":{"category_id":0,"current":"home","current_url":"https://www.ozon.ru/","page_view_id":"2f3d6779-4fd0-48e8-b556-5b0d7b41e58b","previous":"home","previous_page_view_id":"9fc796bb-45f6-4d18-acfb-aea4811db95f","referral_url":""},"timestamp":"2019-10-10 00:12:29 +0300","user":{"ab_group":501,"client_id":0,"full_name":"","google_id":"GA1.2.231942499.1559849202","is_premium":false,"latitude":55.755787,"long_cookie":"cbdb46bccd6117999997b45c2c0fc67f","longitude":37.617634,"region_id":2,"session_id":"hauariikbkdl01ulqqxs1koa","yandex_id":"156258079712708837"},"uuid":"77e0aa3d-6fbc-4ce9-8166-84ad8cec2ef0","version":"2","widget":{"type":"cms","name":"cms.banner","id":242065,"revision_id":191876,"config_revision_id":0,"index":26}},{"action_type":"view","attributes":{"app_version":"2.0.0","namespace":"bx","platform":"site","screen_resolution_x":1920,"screen_resolution_y":1080,"user_agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"},"composer":{"slice":{},"cell":{"index":1,"type":"banner","link":"/highlight/28574/"}},"number":5,"page":{"category_id":0,"current":"home","current_url":"https://www.ozon.ru/","page_view_id":"2f3d6779-4fd0-48e8-b556-5b0d7b41e58b","previous":"home","previous_page_view_id":"9fc796bb-45f6-4d18-acfb-aea4811db95f","referral_url":""},"timestamp":"2019-10-10 00:12:29 +0300","user":{"ab_group":501,"client_id":0,"full_name":"","google_id":"GA1.2.231942499.1559849202","is_premium":false,"latitude":55.755787,"long_cookie":"cbdb46bccd6117999997b45c2c0fc67f","longitude":37.617634,"region_id":2,"session_id":"hauariikbkdl01ulqqxs1koa","yandex_id":"156258079712708837"},"uuid":"0e0d8eb1-7e4c-4633-876d-06600469eb7f","version":"2","widget":{"type":"cms","name":"cms.banner","id":242071,"revision_id":191890,"config_revision_id":0,"index":29}}]}`)

func main() {
	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
}

func worker(id int, wg *sync.WaitGroup) {
	var p interface{}
	json.Unmarshal(msg, &p)
	wg.Done()
}

// Foo aaa
func Foo() {
	var wg sync.WaitGroup
	for i := 0; i <= 1000000; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}

func worker1(id int, wg *sync.WaitGroup, ch chan int) {
	// fmt.Printf("worker %d started\n", id)
	defer wg.Done()
	defer fmt.Printf("worker %d done\n", id)
	for {
		a, ok := <-ch
		if !ok {
			// fmt.Printf("worker %d !ok\n", id)
			break
		}

		if a == 0 {
			// fmt.Printf("worker %d a == 0\n", id)
			close(ch)
			break
		}
		var p interface{}
		json.Unmarshal(msg, &p)
	}
}

// Boo dfd
func Boo() {
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go worker1(i, &wg, ch)
	}
	for i := 0; i <= 1000000; i++ {
		ch <- 1
	}
	ch <- 0
	wg.Wait()
}

// BenchmarkFoo sss
func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// perform the operation we're analyzing
	}
}
