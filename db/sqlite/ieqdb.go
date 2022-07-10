package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3" // the db driver supports database/sql
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

DROP TABLE IF EXISTS metrics;
CREATE TABLE metrics (
	device_id VARCHAR(250),
	created_on datetime,
	temperature float,
	humidity float,
	co2 float,
	voc float,
	pm25 float,
	lighting float,
	noise float
	);

# no FK to rawmetrics yet
DROP TABLE IF EXISTS metricscores;
CREATE TABLE metricscores (
	device_id VARCHAR(250),
	created_on datetime,
	temperature float,
	humidity float,
	co2 float,
	voc float,
	pm25 float,
	lighting float,
	noise float
	);

# no FK to metricscores yet
DROP TABLE IF EXISTS ieqscores;
CREATE TABLE ieqscores (
	device_id VARCHAR(250),
	created_on datetime,
	scheme VARCHAR(250),
	thermal float,
	iaq float,
	lighting float,
	noise float,
	overall float
	);
ALTER TABLE ieqscores
	ADD thermal_weighting float;
ALTER TABLE ieqscores
	ADD iaq_weighting float;
ALTER TABLE ieqscores
	ADD lighting_weighting float;
ALTER TABLE ieqscores
	ADD noise_weighting float;

select * from metrics;
select * from metricscores;
select * from ieqscores;

delete from metrics;
delete from metricscores;
delete from ieqscores;

SELECT rowid, * FROM metrics ORDER BY rowid DESC LIMIT 1;
SELECT rowid, * FROM metricscores ORDER BY rowid DESC LIMIT 1;
SELECT rowid, * FROM ieqscores ORDER BY rowid DESC LIMIT 1;
SELECT datetime(created_on, 'unixepoch', 'localtime'), lighting, noise as created_on FROM metrics ORDER BY rowid DESC LIMIT 10;

# just a error log table. no FK to metrics yet
DROP TABLE IF EXISTS errorlogs;
CREATE TABLE errorlogs (
	device_id VARCHAR(250),
	created_on datetime,
	errordesc VARCHAR(1000),
	);

*/

// Metric database layer struct
type Metric struct {
	DeviceID    string
	CreatedOn   time.Time
	Temperature float64
	Humidity    float64
	CO2         float64
	VOC         float64
	PM25        float64
	Lighting    float64
	Noise       float64
}

// MetricScore database layer struct
type MetricScore struct {
	DeviceID    string
	CreatedOn   time.Time
	Temperature float64
	Humidity    float64
	CO2         float64
	VOC         float64
	PM25        float64
	Lighting    float64
	Noise       float64
}

// IeqScore database layer struct
type IeqScore struct {
	DeviceID          string
	CreatedOn         time.Time
	Scheme            string
	Thermal           float64
	ThermalWeighting  float64
	IAQ               float64
	IAQWeighting      float64
	Lighting          float64
	LightingWeighting float64
	Noise             float64
	NoiseWeighting    float64
	Overall           float64
}

