package store

import "github.com/bakhodur-nazriev/e-shop/internal/app/models"

/* UserRepository ...  */
type UserRepository struct {
	store *Store
}

/* Create User ... */
func (r *UserRepository) Create(u *models.User) (*models.User, error) {
	if err := r.store.db.QueryRow("INSERT INTO users (first_name, last_name, email, gender, birthday) "+
		"VALUES ($1,$2,$3,$4,$5) RETURNING id",
		u.FirstName,
		u.LastName,
		u.Email,
		u.Gender,
		u.Birthday,
	).Scan(&u.Id); err != nil {
		return nil, err
	}
	return u, nil
}

/* Find User By Email ... */
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow("SELECT id, first_name, last_name, email, gender, birthday FROM users WHERE email = $1", email).Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Gender, &u.Birthday); err != nil {
		return nil, err
	}
	return u, nil
}
