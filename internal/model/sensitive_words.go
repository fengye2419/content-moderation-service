package model

// SensitiveWord represents a sensitive word in the database.
type SensitiveWord struct {
	ID   int64  `xorm:"pk autoincr" json:"id"` // 主键，自增
	Word string `xorm:"unique" json:"word"`    // 敏感词，唯一
}

// TableName returns the table name for the SensitiveWord model.
func (SensitiveWord) TableName() string {
	return "sensitive_words"
}

// AddSensitiveWord adds a new sensitive word to the database.
func AddSensitiveWord(word string) error {
	sensitiveWord := &SensitiveWord{Word: word}
	_, err := engine.Insert(sensitiveWord)
	return err
}

// GetSensitiveWords retrieves all sensitive words from the database.
func GetSensitiveWords() ([]SensitiveWord, error) {
	var words []SensitiveWord
	err := engine.Find(&words)
	if err != nil {
		return nil, err
	}
	return words, nil
}

// DeleteSensitiveWord deletes a sensitive word from the database by ID.
func DeleteSensitiveWord(id int64) error {
	_, err := engine.ID(id).Delete(&SensitiveWord{})
	return err
}
