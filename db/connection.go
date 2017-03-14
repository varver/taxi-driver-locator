package db

import (
	"log"

	"gitlab.com/varver/wmd/config"
	"gitlab.com/varver/wmd/logger"
	mgo "gopkg.in/mgo.v2"
)

//DBChecks contains basic checks that has to be done on db
func DBChecks(s *mgo.Session) {
	var c *mgo.Collection
	if config.Setting.EnvMode == "test" {
		c = s.DB(config.Setting.TestDatabase).C(config.Setting.DriversLocationCollection)
	} else {
		c = s.DB(config.Setting.Database).C(config.Setting.DriversLocationCollection)
	}

	//Creating the indexes
	index := mgo.Index{
		Key:  []string{"$2dsphere:location"},
		Bits: 26,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		logger.Crit("Unable to ensure index for driver location collection.")
		panic(err)
	}
}

func Database() *mgo.Session {
	cluster := config.Setting.Cluster // mongodb host
	// connect to mongo
	session, err := mgo.Dial(cluster)
	if err != nil {
		log.Fatal("could not connect to db: ", err)
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	//session.DB(os.Getenv("DB_NAME"))

	DBChecks(session)

	return session

}
