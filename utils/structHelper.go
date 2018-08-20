package utils

import "github.com/goinggo/mapstructure"

//  convert map[string]interface{} to struct
// second parameter must be pointer!!!
func Map2Struct(data map[string]interface{} ,result interface{},tagName string){
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   result,
		TagName:  tagName,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(data)
	if err != nil {
		panic(err)
	}

}