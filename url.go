package census

import (
	"bytes"
	"net/url"
	"path"
	"strconv"
	"strings"
)

const (
	urlBase = "census.daybreakgames.com"

	// DefaultServiceID is the ID which is used when none is specified.
	DefaultServiceID = "example"
)

// A URLOption is an option that is passed to the API.
type URLOption func(url.Values)

// BuildURL builds a request URL from the various components.
//
//     id:   The service ID.
//     verb: The type of request. Usually `get` or `count`.
//     game: The namespace. For example, `ps2`.
//     col:  The collection type. For example, `character`.
//     opts: A list of options. See the example for more information.
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

// SearchOption is a search request. For example,
//
//     SearchOption("name.first_lower", "deedlefaketr")
//
// produces
//
//     name.first_lower=deedlefaketr
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

// ShowOption is a list of fields to show. It cooresponds to `c:show`.
func ShowOption(fields ...string) URLOption {
	return listOption("c:show", fields...)
}

// HideOption is a list of fields to hide. It cooresponds to `c:hide`.
func HideOption(fields ...string) URLOption {
	return listOption("c:hide", fields...)
}

// A Sort is a mapping of field names to sort directions.
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

// A SortDir is a direction to sort a field.
type SortDir int

const (
	// Desc is descending order.
	Desc SortDir = iota - 1

	// DefaultDir is an unspecified order.
	DefaultDir

	// Asc is ascending order.
	Asc
)

// SortOption is a list of fields to sort by. It cooresponds to
// `c:sort`.
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

// HasOption filters objects by the presence of fields. It cooresponds
// to `c:has`.
func HasOption(fields ...string) URLOption {
	return listOption("c:has", fields...)
}

// ResolveOption is a list of resolves. It cooresponds to `c:resolve`.
func ResolveOption(fields ...string) URLOption {
	return listOption("c:resolve", fields...)
}

// IgnoreCaseOption makes searches case-insensetive. It cooresponds to
// `c:case=false`.
func IgnoreCaseOption() URLOption {
	return func(q url.Values) {
		q.Set("c:case", "false")
	}
}

// LimitOption limits the number of returned results. It cooresponds
// to `c:limit`.
func LimitOption(limit int) URLOption {
	str := strconv.FormatInt(int64(limit), 10)
	return func(q url.Values) {
		q.Set("c:limit", str)
	}
}

// LimitPerDBOption is a different type of limit on the number of
// returned result than LimitOption. It cooresponds to `c:limitPerDB`.
func LimitPerDBOption(limit int) URLOption {
	str := strconv.FormatInt(int64(limit), 10)
	return func(q url.Values) {
		q.Set("c:limitPerDB", str)
	}
}

// StartOption specifies the index of the first result. It cooresponds
// to `c:start`.
func StartOption(start int) URLOption {
	str := strconv.FormatInt(int64(start), 10)
	return func(q url.Values) {
		q.Set("c:start", str)
	}
}

// IncludeNullOption enables null field values. It cooresponds to
// `c:includeNull=true`.
func IncludeNullOption() URLOption {
	return func(q url.Values) {
		q.Set("c:includeNull", "true")
	}
}

// LangOption specifies what language to return in multi-lingual
// results. It cooresponds to `c:lang`.
func LangOption(lang string) URLOption {
	return func(q url.Values) {
		q.Set("c:lang", lang)
	}
}

// TODO: Implement JoinOption.

// TODO: Implement TreeOption.

// TimingOption adds a field in the top-level struct that gives the
// time the API servers spent fetching data from the database. It cooresponds to `c:timing=true`.
func TimingOption() URLOption {
	return func(q url.Values) {
		q.Set("c:timing", "true")
	}
}

// ExactMatchFirstOption specifies that an exact match should always
// be the first result. It cooresponds to `c:exactMatchFirst=true`.
func ExactMatchFirstOption() URLOption {
	return func(q url.Values) {
		q.Set("c:exactMatchFirst", "true")
	}
}

// DistinctOption specifies that all distinct values of a field should
// be returned. It cooresponds to `c:distinct`.
func DistinctOption(field string) URLOption {
	return func(q url.Values) {
		q.Set("c:distinct", field)
	}
}

// NoRetryOption specifies that the API shouldn't automatically retry
// on failures. It cooresponds to `c:retry=false`.
func NoRetryOption() URLOption {
	return func(q url.Values) {
		q.Set("c:retry", "false")
	}
}
