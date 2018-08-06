package db

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"

	bolt "github.com/coreos/bbolt"
)

var (
	Db   *bolt.DB
	Mmap map[string]string
	err  error
)

const demo string = `[{
    "id": 1,
    "label": "一级 1",
    "children": [{
      "id": 4,
      "label": "二级 1-1",
      "children": [{
        "id": 9,
        "label": "三级 1-1-1"
      }, {
        "id": 10,
        "label": "三级 1-1-2"
      }]
    }]
  }]`

func init() {
	Mmap = map[string]string{}
	Db, err = bolt.Open(beego.AppConfig.String("init::dbname"), 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	//init tables

	err = CreateBucket(beego.AppConfig.String("init::db"))
	if err != nil {
		// panic(err)
		beego.Critical(err)
	}
	err = CreateBucket(beego.AppConfig.String("init::data"))
	if err != nil {
		// panic(err)
		beego.Critical(err)
	}
	//初始化索引数据 即创建
	ds, _ := GetValueByBucketName(beego.AppConfig.String("init::data"), beego.AppConfig.String("init::data"))
	if len(ds) == 0 {
		AddKeyValueByBucketName(beego.AppConfig.String("init::data"), beego.AppConfig.String("init::data"), demo)
	}
	GetAllTables()
	// log.Println(Mmap)
	log.Println("init db logging success")
}

func CreateBucket(tablename string) error {
	err := Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(tablename))
		if err != nil {
			return fmt.Errorf("create bucket: %s ", err.Error())
		}
		//如果没有index表和index key则立
		return nil
	})
	AddTables(tablename)
	return err
}

func DeleteBucket(tablename string) error {
	err := Db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(tablename))
		if err != nil {
			return fmt.Errorf("delete bucket: %s ", err.Error())
		}
		return nil
	})
	DeleteTables(tablename)
	return err
}

func AddTables(tablename string) error {
	return Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(beego.AppConfig.String("init::db")))
		err := b.Put([]byte(tablename), []byte(tablename))
		Mmap[string(tablename)] = string(tablename)
		return err
	})
}

func GetAllTables() (map[string]string, error) {
	tmp := map[string]string{}
	err := Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(beego.AppConfig.String("init::db")))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tmp[string(k)] = string(v)
			Mmap[string(k)] = string(v)
		}
		return nil
	})
	return tmp, err
}

func DeleteTables(tablename string) error {
	return Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(beego.AppConfig.String("init::db")))
		err := b.Delete([]byte(tablename))
		delete(Mmap, tablename)
		return err
	})
}

func AddKeyValueByBucketName(table, key, value string) error {
	// fmt.Println(Mmap)
	if _, ok := Mmap[table]; !ok {
		fmt.Println(fmt.Printf("%s is not exist\n", table))
		CreateBucket(table)
	}
	// log.Println(Mmap)
	return Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

func AddKeyValueByBucketNameAuto(table, key, value string) error {

	return Db.Update(func(tx *bolt.Tx) error {
		var b *bolt.Bucket
		var err error
		if _, ok := Mmap[table]; !ok {
			b, err = tx.CreateBucket([]byte(table))
			if err != nil {
				return err
			}
		} else {
			b = tx.Bucket([]byte(table))
			if err != nil {
				return err
			}
		}

		err = b.Put([]byte(key), []byte(value))
		return err
	})
}

func GetValueByBucketName(table, key string) ([]byte, error) {
	var value []byte
	var err error
	Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if _, ok := Mmap[table]; ok {
			value = b.Get([]byte(key))
		} else {
			err = errors.New(fmt.Sprintf("table %s not exist", table))
		}
		return err
	})
	return value, err
}

func DeleteKeyValueByBucket(table, key string) error {
	return Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		err := b.Delete([]byte(key))
		return err
	})
}
