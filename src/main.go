package main

func main() {

	app, err := BuildApp()
	if err != nil {
		panic(err)
	}

	app.Server.ListenAndServe(app.Config.Address)
}
