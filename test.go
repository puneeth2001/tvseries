package main

import ("io/ioutil"
		"net/http"
		"fmt"
		"encoding/json"
		"reflect"
	)
	type Initial struct {
		Parse struct {
			Title    string `json:"title"`
			Pageid   int    `json:"pageid"`
			Wikitext struct {
				NAMING_FAILED string `json:"*"`
			} `json:"wikitext"`
		} `json:"parse"`
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
	
	return (body)
}
func main(){
	object := &Initial{}
	var jsonData = FetchSeasonData()
	err := json.Unmarshal([]byte(jsonData), object)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(object.Parse.Wikitext.NAMING_FAILED)
	fmt.Print(reflect.TypeOf(object.Parse.Wikitext.NAMING_FAILED))
}





