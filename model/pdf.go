package model

import (
	"fmt"
	"strings"
	"time"
)

type PDF struct {
	FirstName        string `json:"FirstName" example:"Tony"`
	LastName         string `json:"LastName" example:"Mannino"`
	CompanyName      string `json:"CompanyName" example:"ShareBuilder 401k"`
	PhoneNumber      string `json:"PhoneNumber" example:"+1 (800) 123-4567"`
	AssetsDir        string
	FullName         string
	CompanySignature string
}

func NewPDFModel(workingDir string) *PDF {
	return &PDF{
		AssetsDir: strings.Join([]string{workingDir, "/static"}, ""),
	}
}

// Process some extra fields for our PDF Template
func (p *PDF) FillMissing() {
	p.FullName = strings.Join([]string{p.FirstName, p.LastName}, " ")
	p.CompanySignature = fmt.Sprintf("%s - %s - %v", p.CompanyName, p.PhoneNumber, time.Now().Year())
}
