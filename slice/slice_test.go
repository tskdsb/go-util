package util

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type args struct {
		slice interface{}
		index int
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				slice: []int{1, 2},
				index: 1,
			},
			want:    []int{1},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				slice: []int{1, 2},
				index: 2,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveElement(tt.args.slice, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElement() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSlice(t *testing.T) {
	type (
		args struct {
			slice  interface{}
			filter func(element interface{}) bool
		}
		person struct {
			name string
			sex  bool
		}
	)

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				slice: []person{
					{
						name: "nan",
						sex:  true,
					},
					{
						name: "nv",
						sex:  false,
					},
				},
				filter: func(element interface{}) bool {
					e := element.(person)
					if e.sex {
						return true
					} else {
						return false
					}
				},
			},
			want: []person{
				{
					name: "nan",
					sex:  true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FilterSlice(tt.args.slice, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
