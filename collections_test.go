package collections
import (
	"testing"
	"fmt"
)

type interface_map_test struct {
	source, destination []interface{}
	mapper InterfacesMapper
}

var interface_map_tests = []interface_map_test{
	{
		[]interface{}{ "a", "aaa", "abc" },
		[]interface{}{ "pref_a", "pref_aaa", "pref_abc" },
		func (v interface{}) (interface{}) { return fmt.Sprintf("pref_%s", v.(string)) },
	},
}

func TestInterfacesMap(t *testing.T) {
	for _, v := range interface_map_tests {
		interfaces := NewFromSlice(v.source)
		out := interfaces.Map(v.mapper)
		if len(out.values) != len(v.destination) {
			t.Fatalf("Expected output length %d, got %d", len(v.destination), len(out.values))
		}
		for i := 0; i < len(v.destination); i++ {
			found := false
			for j := 0; j < len(out.values); j++ {
				if out.values[j] == v.destination[i] {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Could not find item %v into output collection", v.destination[i])
			}
		}
	}
}

type interface_filter_test struct {
	source, destination []interface{}
	filter InterfacesFilter
}

var interface_filter_tests = []interface_filter_test{
	{
		[]interface{}{ 1,5,6,8,9,12,23,45,67,78 },
		[]interface{}{ 6, 8, 12, 78},
		func (v interface{}) (bool) { return (v.(int) % 2) == 0 },
	},
}

func TestInterfacesFilter(t* testing.T) {
	for _, v := range interface_filter_tests {
		interfaces := NewFromSlice(v.source)
		out := interfaces.Filter(v.filter)
		if len(out.values) != len(v.destination) {
			t.Fatalf("Expected output length %d, got %d", len(v.destination), len(out.values))
		}
		for i := 0; i < len(v.destination); i++ {
			found := false
			for j := 0; j < len(out.values); j++ {
				if out.values[j] == v.destination[i] {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Could not find item %v into output collection", v.destination[i])
			}
		}
	}
}
