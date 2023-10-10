# Caution-Scheduling

## Files not Included
```.env```

This should be placed in the root directory, and be layed out as such. 
```
PORT= 80
APP_DB_USERNAME = MYSQL_USER
APP_DB_PASSWORD = MYSQL_PASSWORD
APP_DB_NAME = MYSQL_DATABASE
APP_DB = godatabase
```

```adc.json``` (Optional)

This should be placed in the root directory, and be layed out as such. 
```
{
  "client_id": "*client_id*",
  "client_secret": "*client_secret*",
  "refresh_token": "*refresh_token*",
  "type": "authorized_user"
}
```
To get the correct information for this run the command ```gcloud auth application-default login```
