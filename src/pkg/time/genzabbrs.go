// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

//
// usage:
//
// go run genzabbrs.go | gofmt > $GOROOT/src/pkg/time/zoneinfo_abbrs_windows.go
//

package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"text/template"
	"time"
)

// getAbbrs finds timezone abbreviations (standard and daylight saving time)
// for location l.
func getAbbrs(l *time.Location) (st, dt string) {
	t := time.Date(time.Now().Year(), 0, 0, 0, 0, 0, 0, l)
	abbr1, off1 := t.Zone()
	for i := 0; i < 12; i++ {
		t = t.AddDate(0, 1, 0)
		abbr2, off2 := t.Zone()
		if abbr1 != abbr2 {
			if off2-off1 < 0 { // southern hemisphere
				abbr1, abbr2 = abbr2, abbr1
			}
			return abbr1, abbr2
		}
	}
	return abbr1, abbr1
}

type zone struct {
	WinName  string
	UnixName string
	StTime   string
	DSTime   string
}

type zones []*zone

func (zs zones) Len() int           { return len(zs) }
func (zs zones) Swap(i, j int)      { zs[i], zs[j] = zs[j], zs[i] }
func (zs zones) Less(i, j int) bool { return zs[i].UnixName < zs[j].UnixName }

const wzURL = "http://unicode.org/cldr/data/common/supplemental/windowsZones.xml"

type MapZone struct {
	Other     string `xml:"other,attr"`
	Territory string `xml:"territory,attr"`
	Type      string `xml:"type,attr"`
}

type SupplementalData struct {
	Zones []MapZone `xml:"windowsZones>mapTimezones>mapZone"`
}

func readWindowsZones() (zones, error) {
	r, err := http.Get(wzURL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var sd SupplementalData
	err = xml.Unmarshal(data, &sd)
	if err != nil {
		return nil, err
	}
	zs := make(zones, 0)
	for _, z := range sd.Zones {
		if z.Territory != "001" {
			// to avoid dups. I don't know why.
			continue
		}
		l, err := time.LoadLocation(z.Type)
		if err != nil {
			return nil, err
		}
		st, dt := getAbbrs(l)
		zs = append(zs, &zone{
			WinName:  z.Other,
			UnixName: z.Type,
			StTime:   st,
			DSTime:   dt,
		})
	}
	return zs, nil
}

func main() {
	zs, err := readWindowsZones()
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(zs)
	var v = struct {
		URL string
		Zs  zones
	}{
		wzURL,
		zs,
	}
	err = template.Must(template.New("prog").Parse(prog)).Execute(os.Stdout, v)
	if err != nil {
		log.Fatal(err)
	}
}

const prog = `
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// generated by genzabbrs.go from
// {{.URL}}

package time

type abbr struct {
	std string
	dst string
}

var abbrs = map[string]abbr{
{{range .Zs}}	"{{.WinName}}": {"{{.StTime}}", "{{.DSTime}}"}, // {{.UnixName}}
{{end}}}

`
