package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rsh456/reservation-system/db/fixtures"
	"github.com/rsh456/reservation-system/types"
)

func Test_UserGetBooking(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	var (
		noAuthUser = fixtures.AddUser(db.Store, "Jimmy", "watercooler", false)
		user       = fixtures.AddUser(db.Store, "mark", "foo", false)
		hotel      = fixtures.AddHotel(db.Store, "bar hotel", "a", 4, nil)
		room       = fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)

		from           = time.Now()
		till           = from.AddDate(0, 0, 5)
		booking        = fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
		app            = fiber.New()
		route          = app.Group("/", JWTAuthentication(db.User))
		bookingHandler = NewBookingHandler(db.Store)
	)

	// Testing authenticated user
	route.Get("/:id", bookingHandler.HandleGetBooking)
	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("non 200 code got %d", resp.StatusCode)
	}

	var bookingResponse *types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookingResponse); err != nil {
		t.Fatal(err)
	}
	if bookingResponse.ID != booking.ID {
		t.Fatalf("expected %s got %s", booking.ID, bookingResponse.ID)
	}

	if bookingResponse.UserID != booking.UserID {
		t.Fatalf("expected %s got %s", booking.UserID, bookingResponse.UserID)
	}

	//test non admin cannot access the bookings
	req = httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-token", CreateTokenFromUser(noAuthUser))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected non 200 response got %d", resp.StatusCode)
	}

}

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	var (
		adminUser = fixtures.AddUser(db.Store, "admin", "admin", true)
		user      = fixtures.AddUser(db.Store, "mark", "foo", false)
		hotel     = fixtures.AddHotel(db.Store, "bar hotel", "a", 4, nil)
		room      = fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)

		from           = time.Now()
		till           = from.AddDate(0, 0, 5)
		booking        = fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
		app            = fiber.New(fiber.Config{ErrorHandler: ErrorHandler})
		admin          = app.Group("/", JWTAuthentication(db.User), AdminAuth)
		bookingHandler = NewBookingHandler(db.Store)
	)

	// Testing with an admin user
	_ = booking
	admin.Get("/", bookingHandler.HandleGetBookings)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-token", CreateTokenFromUser(adminUser))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("non 200 response %d", resp.StatusCode)
	}
	var bookings []*types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		t.Fatal(err)
	}
	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking got %d", len(bookings))
	}
	have := bookings[0]
	if have.ID != booking.ID {
		t.Fatalf("expected %s got %s", booking.ID, have.ID)
	}

	if have.UserID != booking.UserID {
		t.Fatalf("expected %s got %s", booking.UserID, have.UserID)
	}

	//test non admin cannot access the bookings
	req = httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status unauthorized got %d", resp.StatusCode)
	}
}
