package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/cloudfoundry/bosh-bootloader/storage"
	"github.com/cloudfoundry/multierror"
)

type GCPCreateLBs struct {
	terraformManager          terraformApplier
	cloudConfigManager        cloudConfigManager
	stateStore                stateStore
	environmentValidator      environmentValidator
	logger                    logger
	availabilityZoneRetriever availabilityZoneRetriever
}

type GCPCreateLBsConfig struct {
	LBType       string
	CertPath     string
	KeyPath      string
	Domain       string
	SkipIfExists bool
}

type availabilityZoneRetriever interface {
	GetZones(region string) ([]string, error)
}

func NewGCPCreateLBs(terraformManager terraformApplier,
	cloudConfigManager cloudConfigManager,
	stateStore stateStore, environmentValidator environmentValidator, logger logger,
	availabilityZoneRetriever availabilityZoneRetriever,
) GCPCreateLBs {
	return GCPCreateLBs{
		terraformManager:     terraformManager,
		cloudConfigManager:   cloudConfigManager,
		stateStore:           stateStore,
		environmentValidator: environmentValidator,
		logger:               logger,
		availabilityZoneRetriever: availabilityZoneRetriever,
	}
}

func (c GCPCreateLBs) Execute(config GCPCreateLBsConfig, state storage.State) error {
	err := c.terraformManager.ValidateVersion()
	if err != nil {
		return err
	}

	err = c.checkFastFails(config, state)
	if err != nil {
		return err
	}

	if err := c.environmentValidator.Validate(state); err != nil {
		return err
	}

	state.GCP.Zones, err = c.availabilityZoneRetriever.GetZones(state.GCP.Region)
	if err != nil {
		return err
	}

	if config.SkipIfExists && config.LBType == state.LB.Type {
		c.logger.Step(fmt.Sprintf("lb type %q exists, skipping...", config.LBType))
		return nil
	}

	state.LB.Type = config.LBType

	var cert, key []byte
	if config.LBType == "cf" {
		state.LB.Domain = config.Domain

		cert, err = ioutil.ReadFile(config.CertPath)
		if err != nil {
			return err
		}

		state.LB.Cert = string(cert)

		key, err = ioutil.ReadFile(config.KeyPath)
		if err != nil {
			return err
		}

		state.LB.Key = string(key)
	}

	state, err = c.terraformManager.Apply(state)
	if err != nil {
		return handleTerraformError(err, c.stateStore)
	}

	if err := c.stateStore.Set(state); err != nil {
		return err
	}

	if !state.NoDirector {
		err = c.cloudConfigManager.Update(state)
		if err != nil {
			return err
		}
	}

	return nil
}

func (GCPCreateLBs) checkFastFails(config GCPCreateLBsConfig, state storage.State) error {
	if config.LBType == "" {
		return fmt.Errorf("--type is a required flag")
	}

	if !lbExists(config.LBType) {
		return fmt.Errorf("%q is not a valid lb type, valid lb types are: concourse, cf", config.LBType)
	}

	if config.LBType == "cf" {
		errs := multierror.NewMultiError("create-lbs")
		if err := validateCertOrKeyFlag("cert", config.CertPath); err != nil {
			errs.Add(err)
		}
		if err := validateCertOrKeyFlag("key", config.KeyPath); err != nil {
			errs.Add(err)
		}

		if errs.Length() > 0 {
			return errs
		}
	}

	return nil
}

func validateCertOrKeyFlag(flagName, path string) error {
	if path == "" {
		return fmt.Errorf("--%s is required", flagName)
	} else {
		body, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		if string(body) == "" {
			return fmt.Errorf("provided %s file is empty", flagName)
		}
	}
	return nil
}
