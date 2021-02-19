package offers

import (
	"context"
	"strconv"

	database "github.com/advenjourney/api/internal/pkg/db/postgres"
	"github.com/advenjourney/api/internal/users"
)

// #1
type Offer struct {
	ID            string
	Title         string
	Location      string
	Description   string
	TitleImageURL string
	User          *users.User
}

//#2
func (offer Offer) Save() (int64, error) {
	ctx := context.Background()
	q := "INSERT INTO Offers(Title,Location,Description,TitleImageURL, UserID) VALUES($1,$2,$3,$4,$5) RETURNING id"

	res := database.DB.QueryRow(ctx, q, offer.Title, offer.Location, offer.Description, offer.TitleImageURL, offer.User.ID)
	var newOfferID int64
	if err := res.Scan(&newOfferID); err != nil {
		return 0, err
	}

	return newOfferID, nil
}

func GetAll() ([]Offer, error) {
	ctx := context.Background()
	q := `SELECT o.id, o.title, o.location, o.description, o.titleimageurl, o.UserID, u.Username
          FROM offers o
		  INNER JOIN users u ON o.UserID = u.ID`

	rows, err := database.DB.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []Offer

	for rows.Next() {
		var offer = Offer{}
		var username string
		var userID, offerID int // missmatch? DB-Schema says int...
		err := rows.Scan(&offerID, &offer.Title, &offer.Location, &offer.Description, &offer.TitleImageURL, &userID, &username)
		if err != nil {
			return nil, err
		}

		offer.ID = strconv.Itoa(offerID) // ... but ID is string?
		offer.User = &users.User{
			ID:       strconv.Itoa(userID),
			Username: username,
		}
		offers = append(offers, offer)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return offers, nil
}
