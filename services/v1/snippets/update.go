package snippets

import (
	"rohandhamapurkar/code-executor/core/structs"
)

type updateQuery struct {
	Name     string
	Language string
	Code     string
	Public   bool
}

func UpdateSnippet(reqBody *structs.UpdateSnippetReqBody, userID string) error {

	// update := &updateQuery{
	// 	Name:     reqBody.Name,
	// 	Language: reqBody.Language,
	// 	Code:     reqBody.Code,
	// }

	// snippet := &models.Snippets{}
	// snippet.ID = reqBody.ID
	// snippet.UserID = userID
	// result := db.Postgres.Model(snippet).UpdateColumns()
	// return result.Error

	return nil
}
