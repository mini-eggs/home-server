package database

import (
	"bytes"
	"os"

	"github.com/jinzhu/gorm"
)

// Connect - Get db connection
// must use defer db.Close()
// in client code.
func Connect() (*gorm.DB, error) {
	var details bytes.Buffer
	details.WriteString(os.Getenv("MYSQL_ROOT_USER"))
	details.WriteString(":")
	details.WriteString(os.Getenv("MYSQL_ROOT_PASSWORD"))
	details.WriteString("@tcp(")
	details.WriteString(os.Getenv("MYSQL_HOST"))
	details.WriteString(":")
	details.WriteString(os.Getenv("PMA_PORT"))
	details.WriteString(")/")
	details.WriteString(os.Getenv("MYSQL_DATABASE"))
	details.WriteString("?charset=utf8&parseTime=True&loc=Local")
	return gorm.Open("mysql", details.String())
}
