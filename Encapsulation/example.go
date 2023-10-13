package main

func AddToShoppingCart() {
	checkUserStatus(2)
	checkProductStatus(5)
	productQuantity(5, 4)
	addToCart(5, 2, 4)
}

func checkUserStatus(ID int) {
	//check user status
}

func checkProductStatus(ProductID int) {
	//check product id
}

func productQuantity(ProductID int, Quantity int) {
	//check product quantity
}

func addToCart(ProductID int, ID int, Quantity int) {
	//add to cart
}

//the methods with first lowercase letter are not accessible  from out of the package
