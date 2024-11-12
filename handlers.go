package main

import (
    "encoding/json"
    "github.com/google/uuid"
    "net/http"
    "strings"
)

var receipts = make(map[string]ReceiptData)

func processReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    points := calculatePoints(receipt)
    id := uuid.New().String()
    receipts[id] = ReceiptData{Receipt: receipt, Points: points}

    json.NewEncoder(w).Encode(ReceiptResponse{ID: id})
}

func getPoints(w http.ResponseWriter, r *http.Request) {
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) != 4 || pathParts[2] != "points" {
        http.Error(w, "Invalid URL format", http.StatusNotFound)
        return
    }

    id := pathParts[1]
    receiptData, found := receipts[id]
    if !found {
        http.Error(w, "Receipt not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(PointsResponse{Points: receiptData.Points})
}