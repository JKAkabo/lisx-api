package repos

import (
	"lisxAPI/db"
	"lisxAPI/models"
)

func InsertAnalyzer(
	name string,
	protocol string,
	ip string,
	port int,
	serverMode bool,
	serialPort string,
	baudRate int,
	parity int,
	dataBits int,
	stopBits int,
	startDelimiter string,
	endDelimiter string,
	userId int,
) (int, error) {
	row := db.DB.QueryRow(
		"insert into analyzer (name, protocol, ip, port, server_mode, serial_port, baud_rate, parity, data_bits, stop_bits, start_delimiter, end_delimiter, user_id) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) returning id",
		name,
		protocol,
		ip,
		port,
		serverMode,
		serialPort,
		baudRate,
		parity,
		dataBits,
		stopBits,
		startDelimiter,
		endDelimiter,
		userId,
	)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func SelectAnalyzerById(id int) (analyzer models.Analyzer, err error) {
	err = db.DB.Get(&analyzer, "select * from analyzer where id = $1", id)
	if err != nil {
		return analyzer, err
	}
	return analyzer, nil
}

func SelectAnalyzers() (analyzers []models.Analyzer, err error) {
	analyzers = []models.Analyzer{}
	err = db.DB.Select(&analyzers, "select * from analyzer")
	if err != nil {
		return analyzers, err
	}
	return analyzers, nil
}
