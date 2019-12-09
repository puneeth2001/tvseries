package main

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
	"strconv"
	"encoding/json"
	"packs"
)
type Structure struct {
	ID       string `json:"_id"`
	AirDate  string `json:"air_date"`
	Episodes []struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int     `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int     `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float64 `json:"vote_average"`
		VoteCount      int     `json:"vote_count"`
		Crew           []struct {
			ID          int    `json:"id"`
			CreditID    string `json:"credit_id"`
			Name        string `json:"name"`
			Department  string `json:"department"`
			Job         string `json:"job"`
			Gender      int    `json:"gender"`
			ProfilePath string `json:"profile_path"`
		} `json:"crew"`
		GuestStars []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			CreditID    string `json:"credit_id"`
			Character   string `json:"character"`
			Order       int    `json:"order"`
			Gender      int    `json:"gender"`
			ProfilePath string `json:"profile_path"`
		} `json:"guest_stars"`
	} `json:"episodes"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	ID2           int    `json:"id"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
}
func FetchSeasonData(act string) []byte{

	url := act

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	
	return (body)

}
func ExampleURL(season int64,id int64) string{
	b16 := strconv.FormatInt(season,10)
	fmt.Println(string(b16))
	a16 := strconv.FormatInt(id,10)
	fmt.Println(string(a16))
	u, err := url.Parse("https://api.themoviedb.org/3/tv/"+string(a16)+"/season/"+string(b16)+"?language=jaffa&api_key=cheppa")
	if err != nil {
		log.Fatal(err)
	}
	q:= u.Query()
	q.Set("api_key","85024bf9f2db24e284e8959926cd3226")
	q.Set("language","en-US")
	
	u.RawQuery = q.Encode()
	stringURL := u.String()
	fmt.Println(stringURL)
	
	return stringURL
	// Output: https://google.com/search?q=golang
}
func main(){
	var j int64
	for j= 0;j<8;j++{
		JSONdata := &Structure{}
		
		var example = ExampleURL(j,packs.GetID("game of thrones"))
		var byteData = FetchSeasonData(example)
		err := json.Unmarshal(byteData, JSONdata)
		if err != nil{
			fmt.Println(err)
		}
		for i:= range JSONdata.Episodes{
			fmt.Printf("%s:%d,%s:%d","season",j,"episode",i)
			// fmt.Print(j)
			// fmt.Print("episode:")
			// fmt.Print(i)
			fmt.Println()
			fmt.Println(JSONdata.Episodes[i].Overview)
		}
	} 
		

	
	
}