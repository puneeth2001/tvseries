package main

import("fmt"
		"encoding/json"
	)

type Wheel struct{
	Id string `json:"id"`
	Radius int `json:"radius"`
	Type int `json:"type"`
	attr interface{}
}

type SolidAttributes struct {
	Material string `json:"material"`
   }
   
   type SpokedAttributes struct {
	SpokeCount int `json:"spoke_count"`
	InnerRadius int `json:"inner_radius"`
   }
   
//    type TireAttributes struct {
// 	Fill FillType `json:"fill_type"` // FillType is an enum made of int. See details further down.
// 	InnerRadius int `json:"inner_radius"`
//    }
func (w *Wheel) UnmarshalJSON(b []byte) error {            // Line A
    // populate common attributes in w
    err := json.Unmarshal(b, w)                            // Line B
    ...
    // second pass to get attributes of specific sub-types
    switch w.Type {
    case WheelType_Spoked:
        subTypeAttr := &SpokedAttributes{}
        err := json.Unmarshal([]byte(wheelJson), subTypeAttr)
    ...

func main(){
	var wheelJson = `
	{
	 "id": "bicycle_wheel_8450",
	 "radius": 20,
	 "type": 1,
	 "spoke_count": 30,
	 "inner_radius": 19
	}
	`
	wheelObj := &Wheel{}
	json.Unmarshal([]byte(wheelJson), wheelObj)
	fmt.Println((wheelJson))
	fmt.Println(wheelObj.Radius)
}
