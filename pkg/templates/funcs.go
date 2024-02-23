package templates

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/sboon-gg/svctl/pkg/maplist"
)

func (r *Renderer) FuncMap() template.FuncMap {
	f := sprig.TxtFuncMap()

	extra := template.FuncMap{
		"pyBool":  pyBool,
		"quote":   quote,
		"env":     env,
		"maplist": r.maplist,
	}

	for k, v := range extra {
		f[k] = v
	}

	return f
}

func (t *Renderer) maplist(filterMap map[string]any) (string, error) {
	c, err := json.Marshal(filterMap)
	if err != nil {
		return "", errors.Join(errors.New("failed to marshal filter map"), err)
	}

	var filter maplist.MapInfo

	if err := json.Unmarshal(c, &filter); err != nil {
		return "", errors.Join(errors.New("failed to unmarshal filter"), err)
	}

	allMaps := make([]maplist.MapInfo, 0)
	// for _, f := range filter {
	allMaps = append(allMaps, maplist.Filter(t.maps, filter)...)

	return maplist.Compose(allMaps), nil
}

func pyBool(b bool) (string, error) {
	if b {
		return "True", nil
	}

	return "False", nil
}

func quote(v any) string {
	return fmt.Sprintf("%q", v)
}

func env(s string) (string, error) {
	return os.Getenv(s), nil
}
