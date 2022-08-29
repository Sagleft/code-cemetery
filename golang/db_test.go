package main


func TestGetTableColumns(t *testing.T) {
    // open connection
    db, mock, dbErr := GetTestDB()
    if dbErr != nil {
        t.Fatal(dbErr)
    }
    defer db.Close()
    h := dbHandler{
        Conn:   db.GetConn(),
        DBName: "testdb",
    }

    // get query data
    sqlQuery := getTableColumnsQuery()
    rows := sqlmock.NewRows([]string{"column_name"}).AddRow("id").AddRow("name").AddRow("uid")

    // setup query test
    tableName := "test_table"
    mock.ExpectQuery(regexp.QuoteMeta(sqlQuery)).WithArgs(tableName, h.DBName).
        WillReturnRows(rows)

    // send query
    columns, err := h.getTableColumns(tableName, map[string]struct{}{
        "id":   struct{}{},
        "name": struct{}{},
    })
    if err != nil {
        t.Fatal(err)
    }

    if len(columns) == 0 {
        t.Fatal("0 columns selected")
    }
    if len(columns) != 1 {
        t.Fatal("invalid columns number")
    }
}
