package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetCharacterClassFeatures(t *testing.T) {
	ts.ClearTable("class_features")
	ts.ClearTable("subclass_features")
	ts.ClearTable("features")
	ts.ClearTable("subclasses")
	ts.ClearTable("classes")
	ts.ClearTable("characters")

	class := &models.Class{}
	factories.NewClass(ts.S.Db, class)
	subclass := &models.Subclass{}
	factories.NewSubclass(ts.S.Db, subclass)
	featureOne := &models.Feature{}
	factories.NewFeature(ts.S.Db, featureOne)
	classFeature := &models.ClassFeature{ClassID: class.ID, FeatureID: featureOne.ID}
	factories.NewClassFeature(ts.S.Db, classFeature)
	featureTwo := &models.Feature{}
	factories.NewFeature(ts.S.Db, featureTwo)
	subclassFeature := &models.SubclassFeature{SubclassID: subclass.ID, FeatureID: featureTwo.ID}
	factories.NewSubclassFeature(ts.S.Db, subclassFeature)
	character := &models.Character{ClassID: class.ID, SubclassID: &subclass.ID}
	factories.NewCharacter(ts.S.Db, character)

	noSubclass := &models.Character{ClassID: class.ID}
	factories.NewCharacter(ts.S.Db, noSubclass)

	cases := []helpers.TestCase{
		{
			TestName: "Can get class and subclass features for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/features", character.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, featureOne.Name),
					fmt.Sprintf(`"name":"%v"`, featureTwo.Name),
				},
			},
		},
		{
			TestName: "Can get class for character with no subclass",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/characters/%v/features", noSubclass.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, featureOne.Name),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, featureTwo.Name),
				},
			},
		},
		{
			TestName: "Can get 404 response for invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/features",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}
