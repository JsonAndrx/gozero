package users

import (
    "fmt"
    "time"
	"github.com/jsonandrx/gozero/modelos"
)

func AltaUsuario(){
	user := new(modelos.User)
	user.AddUser(10, "Yeison", time.Now(), true)
	fmt.Println(user)
}