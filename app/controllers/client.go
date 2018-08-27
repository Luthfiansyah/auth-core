package controllers

// func DoPost(data string) (*http.Response, error) {
// 	u := types.RequestEncryptType{ClientId: config.MustGetString(GetRunMode() + ".ci"), Data: data}
// 	b := new(bytes.Buffer)
// 	json.NewEncoder(b).Encode(u)
// 	request, err := http.NewRequest("POST", config.MustGetString(GetRunMode()+".url_api"), b)
// 	request.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{}
// 	resp, err := client.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp, err
// }
