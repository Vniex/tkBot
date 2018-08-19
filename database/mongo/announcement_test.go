package mongo

import "testing"

func TestAnnouncementDB_FindAnns(t *testing.T) {
	db:=NewAnnouncementDB()
	ann_list,err:=db.FindAnns(10)
	t.Log(err)
	t.Log(ann_list)
}
