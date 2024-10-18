package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/anjuna-security/go-nitro-attestation/verifier"
)

type AttestationResult struct {
	Technology string
	Valid bool
	Measurements map[string]string
	UserData string
	Nonce string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <path-to-attestation-report>\n", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	report, err := verifier.NewSignedAttestationReport(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	// Validate the report is signed by the AWS root of trust
	if err = verifier.Validate(report, nil); err != nil {
		panic(err)
	}

	// Access the PCR values as hex strings
	result := AttestationResult{
		Technology: "aws_nitro_enclaves",
		Valid: true,
		Measurements: make(map[string]string),
		UserData: string(report.Document.UserData),
		Nonce: string(report.Document.UserNonce),
	}

	hexPCRs := verifier.ConvertPCRsToHex(report.Document.PCRs)
	for i, v := range hexPCRs {
		result.Measurements[fmt.Sprintf("PCR%d", i)] = v
	}

	jsonByte, err := json.Marshal(result)
	if err != nil{
		panic(err)
	}
	jsonString := string(jsonByte)
	fmt.Println(jsonString)
}
