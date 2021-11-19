package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Responce struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}
type Stats struct {
	Views string `json:"viewCount"`
	Subs  string `json:"subscriberCount"`
}

func GetSubscriber() (Item, error) {

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	q := req.URL.Query()
	q.Add("key", "AIzaSyBs4R4G9FfubRW2SCVNYCnc6GRc34vMNJw")
	// a := os.Getenv("YOUTUBE_KEY")
	// fmt.Println(a)
	q.Add("id", "UCaHjw5yX1_IKW0F7Ygsl4lQ")
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()
	// fmt.Println(req.URL.RawQuery)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Responce Status", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	var responce Responce
	err = json.Unmarshal(body, &responce)
	if err != nil {
		return Item{}, err
	}
	return responce.Items[0], nil
}
