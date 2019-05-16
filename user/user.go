package user

// This is a very basic module that creates users and manages balances
// as of now it doesn't have any functionality that can get aggregate balances of
// users. This will be added later if needed. This whole module should likely also
// save/retrieve data from a database since speed is not crucial for these specific functions.

// User is the basic struct that holds user information and balances
type User struct {
	Username string
	Balance  map[string]float32
}

// CreateUser creates a new user with an empty balance mapping
func CreateUser(username string) User {
	return User{username, make(map[string]float32)}
}

//Balances get's the specified user's balances in all stocks
func (u *User) Balances() map[string]float32 {
	return u.Balance
}

// AddToBalance add a specified amount of a particular stock
// to the user's balance
func (u *User) AddToBalance(stock string, amount float32) {
	if amount > 0 {
		u.Balance[stock] += amount
	}
}

// SubtractFromBalance subtracts a specified amount of a particular stock
// from the user's balance
func (u *User) SubtractFromBalance(stock string, amount float32) {
	if amount > 0 {
		u.Balance[stock] -= amount // check for negative balances?
	}
}
