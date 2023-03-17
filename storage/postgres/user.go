package postgres

import (
	u "TEMPLATE_MICROSERVICE/genproto/user"
	"log"

	_ "github.com/lib/pq"
)

func (r *UserRepo) CreateUser(user *u.UserRequest) (*u.UserResponse, error) {
	var res u.UserResponse
	err := r.db.QueryRow(`
	INSERT INTO 
		users (first_name, last_name, email) 
	VALUES
		($1, $2, $3) 
	RETURNING 
		id, first_name, last_name, email, created_at, updated_at`,
		user.FirstName, user.LastName, user.Email).
		Scan(&res.Id, &res.FirstName, &res.LastName, &res.Email, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		log.Println("Error inserting user info")
		return &u.UserResponse{}, err
	}

	return &res, nil
}

func (r *UserRepo) GetUserById(user *u.UserId) (*u.UserResponse, error) {
	var res u.UserResponse

	err := r.db.QueryRow(`
	SELECT 
		id, first_name, last_name, email, created_at, updated_at
	FROM 
		users
	WHERE 
		id = $1
	`, user.Id).Scan(
		&res.Id, &res.FirstName, &res.LastName, &res.Email, &res.CreatedAt, &res.UpdatedAt,
	)

	if err != nil {
		log.Println("Error gettig users by id")
		return &u.UserResponse{}, err
	}

	return &res, nil
}

func (r *UserRepo) GetUsersAll(users *u.UserListReq) (*u.Users, error) {
	var res u.Users
	query := `
	SELECT 
		id, first_name, last_name, email
	FROM 
		users
	ORDER BY created_at DESC LIMIT $1`

	rows, err := r.db.Query(query, users.Limit)
	if err != nil {
		return &u.Users{}, err
	}

	for rows.Next() {
		temp := u.UserResponse{}
		err := rows.Scan(
			&temp.Id, &temp.FirstName, &temp.LastName, &temp.Email,
		)

		if err != nil {
			return &u.Users{}, err
		}

		res.Users = append(res.Users, &temp)
	}
	return &res, nil
}

func (r *UserRepo) DeleteUser(user *u.UserId) (*u.Users, error) {
	var res u.Users

	query := `
	DELETE FROM 
		users
	WHERE 
		id = $1`
	_, err := r.db.Exec(query, user.Id)
	if err != nil {
		log.Println("Error deleting user", err)
		return &u.Users{}, err
	}
	query = `
	SELECT 
		id, first_name, last_name, email
	FROM 
		users
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return &u.Users{}, err
	}

	for rows.Next() {
		temp := u.UserResponse{}
		err := rows.Scan(
			&temp.Id, &temp.FirstName, &temp.LastName, &temp.Email,
		)

		if err != nil {
			return &u.Users{}, err
		}

		res.Users = append(res.Users, &temp)
	}
	return &res, nil
}

func (r *UserRepo) UpdateUser(user *u.UserUpdateReq) (*u.UserResponse, error) {
	res := u.UserResponse{}

	err := r.db.QueryRow(`
	UPDATE 
		users 
	SET 
		first_name=$1, last_name=$2, email=$3, updated_at=NOW()
	WHERE 
		id=$4 AND deleted_at IS NULL
	RETURNING 
		id, first_name, last_name, email, created_at, updated_at`,
		user.FirstName, user.LastName, user.Email, user.Id).
		Scan(&res.Id, &res.FirstName, &res.LastName, &res.Email, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return &u.UserResponse{}, err
	}
	return &res, nil
}

func (r *UserRepo) SearchUser(user *u.UserSearch) (*u.Users, error) {
	res := u.Users{}

	query := `SELECT id, first_name, last_name, email, created_at, updated_at FROM users WHERE first_name ILIKE '%` + user.Name + "%'"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error SEARCHING USER user", err)
		return &u.Users{}, err
	}

	for rows.Next() {
		temp := u.UserResponse{}
		err := rows.Scan(
			&temp.Id, &temp.FirstName, &temp.LastName, &temp.Email, &temp.CreatedAt, &temp.UpdatedAt,
		)

		if err != nil {
			return &u.Users{}, err
		}

		res.Users = append(res.Users, &temp)

	}

	return &res, nil
}
