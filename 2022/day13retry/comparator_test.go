package day13retry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

	str := "[[9,1],[[[8],2,5],10,5],[[],[[0,10,2,10,2],[6]],5],[[10,[]],[],[[5,1,3,10],5],[[],[],[10,0]]],[[4,3],[[],[],[2,8,0,6]],4,[[8,1,7,1],1,[2,9,0],4,4]]]"

	c := Comparator{}
	err := c.Parse(str)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(c.List))
}
