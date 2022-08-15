package routes

import (
	"API/models"
	"API/utils"
	"encoding/json"
	"net/http"
	"strings"
)

func (route *Route) GetAllQuestion(w http.ResponseWriter, r *http.Request) {
	var question models.Question

	questions, err := question.GetAllQuestion(route.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	var results []models.Question
	for _, item := range questions {
		//item.Answers = []models.Answer{}
		var answer models.Answer
		answer.QuestionID = item.ID
		answers, err := answer.GetAnswerByQuestion(route.DB)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		for _, answer := range answers {
			item.Answers = append(item.Answers, answer)
		}
		results = append(results, item)
	}
	if len(results) == 0 {
		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}
	respondWithJSON(w, http.StatusOK, results)
}

func (route *Route) LoadSaveAnswer(w http.ResponseWriter, r *http.Request) {
	var userNameQuestion models.UserNameQuestion
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userNameQuestion); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	userQuestion := models.UserQuestion{UserName: userNameQuestion.UserName}
	userQuestions, err := userQuestion.GetUserQuestionByUser(route.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(userQuestions) == 0 {
		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}
	var results []models.UserQuestion
	for _, item := range userQuestions {
		var result models.UserQuestion
		result.ID = item.ID
		result.UserID = item.UserID
		result.QuestionID = item.QuestionID
		result.Status = item.Status
		result.QuestionDesc = item.QuestionDesc
		result.UserName = item.UserName
		var userQuestionAnswer models.UserQuestionAnswer
		userQuestionAnswer.UserQuestionID = item.ID
		userQuestionAnswers, err := userQuestionAnswer.GetUserQuestionAnswerByUserQuestion(route.DB)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		result.UserQuestionAnswers = append(result.UserQuestionAnswers, userQuestionAnswers...)
		results = append(results, result)
	}

	respondWithJSON(w, http.StatusOK, results)

}
func (route *Route) SaveAnswers(w http.ResponseWriter, r *http.Request) {
	var userQuestion models.UserQuestion
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userQuestion); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	var user models.User
	user.UserName = userQuestion.UserName
	defer r.Body.Close()
	errors := userQuestion.Validate()
	for _, answer := range userQuestion.UserQuestionAnswers {
		serrors := answer.Validate()
		if len(serrors) > 0 {
			errors = append(errors, serrors...)
		}
	}
	if len(errors) > 0 {
		respondWithError(w, http.StatusBadRequest, strings.Join(errors, "<br/>"))
		return
	}
	err := user.GetUserByUserName(route.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if userQuestion.ID == "" {

		userQuestion.ID = utils.GenerateID()
		userQuestion.UserID = user.ID
		userQuestion.Status = "Save"
		err := userQuestion.CreateUserQuestion(route.DB)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		for _, answer := range userQuestion.UserQuestionAnswers {
			answer.ID = utils.GenerateID()
			answer.UserQuestionID = userQuestion.ID
			if err := answer.CreateUserQuestionAnswer(route.DB); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}

	} else {
		userQuestion.Status = "Save"
		userQuestion.UserID = user.ID
		err := userQuestion.UpdateUserQuestion(route.DB)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		var userQuestionAnswer models.UserQuestionAnswer
		userQuestionAnswer.UserQuestionID = userQuestion.ID
		err = userQuestionAnswer.DeleteUserQuestionAnswersByUserQuestion(route.DB)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		for _, answer := range userQuestion.UserQuestionAnswers {
			answer.ID = utils.GenerateID()
			answer.UserQuestionID = userQuestion.ID
			if err := answer.CreateUserQuestionAnswer(route.DB); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
	respondWithJSON(w, http.StatusOK, userQuestion)

}

func (route *Route) SubmitUserAnswer(w http.ResponseWriter, r *http.Request) {
	var userQuestion models.UserQuestion
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userQuestion); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	var user models.User
	user.UserName = userQuestion.UserName
	defer r.Body.Close()
	errors := userQuestion.Validate()
	for _, answer := range userQuestion.UserQuestionAnswers {
		serrors := answer.Validate()
		if len(serrors) > 0 {
			errors = append(errors, serrors...)
		}
	}
	if len(errors) > 0 {
		respondWithError(w, http.StatusBadRequest, strings.Join(errors, "<br/>"))
		return
	}
	err := user.GetUserByUserName(route.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if userQuestion.ID == "" {
		userQuestion.Status = "Submit"
		userQuestion.ID = utils.GenerateID()
		userQuestion.UserID = user.ID
		err := userQuestion.CreateUserQuestion(route.DB)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		for _, answer := range userQuestion.UserQuestionAnswers {
			answer.ID = utils.GenerateID()
			answer.UserQuestionID = userQuestion.ID
			if err := answer.CreateUserQuestionAnswer(route.DB); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}

	} else {
		userQuestion.Status = "Submit"
		userQuestion.UserID = user.ID
		err := userQuestion.UpdateUserQuestion(route.DB)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		var userQuestionAnswer models.UserQuestionAnswer
		userQuestionAnswer.UserQuestionID = userQuestion.ID
		err = userQuestionAnswer.DeleteUserQuestionAnswersByUserQuestion(route.DB)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		for _, answer := range userQuestion.UserQuestionAnswers {
			answer.ID = utils.GenerateID()
			answer.UserQuestionID = userQuestion.ID
			if err := answer.CreateUserQuestionAnswer(route.DB); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
	//var user models.User
	user.UserName = userQuestion.UserName
	summary, err := user.Summary(route.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, summary)

}
