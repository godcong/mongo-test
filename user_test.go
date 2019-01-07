package mongodb

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"mongodb/util"
	"testing"
)

var user = User{
	Name:          "godcong",
	Username:      "ungodcong",
	Email:         "godcong@ggg.com",
	Mobile:        "123456",
	IDCardFacade:  "/d/d/e/e/d/c/",
	IDCardObverse: "/f/g/h/j/a",
	Password:      "godcong0910",
	Token:         "1212133333",
}

// TestUser_Create ...
func TestUser_Create(t *testing.T) {
	user := NewUser()
	user.Username = "godcong"
	user.Name = util.GenerateRandomString(32)
	t.Log(user.Create())
	t.Log(user)
}

// TestUser_Delete ...
func TestUser_Delete(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("5c2eeb5bb69c469e69c79a26")
	user := User{
		ID: id,
	}
	//t.Log(user.Delete())
	user.SetSoftDelete(true)
	e := user.Delete()
	t.Log(e)
	t.Log(user)

	e = user.Find()
	t.Log(e)
	t.Log(user)

}

// TestUser_Update ...
func TestUser_Update(t *testing.T) {
	user := User{
		ID: ID("5c2eea9a3db6598a9c25c65c"),
	}
	user.softDelete = true
	user.Find()

	user.Username = "SSSSSSSSSSSSSSSSSSSS"
	err := user.Update()
	t.Log(err)
	t.Log(user)
}

// TestUser_Find ...
func TestUser_Find(t *testing.T) {
	user := NewUser()
	user.ID = ID("5c2eeb95761de4f5a13b3b83")
	e := user.Find()
	t.Log(user, e)
}

// TestRoleUser_Find ...
func TestRoleUser_Find(t *testing.T) {
	ru := NewRoleUser()
	ru.UserID = ID("5c2eeb95761de4f5a13b3b83")
	ru.RoleID = ID("5c2f2864451279e9ff6f2128")
	e := ru.Find()
	t.Log(e, ru)
	t.Log(ru.User())

}

// TestFindGenesis ...
func TestFindGenesis(t *testing.T) {
	t.Log(FindGenesis())
}

// TestUser_Find2 ...
func TestUser_Find2(t *testing.T) {
	user := NewUser()
	user.ID = ID("5c33711e06b5362b5f8dccbf")
	cursor, err := C(user._Name()).Aggregate(mgo.TimeOut(),
		mongo.Pipeline{
			[]primitive.E{
				{
					Key: "$lookup",
					Value: &RelateInfo{
						From:         "role_user",
						LocalField:   "_id",
						ForeignField: "userid",
						//Pipeline: mongo.Pipeline{
						//	[]primitive.E{
						//		{
						//			Key:   "$match",
						//			Value: bson.M{"_id": ID("5c3371da40f8748192f0f39e")},
						//		},
						//	},
						//},
						As: "role_user",
					},
				},
			},
		})
	//"$lookup": bson.M{
	//	"from":         "role_user",
	//	"localField":   "_id",
	//	"foreignField": "userid",
	//	"as":           "ru",
	//},
	//ru := NewRoleUser()
	for cursor.Next(mgo.TimeOut()) {
		v := map[string]interface{}{}
		err = cursor.Decode(&v)
		if len(user.RoleUsers) > 0 {
			log.Printf("%+v", user.RoleUsers[0])
		}
		log.Println(v, err)
	}

}