// TODO: refactor to common
// connect opens and return the db object. Caller must close db when done with it.
func connect(dbName string) (db *sql.DB) {
	// mysql
	/*
		dbDriver := "mysql"
		dbUser := "root"
		dbPass := "root"
		dbName := "goblog"
		db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	*/
	// sqlite
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// CreateMetric creates metrics record in database
func CreateMetric(dbname string, data Metric) error {
	db := connect(dbname)
	defer db.Close()

	//log.Printf("Person: %s, %s, %v \n", person.Email, person.Job, person.Interests)
	stmt, err := db.Prepare("INSERT INTO metrics(device_id, created_on, temperature, humidity, co2, voc, pm25, lighting, noise) VALUES(?,strftime('%s', 'now'),?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.DeviceID, data.Temperature, data.Humidity, data.CO2, data.VOC, data.PM25, data.Lighting, data.Noise)
	if err != nil {
		return err
	}
	fmt.Println("metrics created in database")
	return nil
}

// CreateMetricScore creates metricscores record in database
func CreateMetricScore(dbname string, data MetricScore) error {
	db := connect(dbname)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO metricscores(device_id, created_on, temperature, humidity, co2, voc, pm25, lighting, noise) VALUES(?,strftime('%s', 'now'),?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.DeviceID, data.Temperature, data.Humidity, data.CO2, data.VOC, data.PM25, data.Lighting, data.Noise)
	if err != nil {
		return err
	}

	fmt.Println("metricscores created in database")
	return nil
}

// CreateIeqScore creates ieqscores record in database
func CreateIeqScore(dbname string, data IeqScore) error {
	db := connect(dbname)
	defer db.Close()

	//ieqscore.deviceID = "awair-omni_18453"
	//ieqscore.scheme = "SPD"

	stmt, err := db.Prepare("INSERT INTO ieqscores(device_id, created_on, scheme, thermal, iaq, lighting, noise, overall, thermal_weighting, iaq_weighting, lighting_weighting, noise_weighting) VALUES(?,strftime('%s', 'now'),?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		//panic(err.Error())
		return err
	}
	_, err = stmt.Exec(data.DeviceID, data.Scheme, data.Thermal, data.IAQ, data.Lighting, data.Noise, data.Overall, data.ThermalWeighting, data.IAQWeighting, data.LightingWeighting, data.NoiseWeighting)
	if err != nil {
		return err
	}

	fmt.Println("ieqscores created in database")
	return nil
}

// ReadLatestMetrics returns the latest metrics record
func ReadLatestMetrics(dbname string) (data Metric, err error) {
	data = Metric{}
	db := connect(dbname)
	defer db.Close()

	stmt := "SELECT rowid, * FROM metrics ORDER BY rowid DESC LIMIT 1"

	rows, err := db.Query(stmt)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var rowid int
	var results []Metric
	for rows.Next() {
		item := Metric{}
		err2 := rows.Scan(&rowid, &item.DeviceID, &item.CreatedOn, &item.Temperature, &item.Humidity, &item.CO2, &item.VOC, &item.PM25, &item.Lighting, &item.Noise)
		if err2 != nil {
			return data, err
		}
		results = append(results, item)
	}

	// just the latest record
	data = results[0]
	return data, nil
}

// ReadLatestMetricScores returns the latest metricscores record
func ReadLatestMetricScores(dbname string) (data MetricScore, err error) {
	data = MetricScore{}
	db := connect(dbname)
	defer db.Close()

	stmt := "SELECT rowid, * FROM metricscores ORDER BY rowid DESC LIMIT 1"

	rows, err := db.Query(stmt)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var rowid int
	var results []MetricScore
	for rows.Next() {
		item := MetricScore{}
		err2 := rows.Scan(&rowid, &item.DeviceID, &item.CreatedOn, &item.Temperature, &item.Humidity, &item.CO2, &item.VOC, &item.PM25, &item.Lighting, &item.Noise)
		if err2 != nil {
			return data, err
		}
		results = append(results, item)
	}

	// just the latest record
	data = results[0]
	return data, nil
}

// ReadLatestIeqScores returns the latest ieqscores record
func ReadLatestIeqScores(dbname string) (data IeqScore, err error) {
	data = IeqScore{}
	db := connect(dbname)
	defer db.Close()

	stmt := "SELECT rowid, * FROM ieqscores ORDER BY rowid DESC LIMIT 1"

	rows, err := db.Query(stmt)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var rowid int
	var results []IeqScore
	for rows.Next() {
		item := IeqScore{}
		err2 := rows.Scan(&rowid, &item.DeviceID, &item.CreatedOn, &item.Scheme, &item.Thermal, &item.IAQ, &item.Lighting, &item.Noise, &item.Overall, &item.ThermalWeighting, &item.IAQWeighting, &item.LightingWeighting, &item.NoiseWeighting)
		if err2 != nil {
			return data, err
		}
		results = append(results, item)
	}

	// just the latest record
	data = results[0]
	return data, nil
}

// ReadMetrics returns just the metrics record
func ReadMetrics(dbname string, count int) (results []Metric, err error) {
	results = []Metric{}
	db := connect(dbname)
	defer db.Close()

	stmt := "SELECT rowid, * FROM metrics ORDER BY rowid DESC LIMIT " + strconv.Itoa(count)

	rows, err := db.Query(stmt)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	var rowid int
	for rows.Next() {
		item := Metric{}
		err2 := rows.Scan(&rowid, &item.DeviceID, &item.CreatedOn, &item.Temperature, &item.Humidity, &item.CO2, &item.VOC, &item.PM25, &item.Lighting, &item.Noise)
		if err2 != nil {
			return results, err
		}
		results = append(results, item)
	}

	return results, nil
}

/*
func main() {
	args := os.Args
	if len(args) != 2 {
		log.Println("arguments expected: sqlite db file name")
		return
	}

	dbname := args[1]

	log.Printf("Persisting to %s...\n", dbname)
	doPersist(dbname)
	log.Printf("Retrieving to %s...\n", dbname)
	doReadMetrics(dbname)
	doReadMetricScores(dbname)
	doReadIeqScores(dbname)

	log.Println("Done.")
}
*/
