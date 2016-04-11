package census

import (
	"bytes"
	"net/url"
	"path"
	"strconv"
	"strings"
)

const (
	urlBase          = "census.daybreakgames.com"
	DefaultServiceID = "example"
)

type URLOption func(url.Values)

func BuildURL(id, verb, game, col string, opts ...URLOption) *url.URL {
	q := make(url.Values, len(opts))
	for _, opt := range opts {
		opt(q)
	}

	if id == "" {
		id = DefaultServiceID
	}

	return &url.URL{
		Scheme:   "https",
		Host:     urlBase,
		Path:     path.Join("s:"+id, verb, game, col),
		RawQuery: q.Encode(),
	}
}

func SearchOption(name, val string) URLOption {
	return func(q url.Values) {
		q.Set(name, val)
	}
}

func listOption(name string, vals ...string) URLOption {
	str := strings.Join(vals, ",")
	return func(q url.Values) {
		q.Set(name, str)
	}
}

func ShowOption(fields ...string) URLOption {
	return listOption("c:show", fields...)
}

func HideOption(fields ...string) URLOption {
	return listOption("c:hide", fields...)
}

type Sort struct {
	Field string
	Dir   SortDir
}

func (s Sort) String() string {
	buf := bytes.NewBuffer(make([]byte, 0, len(s.Field)+2))
	buf.WriteString(s.Field)

	switch s.Dir {
	case Desc:
		buf.WriteString(":-1")
	case Asc:
		buf.WriteString(":1")
	}

	return buf.String()
}

type SortDir int

const (
	Desc SortDir = iota - 1
	DefaultDir
	Asc
)

func SortOption(sorts ...Sort) URLOption {
	var buf bytes.Buffer
	for _, sort := range sorts {
		buf.WriteString(sort.String())
		buf.WriteByte(',')
	}
	str := buf.String()

	return func(q url.Values) {
		q.Set("c:sort", str[:len(str)-1])
	}
}

func HasOption(fields ...string) URLOption {
	return listOption("c:has", fields...)
}

func ResolveOption(fields ...string) URLOption {
	return listOption("c:resolve", fields...)
}

func IgnoreCaseOption() URLOption {
	return func(q url.Values) {
		q.Set("c:case", "false")
	}
}

func LimitOption(limit int) URLOption {
	str := strconv.FormatInt(int64(limit), 10)
	return func(q url.Values) {
		q.Set("c:limit", str)
	}
}

func LimitPerDBOption(limit int) URLOption {
	str := strconv.FormatInt(int64(limit), 10)
	return func(q url.Values) {
		q.Set("c:limitPerDB", str)
	}
}

func StartOption(start int) URLOption {
	str := strconv.FormatInt(int64(start), 10)
	return func(q url.Values) {
		q.Set("c:start", str)
	}
}

func IncludeNullOption() URLOption {
	return func(q url.Values) {
		q.Set("c:includeNull", "true")
	}
}

func LangOption(lang string) URLOption {
	return func(q url.Values) {
		q.Set("c:lang", lang)
	}
}

// TODO: Implement JoinOption.
// TODO: Implement TreeOption.

func TimingOption() URLOption {
	return func(q url.Values) {
		q.Set("c:timing", "true")
	}
}

func ExactMatchFirstOption() URLOption {
	return func(q url.Values) {
		q.Set("c:exactMatchFirst", "true")
	}
}

func DistinctOption(field string) URLOption {
	return func(q url.Values) {
		q.Set("c:distinct", field)
	}
}

func NoRetryOption() URLOption {
	return func(q url.Values) {
		q.Set("c:retry", "false")
	}
}
