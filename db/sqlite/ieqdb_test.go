package db

import (
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3" // the db driver supports database/sql
)

/*
DROP TABLE IF EXISTS test_table;
CREATE TABLE test_table (
	device_id VARCHAR(250),
	created_on datetime,
	status int,
	status_desc VARCHAR(500)
	);

select rowid,* from test_table;
SELECT rowid, datetime(created_on, 'unixepoch', 'localtime') from test_table order by rowid desc;

go test -count=1
go clean -testcache && go test

check TestTimestamp(t *testing.T) in
https://github.com/mattn/go-sqlite3/blob/master/sqlite3_test.go

*/

func TestDatetime(t *testing.T) {
	db := connect("ieqdb_test.db")
	defer db.Close()

	// using db unix time stored as integer
	stmt, err := db.Prepare("DELETE FROM test_table")
	if err != nil {
		t.Error(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("cleared table")

	// using db unix time stored as integer
	stmt, err = db.Prepare("INSERT INTO test_table(device_id, created_on, status, status_desc) VALUES(?,strftime('%s', 'now'),?,?)")
	if err != nil {
		t.Error(err.Error())
	}
	_, err = stmt.Exec("test device", 1, "db unix time inserted")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("db unix time inserted")

	// using Go Unix time stored as integer
	dtunix := time.Now().Unix()
	stmt, err = db.Prepare("INSERT INTO test_table(device_id, created_on, status, status_desc) VALUES(?,?,?,?)")
	if err != nil {
		t.Error(err.Error())
	}
	_, err = stmt.Exec("test device", dtunix, 1, "Go unix time inserted")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("Go unix time inserted")

	// using Go datetime localtime stored as string
	dt := time.Now()
	stmt, err = db.Prepare("INSERT INTO test_table(device_id, created_on, status, status_desc) VALUES(?,?,?,?)")
	if err != nil {
		t.Error(err.Error())
	}
	_, err = stmt.Exec("test device", dt, 1, "Go localtime inserted")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("Go localtime inserted")

	// using Go datetime UTC stored as string
	dt = time.Now().UTC()
	stmt, err = db.Prepare("INSERT INTO test_table(device_id, created_on, status, status_desc) VALUES(?,?,?,?)")
	if err != nil {
		t.Error(err.Error())
	}
	_, err = stmt.Exec("test device", dt, 1, "Go UTC time inserted")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("Go UTC time inserted")
}

func TestReadMetrics(t *testing.T) {
	db := connect("ieqdb_test.db")
	defer db.Close()

	results, err := ReadMetrics("ieqdb_test.db", 10)
	if err != nil {
		t.Error(err.Error())
	}

	location, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		t.Error(err.Error())
	}

	reversed := []Metric{}
	utc := time.Now().UTC()
	local := utc
	for i, m := range results {
		utc = m.CreatedOn
		local = utc
		local = local.In(location)
		//t.Log("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
		t.Logf("%v %v %g %g\n", m.CreatedOn, local, m.Temperature, m.Humidity)
		n := results[len(results)-1-i]
		reversed = append(reversed, n)
	}

	t.Log("Reversed")
	// reverse order
	for _, r := range reversed {
		utc = r.CreatedOn
		local = utc
		local = local.In(location)
		t.Logf("%v %v %g %g\n", r.CreatedOn.Format("15:04"), local.Format("15:04"), r.Temperature, r.Humidity)
	}

}
