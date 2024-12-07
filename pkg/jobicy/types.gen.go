// This file provides types for the API.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package jobicy

import "net/url"

type ListRemoteJobsParams struct {
	// Number of listings to return (default: 50, range: 1-50)
	Count int
	// Filter by job region (default: all regions)
	Geo string
	// Filter by job category (default: all categories)
	Industry string
	// Search by job title and description (default: all jobs)
	Tag string
}

// JobsList defines a model
type JobsList struct {
	APIVersion       string  `json:"apiVersion"`
	DocumentationURL url.URL `json:"documentationUrl"`
	FriendlyNotice   string  `json:"friendlyNotice"`
	JobCount         int     `json:"jobCount"`
	XRayHash         string  `json:"xRayHash"`
	ClientKey        string  `json:"clientKey"`
	LastUpdate       string  `json:"lastUpdate"`
	Jobs             Jobs    `json:"jobs"`
}

// Jobs defines a model
type Jobs []Job

// Job defines a model
type Job struct {
	ID              int      `json:"id"`
	URL             url.URL  `json:"url"`
	JobSlug         string   `json:"jobSlug"`
	JobTitle        string   `json:"jobTitle"`
	CompanyName     string   `json:"companyName"`
	CompanyLogo     url.URL  `json:"companyLogo"`
	JobIndustry     []string `json:"jobIndustry"`
	JobType         []string `json:"jobType"`
	JobGeo          string   `json:"jobGeo"`
	JobLevel        string   `json:"jobLevel"`
	JobExcerpt      string   `json:"jobExcerpt"`
	JobDescription  string   `json:"jobDescription"`
	PubDate         string   `json:"pubDate"`
	AnnualSalaryMin *string  `json:"annualSalaryMin"`
	AnnualSalaryMax *string  `json:"annualSalaryMax"`
	SalaryCurrency  *string  `json:"salaryCurrency"`
}
