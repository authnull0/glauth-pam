package external
import (
	"fmt"
)

type Authnull struct{

}


func (a Authnull) Init() {
	fmt.Println("Init Authnull")
}

func (a Authnull) FetchUsers() {}
