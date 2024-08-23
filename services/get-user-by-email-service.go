package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/my-little-pet/user-microservice/config"
	"github.com/my-little-pet/user-microservice/models"
)

func GetByEmailUser(email string) (*models.User, error){
	db, err := config.DBConfig()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
		return nil,err
	}
	defer db.Close()
	sqlStatement := `SELECT id,fullname,imageUrl,email,phone,created_at FROM users WHERE email = $1;`
	
	var user models.User

	err = db.QueryRow(sqlStatement,email).Scan(&user.ID, &user.Fullname,&user.ImageUrl,&user.Email,&user.Phone,&user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuário com email %s não encontrado", email)
		}
		return nil,err
	}
	return &user,nil
}