package mongo

import "testing"

func TestAnnouncementDB_FindAnns(t *testing.T) {
	db:=GetAnnouncementDB()
	ann_list,err:=db.FindAnns(-1,16500,10)
	t.Log(err)
	for _,ann:=range ann_list{
		t.Log(ann.Id)
	}

}
