package blockchain

import (
	"encoding/json"

	"github.com/duckcoding00/go-blockchain/constants"
	"github.com/syndtr/goleveldb/leveldb"
)

func Put(bc BlockchainStruct) error {
	db, err := leveldb.OpenFile(constants.Leveldb_Path, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	// save to db
	value, err := json.Marshal(bc)
	if err != nil {
		return err
	}
	if err := db.Put([]byte(constants.Blockchain_key), value, nil); err != nil {
		return err
	}

	return nil
}

func Get() (*BlockchainStruct, error) {
	db, err := leveldb.OpenFile(constants.Leveldb_Path, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var bc BlockchainStruct
	data, err := db.Get([]byte(constants.Blockchain_key), nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &bc); err != nil {
		return nil, err
	}

	return &bc, nil
}

func Exists() (bool, error) {
	db, err := leveldb.OpenFile(constants.Leveldb_Path, nil)
	if err != nil {
		return false, err
	}
	defer db.Close()

	return db.Has([]byte(constants.Blockchain_key), nil)
}
