package validator

import (
	"errors"

	"github.com/Rohith-cyber2617/JS-Intel/internal/config"
)

func Validate(opts *config.Options) error {

	if opts.URL == "" && opts.List == "" && !opts.Update {
		return errors.New("provide a target using -u or -l")
	}

	if opts.Depth < 1 || opts.Depth > 5 {
		return errors.New("depth must be between 1 and 5")
	}

	if opts.Threads < 1 {
		return errors.New("threads must be greater than 0")
	}

	return nil
}
