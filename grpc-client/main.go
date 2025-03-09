package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/andykimchris/grpc-cc/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewInvoicerClient(conn)

	// timeout to stop grpc request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Create a new invoice request
	req := &pb.CreateRequest{
		From:   "Company A",
		To:     "Company B",
		VATNum: "123456789",
		// PaymentAmount: &pb.CreateRequest_SingleAmount{
		// 	SingleAmount: &pb.Amount{
		// 		Amount:   20400,
		// 		Currency: "USD",
		// 	},
		// },
		PaymentAmount: &pb.CreateRequest_MultipleAmounts{
			MultipleAmounts: &pb.MultipleAmounts{
				Amounts: []*pb.Amount{
					{
						Amount:   20400,
						Currency: "USD",
					},
					{
						Amount:   20400,
						Currency: "EUR",
					},
				},
			},
		},
	}

	resp, err := client.Create(ctx, req)
	if err != nil {
		log.Fatalf("could not create invoice: %v", err)
	}
	fmt.Printf("Invoice created witth info: Docx %s, Pdf %s\n", resp.Docx, resp.Pdf)

}
