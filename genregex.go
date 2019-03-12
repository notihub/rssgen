package rssgen

import (
	"regexp"
	"strings"
)

type GenRegex struct {
	Rule string
}

func (g *GenRegex) Regex() (regex string, err error) {
	regex = g.Rule

	// {%:parameter}: The part to parse
	re := regexp.MustCompile(`{%:([a-z]+)}`)
	regex = re.ReplaceAllString(regex, `(?P<$1>[\s\S]+?)`)
	// {%}: The part to parse
	regex = strings.Replace(regex, "{%}", `([\s\S]+?)`, -1)
	// {_}: Blank
	regex = strings.Replace(regex, "{_}", `[ \t]+?`, -1)
	// {|}: Newline
	regex = strings.Replace(regex, "{|}", `\n+?`, -1)
	// {*}: Ignore
	regex = strings.Replace(regex, "{*}", `[\s\S]+?`, -1)

	return regex, nil
}

func (g *GenRegex) Parse(html string) (items []map[string]string, size int, err error) {
	regex, err := g.Regex()
	if err != nil {
		return nil, 0, err
	}
	re := regexp.MustCompile(regex)
	match := re.FindAllStringSubmatch(html, -1)

	size = len(match)
	items = make([]map[string]string, size)
	for i, _ := range items {
		items[i] = make(map[string]string)
	}

	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			for idx, data := range match {
				items[idx][name] = strings.TrimSpace(data[i])
			}
		}
	}

	return items, len(match), nil
}
