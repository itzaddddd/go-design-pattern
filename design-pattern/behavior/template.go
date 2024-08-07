package behavior

import "fmt"

type ReportTemplate interface {
	collectData()
	analyzeData()
	presentData()
}

func generateReport(reportTemplate ReportTemplate) {
	reportTemplate.collectData()
	reportTemplate.analyzeData()
	reportTemplate.presentData()
}

type SalesReport struct{}

func (r *SalesReport) collectData() {
	fmt.Println("collect sale data")
}

func (r *SalesReport) analyzeData() {
	fmt.Println("analyze sale data")
}

func (r *SalesReport) presentData() {
	fmt.Println("present sale data")
}

type InventoryReport struct{}

func (r *InventoryReport) collectData() {
	fmt.Println("collect inventory data")
}

func (r *InventoryReport) analyzeData() {
	fmt.Println("analyze inventory data")
}

func (r *InventoryReport) presentData() {
	fmt.Println("present inventory data")
}

func RunTemplate() {
	salesReport := SalesReport{}
	generateReport(&salesReport)

	inventoryReport := InventoryReport{}
	generateReport(&inventoryReport)
}
