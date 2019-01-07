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

// TestUser_Find3Table ...
func TestUser_Find3Table(t *testing.T) {
	// 3 表查询
	//  select *
	//	from user left join role_user on role_user.userid = user._id left join role on
	//	role._id = role_user.roleid
	//	where user._id = ObjectId('5c33711e06b5362b5f8dccbf')
	//	db.user.aggregate(
	//		[
	//		{
	//			"$project" : {
	//				"_id" : NumberInt(0),
	//				"user" : "$$ROOT"
	//			}
	//		},
	//	{
	//		"$lookup" : {
	//		"localField" : "user._id",
	//			"from" : "role_user",
	//			"foreignField" : "userid",
	//			"as" : "role_user"
	//	}
	//	},
	//	{
	//		"$unwind" : {
	//		"path" : "$role_user",
	//			"preserveNullAndEmptyArrays" : true
	//	}
	//	},
	//	{
	//		"$lookup" : {
	//		"localField" : "role_user.roleid",
	//			"from" : "role",
	//			"foreignField" : "_id",
	//			"as" : "role"
	//	}
	//	},
	//	{
	//		"$unwind" : {
	//		"path" : "$role",
	//			"preserveNullAndEmptyArrays" : true
	//	}
	//	},
	//	{
	//		"$match" : {
	//		"user._id" : ObjectId("5c33711e06b5362b5f8dccbf")
	//	}
	//	}
	//],
	//	{
	//	"allowDiskUse" : true
	//	}
	//	);

	user := NewUser()
	user.ID = ID("5c33711e06b5362b5f8dccbf")
	//find, e := C(user._Name()).Find(mgo.TimeOut())
	//find.
	cursor, err := C(user._Name()).Aggregate(mgo.TimeOut(),
		mongo.Pipeline{
			[]primitive.E{
				{
					Key: "$match",
					Value: primitive.E{
						Key:   "user._id",
						Value: ID("5c33711e06b5362b5f8dccbf"),
					},
				},
				{
					Key: "$lookup",
					Value: &RelateInfo{
						From:         "role_user",
						LocalField:   "user._id",
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
		//if len(user.RoleUsers) > 0 {
		//	log.Printf("%+v", user.RoleUsers[0])
		//}
		log.Println(v, err)
	}
}
