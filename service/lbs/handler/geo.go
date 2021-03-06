package handler

import (
	"context"
	"encoding/json"
	"fmt"

	mongoProxy "traffic-dispatcher/db/mongo"
	"traffic-dispatcher/model"
	"traffic-dispatcher/proto/lbs"

	"github.com/micro/go-micro/v2/util/log"
	h3 "github.com/uber/h3-go/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GeoLocation struct{}

var clientDBMap = map[model.ClientRole]string{
	model.ClientDriver:    "driverInfo",
	model.ClientPassenger: "passengerInfo",
}

func InsertGeo(resolution int, data model.WSMessage) {
	dbCli := mongoProxy.MongoConn()
	// 指定获取要操作的数据集
	collection := dbCli.Database(clientDBMap[data.Role]).Collection("geoInfo")

	geo := h3.GeoCoord{
		Latitude:  data.Geo.Lat,
		Longitude: data.Geo.Lng,
	}
	h3Index := h3.FromGeo(geo, resolution)
	h3IndexStr := fmt.Sprintf("%#x", h3Index)

	doc := bson.M{
		"$set": model.UserLocation{
			Name:    data.User.Name,
			UID:     data.User.UID,
			H3Index: h3IndexStr,
			GeoInfo: bson.M{
				"type":        "Point",
				"coordinates": []float64{data.Geo.Lng, data.Geo.Lat},
			},
		},
	}

	opts := options.Update().SetUpsert(true)
	insertResult, err := collection.UpdateOne(
		context.TODO(),
		bson.D{{"uid", data.User.UID}},
		doc,
		opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Inserted a single document: ", insertResult.UpsertedID)
}

func QueryGeo(lat float64, lng float64, role model.ClientRole) (res []model.UserLocation, err error) {
	dbCli := mongoProxy.MongoConn()
	// 指定获取要操作的数据集
	collection := dbCli.Database(clientDBMap[role]).Collection("geoInfo")

	stages := mongo.Pipeline{}
	getNearbyStage := bson.D{
		{"$geoNear", bson.M{
			"near": bson.M{
				"type":        "Point",
				"coordinates": []float64{lng, lat},
			},
			"maxDistance":   100000,
			"spherical":     true,
			"distanceField": "distance",
			"query":         bson.M{"lat": bson.M{"$ne": ""}, "lng": bson.M{"$ne": ""}},
		}}}

	stages = append(stages, getNearbyStage)

	filterCursor, err := collection.Aggregate(context.TODO(), stages)
	if err != nil {
		log.Error(err)
		return
	}
	for filterCursor.Next(context.TODO()) {
		var elem model.UserLocation
		err = filterCursor.Decode(&elem)
		if err != nil {
			log.Error(err)
			return
		}
		res = append(res, elem)
	}
	return
}

func (g *GeoLocation) ReportGeo(ctx context.Context, req *lbs.ReportRequest, resp *lbs.ReportResponse) error {
	var data model.WSMessage
	if err := json.Unmarshal(req.Data, &data); err == nil {
		InsertGeo(7, data)
		resp.Msg = "Hi " + req.Name
		return nil
	} else {
		resp.Msg = "Oooops..." + req.Name
		return err
	}
}

func (g *GeoLocation) QueryGeoNearby(ctx context.Context, req *lbs.QueryRequest, resp *lbs.QueryResponse) error {
	var data model.WSMessage
	var err error
	if err = json.Unmarshal(req.Data, &data); err == nil {
		if geolist, err := QueryGeo(data.QueryGeo.Lat, data.QueryGeo.Lng, data.QueryRole); err == nil {
			data, _ := json.Marshal(geolist)
			resp.Msg = "Hi " + req.Name
			resp.Data = data
			return nil
		}
	}
	resp.Msg = "Oooops..." + req.Name
	return err
}
