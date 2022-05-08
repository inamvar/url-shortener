package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.alibaba.ir/mag/?utm_source=alibaba&utm_medium=Click_CTA&utm_campaign=inbound&utm_term=category-buttons&utm_content=Alibaba-Maghttps://www.alibaba.ir/mag?utm_source=alibaba&utm_medium=Click_CTA&utm_campaign=inbound&utm_term=category-buttons&utm_content=Alibaba-Mag"
	shortLink_1, err := GenerateShortLink(initialLink_1)
	assert.NoError(t, err)

	initialLink_2 := "https://www.alibaba.ir/mag/category/alibabastories/"
	shortLink_2, err := GenerateShortLink(initialLink_2)
	assert.NoError(t, err)

	initialLink_3 := "https://www.alibaba.ir/tour/search/5129?origin=iran-tehran&from=2022-05-11&to=2022-05-13&rooms=2&destination=turkey-istanbul&step=PDP"
	shortLink_3, err := GenerateShortLink(initialLink_3)
	assert.NoError(t, err)

	assert.Equal(t, shortLink_1, "YWemC5")
	assert.Equal(t, shortLink_2, "ZqRv8j")
	assert.Equal(t, shortLink_3, "gaMBUn")
}
