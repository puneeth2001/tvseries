package main

import ("io/ioutil"
		"net/http"
		"fmt"
	)
func FetchSeasonData() string{


	req,_ := http.NewRequest("Get","https://en.wikipedia.org/w/api.php", nil)
	q := req.URL.Query()
	q.Add("action", "parse")
	q.Add("page", "Game of Thrones (season 1)")
    q.Add("prop", "wikitext")
    q.Add("section", "1")
    q.Add("format", "json")
	req.URL.RawQuery = q.Encode()
	final := req.URL.String()
	resp,_:= http.Get(final)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}
func GetEpisodeData(){
	body := FetchSeasonData()
	fmt.Println(body["parse"])
}
func main(){
	fmt.Println(GetEpisodeData())
}