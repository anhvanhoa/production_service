package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

func inputPaging(reader *bufio.Reader) (int32, int32) {
	fmt.Print("Enter offset (default 1): ")
	offsetStr, _ := reader.ReadString('\n')
	offsetStr = cleanInput(offsetStr)
	offset := int32(1)
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = int32(o)
		}
	}

	fmt.Print("Enter limit (default 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	return offset, limit
}

type CropServiceClient struct {
	harvestRecordClient     proto_harvest_record.HarvestRecordServiceClient
	pestDiseaseRecordClient proto_pest_disease_record.PestDiseaseRecordServiceClient
	conn                    *grpc.ClientConn
}

func NewCropServiceClient(address string) (*CropServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &CropServiceClient{
		harvestRecordClient:     proto_harvest_record.NewHarvestRecordServiceClient(conn),
		pestDiseaseRecordClient: proto_pest_disease_record.NewPestDiseaseRecordServiceClient(conn),
		conn:                    conn,
	}, nil
}

func (c *CropServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Harvest Record Service Tests ==================

func (c *CropServiceClient) TestCreateHarvestRecord() {
	fmt.Println("\n=== Test Create Harvest Record ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter planting cycle ID: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	fmt.Print("Enter harvest date (YYYY-MM-DD): ")
	harvestDateStr, _ := reader.ReadString('\n')
	harvestDateStr = cleanInput(harvestDateStr)
	var harvestDate *timestamppb.Timestamp
	if harvestDateStr != "" {
		if t, err := time.Parse("2006-01-02", harvestDateStr); err == nil {
			harvestDate = timestamppb.New(t)
		}
	}

	fmt.Print("Enter quantity (kg): ")
	quantityStr, _ := reader.ReadString('\n')
	quantityStr = cleanInput(quantityStr)
	quantity := float64(100.0)
	if quantityStr != "" {
		if q, err := strconv.ParseFloat(quantityStr, 64); err == nil {
			quantity = q
		}
	}

	fmt.Print("Enter quality grade: ")
	qualityGrade, _ := reader.ReadString('\n')
	qualityGrade = cleanInput(qualityGrade)

	fmt.Print("Enter size classification: ")
	sizeClassification, _ := reader.ReadString('\n')
	sizeClassification = cleanInput(sizeClassification)

	fmt.Print("Enter market price per kg: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = cleanInput(priceStr)
	price := float64(10.0)
	if priceStr != "" {
		if p, err := strconv.ParseFloat(priceStr, 64); err == nil {
			price = p
		}
	}

	fmt.Print("Enter labor hours: ")
	laborHoursStr, _ := reader.ReadString('\n')
	laborHoursStr = cleanInput(laborHoursStr)
	laborHours := float64(8.0)
	if laborHoursStr != "" {
		if h, err := strconv.ParseFloat(laborHoursStr, 64); err == nil {
			laborHours = h
		}
	}

	fmt.Print("Enter labor cost: ")
	laborCostStr, _ := reader.ReadString('\n')
	laborCostStr = cleanInput(laborCostStr)
	laborCost := float64(50.0)
	if laborCostStr != "" {
		if c, err := strconv.ParseFloat(laborCostStr, 64); err == nil {
			laborCost = c
		}
	}

	fmt.Print("Enter packaging cost: ")
	packagingCostStr, _ := reader.ReadString('\n')
	packagingCostStr = cleanInput(packagingCostStr)
	packagingCost := float64(10.0)
	if packagingCostStr != "" {
		if c, err := strconv.ParseFloat(packagingCostStr, 64); err == nil {
			packagingCost = c
		}
	}

	fmt.Print("Enter storage location: ")
	storageLocation, _ := reader.ReadString('\n')
	storageLocation = cleanInput(storageLocation)

	fmt.Print("Enter storage temperature: ")
	tempStr, _ := reader.ReadString('\n')
	tempStr = cleanInput(tempStr)
	temperature := float64(4.0)
	if tempStr != "" {
		if t, err := strconv.ParseFloat(tempStr, 64); err == nil {
			temperature = t
		}
	}

	fmt.Print("Enter buyer information: ")
	buyerInfo, _ := reader.ReadString('\n')
	buyerInfo = cleanInput(buyerInfo)

	fmt.Print("Enter weather at harvest: ")
	weather, _ := reader.ReadString('\n')
	weather = cleanInput(weather)

	fmt.Print("Enter plant health rating (1-5): ")
	ratingStr, _ := reader.ReadString('\n')
	ratingStr = cleanInput(ratingStr)
	rating := int32(5)
	if ratingStr != "" {
		if r, err := strconv.Atoi(ratingStr); err == nil {
			rating = int32(r)
		}
	}

	fmt.Print("Enter notes: ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	fmt.Print("Enter images: ")
	images, _ := reader.ReadString('\n')
	images = cleanInput(images)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.harvestRecordClient.CreateHarvestRecord(ctx, &proto_harvest_record.CreateHarvestRecordRequest{
		PlantingCycleId:    plantingCycleId,
		HarvestDate:        harvestDate,
		QuantityKg:         quantity,
		QualityGrade:       qualityGrade,
		SizeClassification: sizeClassification,
		MarketPricePerKg:   price,
		LaborHours:         laborHours,
		LaborCost:          laborCost,
		PackagingCost:      packagingCost,
		StorageLocation:    storageLocation,
		StorageTemperature: temperature,
		BuyerInformation:   buyerInfo,
		WeatherAtHarvest:   weather,
		PlantHealthRating:  rating,
		Notes:              notes,
		Images:             images,
		CreatedBy:          createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateHarvestRecord: %v\n", err)
		return
	}

	fmt.Printf("Create Harvest Record result:\n")
	if resp.HarvestRecord != nil {
		fmt.Printf("ID: %s\n", resp.HarvestRecord.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.HarvestRecord.PlantingCycleId)
		fmt.Printf("Quantity: %.2f kg\n", resp.HarvestRecord.QuantityKg)
		fmt.Printf("Quality Grade: %s\n", resp.HarvestRecord.QualityGrade)
		fmt.Printf("Size Classification: %s\n", resp.HarvestRecord.SizeClassification)
		fmt.Printf("Market Price: %.2f per kg\n", resp.HarvestRecord.MarketPricePerKg)
		fmt.Printf("Total Revenue: %.2f\n", resp.HarvestRecord.TotalRevenue)
	}
}

func (c *CropServiceClient) TestGetHarvestRecord() {
	fmt.Println("\n=== Test Get Harvest Record ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter harvest record ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.harvestRecordClient.GetHarvestRecord(ctx, &proto_harvest_record.GetHarvestRecordRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetHarvestRecord: %v\n", err)
		return
	}

	fmt.Printf("Get Harvest Record result:\n")
	if resp.HarvestRecord != nil {
		fmt.Printf("ID: %s\n", resp.HarvestRecord.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.HarvestRecord.PlantingCycleId)
		fmt.Printf("Quantity: %.2f kg\n", resp.HarvestRecord.QuantityKg)
		fmt.Printf("Quality Grade: %s\n", resp.HarvestRecord.QualityGrade)
		fmt.Printf("Size Classification: %s\n", resp.HarvestRecord.SizeClassification)
		fmt.Printf("Market Price: %.2f per kg\n", resp.HarvestRecord.MarketPricePerKg)
		fmt.Printf("Total Revenue: %.2f\n", resp.HarvestRecord.TotalRevenue)
		fmt.Printf("Labor Hours: %.2f\n", resp.HarvestRecord.LaborHours)
		fmt.Printf("Labor Cost: %.2f\n", resp.HarvestRecord.LaborCost)
		fmt.Printf("Packaging Cost: %.2f\n", resp.HarvestRecord.PackagingCost)
		fmt.Printf("Storage Location: %s\n", resp.HarvestRecord.StorageLocation)
		fmt.Printf("Storage Temperature: %.2f°C\n", resp.HarvestRecord.StorageTemperature)
		fmt.Printf("Buyer Information: %s\n", resp.HarvestRecord.BuyerInformation)
		fmt.Printf("Weather at Harvest: %s\n", resp.HarvestRecord.WeatherAtHarvest)
		fmt.Printf("Plant Health Rating: %d\n", resp.HarvestRecord.PlantHealthRating)
		fmt.Printf("Notes: %s\n", resp.HarvestRecord.Notes)
		fmt.Printf("Created By: %s\n", resp.HarvestRecord.CreatedBy)
	}
}

func (c *CropServiceClient) TestListHarvestRecords() {
	fmt.Println("\n=== Test List Harvest Records ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.harvestRecordClient.ListHarvestRecords(ctx, &proto_harvest_record.ListHarvestRecordsRequest{
		Pagination: &proto_common.PaginationRequest{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "created_at",
			SortOrder: "desc",
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListHarvestRecords: %v\n", err)
		return
	}

	fmt.Printf("List Harvest Records result:\n")
	fmt.Printf("Total: %d\n", resp.Pagination.Total)
	fmt.Printf("Harvest Records:\n")
	for i, record := range resp.HarvestRecords {
		fmt.Printf("  [%d] ID: %s, Planting Cycle: %s, Quantity: %.2f kg, Quality: %s\n",
			i+1, record.Id, record.PlantingCycleId, record.QuantityKg, record.QualityGrade)
	}
}

func (c *CropServiceClient) TestGetHarvestRecordsByPlantingCycle() {
	fmt.Println("\n=== Test Get Harvest Records By Planting Cycle ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter planting cycle ID: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.harvestRecordClient.GetHarvestRecordsByPlantingCycle(ctx, &proto_harvest_record.GetHarvestRecordsByPlantingCycleRequest{
		PlantingCycleId: plantingCycleId,
		Pagination: &proto_common.PaginationRequest{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "harvest_date",
			SortOrder: "desc",
		},
	})
	if err != nil {
		fmt.Printf("Error calling GetHarvestRecordsByPlantingCycle: %v\n", err)
		return
	}

	fmt.Printf("Get Harvest Records By Planting Cycle result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Harvest Records:\n")
	for i, record := range resp.HarvestRecords {
		fmt.Printf("  [%d] ID: %s, Quantity: %.2f kg, Quality: %s, Revenue: %.2f\n",
			i+1, record.Id, record.QuantityKg, record.QualityGrade, record.TotalRevenue)
	}
}

// ================== Pest Disease Record Service Tests ==================

func (c *CropServiceClient) TestCreatePestDiseaseRecord() {
	fmt.Println("\n=== Test Create Pest Disease Record ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter planting cycle ID: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	fmt.Print("Enter type (pest/disease/nutrient_deficiency/environmental_stress): ")
	recordType, _ := reader.ReadString('\n')
	recordType = cleanInput(recordType)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter scientific name: ")
	scientificName, _ := reader.ReadString('\n')
	scientificName = cleanInput(scientificName)

	fmt.Print("Enter severity (low/medium/high/critical): ")
	severity, _ := reader.ReadString('\n')
	severity = cleanInput(severity)

	fmt.Print("Enter affected area percentage: ")
	areaStr, _ := reader.ReadString('\n')
	areaStr = cleanInput(areaStr)
	area := float64(10.0)
	if areaStr != "" {
		if a, err := strconv.ParseFloat(areaStr, 64); err == nil {
			area = a
		}
	}

	fmt.Print("Enter affected plant count: ")
	countStr, _ := reader.ReadString('\n')
	countStr = cleanInput(countStr)
	count := int32(10)
	if countStr != "" {
		if c, err := strconv.Atoi(countStr); err == nil {
			count = int32(c)
		}
	}

	fmt.Print("Enter detection date (YYYY-MM-DD): ")
	detectionDateStr, _ := reader.ReadString('\n')
	detectionDateStr = cleanInput(detectionDateStr)
	var detectionDate *timestamppb.Timestamp
	if detectionDateStr != "" {
		if t, err := time.Parse("2006-01-02", detectionDateStr); err == nil {
			detectionDate = timestamppb.New(t)
		}
	}

	fmt.Print("Enter detection method: ")
	detectionMethod, _ := reader.ReadString('\n')
	detectionMethod = cleanInput(detectionMethod)

	fmt.Print("Enter symptoms: ")
	symptoms, _ := reader.ReadString('\n')
	symptoms = cleanInput(symptoms)

	fmt.Print("Enter treatment applied: ")
	treatment, _ := reader.ReadString('\n')
	treatment = cleanInput(treatment)

	fmt.Print("Enter treatment date (YYYY-MM-DD): ")
	treatmentDateStr, _ := reader.ReadString('\n')
	treatmentDateStr = cleanInput(treatmentDateStr)
	var treatmentDate *timestamppb.Timestamp
	if treatmentDateStr != "" {
		if t, err := time.Parse("2006-01-02", treatmentDateStr); err == nil {
			treatmentDate = timestamppb.New(t)
		}
	}

	fmt.Print("Enter treatment cost: ")
	costStr, _ := reader.ReadString('\n')
	costStr = cleanInput(costStr)
	cost := float64(50.0)
	if costStr != "" {
		if c, err := strconv.ParseFloat(costStr, 64); err == nil {
			cost = c
		}
	}

	fmt.Print("Enter treatment duration days: ")
	durationStr, _ := reader.ReadString('\n')
	durationStr = cleanInput(durationStr)
	duration := int32(7)
	if durationStr != "" {
		if d, err := strconv.Atoi(durationStr); err == nil {
			duration = int32(d)
		}
	}

	fmt.Print("Enter recovery status: ")
	recoveryStatus, _ := reader.ReadString('\n')
	recoveryStatus = cleanInput(recoveryStatus)

	fmt.Print("Enter effectiveness rating (1-5): ")
	ratingStr, _ := reader.ReadString('\n')
	ratingStr = cleanInput(ratingStr)
	rating := int32(4)
	if ratingStr != "" {
		if r, err := strconv.Atoi(ratingStr); err == nil {
			rating = int32(r)
		}
	}

	fmt.Print("Enter follow-up date (YYYY-MM-DD): ")
	followUpDateStr, _ := reader.ReadString('\n')
	followUpDateStr = cleanInput(followUpDateStr)
	var followUpDate *timestamppb.Timestamp
	if followUpDateStr != "" {
		if t, err := time.Parse("2006-01-02", followUpDateStr); err == nil {
			followUpDate = timestamppb.New(t)
		}
	}

	fmt.Print("Enter prevention measures: ")
	prevention, _ := reader.ReadString('\n')
	prevention = cleanInput(prevention)

	fmt.Print("Enter environmental factors: ")
	environmental, _ := reader.ReadString('\n')
	environmental = cleanInput(environmental)

	fmt.Print("Enter images: ")
	images, _ := reader.ReadString('\n')
	images = cleanInput(images)

	fmt.Print("Enter notes: ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.pestDiseaseRecordClient.CreatePestDiseaseRecord(ctx, &proto_pest_disease_record.CreatePestDiseaseRecordRequest{
		PlantingCycleId:        plantingCycleId,
		Type:                   recordType,
		Name:                   name,
		ScientificName:         scientificName,
		Severity:               severity,
		AffectedAreaPercentage: area,
		AffectedPlantCount:     count,
		DetectionDate:          detectionDate,
		DetectionMethod:        detectionMethod,
		Symptoms:               symptoms,
		TreatmentApplied:       treatment,
		TreatmentDate:          treatmentDate,
		TreatmentCost:          cost,
		TreatmentDurationDays:  duration,
		RecoveryStatus:         recoveryStatus,
		EffectivenessRating:    rating,
		FollowUpDate:           followUpDate,
		PreventionMeasures:     prevention,
		EnvironmentalFactors:   environmental,
		Images:                 images,
		Notes:                  notes,
		CreatedBy:              createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreatePestDiseaseRecord: %v\n", err)
		return
	}

	fmt.Printf("Create Pest Disease Record result:\n")
	if resp.PestDiseaseRecord != nil {
		fmt.Printf("ID: %s\n", resp.PestDiseaseRecord.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.PestDiseaseRecord.PlantingCycleId)
		fmt.Printf("Type: %s\n", resp.PestDiseaseRecord.Type)
		fmt.Printf("Name: %s\n", resp.PestDiseaseRecord.Name)
		fmt.Printf("Scientific Name: %s\n", resp.PestDiseaseRecord.ScientificName)
		fmt.Printf("Severity: %s\n", resp.PestDiseaseRecord.Severity)
		fmt.Printf("Affected Area: %.2f%%\n", resp.PestDiseaseRecord.AffectedAreaPercentage)
		fmt.Printf("Affected Plants: %d\n", resp.PestDiseaseRecord.AffectedPlantCount)
		fmt.Printf("Detection Method: %s\n", resp.PestDiseaseRecord.DetectionMethod)
		fmt.Printf("Treatment Applied: %s\n", resp.PestDiseaseRecord.TreatmentApplied)
		fmt.Printf("Treatment Cost: %.2f\n", resp.PestDiseaseRecord.TreatmentCost)
		fmt.Printf("Recovery Status: %s\n", resp.PestDiseaseRecord.RecoveryStatus)
		fmt.Printf("Effectiveness Rating: %d\n", resp.PestDiseaseRecord.EffectivenessRating)
	}
}

func (c *CropServiceClient) TestGetPestDiseaseRecord() {
	fmt.Println("\n=== Test Get Pest Disease Record ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter pest disease record ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.pestDiseaseRecordClient.GetPestDiseaseRecord(ctx, &proto_pest_disease_record.GetPestDiseaseRecordRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetPestDiseaseRecord: %v\n", err)
		return
	}

	fmt.Printf("Get Pest Disease Record result:\n")
	if resp.PestDiseaseRecord != nil {
		fmt.Printf("ID: %s\n", resp.PestDiseaseRecord.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.PestDiseaseRecord.PlantingCycleId)
		fmt.Printf("Type: %s\n", resp.PestDiseaseRecord.Type)
		fmt.Printf("Name: %s\n", resp.PestDiseaseRecord.Name)
		fmt.Printf("Scientific Name: %s\n", resp.PestDiseaseRecord.ScientificName)
		fmt.Printf("Severity: %s\n", resp.PestDiseaseRecord.Severity)
		fmt.Printf("Affected Area: %.2f%%\n", resp.PestDiseaseRecord.AffectedAreaPercentage)
		fmt.Printf("Affected Plants: %d\n", resp.PestDiseaseRecord.AffectedPlantCount)
		fmt.Printf("Detection Method: %s\n", resp.PestDiseaseRecord.DetectionMethod)
		fmt.Printf("Symptoms: %s\n", resp.PestDiseaseRecord.Symptoms)
		fmt.Printf("Treatment Applied: %s\n", resp.PestDiseaseRecord.TreatmentApplied)
		fmt.Printf("Treatment Cost: %.2f\n", resp.PestDiseaseRecord.TreatmentCost)
		fmt.Printf("Treatment Duration: %d days\n", resp.PestDiseaseRecord.TreatmentDurationDays)
		fmt.Printf("Recovery Status: %s\n", resp.PestDiseaseRecord.RecoveryStatus)
		fmt.Printf("Effectiveness Rating: %d\n", resp.PestDiseaseRecord.EffectivenessRating)
		fmt.Printf("Prevention Measures: %s\n", resp.PestDiseaseRecord.PreventionMeasures)
		fmt.Printf("Environmental Factors: %s\n", resp.PestDiseaseRecord.EnvironmentalFactors)
		fmt.Printf("Notes: %s\n", resp.PestDiseaseRecord.Notes)
		fmt.Printf("Created By: %s\n", resp.PestDiseaseRecord.CreatedBy)
	}
}

func (c *CropServiceClient) TestListPestDiseaseRecords() {
	fmt.Println("\n=== Test List Pest Disease Records ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.pestDiseaseRecordClient.ListPestDiseaseRecords(ctx, &proto_pest_disease_record.ListPestDiseaseRecordsRequest{
		Pagination: &proto_pest_disease_record.Pagination{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "detection_date",
			SortOrder: "desc",
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListPestDiseaseRecords: %v\n", err)
		return
	}

	fmt.Printf("List Pest Disease Records result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Pest Disease Records:\n")
	for i, record := range resp.PestDiseaseRecords {
		fmt.Printf("  [%d] ID: %s, Type: %s, Name: %s, Severity: %s\n",
			i+1, record.Id, record.Type, record.Name, record.Severity)
	}
}

func (c *CropServiceClient) TestGetPestDiseaseRecordsByPlantingCycle() {
	fmt.Println("\n=== Test Get Pest Disease Records By Planting Cycle ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter planting cycle ID: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.pestDiseaseRecordClient.GetPestDiseaseRecordsByPlantingCycle(ctx, &proto_pest_disease_record.GetPestDiseaseRecordsByPlantingCycleRequest{
		PlantingCycleId: plantingCycleId,
		Pagination: &proto_pest_disease_record.Pagination{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "detection_date",
			SortOrder: "desc",
		},
	})
	if err != nil {
		fmt.Printf("Error calling GetPestDiseaseRecordsByPlantingCycle: %v\n", err)
		return
	}

	fmt.Printf("Get Pest Disease Records By Planting Cycle result:\n")
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Pest Disease Records:\n")
	for i, record := range resp.PestDiseaseRecords {
		fmt.Printf("  [%d] ID: %s, Type: %s, Name: %s, Severity: %s, Status: %s\n",
			i+1, record.Id, record.Type, record.Name, record.Severity, record.RecoveryStatus)
	}
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== gRPC Production Service Test Client ===")
	fmt.Println("1. Harvest Record Service")
	fmt.Println("2. Pest Disease Record Service")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

func printHarvestRecordMenu() {
	fmt.Println("\n=== Harvest Record Service ===")
	fmt.Println("1. Create Harvest Record")
	fmt.Println("2. Get Harvest Record")
	fmt.Println("3. List Harvest Records")
	fmt.Println("4. Get Harvest Records By Planting Cycle")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printPestDiseaseRecordMenu() {
	fmt.Println("\n=== Pest Disease Record Service ===")
	fmt.Println("1. Create Pest Disease Record")
	fmt.Println("2. Get Pest Disease Record")
	fmt.Println("3. List Pest Disease Records")
	fmt.Println("4. Get Pest Disease Records By Planting Cycle")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Connecting to gRPC server at %s...\n", address)
	client, err := NewCropServiceClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Connected successfully!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		switch choice {
		case "1":
			// Harvest Record Service
			for {
				printHarvestRecordMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateHarvestRecord()
				case "2":
					client.TestGetHarvestRecord()
				case "3":
					client.TestListHarvestRecords()
				case "4":
					client.TestGetHarvestRecordsByPlantingCycle()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "2":
			// Pest Disease Record Service
			for {
				printPestDiseaseRecordMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreatePestDiseaseRecord()
				case "2":
					client.TestGetPestDiseaseRecord()
				case "3":
					client.TestListPestDiseaseRecords()
				case "4":
					client.TestGetPestDiseaseRecordsByPlantingCycle()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
