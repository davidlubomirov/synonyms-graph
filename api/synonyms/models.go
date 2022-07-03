package synonyms

type input struct {
	Word    string `json:"word"`
	Synonym string `json:"synonym"`
}

func (i *input) IsValid() bool {
	if i.Word != "" && i.Word != " " && i.Synonym != "" && i.Synonym != " " {
		return true
	}

	return false
}

type output struct {
	Word     string   `json:"word"`
	Synonyms []string `json:"synonyms"`
}
