package main

import ("io/ioutil"
		"net/http"
		"fmt"
		"encoding/json"
	)
type parse struct{
	title string
	pageid int
	wikitext Content2
}
type Content2 struct{
	* string
}

func FetchSeasonData() []byte{

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
	return body
}
func main(){
	var basket parse
	var jsonData = FetchSeasonData()
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(basket.title)
}