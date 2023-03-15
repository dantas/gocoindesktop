package settings

type Repository struct {
	settings *Settings
	storage  SettingsStorage
}

func NewRepository(storage SettingsStorage) *Repository {
	return &Repository{
		storage: storage,
	}
}

func (r *Repository) Load() (Settings, error) {
	if r.settings != nil {
		return *r.settings, nil
	}

	settings, err := r.storage.Load()

	if err != nil {
		settings = newDefaultSettings()
	}

	r.settings = &settings

	return settings, err
}

func (r *Repository) Save(settings Settings) error {
	if e := r.storage.Save(settings); e != nil {
		return e
	}

	r.settings = &settings

	return nil
}
