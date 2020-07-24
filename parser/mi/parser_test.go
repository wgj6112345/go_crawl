package mi

import (
	"fmt"
	"regexp"
	"testing"
)

var str1 = `<div data-v-70283663="" class="product-con"><h2 data-v-70283663=""><img data-v-70283663="">([^<]+)</h2><p data-v-70283663="" class="sale-desc"><font[\d\D]*?>[\d\D]*?</font>([^<]+)</p>[\d\D]*?<div[\d\D]*?class="price-info"><span[\d\D]*?>([^<]+)<!----><!----></span></div>`

func TestParse(t *testing.T) {
	// reList := `<li>[\d\D]*?<a href="([^"]+)"[\d\D]*?><img[\d\D]*?></a>[\d\D]*?<a[\d\D]*?>[\d\D]*?</a>[\d\D]*?</li>`
	reList := `<a class="J_nav_comment " href="([^"]+)"[\d\D]*?>用户评价</a>`

	re := regexp.MustCompile(reList)
	match := re.FindAllSubmatch([]byte(str1), -1)

	for _, m := range match {
		fmt.Println("m[1]: ", string(m[1]))
	}
}
