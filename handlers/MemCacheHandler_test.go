package handlers

import (
	"fmt"
	"memcache/entity"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_createKey(t *testing.T) {
	cases := []struct {
		name   string
		key    string
		body   *entity.PersonalData
		result string
	}{
		{
			name: "ok",
			key:  "test",
			body: &entity.PersonalData{
				FirstName:      "Andres",
				MiddleName:     "Felipe",
				FirstSurname:   "Alfonso",
				SecondLastName: "Ortiz",
			},
			result: "pong",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			x := engine.POST("/ping", nil)
			fmt.Println(x)
		})
	}

}
