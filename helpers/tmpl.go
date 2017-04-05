package helpers

import (
    "regexp"
    "time"

    "github.com/microcosm-cc/bluemonday"
    "github.com/russross/blackfriday"
)

func DateTime(t time.Time, layout string) string {
    return t.Format(layout) //`2006-01-02 15:04:05`
}

func ParseTime(date string) time.Time {
    var layout string
    switch len(date) {
    case 10:
        layout = "2006-01-02"
    case 13:
        layout = "2006-01-02 15"
    case 16:
        layout = "2006-01-02 15:04"
    default:
        layout = "2006-01-02 15:04:05"
    }
    tt, _ := time.ParseInLocation(layout, date, time.Local)
    return tt
}

func Markdown(in string) string {
    unsafe := blackfriday.MarkdownCommon([]byte(in))
    p := bluemonday.UGCPolicy()
    p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code")
    html := p.SanitizeBytes(unsafe)
    return string(html)
}
