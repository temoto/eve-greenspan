package eveapi

import (
	"encoding/xml"
	"testing"
	"time"
)

type testEveTimeXml struct {
	XMLName     xml.Name `xml:"root"`
	CurrentTime EveTime  `xml:"currentTime"`
	CachedUntil EveTime  `xml:"cachedUntil"`
	Garbage     string   `xml:"id,attr"`
}

func TestEveTimeXmlMarshal(t *testing.T) {
	expect := `<root id="foobar">
<currentTime>2014-10-31 01:55:12</currentTime>
<cachedUntil>2014-10-31 13:55:12</cachedUntil>
</root>`
	x := testEveTimeXml{
		CurrentTime: EveTime{time.Date(2014, 10, 31, 1, 55, 12, 0, time.UTC)},
		CachedUntil: EveTime{time.Date(2014, 10, 31, 13, 55, 12, 0, time.UTC)},
		Garbage:     "foobar",
	}
	bytes, err := xml.Marshal(x)
	if err != nil {
		t.Fatal(err)
	}
	str := string(bytes)
	if str != expect {
		t.Fatalf("unexpected marshal output: '%s'", str)
	}
}

func TestEveTimeXmlUnmarshal(t *testing.T) {
	in := `<root id="foobar">
<currentTime>2014-10-31 01:55:12</currentTime>
<cachedUntil>2014-10-31 13:55:12</cachedUntil>
</root>`
	var v testEveTimeXml
	err := xml.Unmarshal([]byte(in), &v)
	if err != nil {
		t.Fatal(err)
	}
	expectCurrentTime := EveTime{time.Date(2014, 10, 31, 1, 55, 12, 0, time.UTC)}
	expectCachedUntil := EveTime{time.Date(2014, 10, 31, 13, 55, 12, 0, time.UTC)}
	if v.CurrentTime != expectCurrentTime {
		t.Fatalf("invalid CurrentTime: %v", v.CurrentTime)
	}
	if v.CachedUntil != expectCachedUntil {
		t.Fatalf("invalid CachedUntil: %v", v.CachedUntil)
	}
}
