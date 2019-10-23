package cli

import (
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// SwaggerFlag is the Public Swagger Flag
	SwaggerFlag string = "swagger"
	// InternalSwaggerFlag is the Internal Swagger Flag
	InternalSwaggerFlag string = "internal-swagger"
	// OrdersSwaggerFlag is the Orders Swagger Flag
	OrdersSwaggerFlag string = "orders-swagger"
	// DPSSwaggerFlag is the DPS Swagger Flag
	DPSSwaggerFlag string = "dps-swagger"
	// AdminSwaggerFlag is the Admin Swagger Flag
	AdminSwaggerFlag string = "admin-swagger"
	// GHCSwaggerFlag is the GHC Swagger Flag
	GHCSwaggerFlag string = "ghc-swagger"
	// ServeSwaggerUIFlag is the Serve Swagger UI Flag
	ServeSwaggerUIFlag string = "serve-swagger-ui"
)

// InitSwaggerFlags initializes the Swagger command line flags
func InitSwaggerFlags(flag *pflag.FlagSet) {
	flag.String(SwaggerFlag, "swagger/api.yaml", "The location of the public API swagger definition")
	flag.String(InternalSwaggerFlag, "swagger/internal.yaml", "The location of the internal API swagger definition")
	flag.String(OrdersSwaggerFlag, "swagger/orders.yaml", "The location of the Orders API swagger definition")
	flag.String(DPSSwaggerFlag, "swagger/dps.yaml", "The location of the DPS API swagger definition")
	flag.String(AdminSwaggerFlag, "swagger/admin.yaml", "The location of the admin API swagger definition")
	flag.String(GHCSwaggerFlag, "swagger/ghc.yaml", "The location of the GHC API swagger definition")
	flag.Bool(ServeSwaggerUIFlag, true, "Whether to serve swagger UI for the APIs")
}

// CheckSwagger validates Swagger command line flags
func CheckSwagger(v *viper.Viper) error {
	swaggerVars := []string{
		SwaggerFlag,
		InternalSwaggerFlag,
		OrdersSwaggerFlag,
		DPSSwaggerFlag,
		GHCSwaggerFlag,
		AdminSwaggerFlag,
	}

	for _, c := range swaggerVars {
		if swaggerFile := v.GetString(c); swaggerFile == "" {
			return errors.Errorf("Swagger file for %s cannot be blank", c)
		}
	}

	return nil
}
