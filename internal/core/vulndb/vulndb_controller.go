package vulndb

import (
	"log/slog"
	"strings"

	"github.com/l3montree-dev/flawfix/internal/core"
	"github.com/l3montree-dev/flawfix/internal/database"
	"github.com/l3montree-dev/flawfix/internal/database/models"
	"github.com/l3montree-dev/flawfix/internal/obj"
	"github.com/labstack/echo/v4"

	gocvss20 "github.com/pandatix/go-cvss/20"
	gocvss30 "github.com/pandatix/go-cvss/30"
	gocvss31 "github.com/pandatix/go-cvss/31"
	gocvss40 "github.com/pandatix/go-cvss/40"
)

type repository interface {
	FindAllListPaged(tx database.DB, pageInfo core.PageInfo, filter []core.FilterQuery, sort []core.SortQuery) (core.Paged[models.CVE], error)
	FindCVE(tx database.DB, cveId string) (any, error)
}

type cveHttpController struct {
	cveRepository repository
}

func NewHttpController(cveRepository repository) *cveHttpController {
	return &cveHttpController{
		cveRepository: cveRepository,
	}
}

func (c cveHttpController) ListPaged(ctx core.Context) error {

	pagedResp, err := c.cveRepository.FindAllListPaged(
		nil,
		core.GetPageInfo(ctx),
		core.GetFilterQuery(ctx),
		core.GetSortQuery(ctx),
	)
	if err != nil {
		return echo.NewHTTPError(500, "could not get CVEs").WithInternal(err)
	}

	env := core.GetEnvironmental(ctx)
	e := core.EnvHandle(env)

	if env.AvailabilityRequirements != "" || env.ConfidentialityRequirements != "" || env.IntegrityRequirements != "" {

		for i, cve := range pagedResp.Data {
			risk, vector := riskCalculation(cve, e)

			pagedResp.Data[i].Vector = vector
			pagedResp.Data[i].Risk = risk

		}
	}

	return ctx.JSON(200, pagedResp)
}

func (c cveHttpController) Read(ctx core.Context) error {
	pagedResp, err := c.cveRepository.FindCVE(
		nil,
		core.GetParam(ctx, "cveId"),
	)

	if err != nil {
		return echo.NewHTTPError(500, "could not get CVEs").WithInternal(err)
	}
	cve := pagedResp.(models.CVE)

	env := core.GetEnvironmental(ctx)
	e := core.EnvHandle(env)

	risk, vector := riskCalculation(cve, e)
	cve.Risk = risk
	cve.Vector = vector

	return ctx.JSON(200, cve)
}

