package models

type AnalyzerCreate struct {
	Name           string `json:"name"`
	Protocol       string `json:"protocol" validate:"oneof='tcp' 'serial'"`
	IP             string `json:"ip"`
	Port           int    `json:"port"`
	ServerMode     bool   `json:"serverMode"`
	SerialPort     string `json:"serialPort"`
	BaudRate       int    `json:"baudRate"`
	Parity         int    `json:"parity"`
	DataBits       int    `json:"dataBits"`
	StopBits       int    `json:"stopBits"`
	StartDelimiter string `json:"startDelimiter"`
	EndDelimiter   string `json:"endDelimiter"`
}
type Analyzer struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Protocol       string `json:"protocol" db:"protocol"`
	IP             string `json:"ip" db:"ip"`
	Port           int    `json:"port" db:"port"`
	ServerMode     bool   `json:"serverMode" db:"server_mode"`
	SerialPort     string `json:"serialPort" db:"serial_port"`
	BaudRate       int    `json:"baudRate" db:"baud_rate"`
	Parity         int    `json:"parity" db:"parity"`
	DataBits       int    `json:"dataBits" db:"data_bits"`
	StopBits       int    `json:"stopBits" db:"stop_bits"`
	StartDelimiter string `json:"startDelimiter" db:"start_delimiter"`
	EndDelimiter   string `json:"endDelimiter" db:"end_delimiter"`
}
