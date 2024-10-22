package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetSubclassFeatures(t *testing.T) {
	ts.ClearTable("subclass_features")
	ts.ClearTable("features")
	ts.ClearTable("subclasses")

	class := &models.Class{}
	factories.NewClass(ts.S.Db, class)
	subclass := &models.Subclass{ClassID: class.ID}
	factories.NewSubclass(ts.S.Db, subclass)
	feature := &models.Feature{}
	factories.NewFeature(ts.S.Db, feature)
	subclassFeature := &models.SubclassFeature{SubclassID: subclass.ID, FeatureID: feature.ID}
	factories.NewSubclassFeature(ts.S.Db, subclassFeature)

	noFeatures := &models.Subclass{ClassID: class.ID}
	factories.NewSubclass(ts.S.Db, noFeatures)

	cases := []helpers.TestCase{
		{
			TestName: "Can get features for subclass",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/subclasses/%v/features", subclass.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, feature.Name),
				},
			},
		},
		{
			TestName: "Can get empty response for subclass with no features",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    fmt.Sprintf("/subclasses/%v/features", noFeatures.ID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "[]",
			},
		},
		{
			TestName: "Can get 404 response for invalid subclass id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/subclasses/invalid-id/features",
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
