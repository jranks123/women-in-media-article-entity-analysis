package secretsanta

import (
	"fmt"
	"testing"
)

func TestReadJsonFile(t *testing.T) {
	res := ReadJsonFile()
	for i := 0; i < len(res.FriendList); i++ {
		fmt.Println("Name", res.FriendList[i].Name)
	}
}