package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// MigrationPathFlag is the migration path flag used for finding files to migrate the DB
	MigrationPathFlag string = "migration-path"
)

var (
	errMissingMigrationPath = errors.New("missing migration path, expected to be set")
)

type errInvalidMigrationPath struct {
	Path string
}

func (e *errInvalidMigrationPath) Error() string {
	return fmt.Sprintf("invalid migration path %q", e.Path)
}

// InitMigrationPathFlags initializes the Migration command line flags
func InitMigrationPathFlags(flag *pflag.FlagSet) {
	flag.StringP(MigrationPathFlag, "p", "file:///migrations/app/schema", "Path to the migrations folder")
}

// CheckMigrationPath validates migration command line flags
func CheckMigrationPath(v *viper.Viper) error {
	// Remove any extra quotes around path
	migrationPath := strings.Trim(v.GetString(MigrationPathFlag), "\"")
	if len(migrationPath) == 0 {
		return errMissingMigrationPath
	}
	for _, p := range strings.Split(migrationPath, ";") {
		if len(p) == 0 {
			continue
		}
		if strings.HasPrefix(p, "file://") {
			filesystemPath := p[len("file://"):]
			if _, err := os.Stat(filesystemPath); os.IsNotExist(err) {
				return errors.Wrapf(&errInvalidMigrationPath{Path: filesystemPath}, "Expected %s to be a path in the filesystem", filesystemPath)
			}
		} else if !strings.HasPrefix(p, "s3://") {
			return errors.Wrapf(&errInvalidMigrationPath{Path: p}, "Expected %s to have prefix file:// or s3://", p)
		}
		if strings.HasSuffix(p, "/") {
			return errors.Wrapf(&errInvalidMigrationPath{Path: p}, "Path %s Cannot end in slash", p)
		}
	}
	return nil
}
