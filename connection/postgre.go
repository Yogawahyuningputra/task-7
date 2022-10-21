package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnect(){
	var err error
	databaseUrl := "postgres://postgres:postgresql@localhost:5432/personal_web"
	
	Conn, err = pgx.Connect(context.Background(),databaseUrl) 

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)

	}
	fmt.Println("Succes Connect to Database")
}