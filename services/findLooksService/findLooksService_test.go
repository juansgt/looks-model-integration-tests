package findLooksService_test

import (
	"context"
	"testing"

	"github.com/juansgt/generics/services"
	"github.com/juansgt/model-test/v3/dataAccess/lookRepository"
	"github.com/juansgt/model-test/v3/services/findLooksService"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestFindLooks_correctCalling_returnExpectedValues(t *testing.T) {

	// Arrange

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://atlasAdmin:Cripto0Virtual@cluster0.yolpv.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	database := client.Database("wap")

	assert := assert.New(t)
	var expectedLook *lookRepository.Look = lookRepository.NewLook("1", "Dress", "Bash")
	var findLooksQueryService services.IQueryServiceNoInput[[]lookRepository.Look] = findLooksService.NewFindLooksQueryService(lookRepository.NewLookRepositoryMongodb(database))

	// Act

	looks := findLooksQueryService.Execute()

	// Assert

	assert.Equal(expectedLook.Id(), looks[0].Id())
	assert.Equal(expectedLook.Name, looks[0].Name)
	assert.Equal(expectedLook.Brand, looks[0].Brand)
}
