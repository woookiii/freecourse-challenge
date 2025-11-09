package mysql

import (
	"encoding/json"
	. "geosqlwithcdn/db/mysql/query"
	//static import
	"log"
)

func (d *DB) RegisterUser(user string, description string, hobby []string, latitude float64, longitude float64) error {
	if tx, err := d.db.Begin(); err != nil {
		return err
	} else if json, err := json.Marshal(hobby); err != nil {
		return err
	} else {
		if result, err := tx.Exec(InsertIgnoreUser, user, description, json); err != nil {
			tx.Rollback()
			return err
		} else {
			count, _ := result.RowsAffected()
			log.Println("Success to insert user", "count", count)
		}

		if result, err := tx.Exec(InsertIgnoreUserLocation, user, latitude, longitude, latitude, longitude); err != nil {
			tx.Rollback()
			return err
		} else {
			count, _ := result.RowsAffected()
			log.Println("Success to insert user location", "count", count)
		}
	}

	return nil
}
