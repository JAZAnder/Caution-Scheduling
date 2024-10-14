package log

func GetAllLogs() ([]Log, error ){

	var logs []Log

	return logs, nil;
}

func GetNumberOfLogs(numberOfRecords int) ([]Log, error ){

	var logs []Log

	return logs, nil;
}

func GetLogsByFilter(startId, endID, numberOfRecords  int, minLevel, maxLevel, Category, SubCategory, User string)([]Log, error ){

	var logs []Log

	return logs, nil;
}

func GetCategories() ([]string, error){
	var Categories []string

	return Categories, nil;
}

func GetSubCategoryFromCategory(category string) ([]string, error){
	var SubCategory []string

	return SubCategory, nil;
}


func GetSubCategory() ([]string, error){
	var SubCategory []string

	return SubCategory, nil;
}


func GetUsers() ([]string, error){
	var Users []string

	return Users, nil;
}



