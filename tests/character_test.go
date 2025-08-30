package tests

import (
	"dnd-api/api/requests"
	"dnd-api/db/factories"
	m "dnd-api/db/models"
	"dnd-api/tests/helpers"
	"dnd-api/tests/mocks"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCharacter_List(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Create class
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// Create race
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	// Create Characters
	character := &m.Character{ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, character)
	character2 := &m.Character{ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, character2)
	namedCharacter := &m.Character{Name: "Test Character", ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, namedCharacter)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	// Create profile pictures
	profilePicture := &m.File{Model: m.FileModelCharacterProfilePicture, ModelId: character.ID}
	factories.NewFile(ts.S.Db, profilePicture)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/characters%v", query),
		}
	}

	permissionRequest := getRequest("")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can get characters",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"filename":"%v"`, profilePicture.Filename),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, differentUserCharacter.Name),
				},
			},
		},
		{
			Name:    "Can get page 0 of characters",
			Request: getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"filename":"%v"`, profilePicture.Filename),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, differentUserCharacter.Name),
				},
			},
		},
		{
			Name:    "Can get page 1 of characters",
			Request: getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"filename":"%v"`, profilePicture.Filename),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, differentUserCharacter.Name),
				},
			},
		},
		{
			Name:    "Can filter characters by name",
			Request: getRequest("?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					`"total_count":1`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"filename":"%v"`, profilePicture.Filename),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, differentUserCharacter.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestCharacter_Get(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")
	ts.ClearTable("files")
	ts.ClearTable("character_proficient_skills")
	ts.ClearTable("character_defenses")
	ts.ClearTable("backgrounds")
	ts.ClearTable("spells")
	ts.ClearTable("character_spells")
	ts.SetupDefaultUsers()

	// Create class
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// Create class images
	backgroundImage := &m.File{Model: m.FileModelClassBackgroundImage, ModelId: class.ID}
	factories.NewFile(ts.S.Db, backgroundImage)

	// Create race
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	// Create background
	background := &m.Background{}
	factories.NewBackground(ts.S.Db, background)

	// Create characters
	character := &m.Character{ClassId: class.ID, RaceId: race.ID, BackgroundId: background.ID}
	factories.NewCharacter(ts.S.Db, character)
	character2 := &m.Character{ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, character2)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	// Create profile pictures
	profilePicture := &m.File{Model: m.FileModelCharacterProfilePicture, ModelId: character.ID}
	factories.NewFile(ts.S.Db, profilePicture)

	// Create proficient skills
	characterProficientSkill := &m.CharacterProficientSkill{CharacterID: character.ID, Skill: m.ValidSkills[0]}
	factories.NewCharacterProficientSkill(ts.S.Db, characterProficientSkill)
	character2ProficientSkill := &m.CharacterProficientSkill{CharacterID: character2.ID, Skill: m.ValidSkills[1]}
	factories.NewCharacterProficientSkill(ts.S.Db, character2ProficientSkill)
	differentUserCharacterProficientSkill := &m.CharacterProficientSkill{CharacterID: differentUserCharacter.ID, Skill: m.ValidSkills[2]}
	factories.NewCharacterProficientSkill(ts.S.Db, differentUserCharacterProficientSkill)

	// Create defenses
	characterDefense := &m.CharacterDefense{CharacterID: character.ID, DamageType: m.DamageTypeFire, DefenseType: m.DefenseTypeResistance}
	factories.NewCharacterDefense(ts.S.Db, characterDefense)
	character2Defense := &m.CharacterDefense{CharacterID: character2.ID, DamageType: m.DamageTypeCold, DefenseType: m.DefenseTypeImmunity}
	factories.NewCharacterDefense(ts.S.Db, character2Defense)
	differentUserCharacterDefense := &m.CharacterDefense{CharacterID: differentUserCharacter.ID, DamageType: m.DamageTypeBludgeoning, DefenseType: m.DefenseTypeVulnerability}
	factories.NewCharacterDefense(ts.S.Db, differentUserCharacterDefense)

	// Create spells
	spell := &m.Spell{}
	factories.NewSpell(ts.S.Db, spell)
	spell2 := &m.Spell{}
	factories.NewSpell(ts.S.Db, spell2)

	// Create character spell links
	characterSpell := &m.CharacterSpell{CharacterID: character.ID, SpellID: spell.ID}
	factories.NewCharacterSpell(ts.S.Db, characterSpell)
	character2Spell := &m.CharacterSpell{CharacterID: character2.ID, SpellID: spell2.ID}
	factories.NewCharacterSpell(ts.S.Db, character2Spell)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/characters/%v", id),
		}
	}

	permissionRequest := getRequest(character.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't get character that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't get character with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't get character that belongs to a different user",
			Request: getRequest(differentUserCharacter.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can get character",
			Request: getRequest(character.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"filename":"%v"`, profilePicture.Filename),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"filename":"%v"`, backgroundImage.Filename),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					characterProficientSkill.Skill,
					characterDefense.DamageType,
					characterDefense.DefenseType,
					fmt.Sprintf(`"name":"%v"`, background.Name),
					fmt.Sprintf(`"name":"%v"`, spell.Name),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, differentUserCharacter.Name),
					character2ProficientSkill.Skill,
					differentUserCharacterProficientSkill.Skill,
					character2Defense.DamageType,
					character2Defense.DefenseType,
					differentUserCharacterDefense.DamageType,
					differentUserCharacterDefense.DefenseType,
					fmt.Sprintf(`"name":"%v"`, spell2.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestCharacter_Create(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")
	ts.SetupDefaultUsers()

	// Create class
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// Create race
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/characters",
	}

	RunNoAuthenticationTests(t, request.Method, request.Url)

	cases := []helpers.TestCase{
		{
			Name:        "Can't create character without required fields",
			Request:     request,
			RequestBody: requests.CreateCharacterRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name is a required field",
					"ClassId is a required field",
					"RaceId is a required field",
					"Level is a required field",
					"Strength is a required field",
					"Dexterity is a required field",
					"Constitution is a required field",
					"Intelligence is a required field",
					"Wisdom is a required field",
					"Charisma is a required field",
					"AdvancementType is a required field",
					"HitPointType is a required field",
				},
			},
		},
		{
			Name:    "Can't create character if fields exceed max length",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:            string(make([]byte, 201)),
				Pronouns:        string(make([]byte, 33)),
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name must be a maximum of 200 characters in length",
					"Pronouns must be a maximum of 32 characters in length",
				},
			},
		},
		{
			Name:    "Can't create character with level that is above 20",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           21,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Level must be 20 or less",
				},
			},
		},
		{
			Name:    "Can't create character with stats that are below 3",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        2,
				Dexterity:       2,
				Constitution:    2,
				Intelligence:    2,
				Wisdom:          2,
				Charisma:        2,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Strength must be 3 or greater",
					"Dexterity must be 3 or greater",
					"Constitution must be 3 or greater",
					"Intelligence must be 3 or greater",
					"Wisdom must be 3 or greater",
					"Charisma must be 3 or greater",
				},
			},
		},
		{
			Name:    "Can't create character with stats that are above 30",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        31,
				Dexterity:       31,
				Constitution:    31,
				Intelligence:    31,
				Wisdom:          31,
				Charisma:        31,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Strength must be 30 or less",
					"Dexterity must be 30 or less",
					"Constitution must be 30 or less",
					"Intelligence must be 30 or less",
					"Wisdom must be 30 or less",
					"Charisma must be 30 or less",
				},
			},
		},
		{
			Name:    "Can't create character with level that is below 1",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           -1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Level must be 1 or greater",
				},
			},
		},
		{
			Name:    "Can't create character with class id",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:            "Test Name",
				ClassId:         1000,
				RaceId:          race.ID,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't create character with race id",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          1000,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't create character with invalid advancement type",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: "invalid advancement type",
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "AdvancementType is not valid",
			},
		},
		{
			Name:    "Can't create character with invalid hit point type",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    "invalid hit point type",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "HitPointType is not valid",
			},
		},
		{
			Name:    "Can create character",
			Request: request,
			RequestBody: requests.CreateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Pronouns:        "He/Him",
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, "Test Name"),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"user_id":%v`, ts.AdminUser.ID),
					`"strength":10`,
					`"dexterity":10`,
					`"constitution":10`,
					`"intelligence":10`,
					`"wisdom":10`,
					`"charisma":10`,
					`"pronouns":"He/Him"`,
					`"level":1`,
					`"advancement_type":"Milestone"`,
					`"hit_point_type":"Fixed"`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Character was created",
					Model: m.Character{
						UserId:          ts.AdminUser.ID,
						Name:            "Test Name",
						ClassId:         class.ID,
						RaceId:          race.ID,
						Pronouns:        "He/Him",
						Level:           1,
						Strength:        10,
						Dexterity:       10,
						Constitution:    10,
						Intelligence:    10,
						Wisdom:          10,
						Charisma:        10,
						AdvancementType: m.AdvancementTypeMilestone,
						HitPointType:    m.HitPointTypeFixed,
					},
					CountExpected: 1,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestCharacter_Update(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")
	ts.SetupDefaultUsers()

	// Create classes
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)
	class2 := &m.Class{}
	factories.NewClass(ts.S.Db, class2)

	// Create races
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)
	race2 := &m.Race{}
	factories.NewRace(ts.S.Db, race2)

	// Create characters
	character := &m.Character{RaceId: race.ID, ClassId: class.ID}
	factories.NewCharacter(ts.S.Db, character)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPut,
			Url:    fmt.Sprintf("/characters/%v", id),
		}
	}

	permissionRequest := getRequest(character.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:        "Can't update character without required fields",
			Request:     getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name is a required field",
					"ClassId is a required field",
					"RaceId is a required field",
					"Level is a required field",
					"Strength is a required field",
					"Dexterity is a required field",
					"Constitution is a required field",
					"Intelligence is a required field",
					"Wisdom is a required field",
					"Charisma is a required field",
					"AdvancementType is a required field",
					"HitPointType is a required field",
				},
			},
		},
		{
			Name:    "Can't update character if fields exceed max length",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            string(make([]byte, 201)),
				Pronouns:        string(make([]byte, 33)),
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Name must be a maximum of 200 characters in length",
					"Pronouns must be a maximum of 32 characters in length",
				},
			},
		},
		{
			Name:    "Can't update character with level that is below 1",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				Pronouns:        "Test Pronouns",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           -1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Level must be 1 or greater",
				},
			},
		},
		{
			Name:    "Can't update character with level that is above 20",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				Pronouns:        "Test Pronouns",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           21,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Level must be 20 or less",
				},
			},
		},
		{
			Name:    "Can't update character with stats that are below 3",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				Pronouns:        "Test Pronouns",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        2,
				Dexterity:       2,
				Constitution:    2,
				Intelligence:    2,
				Wisdom:          2,
				Charisma:        2,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Strength must be 3 or greater",
					"Dexterity must be 3 or greater",
					"Constitution must be 3 or greater",
					"Intelligence must be 3 or greater",
					"Wisdom must be 3 or greater",
					"Charisma must be 3 or greater",
				},
			},
		},
		{
			Name:    "Can't update character with stats that are above 30",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				Pronouns:        "Test Pronouns",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        31,
				Dexterity:       31,
				Constitution:    31,
				Intelligence:    31,
				Wisdom:          31,
				Charisma:        31,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid:",
					"Strength must be 30 or less",
					"Dexterity must be 30 or less",
					"Constitution must be 30 or less",
					"Intelligence must be 30 or less",
					"Wisdom must be 30 or less",
					"Charisma must be 30 or less",
				},
			},
		},
		{
			Name:    "Can't update character that doesn't exist",
			Request: getRequest(1000),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't update character that belongs to a different user",
			Request: getRequest(differentUserCharacter.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't update character with class id that doesn't exist",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				ClassId:         1000,
				RaceId:          race.ID,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't update character with race id that doesn't exist",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          1000,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't update character with invalid advancement type",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: "invalid advancement type",
				HitPointType:    m.HitPointTypeFixed,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "AdvancementType is not valid",
			},
		},
		{
			Name:    "Can't update character with invalid hit point type",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class.ID,
				RaceId:          race.ID,
				Level:           1,
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeMilestone,
				HitPointType:    "invalid hit point type",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "HitPointType is not valid",
			},
		},
		{
			Name:    "Can update character",
			Request: getRequest(character.ID),
			RequestBody: requests.UpdateCharacterRequest{
				Name:            "Test Name",
				ClassId:         class2.ID,
				RaceId:          race2.ID,
				Level:           10,
				Pronouns:        "He/Him",
				Strength:        10,
				Dexterity:       10,
				Constitution:    10,
				Intelligence:    10,
				Wisdom:          10,
				Charisma:        10,
				AdvancementType: m.AdvancementTypeXP,
				HitPointType:    m.HitPointTypeManual,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Test Name"`,
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					`"level":10`,
					`"pronouns":"He/Him"`,
					`"strength":10`,
					`"dexterity":10`,
					`"constitution":10`,
					`"intelligence":10`,
					`"wisdom":10`,
					`"charisma":10`,
					`"advancement_type":"XP"`,
					`"hit_point_type":"Manual"`,
				},
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "Character was updated",
					Model: m.Character{
						ID:              character.ID,
						UserId:          ts.AdminUser.ID,
						Name:            "Test Name",
						ClassId:         class2.ID,
						RaceId:          race2.ID,
						Level:           10,
						Pronouns:        "He/Him",
						Strength:        10,
						Dexterity:       10,
						Constitution:    10,
						Intelligence:    10,
						Wisdom:          10,
						Charisma:        10,
						AdvancementType: m.AdvancementTypeXP,
						HitPointType:    m.HitPointTypeManual,
					},
					CountExpected: 1,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestCharacter_Delete(t *testing.T) {
	ts.ClearTable("characters")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Set mocks
	fileStoreMock := mocks.NewFileStoreMock()
	ts.S.Dependencies.SetFileStore(fileStoreMock)

	setup := func(test *helpers.TestCase) {
		fileStoreMock.Reset()
	}

	// Create characters
	character := &m.Character{}
	factories.NewCharacter(ts.S.Db, character)
	differentUserCharacter := &m.Character{UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	// Create images
	profilePicture := &m.File{Model: m.FileModelCharacterProfilePicture, ModelId: character.ID}
	factories.NewFile(ts.S.Db, profilePicture)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodDelete,
			Url:    fmt.Sprintf("/characters/%v", id),
		}
	}

	permissionRequest := getRequest(character.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't delete character that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't delete character with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't delete character that belongs to a different user",
			Request: getRequest(differentUserCharacter.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can delete character",
			Setup:   setup,
			Request: getRequest(character.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Character deleted successfully",
				DatabaseChecks: []*helpers.DatabaseCheck{
					{
						Name: "Character was deleted",
						Model: m.Character{
							ID:   character.ID,
							Name: character.Name,
						},
						CountExpected: 0,
					},
					{
						Name: "Profile picture was deleted",
						Model: m.File{
							Model:   m.FileModelCharacterProfilePicture,
							ModelId: character.ID,
						},
						CountExpected: 0,
					},
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.DeleteCalls))
					assert.Equal(t, fmt.Sprintf("%v/%v", profilePicture.FileLocation, profilePicture.Filename), fileStoreMock.DeleteCalls[0])
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestCharacter_UploadProfilePicture(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("races")
	ts.ClearTable("characters")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	// Set mocks
	fileStoreMock := mocks.NewFileStoreMock()
	ts.S.Dependencies.SetFileStore(fileStoreMock)

	setup := func(test *helpers.TestCase) {
		ts.ClearTable("files")
		fileStoreMock.Reset()
	}

	// Create classes
	class := &m.Class{}
	factories.NewClass(ts.S.Db, class)

	// Create races
	race := &m.Race{}
	factories.NewRace(ts.S.Db, race)

	// Create characters
	character := &m.Character{ClassId: class.ID, RaceId: race.ID}
	factories.NewCharacter(ts.S.Db, character)
	differentUserCharacter := &m.Character{ClassId: class.ID, RaceId: race.ID, UserId: 1000}
	factories.NewCharacter(ts.S.Db, differentUserCharacter)

	// PNG file
	pngBody, pngMw := createMultipartFile(t, "file", "../assets/example.png")
	// JPG file
	jpgBody, jpgMw := createMultipartFile(t, "file", "../assets/example.jpg")
	// JPEG file
	jpegBody, jpegMw := createMultipartFile(t, "file", "../assets/example.jpeg")
	// WEBP file
	webpBody, webpMw := createMultipartFile(t, "file", "../assets/example.webp")
	// PDF file
	pdfBody, pdfMw := createMultipartFile(t, "file", "../assets/example.pdf")

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodPost,
			Url:    fmt.Sprintf("/characters/%v/upload/profile-picture", id),
		}
	}

	permissionRequest := getRequest(character.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't upload image for character that doesn't exist",
			Request: getRequest(1000),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't upload image for character with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't upload image for character that belongs to a different user",
			Request: getRequest(differentUserCharacter.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't upload image for character if no file is provided",
			Request: getRequest(character.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Unable to read file",
			},
		},
		{
			Name:               "Can't upload image for character with invalid file type",
			Request:            getRequest(character.ID),
			RequestReader:      pdfBody,
			RequestContentType: pdfMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Invalid file type",
			},
		},
		{
			Name:               "Can upload png",
			Setup:              setup,
			Request:            getRequest(character.ID),
			RequestReader:      pngBody,
			RequestContentType: pngMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelCharacterProfilePicture,
						ModelId: character.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("characters/%v", character.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".png")
				},
			},
		},
		{
			Name:               "Can upload jpg",
			Setup:              setup,
			Request:            getRequest(character.ID),
			RequestReader:      jpgBody,
			RequestContentType: jpgMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelCharacterProfilePicture,
						ModelId: character.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("characters/%v", character.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".jpg")
				},
			},
		},
		{
			Name:               "Can upload jpeg",
			Setup:              setup,
			Request:            getRequest(character.ID),
			RequestReader:      jpegBody,
			RequestContentType: jpegMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelCharacterProfilePicture,
						ModelId: character.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("characters/%v", character.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".jpeg")
				},
			},
		},
		{
			Name:               "Can upload webp",
			Setup:              setup,
			Request:            getRequest(character.ID),
			RequestReader:      webpBody,
			RequestContentType: webpMw.FormDataContentType(),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "File uploaded",
				DatabaseCheck: &helpers.DatabaseCheck{
					Name: "File was uploaded",
					Model: m.File{
						Model:   m.FileModelCharacterProfilePicture,
						ModelId: character.ID,
					},
					CountExpected: 1,
				},
				ExpectedCallBack: func(res *httptest.ResponseRecorder) {
					// Ensure file store was called correctly
					assert.Equal(t, 1, len(fileStoreMock.SaveCalls))
					assert.Contains(t, fileStoreMock.SaveCalls[0].Path, fmt.Sprintf("characters/%v", character.ID))
					assert.Contains(t, fileStoreMock.SaveCalls[0].FileName, ".webp")
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}
