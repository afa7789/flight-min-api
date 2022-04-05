package flight

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	type args struct {
		entry [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "single entry",
			args: args{
				entry: [][]string{
					{"a", "b"},
				},
			},
			want:    []string{"a", "b"},
			wantErr: false,
		},
		{
			name: "valid dual entry",
			args: args{
				entry: [][]string{
					{"a", "b"},
					{"b", "c"},
				},
			},
			want:    []string{"a", "c"},
			wantErr: false,
		},
		{
			name: "invalid dual entry one",
			args: args{
				entry: [][]string{
					{"a", "b"},
					{"d", "c"},
				},
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "invalid dual entry two",
			args: args{
				entry: [][]string{
					{"a", "b"},
					{"a", "c"},
				},
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "invalid three entry one",
			args: args{
				entry: [][]string{
					{"a", "b"},
					{"b", "c"},
					{"g", "e"},
				},
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "invalid three entry two",
			args: args{
				entry: [][]string{
					{"a", "b"},
					{"c", "d"},
					{"g", "e"},
				},
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "invalid three entry three",
			args: args{
				entry: [][]string{
					{"a", "b"},
					{"a", "d"},
					{"g", "e"},
				},
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "real case test one",
			args: args{
				entry: [][]string{
					{"SFO", "EWR"},
				},
			},
			want:    []string{"SFO", "EWR"},
			wantErr: false,
		},
		{
			name: "real case test two",
			args: args{
				entry: [][]string{
					{"ATL", "EWR"},
					{"SFO", "ATL"},
				},
			},
			want:    []string{"SFO", "EWR"},
			wantErr: false,
		},
		{
			name: "real case test three",
			args: args{
				entry: [][]string{
					{"IND", "EWR"},
					{"SFO", "ATL"},
					{"GSO", "IND"},
					{"ATL", "GSO"},
				},
			},
			want:    []string{"SFO", "EWR"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parser(tt.args.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser() = %v, want %v", got, tt.want)
			}
		})
	}
}
