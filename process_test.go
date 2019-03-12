package rssgen

import (
	"testing"
)

func TestGenerateRSS(t *testing.T) {
	rule := `<a href="{%:link}">
                      <span class="text-normal">
                        {%:name} / 
                      </span>
                      {%:title}
                    </a>
                  {*}
                  <p class="col-9 d-inline-block text-gray m-0 pr-4">
                            {%:description}
                  </p>`
	rss, err := GenerateRSS(Generator{Url: "https://github.com/trending", Rule: rule})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rss)
}
