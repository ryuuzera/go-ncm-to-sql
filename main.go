package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)


type Nomenclatura struct {
	Codigo      string `json:"Codigo"`
	Descricao   string `json:"Descricao"`
	DataInicio  string `json:"Data_Inicio"`
	DataFim     string `json:"Data_Fim"`
	TipoAtoIni  string `json:"Tipo_Ato_Ini"`
	NumeroAtoIni string `json:"Numero_Ato_Ini"`
	AnoAtoIni   string `json:"Ano_Ato_Ini"`
}

type NCMData struct {
	DataUltimaAtualizacaoNCM string        `json:"Data_Ultima_Atualizacao_NCM"`
	Ato                      string        `json:"Ato"`
	Nomenclaturas            []Nomenclatura `json:"Nomenclaturas"`
}

const batchSize = 100

func main() {
	url := "https://portalunico.siscomex.gov.br/classif/api/publico/nomenclatura/download/json"

	var resp, err = http.Get(url);

	if (err != nil) {
		log.Fatalf("failed to read response bodyÇ %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
	}

	var ncmData NCMData
	err = json.Unmarshal(body, &ncmData)
	if err != nil {
			log.Fatalf("Failed to parse JSON: %v", err)
	}

	db, err := sql.Open("sqlite3", "./ncm.db")
	if err != nil {
			log.Fatalf("Failed to open SQLite database: %v", err)
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS ncm (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		codigo TEXT,
		descricao TEXT,
		data_inicio TEXT,
		data_fim TEXT,
		tipo_ato_ini TEXT,
		numero_ato_ini TEXT,
		ano_ato_ini TEXT
		);`

		_, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }

		insertSQLBase := `INSERT INTO ncm (codigo, descricao, data_inicio, data_fim, tipo_ato_ini, numero_ato_ini, ano_ato_ini) VALUES `
	
		var values []string
	var params []interface{}

	for i, nome := range ncmData.Nomenclaturas {
		values = append(values, "(?, ?, ?, ?, ?, ?, ?)")
		params = append(params, nome.Codigo, nome.Descricao, nome.DataInicio, nome.DataFim, nome.TipoAtoIni, nome.NumeroAtoIni, nome.AnoAtoIni)


		if (i+1)%batchSize == 0 || i+1 == len(ncmData.Nomenclaturas) {
			insertSQL := insertSQLBase + strings.Join(values, ", ")

			_, err = db.Exec(insertSQL, params...)
			if err != nil {
				log.Fatalf("failed to execute batch insert: %v", err)
			}

			values = nil
			params = nil
		}
	}

    fmt.Println("Batch insert completed successfully.")

}