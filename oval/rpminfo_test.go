package oval

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestLookupRPMIntoObject(t *testing.T) {
	var tt = []struct{ Ref, Name string }{
		{"oval:com.redhat.rhba:obj:20070026001", "htdig"},        // This should be the first object.
		{"oval:com.redhat.rhsa:obj:20091206001", "libxml"},       // random one
		{"oval:com.redhat.rhsa:obj:20100720002", "mikmod-devel"}, // This should be the last object.
	}
	f, err := os.Open("../testdata/Red_Hat_Enterprise_Linux_3.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var root Root
	if err := xml.NewDecoder(f).Decode(&root); err != nil {
		t.Fatal(err)
	}
	for _, tc := range tt {
		obj := root.Objects.LookupRPMInfoObject(tc.Ref)
		if obj == nil {
			t.Fatal("exepected to find object, got nil")
		}
		t.Logf("%s: %s (%#+v)", tc.Ref, obj.Name, obj)
		if got, want := obj.Name, tc.Name; got != want {
			t.Fatalf("got: %q, want: %q", got, want)
		}
	}
}