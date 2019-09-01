package harbor

import (
	"github.com/parnurzeal/gorequest"
	"time"
	"fmt"
)

type ChartRepoResp struct {

}

type ChartRepoRecord struct {
	TotalVersions	int64     `json:"total_versions"`
	Name         	string    `json:"name"`
	LatestVersion  	string    `json:"latest_version"`
	Icon		  	string    `json:"icon"`
	Home		  	string    `json:"home"`
	CreationTime 	time.Time `json:"created"`
	UpdateTime   	time.Time `json:"updated"`
	Deprecated 		bool 	  `json:"deprecated"`
}

type ChartVersionRecord struct {
	Name 			string 		`json:"name"`
	Version 		string 		`json:"version"`
	Description 	string 		`json:"description"`
	AppVersion 		string 		`json:"appVersion"`
	Urls 			[]string	`json:"urls"`
	CreationTime 	time.Time 	`json:"created"`
	UpdateTime   	time.Time 	`json:"updated"`
	Digest 			string 		`json:"digest"`
}

type ChartVersionMetadata struct {
	Name 		string 	`json:"name"`
	Version 	string 	`json:"version"`
	Description string	`json:"description"`
	Digest 		string 	`json:"digest"`
}

type ChartVersionDetailRecord struct {
	Metadata		ChartVersionMetadata	`json:"metadata"`
	Values 			map[string]interface{}		`json:"values"`
	Files 			map[string]string 		`json:"files"`
}

type ChartRepositoriesService struct {
	client *Client
}

type ListChartRepositoriesOption struct {
	ListOptions
	ProjectName string `url:"project_name,omitempty" json:"project_name,omitempty"`
	ProjectId int64  `url:"project_id,omitempty" json:"project_id,omitempty"`
	Q         string `url:"q,omitempty" json:"q,omitempty"`
}

func (s *ChartRepositoriesService) ListChartRepositories(projectName string) ([]ChartRepoRecord, *gorequest.Response, []error) {
	var v []ChartRepoRecord
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("chartrepo/%s/charts", projectName)).
		//Query(*opt).
		EndStruct(&v)
	return v, &resp, errs
}

func (s *ChartRepositoriesService) ListChartVersions(projectName, repoName string) ([]ChartVersionRecord, *gorequest.Response, []error) {
	var v []ChartVersionRecord
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("chartrepo/%s/charts/%s", projectName, repoName)).
		//Query(*opt).
		EndStruct(&v)
	return v, &resp, errs
}

func (s *ChartRepositoriesService) GetChartVersionDetail(projectName, repoName, version string) (ChartVersionDetailRecord, *gorequest.Response, []error) {
	var v ChartVersionDetailRecord
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("chartrepo/%s/charts/%s/%s", projectName, repoName, version)).
		//Query(*opt).
		EndStruct(&v)
	return v, &resp, errs
}