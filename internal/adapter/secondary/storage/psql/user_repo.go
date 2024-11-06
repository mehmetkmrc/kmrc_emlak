package psql

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/mehmetkmrc/kmrc_emlak/internal/core/domain/entity"
	"github.com/mehmetkmrc/kmrc_emlak/internal/core/port/db"
	"github.com/mehmetkmrc/kmrc_emlak/internal/core/port/user"
)

var UserRepoSet = wire.NewSet(NewUserRepository)

type(
	UserRepository struct{
		db *sql.DB
	}
)

func NewUserRepository(db db.EngineMaker) user.UserRepositoryPort {
	return &UserRepository{
		db: db.GetDB(),
	}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error){
	userQuery := struct {
		ID		  sql.NullString
		Name	  sql.NullString
		Surname   sql.NullString
		Email	  sql.NullString
		Password  sql.NullString
		CreatedAt sql.NullTime
	}{}
	query := `
	SELECT CAST(SystemUserId AS VARCHAR(64)) as ID, 
       FirstName, 
       LastName, 
       InternalEMailAddress, 
       new_sifre, 
       CreatedOn 
	FROM SystemUserBase 
	WHERE InternalEMailAddress = $1 
  		AND new_sifre IS NOT NULL 
  		AND InternalEMailAddress IS NOT NULL;
	`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&userQuery.ID, &userQuery.Name, &userQuery.Surname, &userQuery.Email, &userQuery.Password, &userQuery.CreatedAt)
	if err != nil{
		return nil, err
	}

	userData := &entity.User{
		ID:        userQuery.ID.String,
		Name:      userQuery.Name.String,
		Surname:   userQuery.Surname.String,
		Email:     userQuery.Email.String,
		Password:  userQuery.Password.String,
		CreatedAt: userQuery.CreatedAt.Time,
	}
	return userData, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	userQuery := struct {
		ID        sql.NullString
		Name      sql.NullString
		Surname   sql.NullString
		Email     sql.NullString
		Password  sql.NullString
		CreatedAt sql.NullTime
	}{}
	query := `SELECT CAST(SystemUserId AS VARCHAR(64)) as ID, 
       FirstName, 
       LastName, 
       InternalEMailAddress, 
       new_sifre, 
       CreatedOn 
	FROM SystemUserBase 
	WHERE SystemUserId = $1 
  		AND new_sifre IS NOT NULL 
  		AND InternalEMailAddress IS NOT NULL;
	`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&userQuery.ID, &userQuery.Name, &userQuery.Surname, &userQuery.Email, &userQuery.Password, &userQuery.CreatedAt)
	if err != nil {
		return nil, err
	}

	userData := &entity.User{
		ID:        userQuery.ID.String,
		Name:      userQuery.Name.String,
		Surname:   userQuery.Surname.String,
		Email:     userQuery.Email.String,
		Password:  userQuery.Password.String,
		CreatedAt: userQuery.CreatedAt.Time,
	}
	return userData, nil
}

func (r *UserRepository) GetUserPassword(ctx context.Context, email string) (string, error) {
	var password string
	query := `SELECT new_sifre 
	FROM SystemUserBase 
	WHERE InternalEmailAddress = $1;
	`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}