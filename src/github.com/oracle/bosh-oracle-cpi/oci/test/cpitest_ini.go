package test

import (
	"github.com/go-ini/ini"
	"os"
	"strings"
)

// CpiTestIni holds the configuration
// used for running the CPI tests
// Default path is ~/.oci/config
type CpiTestIni struct {
	Tenant                 string `ini:"tenancy" mapstructure:"tenant"`
	User                   string `ini:"user" mapstructure:"user"`
	Fingerprint            string `ini:"fingerprint" mapstructure:"fingerprint"`
	Host                   string `ini:"host" mapstructure:"host"`
	KeyFile                string `ini:"key_file" mapstructure:"KeyFile"`
	CompartmentId          string `ini:"compartment" mapstructure:"CompartmentId"`
	Region                 string `ini:"region" mapstructure:"Region"`
	AvailabilityDomain     string `ini:"ad" mapstructure:"AvailabilityDomain"`
	VcnName                string `ini:"vcn" mapstructure:"VcnName"`
	VcnID                  string `ini:"vcn_id" mapstructure:"VcnID"`
	SubnetName             string `ini:"subnet" mapstructure:"SubnetName"`
	SubnetID               string `ini:"subnet_id" mapstructure:"SubnetID"`
	Subnet2Name            string `ini:"subnet2" mapstructure:"Subnet2Name"`
	Subnet2ID              string `ini:"subnet2_id" mapstructure:"Subnet2ID"`
	CpiPrivateKeyPath      string `ini:"cpiPrivateKeyPath" mapstructure:"CpiPrivateKeyPath"`
	CpiPublicKeyPath       string `ini:"cpiPublicKeyPath" mapstructure:"CpiPublicKeyPath"`
	UserPublicKeyPath      string `ini:"userPublicKeyPath" mapstructure:"UserPublicKeyPath"`
	StemcellImageID        string `ini:"stemcellImage" mapstructure:"StemcellImageID"`
	StemcellImageSourceURI string `ini:"stemcellImageSourceURI" mapstructure:"StemcellImageSourceURI"`
	LogLevel               string `ini:"logLevel" mapstructure:"LogLevel"`
	UsePublicIPForSSH      bool   `ini:"usePublicIPForSSH" mapstructure:"UsePublicIPForSSH"`
}

func NewCpiTestIni(filePath string, section string) (CpiTestIni, error) {

	cfg := CpiTestIni{}
	if filePath == "" {
		filePath = defaultIniPath
	}
	if section == "" {
		section = defaultTestSection
	}
	err := (&cfg).loadIni(filePath, section)

	return cfg, err

}

func (b *CpiTestIni) loadIni(filePath string, section string) error {

	path := strings.Replace(filePath, "~", os.Getenv("HOME"), -1)
	t, err := ini.Load(path)
	if err != nil {
		return err
	}
	s, err := t.GetSection(section)
	if err != nil {
		return err
	}
	err = s.MapTo(b)
	b.KeyFile = strings.Replace(b.KeyFile, "~", os.Getenv("HOME"), -1)
	return nil
}
