package models

import (
	"reflect"
	"testing"
)

func init()  {
	ormInit()
}

func TestAddMessge(t *testing.T) {
	type args struct {
		t *Messge
	}
	tests := []struct {
		name    string
		args    args
		wantTid int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTid, err := AddMessge(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddMessge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTid != tt.wantTid {
				t.Errorf("AddMessge() = %v, want %v", gotTid, tt.wantTid)
			}
		})
	}
}

func TestAddMulMessge(t *testing.T) {
	type args struct {
		t []*Messge
	}
	tests := []struct {
		name    string
		args    args
		wantTid int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTid, err := AddMulMessge(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddMulMessge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTid != tt.wantTid {
				t.Errorf("AddMulMessge() = %v, want %v", gotTid, tt.wantTid)
			}
		})
	}
}

func TestGetMessge(t *testing.T) {
	type args struct {
		id       int64
		toUserId string
		status   int
		isDelete bool
	}
	tests := []struct {
		name    string
		args    args
		wantMsg *Messge
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMsg, err := GetMessge(tt.args.id, tt.args.toUserId, tt.args.status, tt.args.isDelete)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMessge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("GetMessge() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestGetMessgeById(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		wantMsg *Messge
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMsg, err := GetMessgeById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMessgeById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("GetMessgeById() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestDeleteMessge(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteMessge(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMessge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAllMessge(t *testing.T) {
	tests := []struct {
		name    string
		wantId  int64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := AllMessge()
			if (err != nil) != tt.wantErr {
				t.Errorf("AllMessge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("AllMessge() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
