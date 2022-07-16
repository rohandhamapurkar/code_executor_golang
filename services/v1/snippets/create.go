package snippets

import (
	"rohandhamapurkar/code-executor/core/db"
	"rohandhamapurkar/code-executor/core/models"
	"rohandhamapurkar/code-executor/core/structs"
)

func CreateSnippet(reqBody *structs.CreateSnippetReqBody, userID string) (uint, error) {
	snippet := &models.Snippets{Name: reqBody.Name, Code: reqBody.Code, Language: reqBody.Language, Public: *reqBody.Public, UserID: userID}
	result := db.Postgres.Create(snippet)

	return snippet.ID, result.Error
}
