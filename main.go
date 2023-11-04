package main

func main() {
	/*
	driver := storage.MySQL
	storage.New(driver)
	
	// GORM perform single creat, update, delete operations 
	// in transactions by default to ensue database data integrity
	// P.S. Foreign keys constraints were created since migration	
	invoice := model.InvoiceHeader{
		Client: "Roberto Guzmán",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
			{ProductID: 2},
		},
	}

	storage.DB().Create(&invoice)
	*/
	/*
	// Delete from db - permanent
	myProduct := model.Product{}
	myProduct.ID = 3

	storage.DB().Unscoped().Delete(&myProduct)
	*/
	/*
	// Delete Soft - Update field "DeletedAt"
	myProduct := model.Product{}
	myProduct.ID = 3

	storage.DB().Delete(&myProduct)
	*/
	/*
	myProduct := model.Product{}
	myProduct.ID = 3

	storage.DB().Model(&myProduct).Updates(
		model.Product{Name: "Curso de Java", Price: 120},
	)
	*/
	/*
	// Get single row
	myProduct := model.Product{}

	storage.DB().First(&myProduct, 3)
	fmt.Println(myProduct)
	*/
	/*
	// Get multiple rows
	products := make([]model.Product, 0)
	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}
	*/
	/*
	// Create product
	product1 := model.Product{
		Name: "Curso de Go",
		Price: 120,
	}

	obs := "testing with go"
	product2 := model.Product{
		Name: "Curso de testing",
		Price: 150,
		Observations: &obs,
	}

	product3 := model.Product{
		Name: "Curso de Python",
		Price: 200,
	}

	storage.DB().Create(&product1)
	storage.DB().Create(&product2)
	storage.DB().Create(&product3)
	*/
	/*
	// Migrations
	storage.DB().AutoMigrate(
		&model.Product{}, 
		&model.InvoiceHeader{}, 
		&model.InvoiceItem{},
	)
	*/
}