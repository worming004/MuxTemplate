package main

type dummy struct {
	Id   int
	Name string
	Body subDummy
}

type subDummy struct {
	Body string
}

func buildDummy() dummy {
	return dummy{1, "DummyName", subDummy{Body: "My trunk"}}
}
