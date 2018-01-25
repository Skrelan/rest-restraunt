package models

// Address is the structure for Addresses data
type Address struct {
	StreetAddress string `json:"street_address,omitempty" db:"street_address"`
	City          string `json:"city,omitempty" db:"city"`
	State         string `json:"state,omitempty" db:"state"`
	ZipCode       string `json:"zip_code,omitempty" db:"zip_code"`
}

// User is the structure for Users data
type User struct {
	ID        int64  `json:"id,omitempty" db:"id"`
	FirstName string `json:"first_name,omitempty" db:"first_name"`
	LastName  string `json:"last_name,omitempty" db:"last_name"`
	Phone     string `json:"phone,omitempty" db:"phone"`
}

// Restaurant is the structure for Restraunts data
type Restaurant struct {
	ID       string `json:"id,omitempty" db:"restaurant_id"`
	Name     string `json:"name,omitempty" db:"name"`
	Category string `json:"category,omitempty" db:"category"`
}

// Venue is the structure that holds information of venues of Restraunt
type Venue struct {
	ID        string `json:"id,omitempty" db:"id"`
	Restraunt *Restaurant
	Address   *Address
}

// Rating is the structure for the Ratings data
type Rating struct {
	Cost               int64  `json:"cost,omitempty" db:"cost"`
	Food               int64  `json:"food,omitempty" db:"food"`
	CleanlinessService int64  `json:"cleanliness_service,omitempty" db:"cleanliness_service"`
	TotalScore         int64  `json:"total_score,omitempty" db:"total_score"`
	RestaurantID       int64  `json:"restaurant_id,omitempty" db:"restaurant_id"`
	UserID             int64  `json:"user_id,omitempty" db:"user_id"`
	DateCreated        string `json:"date_created,omitempty" db:"date_created"`
	DateUpdated        string `json:"date_updated,omitempty" db:"date_updated"`
}
