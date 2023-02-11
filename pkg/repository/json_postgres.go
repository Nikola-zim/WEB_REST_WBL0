package repository

import (
	"WEB_REST_exm0302/static"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type JsonPostgres struct {
	db *sqlx.DB
}

func NewJsonPostgres(db *sqlx.DB) *JsonPostgres {
	return &JsonPostgres{db: db}
}

func (jp *JsonPostgres) WriteInDB(inputJson static.Json) error {
	//Старт транзакции
	tx, err := jp.db.Begin()
	if err != nil {
		return err
	}

	//Запись в order_main_info
	{
		//queryMainInfo := fmt.Sprintf("INSERT INTO %s (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", orderMainInfoTable)
		stmt, err := tx.Prepare(`INSERT INTO order_main_info (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`) //сам запрос по плейсхолдерам
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(inputJson.OrderUid, inputJson.TrackNumber, inputJson.Entry, inputJson.Locale, inputJson.InternalSignature, inputJson.CustomerId, inputJson.DeliveryService, inputJson.Shardkey, inputJson.SmId, inputJson.DateCreated, inputJson.OofShard)
		if err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			logrus.Println("Ошибка записи в orderMainInfoTable  ", err)
			return err
		}
	}

	//Запись в delivery
	{
		stmt, err := tx.Prepare(`INSERT INTO delivery (delivery_id, name, phone, zip, city, address, region, email) values ($1, $2, $3, $4, $5, $6, $7, $8)`) //сам запрос по плейсхолдерам
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(inputJson.OrderUid, inputJson.Delivery.Name, inputJson.Delivery.Phone, inputJson.Delivery.Zip, inputJson.Delivery.City, inputJson.Delivery.Address, inputJson.Delivery.Region, inputJson.Delivery.Email)
		if err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			logrus.Println("Ошибка записи в delivery: ", err)
			return err
		}
	}
	//Запись в payment
	{
		stmt, err := tx.Prepare(`INSERT INTO payment (payment_id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`) //сам запрос по плейсхолдерам
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(inputJson.OrderUid, inputJson.Payment.Transaction, inputJson.Payment.RequestId, inputJson.Payment.Currency, inputJson.Payment.Provider, inputJson.Payment.Amount, inputJson.Payment.PaymentDt, inputJson.Payment.Bank, inputJson.Payment.DeliveryCost, inputJson.Payment.GoodsTotal, inputJson.Payment.CustomFee)
		if err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			logrus.Println("Ошибка записи в delivery: ", err)
			return err
		}
	}

	//Запись в items
	{
		stmt, err := tx.Prepare(`INSERT INTO items (item_id, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) values ($1, $2, $3, $4, $5, $6,$7,$8, $9, $10, $11, $12)`) //сам запрос по плейсхолдерам
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		for index, value := range inputJson.Items {
			_, err = stmt.Exec(inputJson.OrderUid, value.ChrtId, value.TrackNumber, value.Price, value.Rid, value.Name, value.Sale, value.Size, value.TotalPrice, value.NmId, value.Brand, value.Status)
			if err != nil {
				tx.Rollback() // return an error too, we may want to wrap them
				logrus.Println("Ошибка записи в ItemsTable в элементе ", index, " : ", err)
				return err
			}
		}
	}

	tx.Commit()
	return nil
}

