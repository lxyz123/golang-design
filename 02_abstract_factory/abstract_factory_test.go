package main

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderDetailDAO().SaveOrderDetail()
	factory.CreateOrderMainDAO().SaveOrderMain()
}

func ExampleRdbFactory() {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
}

func ExampleXmlFactory() {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
}
