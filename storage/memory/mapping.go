package memory

import "backend/storage"

func (m *Provider) SelectSingleMapping(mapping *storage.Mapping) error {
	return nil
}

func (m *Provider) SelectDriverMapping(id storage.SessionUUId, mapping *[]storage.Mapping) error {
	return nil
}

func (m *Provider) SelectPassengerMapping(id storage.SessionUUId, mapping *[]storage.Mapping) error {
	return nil
}

func (m *Provider) InsertMapping(mapping *storage.Mapping) error {
	return nil
}

func (m *Provider) UpdateMapping(mapping *storage.Mapping) error {
	return nil
}

func (m *Provider) DeleteMapping(mapping *storage.Mapping) error {
	return nil
}