func (jp *JsonPostgres) ReadAllFromDB() (map[string]static.Json, error) {
	jsonMap := make(map[string]static.Json)
	var currentID string
	var queryFormat string

	rowsgetAllOrderUid, errOrderUid := jp.db.Query(`SELECT order_uid From order_main_info`)
	if errOrderUid != nil {
		logrus.Printf("Ошибка чтения столбца всех UID")
	}
	var dummy string     //для записи ключей которых не существует в структуре
	var indexItem uint32 //счетчик для запросов к items

	for rowsgetAllOrderUid.Next() {

		var tempJson static.Json
		errOrderUid := rowsgetAllOrderUid.Scan(&currentID)
		if errOrderUid != nil {
			logrus.Printf("Ошибка чтения UID")
		}
		//Чтение order_main_info
		queryFormat = fmt.Sprintf("SELECT * From %s WHERE order_uid = '%s'", "order_main_info", currentID)
		rowGetOrderUid, errOrderUid := jp.db.Query(queryFormat)
		if errOrderUid != nil {
			logrus.Printf("Ошибка чтения строки по UID в order_main_info")
		}

		for rowGetOrderUid.Next() {
			errOrderMain := rowGetOrderUid.Scan(&tempJson.OrderUid, &tempJson.TrackNumber, &tempJson.Entry, &tempJson.Locale, &tempJson.InternalSignature, &tempJson.CustomerId, &tempJson.DeliveryService, &tempJson.Shardkey, &tempJson.SmId, &tempJson.DateCreated, &tempJson.OofShard)
			if errOrderMain != nil {
				logrus.Printf("Ошибка записи в структуру из таблицы order_main_info. ID = %s", currentID)
			}
		}
		//Чтение delivery
		queryFormat = fmt.Sprintf("SELECT * From %s WHERE delivery_id = '%s'", "delivery", currentID)
		rowGetDelivery, errDeliveryUid := jp.db.Query(queryFormat)
		if errDeliveryUid != nil {
			logrus.Printf("Ошибка чтения строки по UID в delivery")
		}

		for rowGetDelivery.Next() {
			errOrderMain := rowGetDelivery.Scan(&dummy, &tempJson.Delivery.Name, &tempJson.Delivery.Phone, &tempJson.Delivery.Zip, &tempJson.Delivery.City, &tempJson.Delivery.Address, &tempJson.Delivery.Region, &tempJson.Delivery.Email)
			if errOrderMain != nil {
				logrus.Printf("Ошибка записи в структуру из таблицы delivery. ID = %s", currentID)
			}
		}

		//Чтение payment
		queryFormat = fmt.Sprintf("SELECT * From %s WHERE payment_id = '%s'", "payment", currentID)
		rowGetPayment, errPaymentUid := jp.db.Query(queryFormat)
		if errPaymentUid != nil {
			logrus.Printf("Ошибка чтения строки по UID в payment")
		}

		for rowGetPayment.Next() {
			errPaymentMain := rowGetPayment.Scan(&dummy, &tempJson.Payment.Transaction, &tempJson.Payment.RequestId, &tempJson.Payment.Currency, &tempJson.Payment.Provider, &tempJson.Payment.Amount, &tempJson.Payment.PaymentDt, &tempJson.Payment.Bank, &tempJson.Payment.DeliveryCost, &tempJson.Payment.GoodsTotal, &tempJson.Payment.CustomFee)
			if errPaymentMain != nil {
				logrus.Printf("Ошибка записи в структуру из таблицы payment. ID = %s", currentID)
			}
		}

		//Чтение items
		queryFormat = fmt.Sprintf("SELECT * From %s WHERE item_id = '%s'", "items", currentID)
		rowGetItems, errItemsUid := jp.db.Query(queryFormat)
		if errItemsUid != nil {
			logrus.Printf("Ошибка чтения строк по UID в items")
		}
		indexItem = 0
		for rowGetItems.Next() {
			tempJson.Items = append(tempJson.Items, struct {
				ChrtId      int    "json:\"chrt_id\""
				TrackNumber string "json:\"track_number\""
				Price       int    "json:\"price\""
				Rid         string "json:\"rid\""
				Name        string "json:\"name\""
				Sale        int    "json:\"sale\""
				Size        string "json:\"size\""
				TotalPrice  int    "json:\"total_price\""
				NmId        int    "json:\"nm_id\""
				Brand       string "json:\"brand\""
				Status      int    "json:\"status\""
			}{ChrtId: 666})
			errItems := rowGetItems.Scan(&dummy, &tempJson.Items[indexItem].ChrtId, &tempJson.Items[indexItem].TrackNumber, &tempJson.Items[indexItem].Price, &tempJson.Items[indexItem].Rid, &tempJson.Items[indexItem].Name, &tempJson.Items[indexItem].Sale, &tempJson.Items[indexItem].Size, &tempJson.Items[indexItem].TotalPrice, &tempJson.Items[indexItem].NmId, &tempJson.Items[indexItem].Brand, &tempJson.Items[indexItem].Status)
			//errItems := rowGetItems.Scan(&dummy, tempJson.Items[indexItem].ChrtId, tempJson.Items[indexItem].TrackNumber, tempJson.Items[indexItem].Price, tempJson.Items[indexItem].Rid, tempJson.Items[indexItem].Name, tempJson.Items[indexItem].Sale, tempJson.Items[indexItem].Size, tempJson.Items[indexItem].TotalPrice, tempJson.Items[indexItem].NmId, tempJson.Items[indexItem].Brand, tempJson.Items[indexItem].Status)
			if errItems != nil {
				logrus.Printf("Ошибка записи в структуру из таблицы payment. ID = %s", currentID)
			}
			indexItem++
		}
		jsonMap[currentID] = tempJson
	}

	return jsonMap, nil
}
