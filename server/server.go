package server

func Init() {
	rtr := NewRouter()
	rtr.Run(":8000")
}
