package main

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Load(c string) []business {
	opt := option.WithCredentialsFile("onelocal-a9765-firebase-adminsdk-xsl8i-cc7b24903b.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	var businesses []business
	b := business{}
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	iter := client.Collection(c).Documents(ctx)
	for doc, err := iter.Next(); err == nil; doc, err = iter.Next() {
		// doc, err := iter.Next()
		// if err == iterator.Done {
		// 	break
		// }
		// if err != nil {
		// 	panic(err)
		// }
		b = business{
			Name:        doc.Data()["Name"].(string),
			ContactInfo: doc.Data()["Contactinfo"].(string),
			Address:     doc.Data()["Address"].(string),
			Description: doc.Data()["description"].(string),
			ImageLink:   doc.Data()["ImageLink"].(string),
		}

		businesses = append(businesses, b)
	}
	return businesses
}
func LoadAll() []area {
	opt := option.WithCredentialsFile("onelocal-a9765-firebase-adminsdk-xsl8i-cc7b24903b.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	iter := client.Collections(ctx)
	var areas []area
	for {
		collection, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		a := area{
			Name:             collection.ID,
			BusinessesInArea: Load(collection.ID),
		}
		areas = append(areas, a)
	}
	return areas
}
