package snippets

import (
	"rohandhamapurkar/code-executor/core/db"
	"rohandhamapurkar/code-executor/core/models"
)

func GetSnippet(userId string) (*[]models.Snippets, error) {
	query := &models.Snippets{}
	query.UserID = userId
	snippets := &[]models.Snippets{}
	result := db.Postgres.Model(query).Find(snippets)

	return snippets, result.Error
}
