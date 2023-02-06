package repository

import (
	"WEB_REST_exm0302"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type JsonPostgres struct {
	db *sqlx.DB
}

func NewJsonPostgres(db *sqlx.DB) *JsonPostgres {
	return &JsonPostgres{db: db}
}

func (jp *JsonPostgres) WriteInDB(inputJson WEB_REST_exm0302.Json) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id", orderMainInfoTable)
	row := jp.db.QueryRow(query, inputJson.OrderUid, inputJson.TrackNumber, inputJson.Entry, inputJson.Locale, inputJson.InternalSignature, inputJson.CustomerId, inputJson.DeliveryService, inputJson.Shardkey, inputJson.SmId, inputJson.DateCreated, inputJson.OofShard) //сам запрос по плейсхолдерам

	if err := row.Scan(&id); err != nil {
		fmt.Print("Ошибка записи в БД   ")
		fmt.Println(err)
		return err
	}

	return nil
}

func (jp *JsonPostgres) ReadFromDB() {
	//TODO implement me
	panic("implement me")
}
