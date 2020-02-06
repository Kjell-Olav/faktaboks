package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	Kjennsgjerning := "Bjørnedyr er mikroskopopiske dyr med fasong som kan minne om en bjørn. De er utbredt over hele jorden og overlever det meste. De kan gå i en slags dyp dvale de overlever nedfrysning til -200 grader Celsius, vakuum og kosmisk stråling." 
	fmt.Fprintf(w, Kjennsgjerning)
}
