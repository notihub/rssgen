package rssgen

import (
	"bytes"
	"crypto/tls"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yosssi/gohtml"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"regexp"
	"strings"
	"time"
)

func getHttp(link string) (htmlBytes []byte, html string, err error) {
	client := &http.Client{
		// ignore x509: certificate signed by unknown authority
		// https://github.com/golang/go/issues/9586
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(30 * time.Second),
	}

	req, err := http.NewRequest("GET", link, nil)
	req.Close = true
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	htmlBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	html = string(htmlBytes)
	log.Println(len(html))

	// todo: Encoding, EUC-KR
	contentType := resp.Header.Get("Content-Type")
	if len(contentType) > 0 {
		charset := strings.Split(contentType, "charset=")
		if len(charset) == 2 {
			var bufs bytes.Buffer
			var wr *transform.Writer

			switch strings.ToLower(charset[1]) {
			case "euc-kr":
				wr = transform.NewWriter(&bufs, korean.EUCKR.NewDecoder())
				break
			}

			if wr != nil {
				wr.Write(htmlBytes)
				wr.Close()
				html = bufs.String()
			}
		}
	}

	return htmlBytes, html, err
}

func reformatHtml(url, html string) string {
	html = gohtml.Format(html)

	var host string
	u, err := url2.Parse(url)
	if err == nil {
		host = u.Scheme + "://" + u.Host
	} else {
		host = url
		log.Println(url, err)
	}

	path := u.Path[0:strings.LastIndex(u.Path, "/")]

	re := regexp.MustCompile(`((?:href|src)=["'])(.*?)(["'])`)
	//html = re.ReplaceAllString(html, "${1}"+url+"${2}")
	html = re.ReplaceAllStringFunc(html, func(m string) string {
		parts := re.FindStringSubmatch(m)

		var linkUrl *url2.URL

		// Parse the anchor tag URL
		if linkUrl, err = url2.Parse(parts[2]); err != nil {
			return parts[1] + parts[2] + parts[3]
		}

		if linkUrl.IsAbs() {
			return parts[1] + parts[2] + parts[3]
		} else {
			if strings.HasPrefix(parts[2], "//") {
				return parts[1] + u.Scheme + ":" + parts[2] + parts[3]
			}
			if strings.HasPrefix(parts[2], "/") {
				return parts[1] + host + parts[2] + parts[3]
			}
			return parts[1] + host + path + "/" + parts[2] + parts[3]
		}
	})

	return html
}
