package repository

import (
	"encoding/json"
	. "geosqlwithcdn/module/API/repository/mysql/query"
	"geosqlwithcdn/module/API/repository/mysql/types"
	//static import
	"log"
)

func (repository *Repository) RegisterUser(user string, description string, hobby []string, latitude, longitude float64) error {
	if tx, err := repository.db.Begin(); err != nil {
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

		if result, err := tx.Exec(InsertIgnoreUserLocation, user, latitude, longitude, longitude, latitude); err != nil {
			tx.Rollback()
			return err
		} else {
			count, _ := result.RowsAffected()
			log.Println("Success to insert user location", "count", count)
		}

		tx.Commit()
	}

	return nil
}

func (repository *Repository) GetUser(userName string) (*types.User, error) {
	var res types.User

	var image interface{}
	var hobby interface{}

	if err := repository.db.QueryRow(GetUserByNone, userName).Scan(
		&res.UserName,
		&image,
		&res.Description,
		&hobby,
		&res.Latitude,
		&res.Longitude,
	); err != nil {
		return nil, err
	} else if err = unmarshalToFields(
		[]interface{}{image, hobby},
		&res.Image, &res.Hobby,
	); err != nil {
		return nil, err
	} else {
		return &res, nil
	}

}

func (repository *Repository) AroundUser(userName string, latitude, longitude float64, searchRange, limit int64) ([]*types.User, error) {
	if rows, err := repository.db.Query(
		GetAroundUsers,
		userName,
		longitude,
		latitude,
		searchRange,
		longitude,
		latitude,
		limit,
	); err != nil {
		return nil, err
	} else {
		defer rows.Close()

		var result []*types.User

		for rows.Next() {
			var res types.User

			var image interface{}
			var hobby interface{}

			if err = rows.Scan(
				&res.UserName,
				&image,
				&res.Description,
				&hobby,
				&res.Latitude,
				&res.Longitude,
			); err != nil {
				return nil, err
			} else if err = unmarshalToFields(
				[]interface{}{image, hobby},
				&res.Image, &res.Hobby,
			); err != nil {
				return nil, err
			} else {
				result = append(result, &res)
			}
		}
		return result, nil
	}
}
