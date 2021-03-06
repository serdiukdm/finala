package testutils

import (
	"errors"
	"finala/api/storage"
	"time"
)

type MockStorage struct {
	Events int
}

func NewMockStorage() *MockStorage {

	return &MockStorage{
		Events: 0,
	}
}

func (ms *MockStorage) Save(data string) bool {
	ms.Events++
	return true
}

func (ms *MockStorage) GetSummary(executionsID string) (map[string]storage.CollectorsSummary, error) {

	if executionsID == "err" {
		return nil, errors.New("error")
	}
	response := map[string]storage.CollectorsSummary{
		"resource_1": storage.CollectorsSummary{
			ResourceName:  "resource_name_1",
			ResourceCount: 3,
			TotalSpent:    100,
			Status:        1,
			ErrorMessage:  "description",
			EventTime:     123456,
		},
		"resource_2": storage.CollectorsSummary{
			ResourceName:  "resource_name_2",
			ResourceCount: 3,
			TotalSpent:    100,
			Status:        1,
			ErrorMessage:  "description",
			EventTime:     123456,
		},
	}

	return response, nil
}

func (ms *MockStorage) GetExecutions() ([]storage.Executions, error) {
	response := []storage.Executions{
		{
			ID:   "1",
			Name: "Execution 1",
			Time: time.Now(),
		},
		{
			ID:   "2",
			Name: "Execution 2",
			Time: time.Now(),
		},
	}
	return response, nil
}

func (ms *MockStorage) GetResources(resourceType string, executionID string) ([]map[string]interface{}, error) {

	var response []map[string]interface{}

	if executionID == "err" {
		return nil, errors.New("error")
	}

	type tempStruct struct {
		Data string
	}

	rowData := make(map[string]interface{})
	rowData1 := make(map[string]interface{})
	rowData["test"] = tempStruct{Data: "1"}
	rowData1["test1"] = tempStruct{Data: "1"}
	response = append(response, rowData)
	response = append(response, rowData1)

	return response, nil

}
