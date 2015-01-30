package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Bean struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"Name"`
}

type Config struct {
	EnvId   string `bson:"EnvId"`
	Content bson.M `bson:"Content"`
}

func GetList(collection string) []interface{} {
	var list []interface{}
	exeDB(func(db *mgo.Database) {
		err := db.C(collection).Find(nil).All(&list)
		if err != nil {
			panic(err)
		}
	})
	return list
}
func New(collection string, content interface{}) string {
	exeDB(func(db *mgo.Database) {
		err := db.C(collection).Insert(content)
		if err != nil {
			panic(err)
		}
	})
	return "Success"
}
func Update(collection string, id string, content interface{}) string {
	exeDB(func(db *mgo.Database) {
		err := db.C(collection).Update(bson.M{"_id": id}, content)
		if err != nil {
			panic(err)
		}
	})
	return "Success"
}

func Delete(collection string, id string) string {
	exeDB(func(db *mgo.Database) {
		err := db.C(collection).Remove(bson.M{"_id": id})
		if err != nil {
			panic(err)
		}
	})
	return "Success"
}

func GetServices() []Bean {
	var services []Bean
	exeDB(func(db *mgo.Database) {
		err := db.C("Services").Find(nil).All(&services)
		if err != nil {
			panic(err)
		}
	})
	return services
}

func GetEnvs() []Bean {
	var envs []Bean
	exeDB(func(db *mgo.Database) {
		err := db.C("Envs").Find(nil).All(&envs)
		if err != nil {
			panic(err)
		}
	})
	return envs
}
func GetServiceConfig(serviceId string) []Config {
	var configs []Config
	exeDB(func(db *mgo.Database) {
		err := db.C("ServiceConfigs").
			Find(bson.M{"ServiceId": serviceId}).All(&configs)
		if err != nil {
			panic(err)
		}
	})
	return configs
}
func UpdateServiceConfig(serviceId, envId, content interface{}) error {
	log.Println("content", content)
	exeDB(func(db *mgo.Database) {
		err := db.C("ServiceConfigs").Update(
			bson.M{"ServiceId": serviceId, "EnvId": envId},
			bson.M{"$set": bson.M{"Content": content}})
		if err != nil {
			panic(err)
		}
	})
	return nil
}
func exeDB(fn func(*mgo.Database)) {
	session, err := mgo.Dial("chelappews006:27060")
	defer session.Close()
	if err != nil {
		panic("DB connect error.")
	}
	fn(session.DB("AccountMgmtDB"))
}
