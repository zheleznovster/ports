package database

import (
	"reflect"
	"sync"
	"testing"
)

func TestDatabase_Delete(t *testing.T) {
	type fields struct {
		Data map[string]interface{}
		mtx  sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "delete element in nil database",
			fields: fields{},
			args:   args{key: "ABC"},
			want:   false,
		},
		{
			name: "delete non-existent element in non-empty database",
			fields: fields{
				Data: map[string]interface{}{"XYZ": 1},
			},
			args: args{key: "ABC"},
			want: false,
		},
		{
			name: "delete existing element in single element map",
			fields: fields{
				Data: map[string]interface{}{"ABC": 1},
			},
			args: args{key: "ABC"},
			want: true,
		},
		{
			name: "delete existing element in multi element map",
			fields: fields{
				Data: map[string]interface{}{"z": 5, "ABC": 1, "B": 1},
			},
			args: args{key: "ABC"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Database{
				Data: tt.fields.Data,
				mtx:  sync.RWMutex{},
			}
			if got := c.Delete(tt.args.key); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_Get(t *testing.T) {
	type fields struct {
		Data map[string]interface{}
		mtx  sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		{
			name:   "get non-existent element in nil database",
			fields: fields{},
			args:   args{key: "ABC"},
			want:   nil,
			want1:  false,
		},
		{
			name: "get non-existent element in non-empty database",
			fields: fields{
				Data: map[string]interface{}{"XYZ": 1},
			},
			args:  args{key: "ABC"},
			want:  nil,
			want1: false,
		},
		{
			name: "get existing element in single element map",
			fields: fields{
				Data: map[string]interface{}{"ABC": 1},
			},
			args:  args{key: "ABC"},
			want:  1,
			want1: true,
		},
		{
			name: "get existing element in multi element map",
			fields: fields{
				Data: map[string]interface{}{"z": 5, "ABC": 1, "B": 1},
			},
			args:  args{key: "ABC"},
			want:  1,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Database{
				Data: tt.fields.Data,
				mtx:  sync.RWMutex{},
			}
			got, got1 := c.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDatabase_Insert(t *testing.T) {
	type fields struct {
		Data map[string]interface{}
		mtx  sync.RWMutex
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "insert element in nil database",
			fields: fields{},
			args:   args{key: "ABC", value: 1},
			want:   false,
		},
		{
			name: "insert element in non-nil database",
			fields: fields{
				Data: map[string]interface{}{},
			},
			args: args{key: "ABC", value: 1},
			want: true,
		},
		{
			name: "insert already existing element",
			fields: fields{
				Data: map[string]interface{}{"Q": 5, "ABC": 1, "XYZ": 1},
			},
			args: args{key: "XYZ", value: 2},
			want: false,
		},
		{
			name: "insert element with nil value",
			fields: fields{
				Data: map[string]interface{}{"ABC": 1},
			},
			args: args{key: "ZAS", value: nil},
			want: true,
		},
		{
			name: "insert element with empty key",
			fields: fields{
				Data: map[string]interface{}{},
			},
			args: args{key: "", value: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Database{
				Data: tt.fields.Data,
				mtx:  sync.RWMutex{},
			}
			if got := c.Insert(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_Update(t *testing.T) {
	type fields struct {
		Data map[string]interface{}
		mtx  sync.RWMutex
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "update element in empty database",
			fields: fields{},
			args:   args{key: "ABC", value: 1},
			want:   false,
		},
		{
			name: "update already existing element",
			fields: fields{
				Data: map[string]interface{}{"Q": 5, "ABC": 1, "XYZ": 1},
			},
			args: args{key: "XYZ", value: 2},
			want: true,
		},
		{
			name: "update nil element with value",
			fields: fields{
				Data: map[string]interface{}{"ABC": nil},
			},
			args: args{key: "ABC", value: 1},
			want: true,
		},
		{
			name: "update element with empty key",
			fields: fields{
				Data: map[string]interface{}{},
			},
			args: args{key: "", value: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Database{
				Data: tt.fields.Data,
				mtx:  sync.RWMutex{},
			}
			if got := c.Update(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
