package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3" // the db driver supports database/sql
	intf "github.com/sebmaspd/rnd/ieq/interfaces"
)

/*
https://linuxhint.com/install_sqlite_browser_ubuntu_1804/
sudo apt-get update
sudo apt-get install sqlite3
sqlite3 --version

cd to your runtime dir ~/go/src/github.com/sebmaspd/rnd/sql/sqlite
$ sqlite3 ieq.dbreturn owid column using its aliases: _rowid_ and oid.

go get github.com/mattn/go-sqlite3
Data types sqlite3
https://godoc.org/github.com/mattn/go-sqlite3
+------------------------------+
|go        | sqlite3           |
|----------|-------------------|
|nil       | null              |
|int       | integer           |
|int64     | integer           |
|float64   | float             |
|bool      | integer           |
|[]byte    | blob              |
|string    | text              |
|time.Time | timestamp/datetime|
+------------------------------+

DROP TABLE IF EXISTS devicestatus;
CREATE TABLE devicestatus (
	device_id VARCHAR(250),
	created_on datetime,
	status int,
	status_desc VARCHAR(1500)
	);

*/

// CreateDeviceStatus creates new record in database
func CreateDeviceStatus(dbname string, data intf.DeviceInfo) error {
	db := connect(dbname)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO devicestatus(device_id, created_on, status, status_desc) VALUES(?,strftime('%s', 'now'),?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.DeviceID, data.Status, data.StatusDescription)
	if err != nil {
		return err
	}
	fmt.Println("device status created in database")
	return nil
}

// ReadLastDeviceStatus returns the last record
func ReadLastDeviceStatus(dbname string) (data intf.DeviceInfo, err error) {
	//data = DeviceStatus{}
	db := connect(dbname)
	defer db.Close()

	stmt := "SELECT rowid, * FROM devicestatus ORDER BY rowid DESC LIMIT 1"

	rows, err := db.Query(stmt)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var rowid int
	var results []intf.DeviceInfo
	for rows.Next() {
		item := intf.DeviceInfo{}
		err2 := rows.Scan(&rowid, &item.DeviceID, &item.CreatedOn, &item.Status, &item.StatusDescription)
		if err2 != nil {
			return data, err
		}
		results = append(results, item)
	}

	// just the latest record
	data = results[0]
	return data, nil
}
