package ecom

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type Payload struct {
	CustomerID    string   `json:"customerId"`
	Products      []string `json:"products"`
	Offers        []string `json:"offers"`
	PaymentMethod string   `json:"paymentMethod"`
}

type Product struct {
	ID          string `json:"id"`
	Company     string `json:"company" fako:"company"`
	Product     string `json:"product" fako:"product"`
	ProductName string `json:"product_name" fako:"product_name"`
	Available   bool   `json:"available"`
	Cost        int    `json:"cost"`
	Discount    string `json:"discount"`
}

type Price struct {
	OriginalPrice   string `json:"original_price"`
	DiscountedPrice string `json:"discounted_price"`
}

type Pay struct {
	Status bool   `json:"status"`
	ID     string `json:"id"`
}

type Output struct {
	Customer User      `json:"customer"`
	Products []Product `json:"products"`
	Payment  bool      `json:"payment"`
	Shipping bool      `json:"shipping"`
	ID       string    `json:"result_id"`
}

func GetUserActivity(ctx context.Context, userID string) (*User, error) {
	resp, err := http.Get(userServiceServerHostPort + "/user?id=" + userID)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var obj User
	err = json.Unmarshal(body, &obj)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("--------\n", obj)
	}

	// activity.GetLogger(ctx).Info("User details for user.", zap.String("UserID", userID))

	return &obj, nil
}

func GetProductActivitiy(ctx context.Context, productID string) (*Product, error) {

	resp, err := http.Get(productServiceHostPort + "/product?id=" + productID)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var obj Product
	err = json.Unmarshal(body, &obj)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("--------\n", obj)
	}

	// activity.GetLogger(ctx).Info("Product details for user.", zap.String("ProductID", productID))

	return &obj, nil

}

func DiscountActivitiy(ctx context.Context, productID string) (*Price, error) {

	resp, err := http.Get(offerServiceHostPort + "/offer?value=" + productID)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var obj Price
	err = json.Unmarshal(body, &obj)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("--------\n", obj)
	}

	// activity.GetLogger(ctx).Info("Product details for user.", zap.String("ProductID", productID))

	return &obj, nil

}

func PaymentActivitiy(ctx context.Context) (*Pay, error) {

	resp, err := http.Get(paymentHostPort + "/payment")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var obj Pay
	err = json.Unmarshal(body, &obj)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("--------\n", obj)
	}

	// activity.GetLogger(ctx).Info("Product details for user.", zap.String("ProductID", productID))

	return &obj, nil

}
