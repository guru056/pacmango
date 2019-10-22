package pacman

func isWall(e elem) bool {
	return e>=w0 && e<=w24
}

func handleError(e error){
	if(e != nil){
		panic(e)
	}
}

