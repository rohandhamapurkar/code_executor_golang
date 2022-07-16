package snippets

import (
	"rohandhamapurkar/code-executor/core/db"
	"rohandhamapurkar/code-executor/core/models"
	"rohandhamapurkar/code-executor/core/structs"
)

func DeleteSnippet(reqBody *structs.DeleteSnippetReqBody, userID string) error {
	snippet := &models.Snippets{}
	snippet.ID = reqBody.ID
	snippet.UserID = userID

	result := db.Postgres.Delete(snippet)

	return result.Error
}
