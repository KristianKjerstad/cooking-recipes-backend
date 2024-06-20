package database

import (
	"context"
	"cooking-recipes-backend/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client           *mongo.Client
	recipeCollection *mongo.Collection
}

func Connect() *DB {
	uri := "mongodb://admin:password@localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	recipeCollection := client.Database("data").Collection("recipes")

	return &DB{client: client, recipeCollection: recipeCollection}
}

func (db *DB) SaveRecipe(input *model.NewRecipeInput) *model.Recipe {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := db.recipeCollection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	ingredients := make([]*model.Ingredient, len(input.Ingredients))
	for i, ing := range input.Ingredients {
		ingredients[i] = &model.Ingredient{
			Name:     ing.Name,
			Unit:     ing.Unit,
			Quantity: ing.Quantity,
		}
	}

	return &model.Recipe{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Steps:       input.Steps,
		Ingredients: ingredients,
	}
}

func (db *DB) FindByID(ID string) *model.Dog {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := db.recipeCollection.FindOne(ctx, bson.M{"_id": ObjectID})

	dog := model.Dog{}
	res.Decode(&dog)

	return &dog
}

func (db *DB) All() []*model.Dog {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := db.recipeCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var dogs []*model.Dog
	for cur.Next(ctx) {
		var dog *model.Dog
		err := cur.Decode(&dog)
		if err != nil {
			log.Fatal(err)
		}
		dogs = append(dogs, dog)
	}
	return dogs
}
