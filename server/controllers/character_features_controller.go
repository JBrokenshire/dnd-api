package controllers

import (
	"dnd-api/db/models"
	"dnd-api/server"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterFeaturesController struct {
	server.Server
}

func (c *CharacterFeaturesController) GetFeatures(ctx echo.Context) error {
	characterID := ctx.Param("id")

	character, err := c.Server.Stores.Character.Get(characterID)
	if err != nil || character.ID == 0 {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	classFeatures, err := c.Server.Stores.CharacterFeatures.GetClassFeatures(character.ClassID, character.Level)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	subclassFeatures := make([]*models.SubclassFeature, 0)
	if character.SubclassID != nil {
		subclassFeatures, err = c.Server.Stores.CharacterFeatures.GetSubclassFeatures(character.SubclassID, *character.SubclassID)
		if err != nil {
			return res.ErrorResponse(ctx, http.StatusNotFound, err)
		}
	}

	totalFeatureCount := len(classFeatures) + len(subclassFeatures)
	features := make([]*models.Feature, 0)

	// Combine all features into one slice, ordered by level
	for i := 0; i < totalFeatureCount; i++ {
		if len(classFeatures) == 0 {
			for _, subclassFeature := range subclassFeatures {
				features = append(features, &subclassFeature.Feature)
			}
			break
		}

		if len(subclassFeatures) == 0 {
			for _, classFeature := range classFeatures {
				features = append(features, &classFeature.Feature)
			}
			break
		}

		if classFeatures[0].Level <= subclassFeatures[0].Level {
			features = append(features, &classFeatures[0].Feature)
			classFeatures = classFeatures[1:]
		} else {
			features = append(features, &subclassFeatures[0].Feature)
			subclassFeatures = subclassFeatures[1:]
		}
	}

	return ctx.JSON(http.StatusOK, features)

}
