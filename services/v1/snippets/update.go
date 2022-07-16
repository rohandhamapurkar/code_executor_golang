package snippets

import (
	"rohandhamapurkar/code-executor/core/db"
	"rohandhamapurkar/code-executor/core/models"
	"rohandhamapurkar/code-executor/core/structs"
	"rohandhamapurkar/code-executor/core/utils"
)

func UpdateSnippet(reqBody *structs.UpdateSnippetReqBody, userID string) error {

	update := map[string]interface{}{
		"Name":     reqBody.Name,
		"Language": reqBody.Language,
		"Code":     reqBody.Code,
		"Public":   reqBody.Public,
	}
	utils.BuildUpdateQuery(update)

	snippet := &models.Snippets{}
	snippet.ID = reqBody.ID
	snippet.UserID = userID
	result := db.Postgres.Model(snippet).UpdateColumns(update)
	return result.Error

}