func riskCalculation(cve models.CVE, env core.Environmental) (obj.RiskMetrics, string) {
	risk := obj.RiskMetrics{
		BaseScore: float64(cve.CVSS),
	}

	/*
	   //Base Metrics
	   AV : Attack Vector
	   AC : Attack Complexity
	   PR : Privileges Required
	   UI : User Interaction
	   S : Scope
	   C : Confidentiality Impact
	   I : Integrity Impact
	   A : Availability Impact

	   //Environmental (Security Requirements)
	   CR : Confidentiality Requirements
	   IR : Integrity Requirements
	   AR : Availability Requirements

	   //Threat Metrics
	   E : Exploit Maturity
	*/
	/*
		AV := cve.AttackVector
		AC := cve.AttackComplexity
		// AT : AttackRequirements
		PR := cve.PrivilegesRequired
		UI := cve.UserInteraction
		S := cve.Scope
		C := cve.ConfidentialityImpact
		I := cve.IntegrityImpact
		A := cve.AvailabilityImpact

		//Environmental (Security Requirements)
		CR := env.ConfidentialityRequirements
		IR := env.IntegrityRequirements
		AR := env.AvailabilityRequirements

		//Threat Metrics
		E := env.ExploitMaturity
	*/

	vector := cve.Vector
	if vector == "" {
		return risk, vector
	}
	if env.ConfidentialityRequirements != "" {
		vector = vector + "/CR:" + env.ConfidentialityRequirements
	}
	if env.IntegrityRequirements != "" {
		vector = vector + "/IR:" + env.IntegrityRequirements
	}
	if env.AvailabilityRequirements != "" {
		vector = vector + "/AR:" + env.AvailabilityRequirements
	}

	switch {
	case strings.HasPrefix(vector, "CVSS:2.0"):
		// Should be CVSS v2.0 or is invalid
		// build up the temporal score
		// if all affected components have a fixed version, we set it to official fix
		if len(cve.AffectedComponents) > 0 {
			officialFix := true
			for _, component := range cve.AffectedComponents {
				if component.SemverFixed != nil {
					officialFix = false
					break
				}
			}
			if officialFix {
				vector = vector + "/RL:OF"
			} else {
				vector = vector + "/RL:U"
			}
		}

		exploitCodeMaturity := "/E:"
		maturity := "U"
		// check if exploit exist
		if len(cve.Exploits) > 0 {
			// check if there is a verified exploit
			maturity = "POC"
			for _, exploit := range cve.Exploits {
				if exploit.Verified {
					exploitCodeMaturity = "F"
					break
				}
			}
		}
		vector = vector + exploitCodeMaturity + maturity + "/RC:C" // we trust the sources

		cvss, err := gocvss20.ParseVector(vector)
		if err != nil {
			slog.Warn("Error parsing CVSS vector", "vector", vector, "error", err)
			return obj.RiskMetrics{}, vector
		}
		_ = cvss
		risk.WithEnvironment = cvss.EnvironmentalScore()
		risk.WithThreatIntelligence = cvss.TemporalScore()
		risk.WithEnvironmentAndThreatIntelligence = cvss.TemporalScore()
	case strings.HasPrefix(vector, "CVSS:3.0"):
		cvss, err := gocvss30.ParseVector(vector)
		if err != nil {
			slog.Warn("Error parsing CVSS vector", "vector", vector, "error", err)
			return obj.RiskMetrics{}, vector
		}

		vector = cvss.Vector()
		risk.WithEnvironment = cvss.EnvironmentalScore()
		risk.WithThreatIntelligence = cvss.TemporalScore()
		risk.WithEnvironmentAndThreatIntelligence = cvss.TemporalScore()
	case strings.HasPrefix(vector, "CVSS:3.1"):
		// build up the temporal score
		// if all affected components have a fixed version, we set it to official fix
		if len(cve.AffectedComponents) > 0 {
			officialFix := true
			for _, component := range cve.AffectedComponents {
				if component.SemverFixed != nil {
					officialFix = false
					break
				}
			}
			if officialFix {
				vector = vector + "/RL:O"
			} else {
				vector = vector + "/RL:U"
			}
		}

		exploitCodeMaturity := "/E:"
		maturity := "U"
		// check if exploit exist
		if len(cve.Exploits) > 0 {
			// check if there is a verified exploit
			maturity = "F" // functionalv
			for _, exploit := range cve.Exploits {
				if exploit.Verified {
					exploitCodeMaturity = "H"
					break
				}
			}
		}

		vector = vector + exploitCodeMaturity + maturity + "/RC:C" // we trust the sources
		cvss, err := gocvss31.ParseVector(vector)
		if err != nil {
			slog.Warn("Error parsing CVSS vector", "vector", vector, "error", err)
			return obj.RiskMetrics{}, vector
		}

		vector = cvss.Vector()
		risk.WithEnvironment = cvss.EnvironmentalScore()
		risk.WithThreatIntelligence = cvss.TemporalScore()
		risk.WithEnvironmentAndThreatIntelligence = cvss.TemporalScore()
	case strings.HasPrefix(vector, "CVSS:4.0"):
		exploitMaturity := "U"
		if len(cve.Exploits) > 0 {
			exploitMaturity = "P"
		}
		if cve.CISAActionDue != nil {
			exploitMaturity = "A"
		}
		vector = vector + "/E:" + exploitMaturity

		cvss, err := gocvss40.ParseVector(vector)
		if err != nil {
			slog.Warn("Error parsing CVSS vector", "vector", vector, "error", err)
			return obj.RiskMetrics{}, vector
		}

		vector = cvss.Vector()
		risk.WithEnvironmentAndThreatIntelligence = cvss.Score()
	default:
		slog.Warn("Unknown CVSS version", "vector", vector)
		return obj.RiskMetrics{}, vector
	}

	return risk, vector
}
