package goroutines

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const apikey = "xx01a62b0df8e6a51830dc261cbxxxxx"

// code without using go routine, will add go routine in next commit
// Sample output :
// This is the data for: Delhi {Main:{Temp:303.2}}
// This is the data for: Bengaluru {Main:{Temp:297.38}}
// This is the data for: Mumbai {Main:{Temp:298.14}}
// This is the data for: Patna {Main:{Temp:301.11}}
// This operation took :  559.819125ms

//func fetchWeatherData(city string) interface{} {
//
//	var data struct {
//		Main struct {
//			Temp float64 `json:"temp"`
//		} `json:"main"`
//	}
//
//	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apikey)
//	res, err := http.Get(url)
//	if err != nil {
//		fmt.Printf("Error fetching data for city %s: %s ", city, err)
//		return data
//	}
//	defer res.Body.Close()
//	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
//		fmt.Printf("Error decoding data for city %s: %s ", city, err)
//		return data
//	}
//
//	return data
//}

// Demo run with go routines
//This is the city Patna
//This is the city Delhi
//This is the city Banglore
//This is the city Mumbai
//This operation took :  157.35ms

func fetchWeatherData(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apikey)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data for city %s: %s ", city, err)
		return data
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding data for city %s: %s ", city, err)
		return data
	}
	ch <- fmt.Sprintf("This is the city %s", city)

	return data
}

func DemoGoRoutine() {

	startNow := time.Now()
	cities := []string{"Delhi", "Banglore", "Mumbai", "Patna"}
	ch := make(chan string)
	var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		go fetchWeatherData(city, ch, &wg)

	}
	// asking the goroutine to wait before quitting
	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
	fmt.Println("This operation took : ", time.Since(startNow))

	for i := 1; i <= 5; i++ {
		wg.Add(1) //increment the wait group counter
		go tryGoRoutine(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines completed !! ")
}

// The order of execution in goroutines is non-deterministic
func tryGoRoutine(i int, wg *sync.WaitGroup) {
	defer wg.Done() // mark the goroutine as done
	fmt.Printf("Message from %d goroutine\n", i)
}

func EvilNinjaGoRoutine() {
	start := time.Now()
	var wg sync.WaitGroup
	evilNinjas := []string{"Red Women", "Cersei", "Margery", "Joffery"}

	for _, evilNinja := range evilNinjas {
		wg.Add(1)
		go attack(evilNinja, &wg)
	}

	wg.Wait()
	fmt.Println("Time taken in attack ", time.Since(start))

}

func attack(target string, wg *sync.WaitGroup) {
	fmt.Println("Throwing ninja star at evil", target)
	wg.Done()
	//time.Sleep(time.Second)
}
