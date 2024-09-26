package goroutines

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const apikey = "4701a62b0df8e6a51830dc261cb2eaf5"

// code without using go routine, will add go routine in next commit
// Sample output :
// This is the data for: Delhi {Main:{Temp:303.2}}
// This is the data for: Bengaluru {Main:{Temp:297.38}}
// This is the data for: Mumbai {Main:{Temp:298.14}}
// This is the data for: Patna {Main:{Temp:301.11}}
// This operation took :  559.819125ms

func fetchWeatherData(city string) interface{} {

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

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

	return data
}

func main() {
	startNow := time.Now()
	cities := []string{"Delhi", "Banglore", "Mumbai", "Patna"}
	for _, city := range cities {
		data := fetchWeatherData(city)
		fmt.Println("This is the data ", data)

	}
	fmt.Println("This operation took : ", time.Since(startNow))
}
