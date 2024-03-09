package valueobject

import (
	"errors"

	u "github.com/bcicen/go-units"
)

type ExpiryDuration struct {
	value             float64
	unitOrMeasurement string
	unit              u.Unit
}

func (e ExpiryDuration) IsEmpty() bool {

	return e.validation() != nil
}

func (e *ExpiryDuration) validation() error {

	if e.unitOrMeasurement == "" {
		return errors.New("unit or measurement cannot be empty")
	}

	findUnit, err := u.Find(e.unitOrMeasurement)

	if err != nil {
		return err
	}

	e.unit = findUnit

	return nil

}

func setOps(label, short bool, precision int) u.FmtOptions {
	return u.FmtOptions{
		Label:     label,
		Short:     short,
		Precision: precision,
	}
}

func (e ExpiryDuration) Short() string {

	unit := u.NewValue(e.value, e.unit)

	opts := setOps(true, true, 3)

	return unit.Fmt(opts)

}

func (e ExpiryDuration) Long() string {

	unit := u.NewValue(e.value, e.unit)

	opts := setOps(true, false, 3)

	return unit.Fmt(opts)

}

func (e ExpiryDuration) Value() float64 {
	return e.value
}

func (e ExpiryDuration) UnitOrMeasurement() string {
	return e.unitOrMeasurement
}

func NewExpiryDuration(
	value float64,
	unitOrMeasurement string,
) (ExpiryDuration, error) {

	expiryDuration := ExpiryDuration{
		value,
		unitOrMeasurement,
		u.Unit{},
	}

	if err := expiryDuration.validation(); err != nil {
		return ExpiryDuration{}, err
	}

	return expiryDuration, nil

}
