// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	"fmt"

	"github.com/juju/names"
)

func tagFormatError(tag, kind string) error {
	return names.InvalidTagError(tag, kind, fmt.Errorf(`tags must be in the "<kind>-<id>" format`))
}

func unsupportedKindError(tag, kind string) error {
	return names.InvalidTagError(tag, "", fmt.Errorf("unsupported tag kind %q", kind))
}
