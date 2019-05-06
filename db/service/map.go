package db

import "fmt"

func LoadMapRun() {
	//schedule reload map
}

func reloadMap(field string, fieldId interface{}) (err error) {
	if field == "id" {
		data := fieldId.(int32)
		fmt.Print(data)
		//reload map
	} else {
		data := fieldId.(string)
		fmt.Print(data)
		//reload map
	}
	return

}
