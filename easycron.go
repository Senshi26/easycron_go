package easycron

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type EasycronData struct {

	Url string
	Auth_user string
	Auth_pw string
	Cron_job_name string
	Group_id string
	Token string
	Cron_expression string

}

type Cron_expression struct {

Minute int
Hour int
Day_month int
Month int
Day_week int
Year int
Recurring bool

}

type Cronjob_collecttion struct {
	CronJobs []struct {
		AuthPw           string   `json:"auth_pw"`
		AuthUser         string   `json:"auth_user"`
		Created          string   `json:"created"`
		Criterion        string   `json:"criterion"`
		CronExpression   string   `json:"cron_expression"`
		CronJobID        string   `json:"cron_job_id"`
		CronJobName      string   `json:"cron_job_name"`
		CustomTimeout    string   `json:"custom_timeout"`
		Description      string   `json:"description"`
		EmailMe          string   `json:"email_me"`
		EpdsOccupied     string   `json:"epds_occupied"`
		FailureRegexp    string   `json:"failure_regexp"`
		GroupID          string   `json:"group_id"`
		HTTPHeaders      string   `json:"http_headers"`
		HTTPMethod       string   `json:"http_method"`
		NumberFailedTime string   `json:"number_failed_time"`
		Posts            string   `json:"posts"`
		Sensitivity      string   `json:"sensitivity"`
		Status           string   `json:"status"`
		SuccessRegexp    string   `json:"success_regexp"`
		TotalFailures    string   `json:"total_failures"`
		TotalSuccesses   string   `json:"total_successes"`
		Updated          string   `json:"updated"`
		URL              string   `json:"url"`
		UserID           string   `json:"user_id"`
		Wh               string   `json:"wh"`
		WhData           []string `json:"wh_data"`
		WhHTTPMethod     string   `json:"wh_http_method"`
		WhURL            string   `json:"wh_url"`
	} `json:"cron_jobs"`
	Status string `json:"status"`
}
type Easycron interface {

Add() string
Create_Cronjob_expression() string
Edit()  string
List() Cronjob_collecttion
}


func (E EasycronData) Edit(cronjob_name string,ce Cron_expression)string{
	cronjobCollection := E.List()
	var cronjob_id string
	for i :=range cronjobCollection.CronJobs {

		if cronjobCollection.CronJobs[i].CronJobName == cronjob_name{

			cronjob_id = cronjobCollection.CronJobs[i].CronJobID
		}

	}

	resp, err := http.Get("https://www.easycron.com/rest/edit?" + "token=" +   E.Token + "&id=" + cronjob_id + "&cron_expression="+ ce.Create_Cronjob_Expression() + "&auth_user=" + E.Auth_user +
		"&auth_pw=" + E.Auth_pw + "&group_id=" + E.Group_id + "&cron_job_name=" + E.Cron_job_name + "&url=" + E.Url )
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	return string(body)

}


func (E EasycronData) List()Cronjob_collecttion{

	resp, err := http.Get("https://www.easycron.com/rest/list?" + "token=" + E.Token  )
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	 var cronjobCollection Cronjob_collecttion

	err2 := json.Unmarshal(body,&cronjobCollection)
	if err2 != nil{

		log.Println(err)

	}

	return cronjobCollection



}

func (C Cron_expression) Create_Cronjob_Expression()string{

	var expression string
	switch C.Recurring {

	case true:
		if C.Day_week == 0 {

			expression = strconv.Itoa(C.Minute) + " " + strconv.Itoa(C.Hour) + " "  + strconv.Itoa(C.Day_month) + " " + strconv.Itoa(C.Month) + " " + "*" + " "

		} else {
			expression = strconv.Itoa(C.Minute) + " " + strconv.Itoa(C.Hour) + " " + "*" + " " + "*" +  " " +  strconv.Itoa(C.Day_week)

		}

	case false:
		if C.Day_week == 0 {

			expression = strconv.Itoa(C.Minute) + " " + strconv.Itoa(C.Hour) + " "  + strconv.Itoa(C.Day_month) + " " + strconv.Itoa(C.Month) + " " + "*" + " " + strconv.Itoa(C.Year)

		} else {
			expression = strconv.Itoa(C.Minute) + " " + strconv.Itoa(C.Hour) + " "  + "*" + " " + "*" + " " + strconv.Itoa(C.Day_week) + " " + strconv.Itoa(C.Year)

		}
	}
	fmt.Println("Получившееся выражение")
	fmt.Println(expression)
	return expression
}

func (E EasycronData) Add(ce Cron_expression)string   {


	resp, err := http.Get("https://www.easycron.com/rest/add?" + "token=" + E.Token + "&cron_expression="+ ce.Create_Cronjob_Expression() + "&auth_user=" + E.Auth_user +
		"&auth_pw=" + E.Auth_pw + "&group_id=" + E.Group_id + "&cron_job_name=" + E.Cron_job_name + "&url=" + E.Url )
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	return string(body)

}