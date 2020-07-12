# easycron_go
Easycron Go package

# Installation
go get github.com/Senshi26/easycron_go


# Example usage
```Go
test_exp:=easycron.Cron_expression{Minute:0,Hour:11,Day_month:0,Month:0,Day_week:6,Year:0,Recurring:true} 

var test_data *easycron.EasycronData = &easycron.EasycronData{Url:"https://example.com",
		Auth_user:"XXXXXXX",Auth_pw:"XXXXXXXX",Cron_job_name:"TEST_CAMPAIGN_FOR_DELETION", Group_id:"12038","XXXXXXXXXXXXXXXXXXXX",Cron_expression:""} // set authorizattion and cron job settings
```

    
## EasycronData struct
```Go    
type EasycronData struct {

  Url string
	Auth_user string
	Auth_pw string
	Cron_job_name string
	Group_id string
	Token string
	Cron_expression string

}
```

## Cron_expression struct
```Go    
  type Cron_expression struct {

Minute int
Hour int
Day_month int
Month int
Day_week int
Year int
Recurring bool

}  
```
