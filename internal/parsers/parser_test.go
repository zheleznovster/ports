package parsers

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestParser_OpenFile(t *testing.T) {
	type fields struct {
		FilePointer *os.File
		Decoder     *json.Decoder
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "open non existent file",
			fields:  fields{},
			args:    args{path: "test.json"},
			wantErr: true,
		},
		{
			name:    "open existing file",
			fields:  fields{},
			args:    args{path: "../../testdata/empty.json"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &LargeFileParser{
				FilePointer: tt.fields.FilePointer,
				Decoder:     tt.fields.Decoder,
			}
			if err := parser.openFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_ParseNextRecord(t *testing.T) {

	tests := []struct {
		name    string
		path    string
		want    string
		want1   map[string]interface{}
		wantErr bool
	}{
		{
			name: "test parsing good json",
			path: "../../testdata/good.json",
			want: "AEAJM",
			want1: map[string]interface{}{
				"alias":       []interface{}{},
				"city":        "Ajman",
				"code":        "52000",
				"coordinates": []interface{}{55.5136433, 25.4052165},
				"country":     "United Arab Emirates",
				"name":        "Ajman",
				"province":    "Ajman",
				"regions":     []interface{}{},
				"timezone":    "Asia/Dubai",
				"unlocs":      []interface{}{"AEAJM"},
			},
			wantErr: false,
		},
		{
			name:    "test parsing malformed json",
			path:    "../../testdata/malformed.json",
			want:    "",
			want1:   nil,
			wantErr: true,
		},
		{
			name:    "test parsing empty json",
			path:    "../../testdata/empty.json",
			want:    "",
			want1:   nil,
			wantErr: false,
		},
		{
			name:    "test parsing empty json array",
			path:    "../../testdata/emptyarray.json",
			want:    "",
			want1:   nil,
			wantErr: true,
		},
		{
			name:    "test parsing empty json object",
			path:    "../../testdata/emptyobject.json",
			want:    "",
			want1:   nil,
			wantErr: true,
		},

		{
			name:    "test parsing single opening bracket",
			path:    "../../testdata/openingbracket.json",
			want:    "",
			want1:   nil,
			wantErr: true,
		},
		{
			name:    "test parsing single opening brace",
			path:    "../../testdata/openingbrace.json",
			want:    "",
			want1:   nil,
			wantErr: true,
		},
		{
			name:    "test parsing single closing brace",
			path:    "../../testdata/closingbrace.json",
			want:    "",
			want1:   nil,
			wantErr: false,
		},
		{
			name:    "test parsing single random character",
			path:    "../../testdata/randomcharacter.json",
			want:    "",
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			parser := &LargeFileParser{}
			err := parser.openFile(tt.path)
			if err != nil {
				t.Errorf("TestParser_ParseNextRecord error: %v", err)
			}

			got, got1, err := parser.parseNextRecord()
			if (err != nil) != tt.wantErr {
				t.Errorf("parseNextRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseNextRecord() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseNextRecord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_CloseFile(t *testing.T) {
	type fields struct {
		FilePointer *os.File
		Decoder     *json.Decoder
	}
	tests := []struct {
		path    string
		name    string
		fields  fields
		wantErr bool
	}{
		{
			path:    "",
			name:    "test closing unopened file",
			fields:  fields{},
			wantErr: true,
		},
		{
			path:    "../../testdata/good.json",
			name:    "test closing open file",
			fields:  fields{},
			wantErr: false,
		},
	}
	t.Run(tests[0].name, func(t *testing.T) {
		parser := &LargeFileParser{}
		if err := parser.closeFile(); (err != nil) != tests[0].wantErr {
			t.Errorf("closeFile() error = %v, wantErr %v", err, tests[0].wantErr)
		}
	})

	t.Run(tests[1].name, func(t *testing.T) {
		parser := &LargeFileParser{}
		err := parser.openFile(tests[1].path)
		if err != nil {
			t.Errorf("TestParser_CloseFile error: %v", err)
		}
		if err := parser.closeFile(); (err != nil) != tests[1].wantErr {
			t.Errorf("closeFile() error = %v, wantErr %v", err, tests[1].wantErr)
		}
	})
}
