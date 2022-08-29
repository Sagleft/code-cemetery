package main

func rowsToStringArray(
    rows *sql.Rows,
    excludeColumns map[string]struct{},
) ([]string, *constants.APIError) {
    vals := []string{}
    for rows.Next() {
        var newColumn string
        err := rows.Scan(&newColumn)
        if err != nil {
            return nil, constants.Error(
                consts.ErrorNameDataParse,
                err.Error(),
            )
        }

        _, isExclude := excludeColumns[newColumn]
        if !isExclude {
            vals = append(vals, newColumn)
        }
    }
    return vals, nil
}

func getTableColumnsQuery() string {
    return "SELECT column_name FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name=? AND table_schema=?"
}

func (db *dbHandler) getTableColumns(
    tableName string,
    excludeColumns map[string]struct{},
) ([]string, *constants.APIError) {
    rows, err := db.Conn.Query(getTableColumnsQuery(), tableName, db.DBName)
    defer db.rowsClose(rows)
    if err != nil {
        return nil, constants.Error(
            consts.ErrorNameServiceReqFailed,
            err.Error(),
        )
    }

    return rowsToStringArray(rows, excludeColumns)
}
