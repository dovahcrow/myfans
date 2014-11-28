package models

import (
	"testing"
)

func TestDeleteFid(t *testing.T) {
	err := DeleteAllFids()
	if err != nil {
		t.Error(err)
	}
}
func TestInsertFid(t *testing.T) {
	err := InsertFids([]string{"123", "456", `789`, `11`, `12`})
	if err != nil {
		t.Error(err)
	}
}
func TestGetAllFid(t *testing.T) {
	fids, err := GetAllFids()
	if err != nil {
		t.Error(err)
	}
	t.Logf("fids: %v", fids)
}

func TestDeleteTid(t *testing.T) {
	err := DeleteAllTids()
	if err != nil {
		t.Error(err)
	}
}
func TestInsertTid(t *testing.T) {
	err := InsertTids([]string{"123", "456", `789`, `11`, `12`})
	if err != nil {
		t.Error(err)
	}
}
func TestGetAllTid(t *testing.T) {
	tids, err := GetAllTids()
	if err != nil {
		t.Error(err)
	}
	t.Logf("fids: %v", tids)
}
