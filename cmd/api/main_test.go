package main

// func TestMain(t *testing.T) {
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("[WARNING] Not use .env file")
// 	}

// 	env, errENV := environment.Get()
// 	if errENV != nil {
// 		t.Error(errENV)
// 	}

// 	conn, errConn := repository.Get(&repository.GetInput{
// 		Type: repository.ConnectionSQLite,
// 		Env:  env,
// 	})
// 	if errConn != nil {
// 		t.Error(errConn)
// 	}

// 	defer conn.Close()

// 	router := gin.Default()
// 	api.Router(router.Group(env.APIVersion))

// 	if err := router.Run(fmt.Sprintf(":%s", env.APIPort)); err != nil {
// 		t.Error(err)
// 	}
// }
