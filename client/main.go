package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {

	log.Println("client running...")

	ctx := context.TODO()
	var wg sync.WaitGroup

	wg.Add(1)
	go job(ctx, &wg)
	wg.Wait()

	log.Println("job end")

}

func job(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()

		default:
			generateRequest()
		}
	}
}

func generateRequest() {
	jsonString := `{"first_name": "JOPA", "last_name": "PISA", "payment_mode": "CHEQUE", "payment_ref_no": "985", "amount": 985.65}`

	j := strings.NewReader(jsonString)

	resp, err := http.Post("http://127.0.0.1:8080/payments", "application/json", j)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("request generated...")
		log.Println("response: ", resp.StatusCode)
	}

}
