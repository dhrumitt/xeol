package ui

import (
	"fmt"
	"io"

	"github.com/wagoodman/go-partybus"

	xeolEventParsers "github.com/xeol-io/xeol/xeol/event/parsers"
)

func handleEolScanningFinished(event partybus.Event, reportOutput io.Writer) error {
	// show the report to stdout
	pres, err := xeolEventParsers.ParseEolScanningFinished(event)
	if err != nil {
		return fmt.Errorf("bad CatalogerFinished event: %w", err)
	}

	if err := pres.Present(reportOutput); err != nil {
		return fmt.Errorf("unable to show eol report: %w", err)
	}
	return nil
}

func handleNonRootCommandFinished(event partybus.Event, reportOutput io.Writer) error {
	// show the report to stdout
	result, err := xeolEventParsers.ParseNonRootCommandFinished(event)
	if err != nil {
		return fmt.Errorf("bad NonRootCommandFinished event: %w", err)
	}

	if _, err := reportOutput.Write([]byte(*result)); err != nil {
		return fmt.Errorf("unable to show eol report: %w", err)
	}
	return nil
}
