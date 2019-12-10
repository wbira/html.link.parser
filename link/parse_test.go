package link_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wbira/html.link.parser/link"
)

const exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
    A link to another page
    <span> some span  </span>
  </a>
  <a href="/page-two">A link to a second page</a>
</body>
</html>
`

func TestBasicHtmlParsing(t *testing.T) {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)

	if err != nil {
		t.Errorf("Error during parsing")
	}
	fmt.Printf("%+v\n", links)
}
