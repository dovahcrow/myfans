package uwants

import (
	"testing"
)

func Test_post(t *testing.T) {
	cl := New(`doomsplayer`, `1cd3599df`)
	cl.Login()
	t.Log(cl.SendReply(`17318841`, ``, `yes`))
}

//func Test_Newtopic(t *testing.T) {
//	cl := New(`doomsplayer`, `1cd3599df`)
//	cl.Login()
//	t.Log(cl.NewTopic(`401`, `haveatry`, `p`))
//}
