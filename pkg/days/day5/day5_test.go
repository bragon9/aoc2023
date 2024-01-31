package day5

import "testing"

func TestTranslationMap_AddValue(t *testing.T) {
	type fields struct {
		Ranges []int
		Offset []int
	}
	type args struct {
		src  int
		rng  int
		dest int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TranslationMap{
				Ranges: tt.fields.Ranges,
				Offset: tt.fields.Offset,
			}
			tr.AddValue(tt.args.src, tt.args.rng, tt.args.dest)
		})
	}
}
