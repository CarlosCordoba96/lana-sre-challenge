package shop

type Product struct {
	Code	string
	Name 	string
	Price	float64

}

var Products = [] Product{

	{
		"PEN",
		"Lana Pen",
		 5.00,
	},

	{
		"TSHIRT",
		"Lana T-shirt",
		20.00,
	},
	{
		"MUG",
		"Lana Coffe Mug",
		 7.50,
	},
}