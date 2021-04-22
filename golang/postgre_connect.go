// Connect to Postgre
	//example: "postgres://{username}:{password}@{hostname}:{port}/bank?sslmode=require"
	//postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10
	/*var dsn2ConnStr string = "postgres://" + os.Getenv("postgre_user") + ":" +
		os.Getenv("postgre_pass") + "@" +
		os.Getenv("postgre_host") + ":" +
		os.Getenv("postgre_port") + "/" +
		os.Getenv("postgre_db")

	if os.Getenv("postgre_tls") == "1" || os.Getenv("postgre_tls") == "true" {
		dsn2ConnStr += "?sslmode=require"
	}*/
	/*
		config, err := pgx.ParseConfig(dsn2ConnStr)
		if err != nil {
			log.Fatal("error configuring the database: ", err)
		}*/

	//dsn2ConnStr, config.ConnString()
	/*pgxpool, err := pgxpool.Connect(context.Background(), dsn2ConnStr)
	if err != nil {
		return errors.New("Unable to connect to database: " + err.Error())
	}*/