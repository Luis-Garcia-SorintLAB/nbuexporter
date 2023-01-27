package main

import "time"

type Primary struct {
	Fqdn      string `yaml:"fqdn"`
	APIKey    string `yaml:"apiKey"`
	TimeFrame struct {
		Num  int    `yaml:"num"`
		Freq string `yaml:"freq"`
	} `yaml:"timeFrame"`
}

type Entity struct {
	Name      string `yaml:"name"`
	Primaries []struct {
		Name string `yaml:"name"`
	} `yaml:"primaries"`
	Workloads []struct {
		Type                 string   `yaml:"type"`
		WorkloadDisplayName  []string `yaml:"workloadDisplayName,omitempty"`
		InstanceDatabaseName []string `yaml:"instanceDatabaseName,omitempty"`
	} `yaml:"workloads"`
}
type Data struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			JobID                      int       `json:"jobId"`
			ParentJobID                int       `json:"parentJobId"`
			ActiveProcessID            int       `json:"activeProcessId"`
			JobType                    string    `json:"jobType"`
			PolicyName                 string    `json:"policyName"`
			ScheduleName               string    `json:"scheduleName"`
			DestinationStorageUnitName string    `json:"destinationStorageUnitName"`
			DestinationMediaServerName string    `json:"destinationMediaServerName"`
			StreamNumber               int       `json:"streamNumber"`
			CopyNumber                 int       `json:"copyNumber"`
			Priority                   int       `json:"priority"`
			Compression                int       `json:"compression"`
			Status                     int       `json:"status"`
			State                      string    `json:"state"`
			NumberOfFiles              int       `json:"numberOfFiles"`
			EstimatedFiles             int       `json:"estimatedFiles"`
			KilobytesTransferred       int       `json:"kilobytesTransferred"`
			KilobytesToTransfer        int       `json:"kilobytesToTransfer"`
			TransferRate               int       `json:"transferRate"`
			PercentComplete            int       `json:"percentComplete"`
			Restartable                int       `json:"restartable"`
			Suspendable                int       `json:"suspendable"`
			Resumable                  int       `json:"resumable"`
			FrozenImage                int       `json:"frozenImage"`
			DedupRatio                 float64   `json:"dedupRatio"`
			CurrentOperation           int       `json:"currentOperation"`
			SessionID                  int       `json:"sessionId"`
			NumberOfTapeToEject        int       `json:"numberOfTapeToEject"`
			SubmissionType             int       `json:"submissionType"`
			AuditDomainType            int       `json:"auditDomainType"`
			StartTime                  time.Time `json:"startTime"`
			EndTime                    time.Time `json:"endTime"`
			ActiveTryStartTime         time.Time `json:"activeTryStartTime"`
			LastUpdateTime             time.Time `json:"lastUpdateTime"`
			Try                        int       `json:"try"`
			Cancellable                int       `json:"cancellable"`
			JobQueueReason             int       `json:"jobQueueReason"`
			KilobytesDataTransferred   int       `json:"kilobytesDataTransferred"`
			DteMode                    string    `json:"dteMode"`
			DedupSpaceRatio            float64   `json:"dedupSpaceRatio"`
			CompressionSpaceRatio      float64   `json:"compressionSpaceRatio"`
			ElapsedTime                string    `json:"elapsedTime"`
			AssetID                    string    `json:"assetID"`
			AssetDisplayableName       string    `json:"assetDisplayableName"`
			WorkloadDisplayName        string    `json:"workloadDisplayName"`
			PolicyType                 string    `json:"policyType"`
			InstanceDatabaseName       string    `json:"instanceDatabaseName"`
			ScheduleType               string    `json:"scheduleType"`
			Entity                     []string
		} `json:"attributes"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			FileLists struct {
				Href string `json:"href"`
			} `json:"file-lists"`
			TryLogs struct {
				Href string `json:"href"`
			} `json:"try-logs"`
		} `json:"links"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Next  string `json:"next"`
			Limit int    `json:"limit"`
		} `json:"pagination"`
	} `json:"meta"`
	Links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		First struct {
			Href string `json:"href"`
		} `json:"first"`
	} `json:"links"`
}
