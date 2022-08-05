package managers

import (
	"testing"
)

func TestManager_LoadData(t *testing.T) {

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test existing file",
			args:    args{path: "../testdata/good.json"},
			wantErr: false,
		},
		{
			name:    "test existing file with duplicates",
			args:    args{path: "../testdata/duplicatedports.json"},
			wantErr: false,
		},
		{
			name:    "test empty file",
			args:    args{path: "../testdata/empty.json"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager, _ := NewManager(tt.args.path)
			if err := manager.LoadData(); (err != nil) != tt.wantErr {
				t.Errorf("LoadData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
