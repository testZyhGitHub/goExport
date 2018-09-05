package main

import (
	"fmt"
)

const (
	SetEmptyQuoteErrCode = 44500 + iota
	LaborOrderNotExistErrCode
	RejectOperErrCode
	RecruitTmpNotExistErrCode
	RecruitNotExistErrCode
	SubsidyIllegalErrCode
	EnrollFeeIllegalErrCode
	LabourCostIllegalErrCode
	QuoteNotExistErrCode
	QuoteNotHasSubsidyErrCode
	CancelHistoryQuoteErrCode
	LaborOrderSubsidyNotExistErrCode
	searchRecruitAuditFlowFailErrCode
	searchLabourCostFailErrCode
	searchEnrollFeeFailErrCode
	searchRecruitSubsidyFailErrCode
	searchGiftConfFailErrCode
	searchStandardSubsidyFailErrCode
	searchRecruitTmpFailErrCode
	searchLaborOrderFailErrCode
	searchLaborFailErrCode
	searchEmployeeFailErrCode
	searchLaborOrderSubsidyFailErrCode
	searchLaborChargeFailErrCode
	searchCntRecruitAuditFlowFailErrCode
)

func main() {
	// var errcodeVal int 
	// errcodeVal := 0
	
	fmt.Println(SetEmptyQuoteErrCode)
	fmt.Printf("\n")
	
	fmt.Println(SubsidyIllegalErrCode)
	fmt.Printf("\n")
	
	fmt.Println(searchCntRecruitAuditFlowFailErrCode)
	fmt.Printf("\n")	
}