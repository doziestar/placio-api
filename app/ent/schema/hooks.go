package schema

import (
	"errors"
	"fmt"
	gen "placio-app/ent"
	"placio-app/utility"
)

func ProcessLocation(mutation interface{}, oldLocation string) error {
	var location string
	var setLatitude, setLongitude func(string)
	var setMapCoordinates func(map[string]interface{})

	switch m := mutation.(type) {
	case *gen.UserMutation:
		location, _ = m.Location()
		setLatitude = m.SetLatitude
		setLongitude = m.SetLongitude
		setMapCoordinates = m.SetMapCoordinates
	case *gen.PlaceMutation:
		location, _ = m.Location()
		setLatitude = m.SetLatitude
		setLongitude = m.SetLongitude
		setMapCoordinates = m.SetMapCoordinates
	case *gen.BusinessMutation:
		location, _ = m.Location()
		setLatitude = m.SetLatitude
		setLongitude = m.SetLongitude
		setMapCoordinates = m.SetMapCoordinates
	case *gen.EventMutation:
		location, _ = m.Location()
		setLatitude = m.SetLatitude
		setLongitude = m.SetLongitude
		setMapCoordinates = m.SetMapCoordinates
	default:
		return nil
	}
	//
	//if oldLocation != "" && location == oldLocation || location == "" || location == "<string>" {
	//	return nil
	//}

	if location == oldLocation || location == "" || location == "<string>" {
		return nil
	}

	data, err := utility.GetCoordinates(location)
	if err != nil {
		return err
	}

	if len(data.Features) == 0 {
		return errors.New("no coordinates found")
	}

	latitude := fmt.Sprintf("%f", data.Features[0].Geometry.Coordinates[1])
	longitude := fmt.Sprintf("%f", data.Features[0].Geometry.Coordinates[0])

	setLatitude(latitude)
	setLongitude(longitude)

	mapCoordinates, err := utility.StructToMap(data.Features[0])
	if err != nil {
		return err
	}

	setMapCoordinates(mapCoordinates)
	return nil
}
