package main

import "encoding/hex"

type (
    productType string

		car struct {
			 *product
			 name   string
			 price    float64
			 engine bool
			 color bool
			 model bool
			 auto bool
		}

		//product is a product in a store.
		product interface {
        //returns the unique ID of the product.
				ID() productID
				//Type returns the product type.
				Type() productType
				//DisplayName returns the display name of the product.
				DisplayName()  string
				//Description returns brief information about the product.
				Discription() string
				//Price returns the price of the product.
				Price() float64
				//Category returns the category of the product.
				Category() string
				//Images returns a list of image urls of the product.
				Images() string 
				Isvalid()
	}
		
	//Car is a car and a type of product.
	Vehicle interface {
		Product
		Name() string
		make() string
		Mode() string
		BrandNew() bool
		Exteriorcolor() string
		Interiorcolor() string
		Engine() string
		Mileage() string
		Fueltype() string
		Features() string
	}

	Buyer struct {
		Name string
		Amountpaid float64
		Address string
	}

	Order struct {
		id string
		*Buyer
		products []product
	
	}

typecarID [32]byte
	
func (p productID) string() string {
	return hex.EncodeTostring (p[:])
}

func(c *car) Isvalid() bool {
	  return c.auto && c.color && c.engine && c.price !=0 && c.name != ""
}

type product struct {
	id productID
	name string
	price float64
}

//implements all products methods

type vehicle struct {
   *product
	 engine bool
}







