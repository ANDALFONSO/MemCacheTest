package service

import (
	"memcache/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateKey(t *testing.T) {
	cases := []struct {
		name   string
		key    string
		body   *entity.PersonalData
		result string
	}{
		{
			name: "OK",
			key:  "test",
			body: &entity.PersonalData{
				FirstName:      "Johanna",
				MiddleName:     "Lorena",
				FirstSurname:   "Torres",
				SecondLastName: "Marquez",
			},
			result: "creation key test ok",
		},
		{
			name: "key not found",
			body: &entity.PersonalData{
				FirstName:      "Johanna",
				MiddleName:     "Lorena",
				FirstSurname:   "Torres",
				SecondLastName: "Marquez",
			},
			result: "creation key  ok",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			memCacheS := NewMemCache()
			result := memCacheS.CreateKey(c.key, c.body)
			assert.Equal(t, c.result, result)
		})
	}
}
