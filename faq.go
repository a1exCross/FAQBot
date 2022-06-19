package FAQBot

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type GetAnswerResponse struct {
	Answer string `json:"answer"`
}

func GetAnswer(Question string) (string, error) {
	if Question != "" {
		quest := struct {
			Question string `json:"user_question"`
		}{
			Question: Question,
		}

		data, err := json.Marshal(quest)
		if err != nil {
			return "", err
		}

		body, err := request("/answer", data)
		if err != nil {
			return "", err
		}

		jsonn, err := UnicodeToUTF(body)
		if err != nil {
			return "", err
		}

		var ans GetAnswerResponse

		err = json.Unmarshal(jsonn, &ans)
		if err != nil {
			return "", err
		}

		return ans.Answer, nil
	} else {
		return "", errors.New("function parameter is empty")
	}
}

func UnicodeToUTF(j json.RawMessage) (json.RawMessage, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

type NewQuestions struct {
	Questions []string `json:"questions"`
	Answer    string   `json:"answer"`
}

func GetDatasetUpdateParams() *DatasetUpdateParams {
	return &DatasetUpdateParams{}
}

type DatasetUpdateParams struct {
	NewQuestions []NewQuestions `json:"new_questions"`
}

func DatasetUpdate(p *DatasetUpdateParams, j json.RawMessage) (string, error) {
	var data []byte

	if p != nil {
		data, _ = json.Marshal(p)
	} else if j != nil {
		data = j
	} else {
		return "", errors.New("function parameters is empty")
	}

	body, err := request("/update", data)
	if err != nil {
		return "", err
	}

	var up DatasetUpdateResponse

	err = json.Unmarshal(body, &up)
	if err != nil {
		return "", err
	}

	return up.Update, nil
}

type DatasetUpdateResponse struct {
	Update string `json:"update"`
}

func ModelTrain(train bool) (string, error) {
	var t ModelTrainParams

	t.Train = fmt.Sprint(train)

	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	body, err := request("/train", data)
	if err != nil {
		return "", err
	}

	var tr ModelTrainParams

	err = json.Unmarshal(body, &tr)
	if err != nil {
		return "", err
	}

	return tr.Train, nil

}

type ModelTrainParams struct {
	Train string `json:"train"`
}
