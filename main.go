package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/andykimchris/grpc-cc/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type InvoicerServer interface {
	Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error)
	SumNums(context.Context, *invoicer.MultipleAmounts) (*invoicer.SumsResponse, error)
	ExchangeConverter(context.Context, *invoicer.ExchangeRequest) (*invoicer.ExchangeResponse, error)
	UploadInvoices(context.Context, *invoicer.InvoiceRequest) (*invoicer.UploadSummaryResponse, error)
	ListInvoices(context.Context, *invoicer.Empty) (*invoicer.InvoiceRequest, error)
}

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	var totalAmt int64
	var currencyVal string

	switch v := req.PaymentAmount.(type) {
	case *invoicer.CreateRequest_SingleAmount:
		totalAmt = v.SingleAmount.Amount
		currencyVal = v.SingleAmount.Currency
	case *invoicer.CreateRequest_MultipleAmounts:
		for _, amount := range v.MultipleAmounts.Amounts {
			totalAmt += amount.Amount
			currencyVal = amount.Currency
		}
	default:
		return nil, fmt.Errorf("NO AMOUNT PROVIDED")
	}

	fmt.Printf(("Final result %d %s\n"), totalAmt, currencyVal)
	responseMetadata := map[string]string{
		"invoice_id": "1234",
		"status":     "paid",
		"total":      fmt.Sprintf("%d", totalAmt),
		"timestamp":  time.Now().Format(time.RFC3339),
	}

	metadataBytes, err := json.Marshal(responseMetadata)
	if err != nil {
		log.Fatalf("unable to marshal metadata: %s", err)
	}

	return &invoicer.CreateResponse{
		Pdf:      []byte(req.From),
		Docx:     []byte(req.To),
		Metadata: metadataBytes,
	}, nil
}

func (s myInvoicerServer) SumNums(ctx context.Context, req *invoicer.MultipleAmounts) (*invoicer.SumsResponse, error) {
	var count int64
	for _, amount := range req.Amounts {
		count = count + amount.Amount
	}
	return &invoicer.SumsResponse{Total: count}, nil
}

func (s myInvoicerServer) ExchangeConverter(ctx context.Context, req *invoicer.ExchangeRequest) (*invoicer.ExchangeResponse, error) {
	switch req.TargetCurrency {
	case "USD":
		amnt := req.Source.Amount / 120
		return &invoicer.ExchangeResponse{
			Amount:   amnt,
			Currency: req.TargetCurrency,
		}, nil
	case "EUR":
		amnt := req.Source.Amount / 140
		return &invoicer.ExchangeResponse{
			Amount:   amnt,
			Currency: req.TargetCurrency,
		}, nil
	default:
		return nil, nil
	}
}

// UploadInvoices handles client-side streaming of invoices. The server will
// accept a stream of InvoiceRequest messages and send a single
// UploadSummaryResponse message back to the client once the stream is complete.
func (s myInvoicerServer) UploadInvoices(stream invoicer.Invoicer_UploadInvoicesServer) error {
	var totalAmount int64

	for {
		invoice, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Final computed total: %d %s", totalAmount, "USD")
			return stream.SendAndClose(&invoicer.UploadSummaryResponse{TotalAmount: totalAmount, Currency: "USD"})
		}

		if err != nil {
			log.Fatalf("unable to receive stream: %s", err)
		}

		fmt.Printf("Received invoice: %v, %d, \n", invoice, totalAmount)
		totalAmount += invoice.Amount.Amount
	}
}

func (s myInvoicerServer) ListInvoices(req *invoicer.Empty, stream invoicer.Invoicer_ListInvoicesServer) error {
	invoices := []*invoicer.InvoiceRequest{
		{
			InvoiceNumber: "INV-001",
			InvoiceName:   "Andy",
			Amount: &invoicer.Amount{
				Amount:   1500,
				Currency: "USD",
			},
		},
		{
			InvoiceNumber: "INV-002",
			InvoiceName:   "Jamie",
			Amount: &invoicer.Amount{
				Amount:   4500,
				Currency: "USD",
			},
		},
	}

	for _, invoice := range invoices {
		fmt.Println("Sending invoice: ", invoice.InvoiceNumber)
		if err := stream.Send(invoice); err != nil {
			log.Fatalf("could not send invoice: %v", err)
		}
		time.Sleep(time.Millisecond.Abs() * 1350)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot listen to port 8080: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	// enable reflection
	reflection.Register(serverRegistrar)

	if err := serverRegistrar.Serve(lis); err != nil {
		log.Fatalf("unable to serve: %s", err)
	}
	fmt.Println("Server started on port 8080")
}
