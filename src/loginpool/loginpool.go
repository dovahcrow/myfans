package loginpool

import (
	//"fmt"
	// "myfans"
	"uwants"
)

var Uwants map[string]*uwants.Uwants

func init() {
	Uwants = make(map[string]*uwants.Uwants)

}

func GetUwants(username, password string) (*uwants.Uwants, error) {
	//fmt.Printf("|%v|%v|\n", username, password)
	if cl, ok := Uwants[username]; ok {
		return cl, nil
	} else {
		cl := uwants.New(username, password)
		Uwants[username] = cl
		err := cl.Login()
		if err != nil {
			return nil, err
		}
		return cl, nil
	}

}
